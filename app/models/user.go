package models

type User struct {
	// 嵌入公共model
	BaseModel

	// 账号
	UserId string `gorm:"column:user_id;not null;size:255"`

	// 账号类型
	UserType string `gorm:"column:user_type;not null;size:255"`

	// oa简要信息
	Oa map[string]string `gorm:"column:oa;serializer:json"`

	// 被多少人关注
	FollowerCount uint64 `gorm:"column:follower_count;default:0"`

	// 关注多少人
	FollowingCount uint64 `gorm:"column:following_count;default:0"`
}
