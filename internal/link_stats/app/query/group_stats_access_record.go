package query

import (
	"context"
	"log/slog"
	"shortlink/internal/base/decorator"
	"shortlink/internal/base/metrics"
	"shortlink/internal/base/types"
	"time"
)

type groupLinkStatsAccessRecordHandler struct {
	readModel GroupLinkStatsAccessRecordReadModel
}

type GroupLinkStatsAccessRecordHandler decorator.QueryHandler[GroupLinkStatsAccessRecord, *types.PageResp[LinkStatsAccessRecord]]

func NewGroupLinkStatsAccessRecordHandler(
	readModel GroupLinkStatsAccessRecordReadModel,
	logger *slog.Logger,
	metricsClient metrics.Client,
) GroupLinkStatsAccessRecordHandler {
	if readModel == nil {
		panic("nil readModel")
	}

	return decorator.ApplyQueryDecorators[GroupLinkStatsAccessRecord, *types.PageResp[LinkStatsAccessRecord]](
		groupLinkStatsAccessRecordHandler{readModel},
		logger,
		metricsClient,
	)
}

type GroupLinkStatsAccessRecord struct {
	// 分页参数
	types.PageReq
	// 分组ID
	Gid string
	// 开始日期
	StartDate time.Time
	// 结束日期
	EndDate time.Time
}

type GroupLinkStatsAccessRecordReadModel interface {
	// GroupLinkStatsAccessRecord 获取分组指定时间内访问记录监控数据
	GroupLinkStatsAccessRecord(ctx context.Context, param GroupLinkStatsAccessRecord) (*types.PageResp[LinkStatsAccessRecord], error)
}

func (h groupLinkStatsAccessRecordHandler) Handle(ctx context.Context, q GroupLinkStatsAccessRecord) (*types.PageResp[LinkStatsAccessRecord], error) {
	return h.readModel.GroupLinkStatsAccessRecord(ctx, q)
}
