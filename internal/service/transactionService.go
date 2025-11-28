package service

import (
	"github.com/airlangga-hub/ecommerce-go/internal/domain"
	"github.com/airlangga-hub/ecommerce-go/internal/helper"
	"github.com/airlangga-hub/ecommerce-go/internal/repository"
	"github.com/stripe/stripe-go/v78"
)


type TransactionService struct{
	Repo			repository.TransactionRepository
	Auth			*helper.Auth
}


func (s *TransactionService) GetActivePayment(userID uint) (domain.Payment, error) {
	return s.Repo.FindActivePayment(userID)
}


func (s *TransactionService) SavePayment(userID uint, sc *stripe.CheckoutSession, amount float64, orderRef string) error {
	 
	payment := domain.Payment{
		UserID: userID,
		OrderRef: orderRef,
		Amount: amount,
		Status: domain.PaymentStatusInitial,
		PaymentURL: sc.URL,
		PaymentID: sc.ID,
	}
	
	return s.Repo.CreatePayment(payment)
}