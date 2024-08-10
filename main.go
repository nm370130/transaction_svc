package main

import (
	"net/http"
	"transaction_service/repository"
	"transaction_service/service"
	"transaction_service/transport"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize repository and service
	repo := repository.NewTransactionRepository()
	service := service.NewTransactionService(repo)

	// Initialize HTTP handlers
	handler := transport.NewHTTPHandler(service)

	// Set up router
	r := mux.NewRouter()

	r.HandleFunc("/transactionservice/transaction/{transaction_id}", handler.PutTransaction).Methods("PUT")
	r.HandleFunc("/transactionservice/transaction/{transaction_id}", handler.GetTransaction).Methods("GET")
	r.HandleFunc("/transactionservice/types/{type}", handler.GetTransactionsByType).Methods("GET")
	r.HandleFunc("/transactionservice/sum/{transaction_id}", handler.GetSum).Methods("GET")

	// Start the server
	http.ListenAndServe(":8080", r)
}
