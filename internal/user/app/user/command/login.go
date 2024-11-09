package command

import (
	"context"
	"shortlink/internal/user/domain/user"
)

type UserLoginCommand struct {
	Username string
	Password string
	result   string
}

func (c *UserLoginCommand) ExecutionResult() string {
	return c.result
}

type UserLoginHandler struct {
	repo user.Repository
}

func NewUserLoginHandler(repo user.Repository) UserLoginHandler {
	if repo == nil {
		panic("nil repo")
	}

	return UserLoginHandler{repo: repo}
}

func (h UserLoginHandler) Handle(ctx context.Context, cmd *UserLoginCommand) (err error) {
	var token string
	if token, err = h.repo.Login(ctx, cmd.Username, cmd.Password); err != nil {
		return
	}
	cmd.result = token
	return nil
}
