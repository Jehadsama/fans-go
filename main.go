package main

import (
	_ "fans-go/app/config"
	"fans-go/app/proxy"
	"fmt"
	"os"
)

func main() {

	fmt.Println(os.Getenv("MYSQL_USER"))

	user, err := proxy.FindOneOrCreateDB(&proxy.Payload{
		UserId:   "je",
		UserType: "oa",
	})
	// user.GetUsersByIds([]string{"hanyiding"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)

}
