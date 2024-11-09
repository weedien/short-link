package req

import (
	"shortlink/internal/base/types"
	"shortlink/internal/link/domain/link"
)

// LinkCreateReq 创建短链接请求
type LinkCreateReq struct {
	// 原始链接
	OriginalUrl string `json:"original_url,omitempty" validate:"required"`
	// 分组ID
	Gid string `json:"gid,omitempty" validate:"required"`
	// 创建类型 0:接口创建 1:控制台创建
	CreateType *link.CreateType `json:"create_type,omitempty"`
	// 有效期类型 0:永久有效 1:自定义有效期
	ValidType *link.ValidType `json:"valid_type,omitempty"`
	// 有效期 - 开始时间
	StartDate *types.JsonTime `json:"start_date,omitempty" format:"2006-01-02 15:04:05"`
	// 有效期 - 结束时间
	EndDate *types.JsonTime `json:"end_date,omitempty" format:"2006-01-02 15:04:05"`
	// 描述
	Desc string `json:"desc,omitempty"`
}

// LinkBatchCreateReq 批量创建短链接请求
type LinkBatchCreateReq struct {
	// 原始链接集合
	OriginalUrls []string `json:"original_urls,omitempty" validate:"required"`
	// 描述集合
	Descs []string `json:"descs,omitempty"`
	// 分组ID
	Gid string `json:"gid,omitempty" validate:"required"`
	// 创建类型 0:接口创建 1:控制台创建
	CreateType *link.CreateType `json:"create_type,omitempty"`
	// 有效期类型 0:永久有效 1:自定义有效期
	ValidType *link.ValidType `json:"valid_date_type,omitempty"`
	// 有效期 - 开始时间
	StartDate *types.JsonTime `json:"start_date,omitempty" format:"2006-01-02 15:04:05"`
	// 有效期 - 结束时间
	EndDate *types.JsonTime `json:"end_date,omitempty" format:"2006-01-02 15:04:05"`
}

// LinkUpdateReq 更新短链接请求
type LinkUpdateReq struct {
	// 短链接
	ShortUri string `json:"short_uri,omitempty" validate:"required"`
	// 原始分组标识
	//OriginalGid string `json:"original_gid" validate:"required"`
	// 原始链接
	OriginalUrl string `json:"original_url,omitempty"`
	// 分组ID
	Gid string `json:"gid,omitempty"`
	// 状态
	Status link.Status `json:"status,omitempty"`
	// 有效期类型 0:永久有效 1:自定义有效期
	ValidType *link.ValidType `json:"valid_date_type,omitempty"`
	// 有效期 - 开始时间
	StartDate *types.JsonTime `json:"start_date,omitempty" format:"2006-01-02 15:04:05"`
	// 有效期 - 结束时间
	EndDate *types.JsonTime `json:"end_date,omitempty" format:"2006-01-02 15:04:05"`
	// 描述
	Desc *string `json:"desc,omitempty"`
}

// LinkPageReq 分页查询短链接请求
type LinkPageReq struct {
	// 分页参数
	types.PageReq `json:",inline"`
	// 分组ID
	Gid *string `json:"gid,omitempty"`
	// 排序标识
	OrderTag *string `json:"order_tag,omitempty"`
}

// LinkGroupStatReq 分组短链接监控请求
//type LinkGroupStatReq struct {
//	// 分组ID
//	Gid string `json:"gid" validate:"required"`
//	// 开始时间
//	StartTime types.JsonTime `json:"start_time" validate:"required" format:"2006-01-02 15:04:05"`
//	// 结束时间
//	EndTime types.JsonTime `json:"end_time" validate:"required" format:"2006-01-02 15:04:05"`
//}

// LinkStatsAccessRecordReq 短链接监控访问记录请求
//type LinkStatsAccessRecordReq struct {
//	// 分页参数
//	types.PageReq `json:",inline" validate:"required"`
//	// 短链接
//	ShortUri string `json:"short_uri" validate:"required"`
//	// 分组标识
//	Gid string `json:"gid" validate:"required"`
//	// 开始时间
//	StartTime types.JsonTime `json:"start_time" validate:"required" format:"2006-01-02 15:04:05"`
//	// 结束时间
//	EndTime types.JsonTime `json:"end_time" validate:"required" format:"2006-01-02 15:04:05"`
//	// 启用标识
//	Status int `json:"enable_status" validate:"required"`
//}

// LinkStatsReq 短链接监控请求
//type LinkStatsReq struct {
//	// 完整短链接
//	ShortUri string `json:"short_uri" validate:"required"`
//	// 分组标识
//	Gid string `json:"gid" validate:"required"`
//	// 开始时间
//	StartTime types.JsonTime `json:"start_time" validate:"required" format:"2006-01-02 15:04:05"`
//	// 结束时间
//	EndTime types.JsonTime `json:"end_time" validate:"required" format:"2006-01-02 15:04:05"`
//	// 启用标识
//	Status int `json:"enable_status" validate:"required"`
//}

// LinkGroupStatAccessRecordReq 分组短链接监控访问记录请求
//type LinkGroupStatAccessRecordReq struct {
//	// 分页参数
//	types.PageReq `json:",inline" validate:"required"`
//	// 分组ID
//	Gid string `json:"gid" validate:"required"`
//	// 开始时间
//	StartTime types.JsonTime `json:"start_time" validate:"required" format:"2006-01-02 15:04:05"`
//	// 结束时间
//	EndTime types.JsonTime `json:"end_time" validate:"required" format:"2006-01-02 15:04:05"`
//}
