package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"os"
	"stock-wallet/cmd/api/routes"
)

func main() {
	e := echo.New()
	routes.SetupRoutes(e)
	connPool, err := openDBConnection()
	if err != nil {
		e.Logger.Fatal(err)
		os.Exit(1)
	}
	defer connPool.Close()

	e.Logger.Fatal(e.Start(":8080"))
}

func openDBConnection() (*pgxpool.Pool, error) {
	connPoolStr := "postgres://stock-wallet-user:stock-wallet-password@localhost/postgres?sslmode=disable"

	connPool, err := pgxpool.New(context.Background(), connPoolStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connPoolect to database: %v\n", err)
		os.Exit(1)
	}

	return connPool, nil
}
