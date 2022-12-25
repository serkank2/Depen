package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func Env() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func GetPort() string {
	return os.Getenv("PORT")
}
