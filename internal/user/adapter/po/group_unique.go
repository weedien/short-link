package po

import "gorm.io/gorm"

// GroupUnique mapped from table <group_unique>
type GroupUnique struct {
	ID         int            `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	Gid        string         `gorm:"column:gid;not null;comment:分组标识" json:"gid"`                  // 分组标识
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time;comment:删除时间" json:"delete_time"`           // 删除时间
}
