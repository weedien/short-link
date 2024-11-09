package po

import "gorm.io/gorm"

const TableNameLinkGoto = "link_goto"

// LinkGoto mapped from table <link_goto>
//
// 为什么需要 link_goto 表？
// 为了避免在分库分表场景下的扩散查询
// 首先，项目基于 gid 进行分库分表，分表键的候选项为 shortUri、username、gid
// 如果使用 shortUri 作为分表键，就需要查询所有分片才能获取用户某个分组下的短链接
// 如果使用 username 作为分表键，如果用户创建了大量短链接，也会导致查询性能下降
// 如果使用 gid 作为分表键，限制每个分组下的短链接数量，可以较好地控制查询性能
type LinkGoto struct {
	ID         int            `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`
	Gid        string         `gorm:"column:gid;not null;comment:分组标识" json:"gid"`
	ShortUri   string         `gorm:"column:short_uri;not null;comment:短链接" json:"short_uri"`
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time;comment:删除时间" json:"delete_time"`
}

func (*LinkGoto) TableName() string {
	return TableNameLinkGoto
}
