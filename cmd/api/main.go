package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"os"
	"stock-wallet/cmd/api/routes"
	"stock-wallet/internal/domain"
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

	domain.InitDB(connPool)

	client, err := openRedisConnection()
	if err != nil {
		e.Logger.Fatal(err)
		os.Exit(1)
	}

	domain.InitRedis(client)

	e.Logger.Fatal(e.Start(":8080"))
}

func openDBConnection() (*pgxpool.Pool, error) {
	connPoolStr := "postgres://stock-wallet-user:stock-wallet-password@localhost/stock-wallet-db?sslmode=disable"

	connPool, err := pgxpool.New(context.Background(), connPoolStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connPoolect to database: %v\n", err)
		os.Exit(1)
	}

	return connPool, nil
}

func openRedisConnection() (*redis.Client, error) {
	opt, err := redis.ParseURL("redis://localhost:6379/0")
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)

	return client, nil
}
