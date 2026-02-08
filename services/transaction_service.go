package services

import (
	"github.com/heru-oktafian/tugas-02/models"
	"github.com/heru-oktafian/tugas-02/repositories"
)

type TransactionService struct {
	repo *repositories.TransactionRepository
}

func NewTransactionService(repo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Checkout(items []models.CheckoutItem) (*models.Transaction, error) {
	return s.repo.CreateTransaction(items)
}

func (s *TransactionService) GetReport(startDate, endDate string) (*models.ReportResponse, error) {
	return s.repo.GetReport(startDate, endDate)
}
