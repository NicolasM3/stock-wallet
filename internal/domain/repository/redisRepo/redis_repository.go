package redisRepo

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"stock-wallet/internal/.gen/stock-wallet-db/public/model"
	"strconv"
	"time"
)

type Redis struct {
	Conn *redis.Client
}

func New(config Config) (*Redis, error) {
	opt, err := redis.ParseURL(fmt.Sprintf("redis://%s:%s/0", config.Host, config.Port))
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)

	return &Redis{
		Conn: client,
	}, nil
}

func (r *Redis) StoreStockInRedis(stock model.Stock) error {
	stockJson, err := json.Marshal(stock)
	if err != nil {
		return err
	}

	err = r.Conn.Set(context.Background(), strconv.Itoa(int(stock.ID)), stockJson, 24*time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *Redis) GetStockByIdInRedis(id string) (*model.Stock, error) {
	stockJson, err := r.Conn.Get(context.Background(), id).Result()
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
