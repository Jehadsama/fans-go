package main

import (
	_ "fans-go/app/config"
	"fans-go/app/proxy"
	"fmt"
	"os"
)

func main() {

	fmt.Println(os.Getenv("MYSQL_USER"))

	// user, err := proxy.FindOneOrCreateDB(&proxy.Payload{
	// 	UserId:   "je",
	// 	UserType: "oa",
	// })
	// user := &models.User{}
	// user.GetUsersByIds([]string{"hanyiding"})
	// if err != nil {
	// 	fmt.Println(err)
	// }

	payload := &proxy.Payload{UserId: "lixunhuan", UserType: "portal"}
	user := proxy.FindOneOrCreateDB(payload)

	fmt.Printf("%#v", user)
	fmt.Println("=========")
	fmt.Println()

}
