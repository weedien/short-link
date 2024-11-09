package query

import (
	"shortlink/internal/base/types"
	"shortlink/internal/link/domain/link"
)

type GroupLinkCount struct {
	// 分组ID
	Gid string
	// 短链接数量
	LinkCount int
}

type Link struct {
	ID           int             `json:"id,omitempty"`
	Domain       string          `json:"domain,omitempty"`
	ShortUri     string          `json:"short_uri,omitempty"`
	FullShortUrl string          `json:"full_short_url,omitempty"`
	OriginalUrl  string          `json:"original_url,omitempty"`
	Gid          string          `json:"gid,omitempty"`
	Status       link.Status     `json:"status,omitempty"`
	CreateType   link.CreateType `json:"create_type"`
	ValidType    link.ValidType  `json:"valid_type"`
	StartDate    types.JsonTime  `json:"start_date,omitempty"`
	EndDate      types.JsonTime  `json:"end_date,omitempty"`
	Desc         string          `json:"desc,omitempty"`
	Favicon      string          `json:"favicon,omitempty"`
	ClickNum     int             `json:"click_num,omitempty"`
	TotalPv      int             `json:"total_pv,omitempty"`
	TotalUv      int             `json:"total_uv,omitempty"`
	TotalUip     int             `json:"total_uip,omitempty"`
	TodayPv      int             `json:"today_pv,omitempty"`
	TodayUv      int             `json:"today_uv,omitempty"`
	TodayUip     int             `json:"today_uip,omitempty"`
}
