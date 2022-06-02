package models

import "time"

// gorm.Model 的定义
type BaseModel struct {
	// 其实是mongo objectid
	ID string `gorm:"primaryKey;column:_id;size:255"`
	// 创建时间
	CreatedAt time.Time `gorm:"column:created_at"`
	// 修改时间
	UpdatedAt time.Time `gorm:"column:updated_at"`
	// 是否删除
	IsDel bool `gorm:"column:is_del;default:false"`
}
