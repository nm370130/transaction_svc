package transport

import (
	"encoding/json"
	"net/http"
	"strconv"
	"transaction_service/models"
	"transaction_service/service"

	"github.com/gorilla/mux"
)

type HTTPHandler struct {
	Service *service.TransactionService
}

func NewHTTPHandler(service *service.TransactionService) *HTTPHandler {
	return &HTTPHandler{Service: service}
}

func (h *HTTPHandler) PutTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["transaction_id"], 10, 64)

	var txn models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&txn); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.Service.CreateTransaction(id, txn)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func (h *HTTPHandler) GetTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["transaction_id"], 10, 64)

	txn, ok := h.Service.GetTransaction(id)
	if !ok {
		http.Error(w, "Transaction not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(txn)
}

func (h *HTTPHandler) GetTransactionsByType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	txnType := vars["type"]

	txnIds := h.Service.GetTransactionsByType(txnType)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(txnIds)
}

func (h *HTTPHandler) GetSum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["transaction_id"], 10, 64)

	sum := h.Service.CalculateSum(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]float64{"sum": sum})
}
