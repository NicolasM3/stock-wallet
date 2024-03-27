package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateStock(c echo.Context) error {
	return nil
}

func GetStocks(c echo.Context) error {
	return c.JSON(http.StatusOK, []string{"Stock 1", "Stock 2"})
}

func GetStockById(c echo.Context) error {
	return nil
}

func UpdateStock(c echo.Context) error {
	return nil
}

func DeleteStock(c echo.Context) error {
	return nil
}
