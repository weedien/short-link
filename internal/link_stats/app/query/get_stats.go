package query

import (
	"context"
	"log/slog"
	"shortlink/internal/base/decorator"
	"shortlink/internal/base/metrics"
	"time"
)

type getLinkStatsHandler struct {
	readModel GetLinkStatsReadModel
}

type GetLinkStatsHandler decorator.QueryHandler[GetLinkStats, *LinkStats]

func NewGetLinkStatsHandler(
	readModel GetLinkStatsReadModel,
	logger *slog.Logger,
	metricsClient metrics.Client,
) GetLinkStatsHandler {
	if readModel == nil {
		panic("nil readModel")
	}

	return decorator.ApplyQueryDecorators[GetLinkStats, *LinkStats](
		getLinkStatsHandler{readModel},
		logger,
		metricsClient,
	)
}

type GetLinkStats struct {
	// 完整短链接
	FullShortUrl string
	// 分组ID
	Gid string
	// 开始日期
	StartDate time.Time
	// 结束日期
	EndDate time.Time
	// 启用标识
	Status string
}

type GetLinkStatsReadModel interface {
	// GetLinkStats 获取单个短链接监控数据
	GetLinkStats(ctx context.Context, param GetLinkStats) (res *LinkStats, err error)
}

func (h getLinkStatsHandler) Handle(ctx context.Context, q GetLinkStats) (res *LinkStats, err error) {
	return h.readModel.GetLinkStats(ctx, q)
}
