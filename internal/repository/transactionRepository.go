package repository

import (
	"errors"
	"log"

	"github.com/airlangga-hub/ecommerce-go/internal/domain"
	"gorm.io/gorm"
)


type TransactionRepository interface{
	CreatePayment(payment *domain.Payment) error
	FindActivePayment(userID uint) (*domain.Payment, error)
	UpdatePayment(payment *domain.Payment) error
	
	FindOrders(userID uint) ([]*domain.Order, error)
	FindOrderByID(id, userID uint) (domain.Order, error)
	
	FindOrderItems(sellerID uint) ([]*domain.OrderItem, error)
	FindOrderItemByID(id uint) (domain.OrderItem, error)
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
		
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		
		log.Println("--> db_err FindActivePayment: ", err)
		return nil, errors.New("error finding payment")
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


func (r *transactionRepository) FindOrderByID(id, userID uint) (domain.Order, error) {
	
	order := domain.Order{ID: id}
	
	if err := r.db.First(&order, "user_id=?", userID).Error; err != nil {
		log.Println("--> db_err FindOrderByID: ", err)
		return domain.Order{}, errors.New("error finding order by id")
	}
	
	return order, nil
}


func (r *transactionRepository) FindOrderItems(sellerID uint) ([]*domain.OrderItem, error) {
	
	orderItems := []*domain.OrderItem{}
	
	tx := r.db.Find(&orderItems, "seller_id=?", sellerID)
	
	if err := tx.Error; err != nil {
		log.Println("--> db_err FindOrderItems: ", err)
		return nil, errors.New("couldn't find order items")
	}
	
	if tx.RowsAffected == 0 {
		return nil, errors.New("no order items found")
	}
	
	return orderItems, nil
}


func (r *transactionRepository) FindOrderItemByID(id uint) (domain.OrderItem, error) {

	orderItem := domain.OrderItem{ID: id}
	
	if err := r.db.First(&orderItem).Error; err != nil {
		log.Println("--> db_err FindOrderItemByID: ", err)
		return domain.OrderItem{}, errors.New("couldn't find order item")
	}
	
	return orderItem, nil
}