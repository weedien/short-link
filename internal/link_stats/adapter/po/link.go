package po

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

const TableNameLink = "link"

// Link mapped from table <link>
type Link struct {
	ID          int            `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`
	Gid         string         `gorm:"column:gid;not null;comment:分组标识" json:"gid"`
	ShortUri    string         `gorm:"column:short_uri;not null;comment:短链接" json:"short_uri"`
	OriginalUrl string         `gorm:"column:origin_url;not null;comment:原始链接" json:"origin_url"`
	Favicon     string         `gorm:"column:favicon;comment:网站图标" json:"favicon"`
	Status      string         `gorm:"column:status;not null;default:active;comment:可选值:,active,expired,disabled,pending,deleted,reserved" json:"status"`
	CreateType  int            `gorm:"column:created_type;not null;comment:创建类型 0：接口创建 1：控制台创建" json:"created_type"`
	ValidType   int            `gorm:"column:valid_date_type;not null;comment:有效期类型 0：永久有效 1：自定义" json:"valid_date_type"`
	StartDate   time.Time      `gorm:"column:valid_start_date;not null;comment:有效期开始时间" json:"valid_start_date"`
	EndDate     time.Time      `gorm:"column:valid_end_date;not null;comment:有效期结束时间" json:"valid_end_date"`
	Desc        string         `gorm:"column:desc;comment:描述" json:"desc"`
	ClickNum    int            `gorm:"column:click_num;comment:点击量" json:"click_num"`
	TotalPv     int            `gorm:"column:total_pv;comment:历史PV" json:"total_pv"`
	TotalUv     int            `gorm:"column:total_uv;comment:历史UV" json:"total_uv"`
	TotalUip    int            `gorm:"column:total_uip;comment:历史UIP" json:"total_uip"`
	CreateTime  time.Time      `gorm:"column:create_time;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"`
	UpdateTime  time.Time      `gorm:"column:update_time;default:CURRENT_TIMESTAMP;comment:修改时间" json:"update_time"`
	RecycleTime sql.NullTime   `gorm:"column:recycle_time;comment:回收时间" json:"recycle_time"`
	DeleteTime  gorm.DeletedAt `gorm:"column:delete_time;comment:删除时间戳" json:"delete_time"`
}

func (*Link) TableName() string {
	return TableNameLink
}
