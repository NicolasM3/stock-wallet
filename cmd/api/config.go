package main

import (
	"context"
	"gopkg.in/yaml.v3"
	"os"
	"stock-wallet/internal/domain/repository/database"
	"stock-wallet/internal/domain/repository/redisRepo"
)

type Config struct {
	Database database.Config  `yaml:"database"`
	Cache    redisRepo.Config `yaml:"cache"`
}

func NewConfig(ctx context.Context) Config {
	data, err := os.ReadFile("deployments/config.yml")
	if err != nil {
		println(err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		println(err)
	}

	return config
}
