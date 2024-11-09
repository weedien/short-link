package query

import "time"

// LinkStatsAccessRecord 短链接监控访问记录
type LinkStatsAccessRecord struct {
	// 访客类型
	UvType string
	// 浏览器
	Browser string
	// 操作系统
	Os string
	// ip
	Ip string
	// 访问网络
	Network string
	// 访问设备
	Device string
	// 地区
	Locale string
	// 用户信息
	User string
	// 访问时间
	AccessTime time.Time
}

// LinkStats 短链接监控统计
type LinkStats struct {
	// PV
	Pv int `json:"pv"`
	// UV
	Uv int `json:"uv"`
	// UIP
	Uip int `json:"uip"`
	// 小时访问统计
	Hourly []int `json:"hourly"`
	// 日访问统计
	Daily []LinkStatsAccessDaily `json:"daily"`
	// 周访问统计
	Weekly []int `json:"weekly"`
	// 地区访问统计（仅国内）
	LocationCnStat []LinkStatsLocale `json:"locationCnStat"`
	// 高频访问IP统计
	TopIpStat []LinkStatsTopIp `json:"topIpStat"`
	// 浏览器统计
	BrowserStat []LinkStatsBrowser `json:"browserStat"`
	// 操作系统统计
	OsStat []LinkStatsOs `json:"osStat"`
	// 访客类型统计
	VisitorTypeStat []LinkStatsUv `json:"visitorTypeStat"`
	// 设备统计
	DeviceStat []LinkStatsDevice `json:"deviceStat"`
	// 网络统计
	NetworkStat []LinkStatsNetwork `json:"networkStat"`
}

// LinkStatsAccessDaily 短链接监控访问统计基础响应
type LinkStatsAccessDaily struct {
	// 日期
	Date string `json:"date"`
	// PV
	Pv int `json:"pv"`
	// UV
	Uv int `json:"uv"`
	// UIP
	Uip int `json:"uip"`
}

// LinkStatsBrowser 浏览器统计响应
type LinkStatsBrowser struct {
	// 统计
	Cnt int `json:"count"`
	// 浏览器
	Browser string `json:"browser"`
	// 占比
	Ratio float64 `json:"ratio"`
}

// LinkStatsDevice 设备统计响应
type LinkStatsDevice struct {
	// 统计
	Cnt int `json:"count"`
	// 设备
	Device string `json:"device"`
	// 占比
	Ratio float64 `json:"ratio"`
}

// LinkStatsLocale 地区统计响应
type LinkStatsLocale struct {
	// 统计
	Cnt int `json:"count"`
	// 地区
	Locale string `json:"locale"`
	// 占比
	Ratio float64 `json:"ratio"`
}

// LinkStatsNetwork 网络统计响应
type LinkStatsNetwork struct {
	// 统计
	Cnt int `json:"count"`
	// 网络
	Network string `json:"network"`
	// 占比
	Ratio float64 `json:"ratio"`
}

// LinkStatsOs 操作系统统计响应
type LinkStatsOs struct {
	// 统计
	Cnt int `json:"count"`
	// 操作系统
	Os string `json:"os"`
	// 占比
	Ratio float64 `json:"ratio"`
}

// LinkStatsTopIp 短链接高频访问IP统计响应
type LinkStatsTopIp struct {
	// 统计
	Cnt int `json:"count"`
	// IP
	Ip string `json:"ip"`
}

// LinkStatsUv 短链接监控访问统计UV响应
type LinkStatsUv struct {
	// 统计
	Cnt int `json:"count"`
	// 访客类型
	VisitorType string `json:"visitorType"`
	// 占比
	Ratio float64 `json:"ratio"`
}
