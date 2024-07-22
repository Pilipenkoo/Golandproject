package tests

import (
	"Microservice_TS_1/mocks"
	"Microservice_TS_1/models"
	"Microservice_TS_1/repositories"
	"Microservice_TS_1/services"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestGetOrderHistory(t *testing.T) {
	mockDB := new(mocks.MockDB)
	mockRepo := repositories.NewOrderHistoryRepository(mockDB)
	service := services.NewOrderHistoryService(mockRepo)

	client := &models.Client{
		ClientName:   "testClient",
		ExchangeName: "testExchange",
		Label:        "testLabel",
		Pair:         "BTC/USDT",
	}

	expectedHistory := []*models.HistoryOrder{
		{
			ClientName:   "testClient",
			ExchangeName: "testExchange",
			Label:        "testLabel",
			Pair:         "BTC/USDT",
			Side:         "buy",
			Type:         "limit",
			BaseQty:      1.0,
			Price:        1000.0,
		},
	}

	mockRows := new(mocks.MockRows)
	mockRows.On("Next").Return(true).Once()
	mockRows.On("Scan", mock.AnythingOfType("*string"), mock.AnythingOfType("*string"), mock.AnythingOfType("*string"), mock.AnythingOfType("*string"), mock.AnythingOfType("*string"), mock.AnythingOfType("*string"), mock.AnythingOfType("*float64"), mock.AnythingOfType("*float64")).Run(func(args mock.Arguments) {
		*(args[0].(*string)) = "testClient"
		*(args[1].(*string)) = "testExchange"
		*(args[2].(*string)) = "testLabel"
		*(args[3].(*string)) = "BTC/USDT"
		*(args[4].(*string)) = "buy"
		*(args[5].(*string)) = "limit"
		*(args[6].(*string)) = ""
		*(args[7].(*float64)) = 1000.0
		*(args[8].(*float64)) = 1.0
	}).Return(nil).Once()
	mockRows.On("Next").Return(false).Once()
	mockRows.On("Err").Return(nil).Once()

	mockDB.On("Query", mock.Anything, "SELECT ...", mock.AnythingOfType("*models.Client")).Return(mockRows, nil)

	ctx := context.Background()
	history, err := service.GetOrderHistory(ctx, client)

	assert.NoError(t, err)
	assert.Equal(t, expectedHistory, history)
	mockDB.AssertExpectations(t)
	mockRows.AssertExpectations(t)
}

func TestSaveOrder(t *testing.T) {
	mockDB := new(mocks.MockDB)
	mockRepo := repositories.NewOrderHistoryRepository(mockDB)
	service := services.NewOrderHistoryService(mockRepo)

	client := &models.Client{
		ClientName:   "testClient",
		ExchangeName: "testExchange",
		Label:        "testLabel",
		Pair:         "BTC/USDT",
	}

	order := &models.HistoryOrder{
		ClientName:   "testClient",
		ExchangeName: "testExchange",
		Label:        "testLabel",
		Pair:         "BTC/USDT",
		Side:         "buy",
		Type:         "limit",
		BaseQty:      1.0,
		Price:        1000.0,
	}

	mockDB.On("Exec", mock.Anything, "INSERT INTO ...", mock.AnythingOfType("*models.Client"), mock.AnythingOfType("*models.HistoryOrder")).Return(nil, nil)

	ctx := context.Background()
	err := service.SaveOrder(ctx, client, order)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}
