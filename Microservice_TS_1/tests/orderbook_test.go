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

func TestGetOrderBook(t *testing.T) {
	mockDB := new(mocks.MockDB)
	mockRepo := repositories.NewOrderBookRepository(mockDB)
	service := services.NewOrderBookService(mockRepo)

	expectedOrderBook := []*models.DepthOrder{
		{Price: 1000.0, BaseQty: 1.0},
	}

	mockRows := new(mocks.MockRows)
	mockRows.On("Next").Return(true).Once()
	mockRows.On("Scan", mock.AnythingOfType("*float64"), mock.AnythingOfType("*float64")).Run(func(args mock.Arguments) {
		*(args[0].(*float64)) = 1000.0
		*(args[1].(*float64)) = 1.0
	}).Return(nil).Once()
	mockRows.On("Next").Return(false).Once()
	mockRows.On("Err").Return(nil).Once()

	mockDB.On("Query", mock.Anything, "SELECT ...", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(mockRows, nil)

	ctx := context.Background()
	orderBook, err := service.GetOrderBook(ctx, "Binance", "BTC/USDT")

	assert.NoError(t, err)
	assert.Equal(t, expectedOrderBook, orderBook)
	mockDB.AssertExpectations(t)
	mockRows.AssertExpectations(t)
}

func TestSaveOrderBook(t *testing.T) {
	mockDB := new(mocks.MockDB)
	mockRepo := repositories.NewOrderBookRepository(mockDB)
	service := services.NewOrderBookService(mockRepo)

	orderBook := []*models.DepthOrder{
		{Price: 1000.0, BaseQty: 1.0},
	}

	mockDB.On("Exec", mock.Anything, "INSERT INTO ...", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.Anything).Return(nil, nil)

	ctx := context.Background()
	err := service.SaveOrderBook(ctx, "Binance", "BTC/USDT", orderBook)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}
