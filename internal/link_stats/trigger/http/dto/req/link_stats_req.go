package req

import (
	"shortlink/internal/base/types"
	"time"
)

// LinkGroupStatsAccessRecordReq 分组短链接监控访问记录请求
type LinkGroupStatsAccessRecordReq struct {
	// 分页参数
	types.PageReq `json:",inline" validate:"required"`
	// 分组ID
	Gid string `json:"gid" validate:"required"`
	// 开始时间
	StartTime time.Time `json:"start_time" validate:"required" format:"2006-01-02 15:04:05"`
	// 结束时间
	EndTime time.Time `json:"end_time" validate:"required" format:"2006-01-02 15:04:05"`
}

// LinkGroupStatsReq 分组短链接监控请求
type LinkGroupStatsReq struct {
	// 分组ID
	Gid string `json:"gid" validate:"required"`
	// 开始时间
	StartTime time.Time `json:"start_time" validate:"required" format:"2006-01-02 15:04:05"`
	// 结束时间
	EndTime time.Time `json:"end_time" validate:"required" format:"2006-01-02 15:04:05"`
}

// LinkPageReq 分页查询短链接请求
type LinkPageReq struct {
	// 分页参数
	types.PageReq `json:",inline" validate:"required"`
	// 分组ID
	Gid string `json:"gid" validate:"required"`
	// 排序标识
	OrderTag string `json:"order_tag" validate:"required"`
}

// LinkStatsAccessRecordReq 短链接监控访问记录请求
type LinkStatsAccessRecordReq struct {
	// 分页参数
	types.PageReq `json:",inline" validate:"required"`
	// 完整短链接
	FullShortUrl string `json:"full_short_url" validate:"required"`
	// 分组标识
	Gid string `json:"gid" validate:"required"`
	// 开始时间
	StartTime time.Time `json:"start_time" validate:"required" format:"2006-01-02 15:04:05"`
	// 结束时间
	EndTime time.Time `json:"end_time" validate:"required" format:"2006-01-02 15:04:05"`
	// 启用标识
	Status int `json:"enable_status" validate:"required"`
}

// LinkStatsReq 短链接监控请求
type LinkStatsReq struct {
	// 完整短链接
	FullShortUrl string `json:"full_short_url" validate:"required"`
	// 分组标识
	Gid string `json:"gid" validate:"required"`
	// 开始时间
	StartTime time.Time `json:"start_time" validate:"required" format:"2006-01-02 15:04:05"`
	// 结束时间
	EndTime time.Time `json:"end_time" validate:"required" format:"2006-01-02 15:04:05"`
	// 启用标识
	Status int `json:"enable_status" validate:"required"`
}
