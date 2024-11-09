package po

import (
	"gorm.io/gorm"
	"time"
)

const TableNameLinkAccessStat = "link_access_stats"

// LinkAccessStat mapped from table <link_access_stats>
type LinkAccessStat struct {
	ID         int            `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`                 // ID
	ShortUri   string         `gorm:"column:short_uri;not null;comment:短链接" json:"short_uri"`                       // 完整短链接
	Date       time.Time      `gorm:"column:date;default:CURRENT_DATE;comment:日期" json:"date"`                      // 日期
	Pv         int            `gorm:"column:pv;comment:访问量" json:"pv"`                                              // 访问量
	Uv         int            `gorm:"column:uv;comment:独立访客数" json:"uv"`                                            // 独立访客数
	Uip        int            `gorm:"column:uip;comment:独立IP数" json:"uip"`                                          // 独立IP数
	Hour       int            `gorm:"column:hour;comment:小时" json:"hour"`                                           // 小时
	Week       int            `gorm:"column:week;comment:星期" json:"week"`                                           // 星期
	CreateTime time.Time      `gorm:"column:create_time;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime time.Time      `gorm:"column:update_time;default:CURRENT_TIMESTAMP;comment:修改时间" json:"update_time"` // 修改时间
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time;comment:删除时间" json:"delete_time"`                           // 删除时间
}

// TableName LinkAccessStat's table name
func (*LinkAccessStat) TableName() string {
	return TableNameLinkAccessStat
}
