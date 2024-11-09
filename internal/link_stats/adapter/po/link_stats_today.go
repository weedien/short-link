package po

import (
	"gorm.io/gorm"
	"time"
)

const TableNameLinkStatsToday = "link_stats_today"

// LinkStatsToday mapped from table <link_stats_today>
type LinkStatsToday struct {
	ID         int            `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`                 // ID
	ShortUri   string         `gorm:"column:short_uri;not null;comment:短链接" json:"short_uri"`                       // 短链接
	Date       time.Time      `gorm:"column:date;not null;default:CURRENT_DATE;comment:日期" json:"date"`             // 日期
	TodayPv    int            `gorm:"column:today_pv;comment:今日PV" json:"today_pv"`                                 // 今日PV
	TodayUv    int            `gorm:"column:today_uv;comment:今日UV" json:"today_uv"`                                 // 今日UV
	TodayUip   int            `gorm:"column:today_uip;comment:今日IP数" json:"today_uip"`                              // 今日IP数
	CreateTime time.Time      `gorm:"column:create_time;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime time.Time      `gorm:"column:update_time;default:CURRENT_TIMESTAMP;comment:修改时间" json:"update_time"` // 修改时间
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time;comment:删除时间" json:"delete_time"`                           // 删除时间
}

// TableName LinkStatsToday's table name
func (*LinkStatsToday) TableName() string {
	return TableNameLinkStatsToday
}
