package database

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID         uint           `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`
	CreateTime time.Time      `gorm:"column:create_time;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"`
	UpdateTime time.Time      `gorm:"column:update_time;default:CURRENT_TIMESTAMP;comment:修改时间" json:"update_time"`
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time;comment:删除时间戳" json:"delete_time"`
	TenantID   string         `gorm:"column:tenant_id;comment:租户ID" json:"tenant_id"`
}

// BeforeCreate 在创建记录之前设置 TenantID
func (base *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	// 从上下文中获取 TenantID
	tenantID, ok := tx.Statement.Context.Value("tenant_id").(string)
	if !ok {
		tenantID = "default"
	}
	base.TenantID = tenantID
	return
}
