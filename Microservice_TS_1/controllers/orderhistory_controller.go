package controllers

import (
	"Microservice_TS_1/models"
	"Microservice_TS_1/services"
	"encoding/json"
	"net/http"
)

type OrderHistoryController struct {
	service *services.OrderHistoryService
}

func NewOrderHistoryController(service *services.OrderHistoryService) *OrderHistoryController {
	return &OrderHistoryController{service: service}
}

func (c *OrderHistoryController) GetOrderHistory(w http.ResponseWriter, r *http.Request) {
	var client models.Client
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	history, err := c.service.GetOrderHistory(&client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(history)
}

func (c *OrderHistoryController) SaveOrder(w http.ResponseWriter, r *http.Request) {
	var order models.HistoryOrder
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var client models.Client
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.service.SaveOrder(&client, &order); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
