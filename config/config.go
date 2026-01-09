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

	if err := godotenv.Load(PathEnv + "/.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func AppPort() int {
	serverPort := os.Getenv("SERVER_PORT")

	Port, err := strconv.Atoi(serverPort)
	if err != nil {
		return 6000
	}

	return Port
}
