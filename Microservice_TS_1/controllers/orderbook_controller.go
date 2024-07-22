package controllers

import (
	"Microservice_TS_1/models"
	"Microservice_TS_1/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type OrderBookController struct {
	service *services.OrderBookService
}

func NewOrderBookController(service *services.OrderBookService) *OrderBookController {
	return &OrderBookController{service: service}
}

func (c *OrderBookController) GetOrderBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	exchangeName := vars["exchange"]
	pair := vars["pair"]

	orderBook, err := c.service.GetOrderBook(exchangeName, pair)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(orderBook); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *OrderBookController) SaveOrderBook(w http.ResponseWriter, r *http.Request) {
	var orderBook []*models.DepthOrder
	if err := json.NewDecoder(r.Body).Decode(&orderBook); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	exchangeName := vars["exchange"]
	pair := vars["pair"]

	if err := c.service.SaveOrderBook(exchangeName, pair, orderBook); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
