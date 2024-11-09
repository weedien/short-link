package query

import (
	"context"
	"log/slog"
	"shortlink/internal/base/decorator"
	"shortlink/internal/base/metrics"
	"shortlink/internal/base/types"
)

type pageRecycleBinHandler struct {
	readModel PageRecycleBinReadModel
}

type PageRecycleBinHandler decorator.QueryHandler[PageRecycleBin, *types.PageResp[Link]]

func NewPageRecycleBinHandler(
	readModel PageRecycleBinReadModel,
	logger *slog.Logger,
	metricsClient metrics.Client,
) PageRecycleBinHandler {
	if readModel == nil {
		panic("nil readModel")
	}

	return decorator.ApplyQueryDecorators[PageRecycleBin, *types.PageResp[Link]](
		pageRecycleBinHandler{readModel},
		logger,
		metricsClient,
	)
}

type PageRecycleBin struct {
	types.PageReq
	Gids []string `json:"gids"`
}

type PageRecycleBinReadModel interface {
	PageRecycleBin(ctx context.Context, param PageRecycleBin) (*types.PageResp[Link], error)
}

func (h pageRecycleBinHandler) Handle(ctx context.Context, cmd PageRecycleBin) (*types.PageResp[Link], error) {
	return h.readModel.PageRecycleBin(ctx, cmd)
}
