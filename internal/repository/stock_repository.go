package repository

import (
	"context"
	"stock-wallet/internal/domain"
	"stock-wallet/internal/domain/.gen/stock-wallet-db/public/model"
)

func GetStock(ctx context.Context, page int, pageSize int) ([]*model.Stock, error) {
	dbpool := domain.DB

	offset := (page - 1) * pageSize
	rows, err := dbpool.Query(ctx, "select * from stock LIMIT $1 OFFSET $2", pageSize, offset)
	if err != nil {
		return nil, err
	}

	var stocks []*model.Stock
	for rows.Next() {
		var stock model.Stock
		if err := rows.Scan(&stock.ID, &stock.Code, &stock.Name, &stock.CurrentPrice, &stock.CreatedAt, &stock.LastUpdate, &stock.Deleted); err != nil {
			return nil, err
		}
		stocks = append(stocks, &stock)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return stocks, nil
}

func GetStockById(ctx context.Context, id string) (model.Stock, error) {
	dbpool := domain.DB

	var stock model.Stock
	err := dbpool.QueryRow(ctx, "select * from stock where id = $1", id).Scan(&stock.ID, &stock.Code, &stock.Name, &stock.CurrentPrice, &stock.CreatedAt, &stock.LastUpdate, &stock.Deleted)
	if err != nil {
		return stock, err
	}

	return stock, nil

}

func CreateStock(ctx context.Context, stock model.Stock) (model.Stock, error) {
	dbpool := domain.DB

	err := dbpool.QueryRow(ctx, "insert into stock(code, name, current_price, deleted) values($1, $2, $3, $4) RETURNING id", stock.Code, stock.Name, stock.CurrentPrice, false).Scan(&stock.ID)
	if err != nil {
		return stock, err
	}

	return stock, nil
}

func DeleteStock(ctx context.Context, id string) error {
	dbpool := domain.DB

	_, err := dbpool.Exec(ctx, "delete from stock where id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateStock(ctx context.Context, id string, stock model.Stock) (model.Stock, error) {
	dbpool := domain.DB

	_, err := dbpool.Exec(ctx, "update stock set code = $1, name = $2, current_price = $3, last_update = now() where id = $4", stock.Code, stock.Name, stock.CurrentPrice, id)
	if err != nil {
		return stock, err
	}

	return stock, nil
}
