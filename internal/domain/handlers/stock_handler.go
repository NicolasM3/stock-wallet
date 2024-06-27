package handlers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"stock-wallet/internal/.gen/stock-wallet-db/public/model"
	"stock-wallet/internal/domain/repository/database"
	redisRepo "stock-wallet/internal/domain/repository/redisRepo"
	"strconv"
)

type Service struct {
	database *database.Database
	redis    *redisRepo.Redis
}

func NewStockService(db *database.Database, r *redisRepo.Redis) *Service {
	return &Service{db, r}
}

func (s Service) CreateStock(c echo.Context) error {
	ctx := c.Request().Context()

	var stock model.Stock
	if err := c.Bind(&stock); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	stock, err := s.database.CreateStock(ctx, stock)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, stock)
}

func (s Service) GetStocks(c echo.Context) error {
	ctx := c.Request().Context()
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	stocks, err := s.database.GetStock(ctx, page, pageSize)
	if err != nil {
		// Handle error
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, stocks)
}

func (s Service) GetStockById(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, "id is required")
	}

	redisStock, _ := s.redis.GetStockByIdInRedis(id)
	if redisStock != nil {
		return c.JSON(http.StatusOK, redisStock)
	}

	stock, err := s.database.GetStockById(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = s.redis.StoreStockInRedis(stock)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, stock)
}

func (s Service) UpdateStock(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, "id is required")
	}

	storedStock, err := s.database.GetStockById(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	updatedStock := updateFields(c, storedStock)

	stock, err := s.database.UpdateStock(ctx, id, updatedStock)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = s.redis.StoreStockInRedis(stock)
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

func (s Service) DeleteStock(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, "id is required")
	}

	err := s.database.DeleteStock(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Stock deleted successfully")
}
