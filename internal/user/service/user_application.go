package service

import (
	"gorm.io/gorm"
	"shortlink/internal/base/lock"
	"shortlink/internal/user/adapter"
	"shortlink/internal/user/app/user"
	"shortlink/internal/user/app/user/command"
	"shortlink/internal/user/app/user/query"
)

func NewUserApplication(
	db *gorm.DB,
	rdb *redis.Client,
	groupService command.GroupService,
) user.Application {

	repository := adapter.NewUserRepositoryImpl(db, rdb)
	locker := lock.NewRedisLock(rdb)

	// 对 group 领域的依赖项
	//groupRepository := adapter.NewGroupRepositoryImpl(db)
	//groupService := groupcommand.NewCreateGroupHandler(groupRepository)

	a := user.Application{
		Commands: user.Commands{
			UserRegister: command.NewUserRegisterHandler(repository, locker, groupService),
			UserLogin:    command.NewUserLoginHandler(repository),
			UserLogout:   command.NewUserLogoutHandler(repository),
			UpdateUser:   command.NewUpdateUserHandler(repository),
			DeleteUser:   command.NewDeleteUserHandler(repository),
		},
		Queries: user.Queries{
			GetUser:        query.NewGetUserHandler(repository),
			CheckLogin:     query.NewCheckLoginHandler(repository),
			CheckUserExist: query.NewCheckUserExistHandler(repository),
		},
	}

	return a
}
