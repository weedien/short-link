package command

import (
	"context"
	"log/slog"
	"shortlink/internal/base/decorator"
	"shortlink/internal/base/metrics"
	"shortlink/internal/link/domain"
	"shortlink/internal/link/domain/link"
	"time"
)

type createLinkBatchHandler struct {
	repo        domain.Repository
	linkFactory *link.Factory
}

type CreateLinkBatchHandler decorator.CommandHandler[*CreateLinkBatch]

func NewCreateLinkBatchHandler(
	linkFactory *link.Factory,
	repo domain.Repository,
	logger *slog.Logger,
	metricsClient metrics.Client,
) CreateLinkBatchHandler {
	if linkFactory == nil {
		panic("linkFactory is nil")
	}
	if repo == nil {
		panic("repo is nil")
	}

	return decorator.ApplyCommandDecorators[*CreateLinkBatch](
		createLinkBatchHandler{repo: repo, linkFactory: linkFactory},
		logger,
		metricsClient,
	)
}

type CreateLinkBatch struct {
	// 原始链接
	OriginalUrls []string
	// 描述
	Descs []string
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
	// 执行结果
	result *CreateLinkBatchResult
}

type CreateLinkBatchResult struct {
	SuccessCount int
	LinkInfos    []CreateLinkResult
}

func (c CreateLinkBatch) ExecutionResult() *CreateLinkBatchResult {
	return c.result
}

func (h createLinkBatchHandler) Handle(
	ctx context.Context,
	cmd *CreateLinkBatch,
) (err error) {

	lks := make([]*link.Link, 0)
	linkInfos := make([]CreateLinkResult, len(cmd.OriginalUrls))
	for idx, originalUrl := range cmd.OriginalUrls {
		desc := ""
		if idx < len(cmd.Descs) {
			desc = cmd.Descs[idx]
		}

		lk := &link.Link{}
		lk, err = h.linkFactory.NewAvailableLink(
			originalUrl, cmd.Gid, cmd.CreateType, cmd.ValidType, cmd.StartDate, cmd.EndDate, desc,
			func(shortUri string) (exists bool, err error) {
				if exists, err = h.repo.ShortUriExists(ctx, shortUri); err != nil {
					return exists, err
				}
				return exists, nil
			},
		)
		if err != nil {
			return err
		}

		lks = append(lks, lk)

		linkInfos[idx] = CreateLinkResult{
			Gid:          lk.Gid(),
			FullShortUrl: lk.FullShortUrl(),
			OriginalUrl:  lk.OriginalUrl(),
		}
	}

	if err = h.repo.CreateLinkBatch(ctx, lks); err != nil {
		return err
	}

	cmd.result = &CreateLinkBatchResult{
		SuccessCount: len(linkInfos),
		LinkInfos:    linkInfos,
	}

	return nil
}
