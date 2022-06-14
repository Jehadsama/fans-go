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

// const findOneOrCreateDB = async (payload) => {
// 	const baseQuery = _.pick(payload, ['user_id', 'user_type']);
// 	const user = await models.td_user.findOne({ where: baseQuery, raw: true });
// 	if (user) {
// 	  return user;
// 	}

// 	const { user_id: userID, user_type: userType } = payload;
// 	// validate whether oaID(userID) was correct
// 	if (userType === 'oa') {
// 	  const oaUsername = await models.oa_p_gf_user.usernameByOAID(userID);
// 	  baseQuery.oa = {
// 		username: oaUsername,
// 	  };
// 	}

// 	// upsert: MySQL - Implemented with ON DUPLICATE KEY UPDATE
// 	await models.td_user.upsert(baseQuery);
// 	return models.td_user.findOne({ where: baseQuery, raw: true });
//   };

type Payload struct {
	UserId   string `json:"user_id"`
	UserType string `json:"user_type"`
}

func FindOneOrCreateDB(payload *Payload) (*models.User, error) {
	var user *models.User
	result := models.Model(&user).Where("user_id = ? AND user_type = ?", payload.UserId, payload.UserType).First(&user)
	if result.RowsAffected == 1 {
		return user, nil
	}
	return nil, result.Error
}
