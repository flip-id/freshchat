package freshchat

import (
	"github.com/joho/godotenv"
)


func init() {
	loadEnv()
}

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		_ = godotenv.Load("./../../.env")
	}
}
