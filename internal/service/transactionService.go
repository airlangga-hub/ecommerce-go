package service

import (
	"github.com/airlangga-hub/ecommerce-go/internal/domain"
	"github.com/airlangga-hub/ecommerce-go/internal/dto"
	"github.com/airlangga-hub/ecommerce-go/internal/helper"
	"github.com/airlangga-hub/ecommerce-go/internal/repository"
)


type TransactionService struct{
	Repo			repository.TransactionRepository
	Auth			*helper.Auth
}


func (s *TransactionService) SavePayment(input dto.CreatePayment) error {
	 
	payment := &domain.Payment{
		UserID: input.UserID,
		OrderRef: input.OrderRef,
		Amount: input.Amount,
		Status: domain.PaymentStatusInitial, 
		PaymentID: input.PaymentID,
		ClientSecret: input.ClientSecret,
	}
	
	return s.Repo.CreatePayment(payment)
}


func (s *TransactionService) GetActivePayment(userID uint) (*domain.Payment, error) {
	return s.Repo.FindActivePayment(userID)
}


func (s *TransactionService) UpdatePayment(payment *domain.Payment) error {
		
	return s.Repo.UpdatePayment(payment)
}