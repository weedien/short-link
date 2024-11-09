package service

import (
	"gorm.io/gorm"
	"shortlink/internal/base/lock"
	"shortlink/internal/user/adapter"
	"shortlink/internal/user/app/group"
	"shortlink/internal/user/app/group/command"
	"shortlink/internal/user/app/group/query"
)

func NewGroupApplication(db *gorm.DB, rdb *redis.Client, linkService query.LinkService) group.Application {

	repository := adapter.NewGroupRepositoryImpl(db, rdb)
	locker := lock.NewRedisLock(rdb)

	a := group.Application{
		Commands: group.Commands{
			CreateGroup: command.NewCreateGroupHandler(repository, locker),
			UpdateGroup: command.NewUpdateGroupHandler(repository),
			DeleteGroup: command.NewDeleteGroupHandler(repository),
			SortGroup:   command.NewSortGroupHandler(repository),
		},
		Queries: group.Queries{
			ListGroup: query.NewListGroupHandler(repository, linkService),
		},
	}

	return a
}
