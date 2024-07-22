package services

import (
	"Microservice_TS_1/models"
	"Microservice_TS_1/repositories"
	"context"
)

type OrderBookService struct {
	repository *repositories.OrderBookRepository
}

func NewOrderBookService(repository *repositories.OrderBookRepository) *OrderBookService {
	return &OrderBookService{repository: repository}
}

func (s *OrderBookService) GetOrderBook(ctx context.Context, exchangeName, pair string) ([]*models.DepthOrder, error) {
	return s.repository.GetOrderBook(ctx, exchangeName, pair)
}

func (s *OrderBookService) SaveOrderBook(ctx context.Context, exchangeName, pair string, orderBook []*models.DepthOrder) error {
	return s.repository.SaveOrderBook(ctx, exchangeName, pair, orderBook)
}
