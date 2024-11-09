package command

import (
	"context"
	"shortlink/internal/base/lock"
	"shortlink/internal/link/constant"
	"shortlink/internal/user/domain/group"
	"time"
)

type CreateGroupCommand struct {
	Username  string
	GroupName string
}

type CreateGroupHandler struct {
	repo   group.Repository
	locker lock.DistributedLock
}

func NewCreateGroupHandler(repo group.Repository, locker lock.DistributedLock) CreateGroupHandler {
	if repo == nil {
		panic("nil repo service")
	}

	if locker == nil {
		panic("nil locker service")
	}

	return CreateGroupHandler{repo: repo, locker: locker}
}

func (h CreateGroupHandler) Handle(ctx context.Context, cmd CreateGroupCommand) (err error) {
	lockKey := constant.LockGroupCreateKey + cmd.GroupName
	if _, err = h.locker.Acquire(ctx, lockKey, 1*time.Hour); err != nil {
		return err
	}
	defer func(ctx context.Context, lockKey string) {
		_ = h.locker.Release(ctx, lockKey)
	}(ctx, lockKey)
	// 最大分组数限制
	var size int
	if size, err = h.repo.GetGroupSize(ctx, cmd.Username); err != nil {
		return err
	}
	if size >= group.MaxGroupSize {
		return group.ErrGroupSizeExceed
	}
	// 生成唯一分组ID
	retryCount, maxRetries := 0, 10
	var gid string
	for retryCount < maxRetries {
		if gid, err = h.repo.UniqueReturnGid(ctx); err != nil {
			return err
		}
		if gid != "" {
			g := group.NewGroup(gid, cmd.Username, cmd.GroupName, size)
			if err = h.repo.CreateGroup(ctx, g); err != nil {
				return err
			}
			break
		}
		retryCount++
	}
	if gid == "" {
		return group.ErrGenGroupUniqueID
	}
	return nil
}

// CreateGroup 创建分组
// 为其他领域提供服务
func (h CreateGroupHandler) CreateGroup(
	ctx context.Context,
	username, groupName string,
) error {
	return h.Handle(ctx, CreateGroupCommand{
		Username:  username,
		GroupName: groupName,
	})
}
