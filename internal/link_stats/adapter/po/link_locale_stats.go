package po

import (
	"gorm.io/gorm"
	"time"
)

const TableNameLinkLocaleStat = "link_locale_stats"

// LinkLocaleStat mapped from table <link_locale_stats>
type LinkLocaleStat struct {
	ID         int            `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`                 // ID
	ShortUri   string         `gorm:"column:short_uri;not null;comment:短链接" json:"short_uri"`                       // 完整短链接
	Date       time.Time      `gorm:"column:date;not null;default:CURRENT_DATE;comment:日期" json:"date"`             // 日期
	Cnt        int            `gorm:"column:cnt;comment:访问量" json:"cnt"`                                            // 访问量
	Province   string         `gorm:"column:province;not null;comment:省份" json:"province"`                          // 省份
	City       string         `gorm:"column:city;not null;comment:市区" json:"city"`                                  // 市区
	Country    string         `gorm:"column:country;not null;comment:国家" json:"country"`                            // 国家
	CreateTime time.Time      `gorm:"column:create_time;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime time.Time      `gorm:"column:update_time;default:CURRENT_TIMESTAMP;comment:修改时间" json:"update_time"` // 修改时间
	Coords     string         `gorm:"column:coords;comment:坐标（经纬度）" json:"coords"`                                  // 坐标（经纬度）
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time;comment:删除时间" json:"delete_time"`                           // 删除时间
}

// TableName LinkLocaleStat's table name
func (*LinkLocaleStat) TableName() string {
	return TableNameLinkLocaleStat
}
