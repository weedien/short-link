package group

import (
	"errors"
	"fmt"
)

const (
	MaxGroupSize = 16
)

var (
	ErrGroupSizeExceed  = fmt.Errorf("group size exceed %d", MaxGroupSize)
	ErrGenGroupUniqueID = errors.New("generate unique group id failed")
)

type Group struct {
	gid       string
	username  string
	name      string
	sortOrder int
}

func NewGroup(gid, username, name string, sortOrder int) Group {
	return Group{
		gid:       gid,
		username:  username,
		name:      name,
		sortOrder: sortOrder,
	}
}

func NewGroupWithName(gid string, name string) Group {
	return Group{
		gid:  gid,
		name: name,
	}
}

func NewGroupWithSortOrder(gid string, sortOrder int) Group {
	return Group{
		gid:       gid,
		sortOrder: sortOrder,
	}
}

func (g Group) Gid() string {
	return g.gid
}

func (g Group) Username() string {
	return g.username
}

func (g Group) Name() string {
	return g.name
}

func (g Group) SortOrder() int {
	return g.sortOrder
}
