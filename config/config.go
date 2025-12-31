package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Port int
var PathEnv string = "/golangfromzero/config"

func LoadEnv() {
	var err error

	if err := godotenv.Load(PathEnv + "/.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	Port, err = strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		Port = 6000
	}
}
