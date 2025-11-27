package repository

import (
	"github.com/airlangga-hub/ecommerce-go/internal/domain"
	"gorm.io/gorm"
)


type TransactionRepository interface{
	CreatePayment(payment *domain.Payment) error
	FindOrders(userID uint) ([]*domain.Order, error)
	FindOrderByID(id, userID uint) (domain.Order, error)
}


type transactionRepository struct {
	db	*gorm.DB
}


func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}


func (r *transactionRepository) CreatePayment(payment *domain.Payment) error


func (r *transactionRepository) FindOrders(userID uint) ([]*domain.Order, error)


func (r *transactionRepository) FindOrderByID(id, userID uint) (domain.Order, error)