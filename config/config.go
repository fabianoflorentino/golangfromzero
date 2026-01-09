package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	var PathEnv string = "/golangfromzero/config"

	if err := godotenv.Load(PathEnv + "/.env"); err != nil {
		return ErrLoadEnv
	}

	return nil
}

func AppPort() int {
	serverPort := os.Getenv("SERVER_PORT")

	Port, err := strconv.Atoi(serverPort)
	if err != nil {
		return 6000
	}

	return Port
}
