package repository

import (
	"context"
	"encoding/json"
	"stock-wallet/internal/domain"
	"stock-wallet/internal/domain/.gen/stock-wallet-db/public/model"
	"strconv"
	"time"
)

func StoreStockInRedis(stock model.Stock) error {
	stockJson, err := json.Marshal(stock)
	if err != nil {
		return err
	}

	err = domain.RedisClient.Set(context.Background(), strconv.Itoa(int(stock.ID)), stockJson, 24*time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}

func GetStockByIdInRedis(id string) (*model.Stock, error) {
	stockJson, err := domain.RedisClient.Get(context.Background(), id).Result()
	if err != nil {
		return nil, err
	}

	var stock model.Stock
	err = json.Unmarshal([]byte(stockJson), &stock)
	if err != nil {
		return nil, err
	}

	return &stock, nil
}
