package command

import (
	"context"
	"shortlink/internal/user/domain/user"
)

//type DeleteUserCommand struct {
//	UserId string
//}

type DeleteUserHandler struct {
	repo user.Repository
}

func NewDeleteUserHandler(repo user.Repository) DeleteUserHandler {
	if repo == nil {
		panic("nil repo service")
	}

	return DeleteUserHandler{repo: repo}
}

func (h DeleteUserHandler) Handle(ctx context.Context, username string) error {
	return h.repo.DeleteUser(username)
}
