package command

import (
	"context"
	"shortlink/internal/user/domain/group"
)

type UpdateGroupCommand struct {
	Gid  string
	Name string
}

// UpdateGroupHandler 只能修改组名
type UpdateGroupHandler struct {
	repo group.Repository
}

func NewUpdateGroupHandler(repo group.Repository) UpdateGroupHandler {
	if repo == nil {
		panic("nil repo service")
	}

	return UpdateGroupHandler{repo: repo}
}

func (h UpdateGroupHandler) Handle(ctx context.Context, cmd UpdateGroupCommand) (err error) {
	g := group.NewGroupWithName(cmd.Gid, cmd.Name)
	return h.repo.UpdateGroupName(ctx, g)
}
