package services

import (
	"Microservice_TS_1/models"
	"Microservice_TS_1/repositories"
	"context"
)

type OrderHistoryService struct {
	repository *repositories.OrderHistoryRepository
}

func NewOrderHistoryService(repository *repositories.OrderHistoryRepository) *OrderHistoryService {
	return &OrderHistoryService{repository: repository}
}

func (s *OrderHistoryService) GetOrderHistory(ctx context.Context, client *models.Client) ([]*models.HistoryOrder, error) {
	return s.repository.GetOrderHistory(ctx, client)
}

func (s *OrderHistoryService) SaveOrder(ctx context.Context, client *models.Client, order *models.HistoryOrder) error {
	return s.repository.SaveOrder(ctx, client, order)
}
