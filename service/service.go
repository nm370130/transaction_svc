package service

import (
	"transaction_service/models"
	"transaction_service/repository"
)

type TransactionService struct {
	repo *repository.TransactionRepository
}

func NewTransactionService(repo *repository.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) CreateTransaction(transactionID int64, transaction models.Transaction) {
	s.repo.Save(transactionID, transaction)
}

func (s *TransactionService) GetTransaction(transactionID int64) (models.Transaction, bool) {
	return s.repo.FindByID(transactionID)
}

func (s *TransactionService) GetTransactionsByType(transactionType string) []int64 {
	return s.repo.FindByType(transactionType)
}

func (s *TransactionService) CalculateSum(transactionID int64) float64 {
	transaction, exists := s.repo.FindByID(transactionID)
	if !exists {
		return 0
	}

	totalSum := transaction.Amount

	for id, t := range s.repo.GetAll() {
		if t.ParentID != nil && *t.ParentID == transactionID {
			totalSum += s.CalculateSum(id)
		}
	}

	return totalSum
}
