package models

import "time"

type Relationship struct {
	// 嵌入公共model
	BaseModel

	// 关注的目标,其实是mongo objectid
	TargetUser string `gorm:"type:varchar(255);not null;" json:"target_user"`

	// 发起关注的人,其实是mongo objectid
	SourceUser string `gorm:"type:varchar(255);not null;" json:"source_user"`

	// user_type.user_id
	TargetUserIndex string `gorm:"type:varchar(255);not null;" json:"target_user_index"`

	// user_type.user_id
	SourceUserIndex string `gorm:"type:varchar(255);not null;" json:"source_user_index"`

	// 是否关注
	IsFollow bool `gorm:"default:true" json:"is_follow"`

	// 关注时间
	FollowedAt time.Time `json:"followed_at"`
}
