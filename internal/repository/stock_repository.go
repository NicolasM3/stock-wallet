package repository

import (
	"context"
	"stock-wallet/internal/domain"
)

func GetStock(ctx context.Context) (*string, error) {
	dbpool := domain.DB

	var rows string
	err := dbpool.QueryRow(ctx, "SELECT * FROM stock").Scan(&rows)
	if err != nil {
		return nil, err
	}

	return &rows, nil
}
