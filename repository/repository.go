package repository

import (
	"sync"
	"transaction_service/models"
)

type TransactionRepository struct {
	store map[int64]models.Transaction
	mu    *sync.RWMutex
}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{
		store: make(map[int64]models.Transaction),
		mu:    &sync.RWMutex{},
	}
}

func (r *TransactionRepository) Save(transactionID int64, transaction models.Transaction) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.store[transactionID] = transaction
}

func (r *TransactionRepository) FindByID(transactionID int64) (models.Transaction, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	transaction, exists := r.store[transactionID]
	return transaction, exists
}

func (r *TransactionRepository) FindByType(transactionType string) []int64 {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var ids []int64
	for id, transaction := range r.store {
		if transaction.Type == transactionType {
			ids = append(ids, id)
		}
	}
	return ids
}

func (r *TransactionRepository) GetAll() map[int64]models.Transaction {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.store
}
