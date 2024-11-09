package group

import (
	"shortlink/internal/user/app/group/command"
	"shortlink/internal/user/app/group/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateGroup command.CreateGroupHandler
	UpdateGroup command.UpdateGroupHandler
	SortGroup   command.SortGroupHandler
	DeleteGroup command.DeleteGroupHandler
}

type Queries struct {
	ListGroup query.ListGroupHandler
}
