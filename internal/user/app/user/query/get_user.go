package query

import (
	"context"
	"shortlink/internal/base/error_no"
	"shortlink/internal/user/domain/user"
)

type GetUserHandler struct {
	repo user.Repository
}

func NewGetUserHandler(repo user.Repository) GetUserHandler {
	return GetUserHandler{repo: repo}
}

func (h GetUserHandler) Handle(ctx context.Context, username string) (res *user.User, err error) {
	if res, err = h.repo.GetUser(ctx, username); err != nil {
		return
	}
	if res == nil {
		return nil, error_no.UserNotExist
	}
	return
}
