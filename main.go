package main

import (
	_ "fans-go/app/config"
	"fans-go/app/models"
	"fmt"
	"os"
)

func main() {

	fmt.Println(os.Getenv("MYSQL_USER"))

	user := &models.User{
		// BaseModel: models.BaseModel{},
		// UserType: "j",
	}
	user.GetUsersByIds([]string{"hanyiding"})

	fmt.Println(user)

}
