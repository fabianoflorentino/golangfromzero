package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fabianoflorentino/golangfromzero/config"
	"github.com/fabianoflorentino/golangfromzero/src/router"
)

func init() {
	config.LoadEnv()
}

func main() {
	r := router.NewRouter()
	app_port := fmt.Sprintf(":%d", config.Port)

	if err := http.ListenAndServe(app_port, r); err != nil {
		log.Fatal(err)
	}
}
