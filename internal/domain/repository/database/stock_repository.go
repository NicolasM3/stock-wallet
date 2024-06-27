package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"stock-wallet/internal/.gen/stock-wallet-db/public/model"
)

type StockRepositoryInterface interface {
	GetStock(ctx context.Context, page int, pageSize int) ([]*model.Stock, error)
	GetStockById(ctx context.Context, id string) (model.Stock, error)
	CreateStock(ctx context.Context, stock model.Stock) (model.Stock, error)
	DeleteStock(ctx context.Context, id string) error
	UpdateStock(ctx context.Context, id string, stock model.Stock) (model.Stock, error)
}

type Database struct {
	Conn *pgxpool.Pool
}

func New(config Config) (*Database, error) {
	connPoolStr := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", config.User, config.Password, config.Database)

	connPool, err := pgxpool.New(context.Background(), connPoolStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connPoolect to Conn: %v\n", err)
		os.Exit(1)
	}

	return &Database{
		Conn: connPool,
	}, nil
}

func (r *Database) GetStock(ctx context.Context, page int, pageSize int) (model.StockListReponse, error) {
	offset := (page - 1) * pageSize
	rows, err := r.Conn.Query(ctx, "select * from stock LIMIT $1 OFFSET $2", pageSize, offset)
	if err != nil {
		return model.StockListReponse{}, err
	}

	var stocks []*model.Stock
	for rows.Next() {
		var stock model.Stock
		if err := rows.Scan(&stock.ID, &stock.Code, &stock.Name, &stock.CurrentPrice, &stock.CreatedAt, &stock.LastUpdate, &stock.Deleted); err != nil {
			return model.StockListReponse{}, err
		}
		stocks = append(stocks, &stock)
	}

	if err := rows.Err(); err != nil {
		return model.StockListReponse{}, err
	}

	var total int
	err = r.Conn.QueryRow(ctx, "SELECT COUNT(*) FROM stock").Scan(&total)
	if err != nil {
		return model.StockListReponse{}, err
	}

	return model.StockListReponse{
		Data: stocks,
		Pagination: model.Pagination{
			Page:     page,
			PageSize: pageSize,
			Total:    total,
		},
	}, nil
}

func (r *Database) GetStockById(ctx context.Context, id string) (model.Stock, error) {
	var stock model.Stock
	err := r.Conn.QueryRow(ctx, "select * from stock where id = $1", id).Scan(&stock.ID, &stock.Code, &stock.Name, &stock.CurrentPrice, &stock.CreatedAt, &stock.LastUpdate, &stock.Deleted)
	if err != nil {
		return stock, err
	}

	return stock, nil

}

func (r *Database) CreateStock(ctx context.Context, stock model.Stock) (model.Stock, error) {
	err := r.Conn.QueryRow(ctx, "insert into stock(code, name, current_price, deleted) values($1, $2, $3, $4) RETURNING id", stock.Code, stock.Name, stock.CurrentPrice, false).Scan(&stock.ID)
	if err != nil {
		return stock, err
	}

	return stock, nil
}

func (r *Database) DeleteStock(ctx context.Context, id string) error {
	_, err := r.Conn.Exec(ctx, "delete from stock where id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (r *Database) UpdateStock(ctx context.Context, id string, stock model.Stock) (model.Stock, error) {
	_, err := r.Conn.Exec(ctx, "update stock set code = $1, name = $2, current_price = $3, last_update = now() where id = $4", stock.Code, stock.Name, stock.CurrentPrice, id)
	if err != nil {
		return stock, err
	}

	return stock, nil
}
