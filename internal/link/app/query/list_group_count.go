package query

import (
	"context"
	"log/slog"
	"shortlink/internal/base/decorator"
	"shortlink/internal/base/metrics"
)

type listGroupCountHandler struct {
	readModel ListGroupCountReadModel
}

type ListGroupLinksCount struct {
	Gids []string `json:"gids"`
}

type ListGroupCountHandler decorator.QueryHandler[ListGroupLinksCount, []GroupLinkCount]

func NewListGroupCountHandler(
	readModel ListGroupCountReadModel,
	logger *slog.Logger,
	metrics metrics.Client,
) ListGroupCountHandler {
	if readModel == nil {
		panic("nil readModel")
	}

	return decorator.ApplyQueryDecorators[ListGroupLinksCount, []GroupLinkCount](
		listGroupCountHandler{readModel: readModel},
		logger,
		metrics,
	)
}

type ListGroupCountReadModel interface {
	ListGroupLinkCount(ctx context.Context, gidList []string) ([]GroupLinkCount, error)
}

func (h listGroupCountHandler) Handle(ctx context.Context, q ListGroupLinksCount) ([]GroupLinkCount, error) {
	return h.readModel.ListGroupLinkCount(ctx, q.Gids)
}
