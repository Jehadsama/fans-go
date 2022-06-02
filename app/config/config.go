package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	dirname, _ := os.Getwd()
	env := os.Getenv("FANSGO_ENV")
	if env == "" {
		env = "local"
	}
	err := godotenv.Load(dirname + "/app/config/" + env)
	if err != nil {
		log.Fatal(err)
	}

}
