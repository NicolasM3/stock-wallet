package routes

import (
	"github.com/labstack/echo/v4"
	"stock-wallet/internal/domain/handlers"
)

func SetupRoutes(e *echo.Echo, s handlers.Service) {
	e.GET("v1/stock", s.GetStocks)
	e.GET("v1/stock/:id", s.GetStockById)
	e.POST("v1/stock", s.CreateStock)
	e.PATCH("v1/stock/:id", s.UpdateStock)
	e.DELETE("v1/stock/:id", s.DeleteStock)
}
