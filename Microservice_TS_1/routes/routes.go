package routes

import (
	"Microservice_TS_1/controllers"
	"github.com/gorilla/mux"
)

func SetupRouter(
	orderBookController *controllers.OrderBookController,
	orderHistoryController *controllers.OrderHistoryController,
) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/orderbook/{exchange}/{pair}", orderBookController.GetOrderBook).Methods("GET")
	router.HandleFunc("/orderbook/{exchange}/{pair}", orderBookController.SaveOrderBook).Methods("POST")
	router.HandleFunc("/orderhistory", orderHistoryController.GetOrderHistory).Methods("POST")
	router.HandleFunc("/orderhistory", orderHistoryController.SaveOrder).Methods("POST")

	return router
}
