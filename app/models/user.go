package models

type User struct {
	// 嵌入公共model
	BaseModel

	// 账号
	UserId string `gorm:"type:varchar(255);not null;" json:"user_id"`

	// 账号类型
	UserType string `gorm:"type:varchar(255);not null;type:enum('oa','portal','virtual')" json:"user_type"`

	// oa简要信息
	Oa string `json:"oa"`

	// 被多少人关注
	FollowerCount uint64 `gorm:"type:int(10);unsigned;default:0" json:"follower_count"`

	// 关注多少人
	FollowingCount uint64 `gorm:"type:int(10);unsigned;default:0" json:"following_count"`
}
