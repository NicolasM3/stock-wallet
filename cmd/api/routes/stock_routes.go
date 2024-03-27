package routes

import (
	"github.com/labstack/echo/v4"
	"stock-wallet/cmd/api/handlers"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("v1/stock", handlers.GetStocks)
	e.GET("v1/stock/:id", handlers.GetStockById)
	e.POST("v1/stock", handlers.CreateStock)
	e.PATCH("v1/stock/:id", handlers.UpdateStock)
	e.DELETE("v1/stock/:id", handlers.DeleteStock)
}
