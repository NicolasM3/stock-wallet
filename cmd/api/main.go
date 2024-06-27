package main

import (
	"context"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"os"
	"stock-wallet/internal/domain/handlers"
	"stock-wallet/internal/domain/repository/database"
	"stock-wallet/internal/domain/repository/redisRepo"
	"stock-wallet/internal/domain/routes"
)

func main() {
	e := echo.New()

	cfg := NewConfig(context.Background())

	dbConn, err := database.New(cfg.Database)
	if err != nil {
		e.Logger.Fatal(err)
		os.Exit(1)
	}

	redisClient, err := redisRepo.New(cfg.Cache)
	if err != nil {
		e.Logger.Fatal(err)
		os.Exit(1)
	}

	svc := handlers.NewStockService(dbConn, redisClient)
	routes.SetupRoutes(e, *svc)

	e.Logger.Fatal(e.Start(":8080"))
}
