package main

import (
	"log"
	"net/http"

	"github.com/fabianoflorentino/golangfromzero/src/router"
)

func main() {
	r := router.NewRouter()

	if err := http.ListenAndServe(":6000", r); err != nil {
		log.Fatal(err)
	}
}
