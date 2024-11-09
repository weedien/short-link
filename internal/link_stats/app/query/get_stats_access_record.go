package query

import (
	"context"
	"log/slog"
	"shortlink/internal/base/decorator"
	"shortlink/internal/base/metrics"
	"shortlink/internal/base/types"
	"time"
)

// getLinkStatsAccessRecordHandler 获取单个短链接指定时间内访问记录监控数据
type getLinkStatsAccessRecordHandler struct {
	readModel GetLinkStatsAccessRecordReadModel
}

type GetLinkStatsAccessRecordHandler decorator.QueryHandler[GetLinkStatsAccessRecord, *types.PageResp[LinkStatsAccessRecord]]

func NewGetLinkStatsAccessRecordHandler(
	readModel GetLinkStatsAccessRecordReadModel,
	logger *slog.Logger,
	metricsClient metrics.Client,
) GetLinkStatsAccessRecordHandler {
	if readModel == nil {
		panic("nil readModel")
	}

	return decorator.ApplyQueryDecorators[GetLinkStatsAccessRecord, *types.PageResp[LinkStatsAccessRecord]](
		getLinkStatsAccessRecordHandler{readModel},
		logger,
		metricsClient,
	)
}

type GetLinkStatsAccessRecord struct {
	types.PageReq
	// 完整短链接
	FullShortUrl string
	// 分组ID
	Gid string
	// 开始日期
	StartDate time.Time
	// 结束日期
	EndDate time.Time
	// 启用标识
	Status int
}

type GetLinkStatsAccessRecordReadModel interface {
	// GetLinkStatsAccessRecord 获取单个短链接指定时间内访问记录监控数据
	GetLinkStatsAccessRecord(ctx context.Context, param GetLinkStatsAccessRecord) (*types.PageResp[LinkStatsAccessRecord], error)
}

func (h getLinkStatsAccessRecordHandler) Handle(ctx context.Context, query GetLinkStatsAccessRecord) (d *types.PageResp[LinkStatsAccessRecord], err error) {
	return h.readModel.GetLinkStatsAccessRecord(ctx, query)
}
