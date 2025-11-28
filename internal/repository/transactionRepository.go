package repository

import (
	"errors"
	"log"

	"github.com/airlangga-hub/ecommerce-go/internal/domain"
	"github.com/airlangga-hub/ecommerce-go/internal/dto"
	"gorm.io/gorm"
)


type TransactionRepository interface{
	CreatePayment(payment *domain.Payment) error
	FindActivePayment(userID uint) (*domain.Payment, error)
	UpdatePayment(payment *domain.Payment) error
	
	FindOrders(userID uint) ([]*domain.Order, error)
	FindOrderByID(id, userID uint) (dto.SellerOrderDetails, error)
}


type transactionRepository struct {
	db	*gorm.DB
}


func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}


func (r *transactionRepository) CreatePayment(payment *domain.Payment) error {
	
	if err := r.db.Create(payment).Error; err != nil {
		log.Println("--> db_err CreatePayment: ", err)
		return errors.New("error creating payment")
	}
	
	return nil
}


func (r *transactionRepository) FindActivePayment(userID uint) (*domain.Payment, error) {
	
	payment := &domain.Payment{}
	
	if err := r.db.First(payment, "user_id=? and status='initial'", userID).Error; err != nil {
		log.Println("--> db_err FindActivePayment: ", err)
		return nil, errors.New("payment does not exist")
	}
	
	return payment, nil
}


func (r *transactionRepository) UpdatePayment(payment *domain.Payment) error {
	
	tx := r.db.Updates(payment)
	
	if err := tx.Error; err != nil {
		log.Println("--> db_err UpdatePayment: ", err)
		return errors.New("error updating payment")
	}
	
	if tx.RowsAffected == 0 {
		return errors.New("payment not found, error updating payment")
	}
	
	return nil
}


func (r *transactionRepository) FindOrders(userID uint) ([]*domain.Order, error) {
	
	orders := []*domain.Order{}
	
	tx := r.db.Find(&orders, "user_id=?", userID)
	
	if err := tx.Error; err != nil {
		log.Println("--> db_err FindOrders: ", err)
		return nil, errors.New("error finding orders")
	}
	
	if len(orders) == 0 {
		return nil, errors.New("user_id not found, couldn't find orders")
	}
	
	return orders, nil
}


func (r *transactionRepository) FindOrderByID(id, userID uint) (dto.SellerOrderDetails, error) {
	
	order := domain.Order{}
	
	if err := r.db.First(&order, "id=? and user_id=?", id, userID).Error; err != nil {
		log.Println("--> db_err FindOrderByID: ", err)
		return dto.SellerOrderDetails{}, errors.New("error finding order by id")
	}
	
	return dto.SellerOrderDetails{}, nil
}