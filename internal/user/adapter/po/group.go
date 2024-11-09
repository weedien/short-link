package po

import (
	"gorm.io/gorm"
	"time"
)

// Group mapped from table <group>
type Group struct {
	ID         int            `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`                 // ID
	Gid        string         `gorm:"column:gid;not null;comment:分组标识" json:"gid"`                                  // 分组标识
	Name       string         `gorm:"column:name;not null;comment:分组名称" json:"name"`                                // 分组名称
	Username   string         `gorm:"column:username;not null;comment:创建分组用户名" json:"username"`                     // 创建分组用户名
	SortOrder  int            `gorm:"column:sort_order;comment:分组排序" json:"sort_order"`                             // 分组排序
	CreateTime time.Time      `gorm:"column:create_time;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime time.Time      `gorm:"column:update_time;default:CURRENT_TIMESTAMP;comment:修改时间" json:"update_time"` // 修改时间
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time;comment:删除时间" json:"delete_time"`                           // 删除时间
}
