package main

import (
	"log"

	"github.com/fabianoflorentino/golangfromzero/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	if err := app.Run(); err != nil {
		log.Fatalf("application failed to start: %v", err)
	}
}
