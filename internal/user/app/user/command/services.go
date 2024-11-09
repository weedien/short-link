package command

import "context"

type GroupService interface {
	CreateGroup(ctx context.Context, username, groupName string) error
}
