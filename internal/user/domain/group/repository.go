package group

import "context"

type Repository interface {
	GetGroupSize(ctx context.Context, username string) (int, error)
	UniqueReturnGid(ctx context.Context) (string, error)
	CreateGroup(ctx context.Context, group Group) error
	UpdateGroupName(ctx context.Context, g Group) error
	UpdateGroupSortOrder(ctx context.Context, g Group) error
	DeleteGroup(ctx context.Context, gid string) error
}
