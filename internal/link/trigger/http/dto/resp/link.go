package resp

import (
	"time"
)

// LinkCreateResp 短链接创建响应
type LinkCreateResp struct {
	// 分组标识
	Gid string `json:"gid"`
	// 原始链接
	OriginalUrl string `json:"originalUrl"`
	// 短链接
	FullShortUrl string `json:"fullShortUrl"`
}

// LinkBatchCreateResp 短链接批量创建响应
type LinkBatchCreateResp struct {
	// 成功数量
	SuccessCount int `json:"successCount"`
	// 批量创建返回参数
	LinkInfos []LinkBaseInfoDTO `json:"linkInfos"`
}

// LinkGroupCountQueryResp 短链接分组数量查询响应
type LinkGroupCountQueryResp []GroupCountDTO

type GroupCountDTO struct {
	// 分组标识
	Gid string `json:"gid"`
	// 短链接数量
	LinkCount int `json:"link_count"`
}

// LinkPageResp 短链接分页查询响应
type LinkPageResp struct {
	// ID
	ID int64 `json:"id"`
	// 域名
	Domain string `json:"domain"`
	// 短链接
	ShortUri string `json:"shortUri"`
	// 完整短链接
	FullShortUrl string `json:"fullShortUrl"`
	// 原始链接
	OriginalUrl string `json:"originalUrl"`
	// 分组标识
	Gid string `json:"gid"`
	// 有效期类型 0：永久有效 1：自定义
	ValidDateType int `json:"validDateType"`
	// 启用标识 0：启用 1：未启用
	Status int `json:"enableStatus"`
	// 有效期
	ValidDate time.Time `json:"validDate" format:"2006-01-02 15:04:05"`
	// 创建时间
	CreateTime time.Time `json:"createTime" format:"2006-01-02 15:04:05"`
	// 描述
	Description string `json:"description"`
	// 网站标识
	Favicon string `json:"favicon"`
	// 历史PV
	TotalPv int `json:"totalPv"`
	// 今日PV
	TodayPv int `json:"todayPv"`
	// 历史UV
	TotalUv int `json:"totalUv"`
	// 今日UV
	TodayUv int `json:"todayUv"`
	// 历史UIP
	TotalUip int `json:"totalUip"`
	// 今日UIP
	TodayUip int `json:"todayUip"`
}

// LinkBaseInfoDTO 短链接基本信息响应
type LinkBaseInfoDTO struct {
	// 分组标识
	Gid string `json:"gid"`
	// 原始链接
	OriginalUrl string `json:"originalUrl"`
	// 短链接
	FullShortUrl string `json:"fullShortUrl"`
}
