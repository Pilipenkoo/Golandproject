package repositories

import (
	"Microservice_TS_1/models"
	"context"
	"github.com/jackc/pgx/v4"
)

type OrderBookRepository struct {
	db pgx.Tx // pgx.Tx - интерфейс для транзакций, можно заменить на другой подходящий интерфейс
}

func NewOrderBookRepository(db pgx.Tx) *OrderBookRepository {
	return &OrderBookRepository{db: db}
}

func (r *OrderBookRepository) GetOrderBook(ctx context.Context, exchangeName, pair string) ([]*models.DepthOrder, error) {
	// Реализация запроса к базе данных для получения данных order book
	return nil, nil
}

func (r *OrderBookRepository) SaveOrderBook(ctx context.Context, exchangeName, pair string, orderBook []*models.DepthOrder) error {
	// Реализация запроса к базе данных для сохранения данных order book
	return nil
}
