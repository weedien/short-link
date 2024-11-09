package resp

import (
	"shortlink/internal/base/types"
	"time"
)

// LinkStatsResp 短链接监控统计响应
type LinkStatsResp struct {
	// PV
	Pv int `json:"pv"`
	// UV
	Uv int `json:"uv"`
	// UIP
	Uip int `json:"uip"`
	// 小时访问统计
	Hourly []LinkStatsAccessBaseDTO `json:"hourly"`
	// 日访问统计
	Daily []LinkStatsAccessBaseDTO `json:"daily"`
	// 周访问统计
	Weekly []LinkStatsAccessBaseDTO `json:"weekly"`
	// 地区访问统计（仅国内）
	LocationCnStat []LinkStatsLocationDTO `json:"locationCnStat"`
	// 高频访问IP统计
	TopIpStat []LinkStatsTopIpDTO `json:"topIpStat"`
	// 浏览器统计
	BrowserStat []LinkStatsBrowserDTO `json:"browserStat"`
	// 操作系统统计
	OsStat []LinkStatsOsDTO `json:"osStat"`
	// 访客类型统计
	VisitorTypeStat []LinkStatsUvDTO `json:"visitorTypeStat"`
	// 设备统计
	DeviceStat []LinkStatsDeviceDTO `json:"deviceStat"`
	// 网络统计
	NetworkStat []LinkStatsNetworkDTO `json:"networkStat"`
}

// LinkStatsAccessBaseDTO 短链接监控访问统计基础响应
type LinkStatsAccessBaseDTO struct {
	// 日期
	Date time.Time `json:"date" format:"2006-01-02"`
	// PV
	Pv int `json:"pv"`
	// UV
	Uv int `json:"uv"`
	// UIP
	Uip int `json:"uip"`
}

// LinkStatsBrowserDTO 浏览器统计响应
type LinkStatsBrowserDTO struct {
	// 统计
	Count int `json:"count"`
	// 浏览器
	Browser string `json:"browser"`
	// 占比
	Ratio float64 `json:"ratio"`
}

// LinkStatsDeviceDTO 设备统计响应
type LinkStatsDeviceDTO struct {
	// 统计
	Count int `json:"count"`
	// 设备
	Device string `json:"device"`
	// 占比
	Ratio float64 `json:"ratio"`
}

// LinkStatsLocationDTO 地区统计响应
type LinkStatsLocationDTO struct {
	// 统计
	Count int `json:"count"`
	// 地区
	Location string `json:"location"`
	// 占比
	Ratio float64 `json:"ratio"`
}

// LinkStatsNetworkDTO 网络统计响应
type LinkStatsNetworkDTO struct {
	// 统计
	Count int `json:"count"`
	// 网络
	Network string `json:"network"`
	// 占比
	Ratio float64 `json:"ratio"`
}

// LinkStatsOsDTO 操作系统统计响应
type LinkStatsOsDTO struct {
	// 统计
	Count int `json:"count"`
	// 操作系统
	Os string `json:"os"`
	// 占比
	Ratio float64 `json:"ratio"`
}

// LinkStatsTopIpDTO 短链接高频访问IP统计响应
type LinkStatsTopIpDTO struct {
	// 统计
	Count int `json:"count"`
	// IP
	Ip string `json:"ip"`
}

// LinkStatsUvDTO 短链接监控访问统计UV响应
type LinkStatsUvDTO struct {
	// 统计
	Count int `json:"count"`
	// 访客类型
	VisitorType string `json:"visitorType"`
	// 占比
	Ratio float64 `json:"ratio"`
}

// LinkStatsAccessRecordResp 短链接监控访问统计记录响应
type LinkStatsAccessRecordResp types.PageResp[LinkStatsAccessRecordDTO]

type LinkStatsAccessRecordDTO struct {
	// 访客类型
	UvType string `json:"UvType"`
	// 浏览器
	Browser string `json:"browser"`
	// 操作系统
	Os string `json:"os"`
	// IP
	Ip string `json:"ip"`
	// 访问网络
	Network string `json:"network"`
	// 访问设备
	Device string `json:"device"`
	// 地区
	Location string `json:"location"`
	// 用户信息
	UserAgent string `json:"userAgent"`
	// 访问时间
	AccessTime time.Time `json:"accessTime" format:"2006-01-02 15:04:05"`
}
