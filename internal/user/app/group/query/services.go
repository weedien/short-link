package query

import "context"

type GroupLinkCountDto struct {
	Gid       string `json:"gid"`
	LinkCount int    `json:"short_link_count"`
}

type LinkService interface {
	ListGroupLinkCount(ctx context.Context, gids []string) ([]GroupLinkCountDto, error)
}
