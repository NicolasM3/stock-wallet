package handlers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"stock-wallet/internal/domain/.gen/stock-wallet-db/public/model"
	"stock-wallet/internal/repository"
	"strconv"
)

func CreateStock(c echo.Context) error {
	ctx := c.Request().Context()

	var stock model.Stock
	if err := c.Bind(&stock); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	stock, err := repository.CreateStock(ctx, stock)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, stock)
}

func GetStocks(c echo.Context) error {
	ctx := c.Request().Context()
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	stocks, err := repository.GetStock(ctx, page, pageSize)
	if err != nil {
		// Handle error
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, stocks)
}

func GetStockById(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, "id is required")
	}

	redisStock, _ := repository.GetStockByIdInRedis(id)
	if redisStock != nil {
		return c.JSON(http.StatusOK, redisStock)
	}

	stock, err := repository.GetStockById(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = repository.StoreStockInRedis(stock)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, stock)
}

func UpdateStock(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, "id is required")
	}

	storedStock, err := repository.GetStockById(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	updatedStock := updateFields(c, storedStock)

	stock, err := repository.UpdateStock(ctx, id, updatedStock)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = repository.StoreStockInRedis(stock)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, stock)
}

func updateFields(c echo.Context, stock model.Stock) model.Stock {
	var body map[string]interface{}
	err := json.NewDecoder(c.Request().Body).Decode(&body)
	if err != nil {
		return stock
	}

	if body["code"] != nil {
		stock.Code = body["code"].(string)
	}
	if body["name"] != nil {
		stock.Name = body["name"].(string)
	}
	if body["currentPrice"] != nil {
		stock.CurrentPrice = float32(body["currentPrice"].(float64))
	}
	if body["deleted"] != nil {
		stock.Deleted = body["deleted"].(bool)
	}

	return stock
}

func DeleteStock(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, "id is required")
	}

	err := repository.DeleteStock(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Stock deleted successfully")
}
