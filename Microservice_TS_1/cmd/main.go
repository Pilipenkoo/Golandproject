package main

import (
	"Microservice_TS_1/controllers"
	"Microservice_TS_1/database"
	"Microservice_TS_1/repositories"
	"Microservice_TS_1/routes"
	"Microservice_TS_1/services"
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
)

func main() {
	db := database.Connect()
	defer db.Close(context.Background())

	orderBookRepo := repositories.NewOrderBookRepository(db)
	orderHistoryRepo := repositories.NewOrderHistoryRepository(db)

	orderBookService := services.NewOrderBookService(orderBookRepo)
	orderHistoryService := services.NewOrderHistoryService(orderHistoryRepo)

	orderBookController := controllers.NewOrderBookController(orderBookService)
	orderHistoryController := controllers.NewOrderHistoryController(orderHistoryService)

	router := routes.SetupRouter(orderBookController, orderHistoryController)

	srv := &http.Server{
		Handler:      handlers.LoggingHandler(os.Stdout, router),
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Starting server on :8080")
	log.Fatal(srv.ListenAndServe())
}
