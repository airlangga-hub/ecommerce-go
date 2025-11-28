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


func (s *TransactionService) GetPayment(userID uint) (domain.Payment, error) {
	return s.Repo.FindPayment(userID)
}