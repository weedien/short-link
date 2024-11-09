package command

import (
	"context"
	"shortlink/internal/user/domain/group"
)

type DeleteGroupHandler struct {
	repo group.Repository
}

func NewDeleteGroupHandler(repo group.Repository) DeleteGroupHandler {
	if repo == nil {
		panic("nil repo service")
	}

	return DeleteGroupHandler{repo: repo}
}

func (h DeleteGroupHandler) Handle(ctx context.Context, gid string) (err error) {
	return h.repo.DeleteGroup(ctx, gid)
}
