package query

import (
	"context"
	"log/slog"
	"shortlink/internal/base/decorator"
	"shortlink/internal/base/metrics"
	"time"
)

type groupLinkStatsHandler struct {
	readModel GroupLinkStatsReadModel
}

type GroupLinkStatsHandler decorator.QueryHandler[GroupLinkStats, *LinkStats]

func NewGroupLinkStatsHandler(
	readModel GroupLinkStatsReadModel,
	logger *slog.Logger,
	metricsClient metrics.Client,
) GroupLinkStatsHandler {
	if readModel == nil {
		panic("nil readModel")
	}

	return decorator.ApplyQueryDecorators[GroupLinkStats, *LinkStats](
		groupLinkStatsHandler{readModel},
		logger,
		metricsClient,
	)
}

type GroupLinkStats struct {
	// 分组ID
	Gid string
	// 开始日期
	StartDate time.Time
	// 结束日期
	EndDate time.Time
}

type GroupLinkStatsReadModel interface {
	// GroupLinkStats 获取分组短链接监控数据
	GroupLinkStats(ctx context.Context, param GroupLinkStats) (*LinkStats, error)
}

func (h groupLinkStatsHandler) Handle(ctx context.Context, q GroupLinkStats) (*LinkStats, error) {
	return h.readModel.GroupLinkStats(ctx, q)
}
