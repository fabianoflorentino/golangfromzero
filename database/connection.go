package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var ConnectionString string = ""

func Connect() (*pgxpool.Pool, error) {
	var (
		user     string = os.Getenv("POSTGRES_USER")
		password string = os.Getenv("POSTGRES_PASSWORD")
		host     string = os.Getenv("POSTGRES_HOST")
		database string = os.Getenv("POSTGRES_DBNAME")
		sslmode  string = os.Getenv("POSTGRES_SSLMODE")
		port     string = os.Getenv("POSTGRES_PORT")
	)

	ctx := context.Background()
	ConnectionString = connectionString(user, password, host, port, database, sslmode)

	pool, err := pgxpool.New(ctx, ConnectionString)
	if err != nil {
		return nil, err
	}

	return pool, nil

}

func connectionString(user, password, host, port, dbname, sslmode string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user, password, host, port, dbname, sslmode)
}
