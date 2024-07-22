package repositories

import (
	"Microservice_TS_1/models"
	"context"
	"github.com/jackc/pgx/v4"
)

type OrderHistoryRepository struct {
	db pgx.Tx // pgx.Tx - интерфейс для транзакций, можно заменить на другой подходящий интерфейс
}

func NewOrderHistoryRepository(db pgx.Tx) *OrderHistoryRepository {
	return &OrderHistoryRepository{db: db}
}

func (r *OrderHistoryRepository) GetOrderHistory(ctx context.Context, client *models.Client) ([]*models.HistoryOrder, error) {
	// Реализация запроса к базе данных для получения истории ордеров
	return nil, nil
}

func (r *OrderHistoryRepository) SaveOrder(ctx context.Context, client *models.Client, order *models.HistoryOrder) error {
	// Реализация запроса к базе данных для сохранения ордера
	return nil
}
