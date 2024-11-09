package command

import (
	"context"
	"shortlink/internal/base/error_no"
	"shortlink/internal/user/domain/user"
)

type UserLogoutCommand struct {
	Username string
	Token    string
}

type UserLogoutHandler struct {
	repo user.Repository
}

func NewUserLogoutHandler(repo user.Repository) UserLogoutHandler {
	if repo == nil {
		panic("nil repo")
	}
	return UserLogoutHandler{repo: repo}
}

func (h UserLogoutHandler) Handle(ctx context.Context, cmd UserLogoutCommand) (err error) {
	var login bool
	if login, err = h.repo.CheckLogin(ctx, cmd.Username, cmd.Token); err != nil {
		return err
	}
	if !login {
		return error_no.InvalidTokenOrUnloggedLoginUser
	}
	return h.repo.InvalidateToken(ctx, cmd.Username, cmd.Token)
}
