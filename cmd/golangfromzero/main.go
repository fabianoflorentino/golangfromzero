package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/fabianoflorentino/golangfromzero/config"
	"github.com/fabianoflorentino/golangfromzero/database"
	"github.com/fabianoflorentino/golangfromzero/src/router"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config.LoadEnv()

	ctx := context.Background()
	dsn := database.ConnectionString

	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close()

	r := router.NewRouter()
	appPort := fmt.Sprintf(":%d", config.Port)

	if err := http.ListenAndServe(appPort, r); err != nil {
		log.Fatal(err)
	}
}
