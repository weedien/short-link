package po

import (
	"gorm.io/gorm"
	"time"
)

const TableNameLinkBrowserStat = "link_browser_stats"

// LinkBrowserStat mapped from table <link_browser_stats>
type LinkBrowserStat struct {
	ID             int            `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`                 // ID
	ShortUri       string         `gorm:"column:short_uri;not null;comment:短链接" json:"short_uri"`                       // 完整短链接
	Date           time.Time      `gorm:"column:date;default:CURRENT_DATE;comment:日期" json:"date"`                      // 日期
	Cnt            int            `gorm:"column:cnt;comment:访问量" json:"cnt"`                                            // 访问量
	Browser        string         `gorm:"column:browser;comment:浏览器，Edge,Chrome,FireFox" json:"browser"`                // 浏览器，Edge,Chrome,FireFox
	CreateTime     time.Time      `gorm:"column:create_time;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime     time.Time      `gorm:"column:update_time;default:CURRENT_TIMESTAMP;comment:修改时间" json:"update_time"` // 修改时间
	BrowserVersion string         `gorm:"column:browser_version;not null;comment:浏览器版本" json:"browser_version"`         // 浏览器版本
	DeleteTime     gorm.DeletedAt `gorm:"column:delete_time;comment:删除时间" json:"delete_time"`                           // 删除时间
}

// TableName LinkBrowserStat's table name
func (*LinkBrowserStat) TableName() string {
	return TableNameLinkBrowserStat
}
