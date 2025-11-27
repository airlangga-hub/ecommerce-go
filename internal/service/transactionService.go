package service

import (
	"github.com/airlangga-hub/ecommerce-go/internal/helper"
	"github.com/airlangga-hub/ecommerce-go/internal/repository"
	"github.com/airlangga-hub/ecommerce-go/pkg/payment"
)


type TransactionService struct{
	Repo			repository.TransactionRepository
	Auth			*helper.Auth
	PaymentClient	payment.PaymentClient
}