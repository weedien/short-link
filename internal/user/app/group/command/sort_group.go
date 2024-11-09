package command

import (
	"context"
	"shortlink/internal/user/domain/group"
)

type SortGroupCommand []GroupSortEntry

type GroupSortEntry struct {
	Gid       string
	SortOrder int
}

type SortGroupHandler struct {
	repo group.Repository
}

func NewSortGroupHandler(repo group.Repository) SortGroupHandler {
	if repo == nil {
		panic("nil repo service")
	}

	return SortGroupHandler{repo: repo}
}

func (h SortGroupHandler) Handle(ctx context.Context, cmd SortGroupCommand) (err error) {
	for _, entry := range cmd {
		g := group.NewGroupWithSortOrder(entry.Gid, entry.SortOrder)
		if err = h.repo.UpdateGroupSortOrder(ctx, g); err != nil {
			return err
		}
	}
	return nil
}
