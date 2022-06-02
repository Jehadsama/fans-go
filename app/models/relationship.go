package models

import "time"

type Relationship struct {
	// 嵌入公共model
	BaseModel

	// 关注的目标,其实是mongo objectid
	TargetUser string `gorm:"column:target_user;not null;"`

	// 发起关注的人,其实是mongo objectid
	SourceUser string `gorm:"column:source_user;not null;"`

	// user_type.user_id
	TargetUserIndex string `gorm:"column:target_user_index;not null;"`

	// user_type.user_id
	SourceUserIndex string `gorm:"column:source_user_index;not null;"`

	// 是否关注
	IsFollow bool `gorm:"column:is_follow;default:true"`

	// 关注时间
	FollowedAt time.Time `gorm:"column:followed_at"`
}
