package service

import (
	"github.com/airlangga-hub/ecommerce-go/internal/domain"
	"github.com/airlangga-hub/ecommerce-go/internal/helper"
	"github.com/airlangga-hub/ecommerce-go/internal/repository"
)


type TransactionService struct{
	Repo			repository.TransactionRepository
	Auth			*helper.Auth
}


func (s *TransactionService) SavePayment(payment *domain.Payment) error {
	
	return s.Repo.CreatePayment(payment)
}


func (s *TransactionService) GetActivePayment(userID uint) (*domain.Payment, error) {
	return s.Repo.FindActivePayment(userID)
}


func (s *TransactionService) UpdatePayment(payment *domain.Payment) error {
		
	return s.Repo.UpdatePayment(payment)
}


func (s *TransactionService) GetOrderItems(sellerID uint) ([]*domain.OrderItem, error) {
	
	return s.Repo.FindOrderItems(sellerID)
}


func (s *TransactionService) GetOrderItemByID(id uint) (domain.OrderItem, error) {
	
	return s.Repo.FindOrderItemByID(id)
}