package proxy

import (
	"fans-go/app/models"
	"log"
	"sync"
)

// type Payload struct {
// 	UserId   string `json:"user_id"`
// 	UserType string `json:"user_type"`
// }

func BuildUserIndex(user Payload) string {
	return user.UserType + "." + user.UserId
}

var RoleOpposite = map[string]string{
	"source": "target",
	"target": "source",
}

func resolveSourceAndTargetUsers(sourceUser, targetUser *Payload) (*models.User, *models.User) {

	sourceUserChannel := make(chan *models.User)
	targetUserChannel := make(chan *models.User)
	go func() {
		defer close(sourceUserChannel)
		sourceUserChannel <- FindOneOrCreateDB(sourceUser)
	}()
	go func() {
		defer close(targetUserChannel)
		targetUserChannel <- FindOneOrCreateDB(targetUser)

	}()

	sourceUserResult := <-sourceUserChannel
	targetUserResult := <-targetUserChannel

	return sourceUserResult, targetUserResult
}

// const simpleUpdate = async (payload) => {
// 	const { sourceUser, targetUser } = await resolveSourceAndTargetUsers(payload);
// 	const { updatePatch, eventlandAction } = payload;
// 	const finalUpdatePatch = {
// 	  target_user_index: buildUserIndex(targetUser),
// 	  source_user_index: buildUserIndex(sourceUser),
// 	  ...updatePatch,
// 	};

// 	await models.td_relationship.upsert({
// 	  source_user: sourceUser._id,
// 	  target_user: targetUser._id,
// 	  ...finalUpdatePatch,
// 	});
// 	const relationship = await models.td_relationship.findOne({
// 	  where: {
// 		source_user: sourceUser._id,
// 		target_user: targetUser._id,
// 	  },
// 	  raw: true,
// 	});

// 	if (eventlandAction) {
// 	  eventland.fire({
// 		collection: 'relationship',
// 		action: eventlandAction,
// 		data: relationship,
// 	  });
// 	}

// 	return relationship;
//   };

func Del(user Payload) bool {
	userIndex := BuildUserIndex(user)
	log.Println("proxy,user,del,start", userIndex)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		models.Model("user").Where("is_del = ? AND user_id = ? AND user_type = ?", false, user.UserId, user.UserType).Update("is_del", true)
	}()
	go func() {
		defer wg.Done()
		models.Model("relationship").Where("is_del = ? AND target_user_index = ?", false, userIndex).Or("is_del = ? AND source_user_index = ?", false, userIndex).Update("is_del", true)
	}()
	wg.Wait()
	log.Println("proxy,user,del,end", userIndex)
	return true
}
