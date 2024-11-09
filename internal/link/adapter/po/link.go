package po

import (
	"database/sql"
	"shortlink/internal/base/database"
	"shortlink/internal/link/domain/link"
)

const TableNameLink = "t_link"

// Link mapped from table <link>
type Link struct {
	database.BaseModel
	Gid         string          `gorm:"column:gid;not null;comment:分组标识" json:"gid"`
	ShortUri    string          `gorm:"column:short_uri;not null;comment:短链接" json:"short_uri"`
	OriginalUrl string          `gorm:"column:original_url;not null;comment:原始链接" json:"origin_url"`
	Favicon     string          `gorm:"column:favicon;comment:网站图标" json:"favicon"`
	Status      link.Status     `gorm:"column:status;not null;default:active;comment:可选值:,active,expired,disabled,pending,deleted,reserved" json:"status"`
	CreateType  link.CreateType `gorm:"column:created_type;not null;comment:创建类型 0：接口创建 1：控制台创建" json:"created_type"`
	ValidType   link.ValidType  `gorm:"column:valid_type;not null;comment:有效期类型 0：永久有效 1：自定义" json:"valid_date_type"`
	StartDate   sql.NullTime    `gorm:"column:start_date;not null;default:CURRENT_TIMESTAMP;comment:有效期开始时间" json:"valid_start_date"`
	EndDate     sql.NullTime    `gorm:"column:end_date;comment:有效期结束时间" json:"valid_end_date"`
	Desc        string          `gorm:"column:desc;comment:描述" json:"desc"`
	RecycleTime sql.NullTime    `gorm:"column:recycle_time;comment:回收时间" json:"recycle_time"`
}

func (*Link) TableName() string {
	return TableNameLink
}
