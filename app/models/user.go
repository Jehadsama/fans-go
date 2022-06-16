package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type JSON json.RawMessage

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (j *JSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}

// 实现 driver.Valuer 接口，Value 返回 json value
func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}

type User struct {
	// 嵌入公共model
	BaseModel

	// 账号
	UserId string `gorm:"type:varchar(255);not null;" json:"user_id"`

	// 账号类型
	UserType string `gorm:"type:varchar(255);not null;type:enum('oa','portal','virtual')" json:"user_type"`

	// oa简要信息
	Oa map[string]string `gorm:"serialize:json" json:"oa"`

	// 被多少人关注
	FollowerCount uint64 `gorm:"type:int(10);unsigned;default:0" json:"follower_count"`

	// 关注多少人
	FollowingCount uint64 `gorm:"type:int(10);unsigned;default:0" json:"following_count"`
}

func (user *User) GetUsersByIds(ids []string) *gorm.DB {
	return Model("user").Where("user_id IN ? AND user_type IN ?", ids, []string{"oa", "virtual"}).Find(&user)
}
