package po

import (
	"gorm.io/gorm"
	"time"
)

const TableNameLinkStats = "link_stats"

// LinkStats mapped from table <link>
type LinkStats struct {
	ID          int            `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`
	Gid         string         `gorm:"column:gid;not null;comment:分组标识" json:"gid"`
	ShortUri    string         `gorm:"column:short_uri;not null;comment:短链接" json:"short_uri"`
	ClickNum    int            `gorm:"column:click_num;comment:点击量" json:"click_num"`
	TotalPv     int            `gorm:"column:total_pv;comment:历史PV" json:"total_pv"`
	TotalUv     int            `gorm:"column:total_uv;comment:历史UV" json:"total_uv"`
	TotalUip    int            `gorm:"column:total_uip;comment:历史UIP" json:"total_uip"`
	StatsDetail string         `gorm:"column:stats_detail;comment:统计详情" json:"stats_detail"`
	CreateTime  time.Time      `gorm:"column:create_time;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"`
	UpdateTime  time.Time      `gorm:"column:update_time;default:CURRENT_TIMESTAMP;comment:修改时间" json:"update_time"`
	DeleteTime  gorm.DeletedAt `gorm:"column:delete_time;comment:删除时间戳" json:"delete_time"`
}

func (*LinkStats) TableName() string {
	return TableNameLink
}
