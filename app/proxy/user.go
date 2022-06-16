package proxy

import (
	"fans-go/app/libs/utils"
	"fans-go/app/models"
)

type Payload struct {
	UserId   string            `json:"user_id"`
	UserType string            `json:"user_type"`
	Oa       map[string]string `json:"oa"`
}

func FindOneOrCreateDB(payload *Payload) *models.User {
	baseQuery := map[string]interface{}{"user_id": payload.UserId, "user_type": payload.UserType}
	var user *models.User
	result := models.Model("user").Where(baseQuery).First(&user)
	if result.RowsAffected == 1 {
		return user
	}
	if payload.UserType == "oa" {
		oaUsername := (&models.OaPGfUser{}).UsernameByOAID(payload.UserId)
		payload.Oa = map[string]string{"username": oaUsername}
	}

	models.Model("user").Where(baseQuery).FirstOrCreate(&user)
	return user
}

var followKeysMapping = map[string](map[string]string){
	"source": {
		"relationship": "source_user",
		"user":         "following_count",
	},
	"target": {
		"relationship": "target_user",
		"user":         "follower_count",
	},
}

func UpdateFollowCountByID(userObjectID, followType string) *models.User {
	temp := followKeysMapping[followType]
	var followCount int64
	result := models.Model("relationship").Where(map[string]interface{}{"is_follow": true, temp["relationship"]: userObjectID}).Count(&followCount)
	if result.Error != nil {
		utils.CheckError("proxy,user,UpdateFollowCountByID", result.Error)
	}
	models.Model("user").Where("_id = ?", userObjectID).Update(temp["user"], followCount)
	var user *models.User
	models.Model("user").Where("_id = ?", userObjectID).First(&user)
	return user
}
