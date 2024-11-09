package query

import (
	"context"
	"shortlink/internal/user/domain/user"
)

type CheckLogin struct {
	Username string
	Token    string
}

type CheckLoginHandler struct {
	repo user.Repository
}

func NewCheckLoginHandler(repo user.Repository) CheckLoginHandler {
	if repo == nil {
		panic("nil repo")
	}
	return CheckLoginHandler{repo: repo}
}

func (h CheckLoginHandler) Handle(ctx context.Context, q CheckLogin) (bool, error) {
	return h.repo.CheckLogin(ctx, q.Username, q.Token)
}
