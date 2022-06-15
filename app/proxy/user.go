package proxy

import (
	"fans-go/app/models"
)

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

type Payload struct {
	UserId   string            `json:"user_id"`
	UserType string            `json:"user_type"`
	Oa       map[string]string `json:"oa"`
}

func FindOneOrCreateDB(payload *Payload) *models.User {
	var user *models.User
	result := models.Model(&user).Where("user_id = ? AND user_type = ?", payload.UserId, payload.UserType).First(&user)
	if result.RowsAffected == 1 {
		return user
	}
	if payload.UserType == "oa" {
		oaUsername := (&models.OaPGfUser{}).UsernameByOAID(payload.UserId)
		payload.Oa = map[string]string{"username": oaUsername}
	}

	models.Model(&user).Where(models.User{UserId: payload.UserId, UserType: payload.UserType}).FirstOrCreate(&user)
	return user
}


const updateFollowCountByID = async ({
	userObjectID,
	followType = 'source',
  }) => {
	const followCount = await models.td_relationship.count({
	  where: {
		[followKeysMapping[followType].relationship]: userObjectID,
		is_follow: true,
	  },
	});

	await models.td_user.update(
	  {
		[followKeysMapping[followType].user]: followCount,
	  },
	  { where: { _id: userObjectID } }
	);

	return models.td_user.findOne({ where: { _id: userObjectID } });
  };

func UpdateFollowCountByID (userObjectID, followType string) *models.User{
	if followType == "" {
		followType= "source"
	}


}