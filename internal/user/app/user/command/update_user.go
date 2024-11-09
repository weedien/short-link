package command

import (
	"context"
	"shortlink/internal/base/error_no"
	"shortlink/internal/user/domain/user"
)

type UpdateUserCommand struct {
	Username string
	Password string
	RealName string
	Email    string
	Phone    string
}

type UpdateUserHandler struct {
	repo user.Repository
}

func NewUpdateUserHandler(repo user.Repository) UpdateUserHandler {
	if repo == nil {
		panic("nil repo")
	}
	return UpdateUserHandler{repo: repo}
}

func (h UpdateUserHandler) Handle(ctx context.Context, cmd UpdateUserCommand) error {
	currentUsername := ctx.Value("username").(string)
	if currentUsername != cmd.Username {
		return error_no.UserForbidden
	}

	u := user.NewUser(cmd.Username, cmd.Password, cmd.RealName, cmd.Email, cmd.Phone)
	return h.repo.UpdateUser(ctx, &u)
}
