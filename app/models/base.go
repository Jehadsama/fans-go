package models

import (
	"fans-go/app/connectors/db"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

type BaseModel struct {
	// 其实是mongo objectid, `gorm:"type:varchar(255);column:_id;primaryKey;not null;" json:"_id"`
	ID string `gorm:"type:varchar(255);column:_id;primaryKey;not null;" json:"_id"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 修改时间
	UpdatedAt time.Time `json:"updated_at"`
	// 是否删除
	IsDel bool `gorm:"default:false" json:"is_del"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	id := primitive.NewObjectID().Hex()
	b.ID = id
	return
}

func Model(value interface{}) *gorm.DB {
	return db.DB.Model(value)
}
