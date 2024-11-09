package query

import (
	"context"
	"shortlink/internal/user/domain/user"
)

type CheckUserExistHandler struct {
	repo user.Repository
}

func NewCheckUserExistHandler(repo user.Repository) CheckUserExistHandler {
	return CheckUserExistHandler{repo: repo}
}

func (h CheckUserExistHandler) Handle(ctx context.Context, username string) (bool, error) {
	return h.repo.CheckUserExist(ctx, username)
}
