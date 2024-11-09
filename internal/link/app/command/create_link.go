package command

import (
	"context"
	"fmt"
	"log/slog"
	"shortlink/internal/base/cache"
	"shortlink/internal/base/decorator"
	"shortlink/internal/base/lock"
	"shortlink/internal/base/metrics"
	"shortlink/internal/link/common/constant"
	"shortlink/internal/link/domain"
	"shortlink/internal/link/domain/link"
	"time"
)

type createLinkHandler struct {
	repo             domain.Repository
	locker           lock.DistributedLock
	linkFactory      *link.Factory
	distributedCache cache.DistributedCache
}

type CreateLink struct {
	// 原始链接
	OriginalUrl string
	// 分组ID
	Gid string
	// 创建类型 0:接口创建 1:控制台创建
	CreateType *link.CreateType
	// 有效期类型 0:永久有效 1:自定义有效期
	ValidType *link.ValidType
	// 有效期 - 开始时间
	StartDate *time.Time
	// 有效期 - 结束时间
	EndDate *time.Time
	// 描述
	Desc string
	// 是否加锁
	WithLock bool
	// 执行结果
	result *CreateLinkResult
}

type CreateLinkResult struct {
	Gid          string
	FullShortUrl string
	OriginalUrl  string
}

func (c *CreateLink) ExecutionResult() *CreateLinkResult {
	return c.result
}

type CreateLinkHandler decorator.CommandHandler[*CreateLink]

func NewCreateLinkHandler(
	linkFactory *link.Factory,
	repo domain.Repository,
	locker lock.DistributedLock,
	logger *slog.Logger,
	metrics metrics.Client,
) CreateLinkHandler {
	if repo == nil {
		panic("nil repo")
	}
	if locker == nil {
		panic("nil locker")
	}

	return decorator.ApplyCommandDecorators[*CreateLink](
		createLinkHandler{repo: repo, locker: locker, linkFactory: linkFactory},
		logger,
		metrics,
	)
}

func (h createLinkHandler) Handle(
	ctx context.Context,
	cmd *CreateLink,
) (err error) {

	// 获取分布式锁
	if cmd.WithLock {
		lockKey := fmt.Sprintf(constant.LinkCreateLockKey, cmd.OriginalUrl)
		if _, err = h.locker.Acquire(ctx, lockKey, constant.DefaultTimeOut); err != nil {
			return
		}
		defer func() {
			if releaseErr := h.locker.Release(ctx, lockKey); releaseErr != nil {
				err = fmt.Errorf("释放锁异常: %w", releaseErr)
			}
		}()
	}

	// 查询分组短链接数量
	var count int
	if count, err = h.repo.CountLinksByGid(ctx, cmd.Gid); err != nil {
		return err
	}

	// linkFactory 校验数量是否超限
	if err = h.linkFactory.CheckGroupLinkCount(count); err != nil {
		return err
	}

	// 创建短链接实体
	lk := &link.Link{}
	lk, err = h.linkFactory.NewAvailableLink(
		cmd.OriginalUrl, cmd.Gid, cmd.CreateType, cmd.ValidType, cmd.StartDate, cmd.EndDate, cmd.Desc,
		func(shortUri string) (exists bool, err error) {
			if exists, err = h.repo.ShortUriExists(ctx, shortUri); err != nil {
				return exists, err
			}
			return exists, nil
		},
	)
	if err != nil {
		return
	}

	// 持久化短链接
	if err = h.repo.CreateLink(ctx, lk); err != nil {
		return
	}

	// 返回结果
	cmd.result = &CreateLinkResult{
		Gid:          lk.Gid(),
		FullShortUrl: lk.FullShortUrl(),
		OriginalUrl:  lk.OriginalUrl(),
	}
	return
}
