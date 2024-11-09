package command

import (
	"context"
	"shortlink/internal/base/error_no"
	"shortlink/internal/base/lock"
	"shortlink/internal/link/constant"
	"shortlink/internal/user/domain/user"
	"time"
)

type UserRegisterHandler struct {
	repo         user.Repository
	locker       lock.DistributedLock
	groupService GroupService
}

type UserRegisterCommand struct {
	Username string
	Password string
	RealName string
	Phone    string
	Email    string
}

func NewUserRegisterHandler(
	repo user.Repository,
	locker lock.DistributedLock,
	groupService GroupService,
) UserRegisterHandler {
	if repo == nil {
		panic("nil repo service")
	}
	if locker == nil {
		panic("nil locker service")
	}
	if groupService == nil {
		panic("nil group service")
	}
	return UserRegisterHandler{repo: repo, locker: locker, groupService: groupService}
}

func (h UserRegisterHandler) Handle(ctx context.Context, cmd UserRegisterCommand) error {
	if exist, err := h.repo.CheckUserExist(ctx, cmd.Username); err != nil {
		return err
	} else if exist {
		return error_no.UserExist
	}
	// 获取分布式锁
	lockKey := constant.LockUserRegisterKey + cmd.Username
	if _, err := h.locker.Acquire(ctx, lockKey, 1*time.Hour); err != nil {
		return error_no.LockAcquireFailed
	}
	defer func(ctx context.Context, lockKey string) {
		_ = h.locker.Release(ctx, lockKey)
	}(ctx, lockKey)
	// 再次检查用户是否存在
	if exist, err := h.repo.CheckUserExist(ctx, cmd.Username); err != nil {
		return err
	} else if exist {
		return error_no.UserExist
	}
	// 创建用户
	u := user.NewUser(cmd.Username, cmd.Password, cmd.RealName, cmd.Email, cmd.Phone)
	if err := h.repo.CreateUser(ctx, &u); err != nil {
		return err
	}
	// 加入分组
	if err := h.groupService.CreateGroup(ctx, u.Name(), "默认分组"); err != nil {
		return err
	}
	// 加入布隆过滤器
	//if err := h.repo.AddUserToBloomFilter(ctx, u.Tag()); err != nil {
	//	return err
	//}
	return nil
}
