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


func (s *TransactionService) SavePayment(payment *domain.Payment) error {
	
	return s.Repo.CreatePayment(payment)
}


func (s *TransactionService) GetActivePayment(userID uint) (*domain.Payment, error) {
	return s.Repo.FindActivePayment(userID)
}


func (s *TransactionService) UpdatePayment(payment *domain.Payment) error {
		
	return s.Repo.UpdatePayment(payment)
}


func (s *TransactionService) GetOrderItems(sellerID uint) ([]*dto.OrderItemResponse, error) {
	
	orderItems, err := s.Repo.FindOrderItems(sellerID)
	if err != nil {
		return nil, err
	}
	
	orderItemResp := []*dto.OrderItemResponse{}
	
	for _, item := range orderItems {
		
		orderItemDto := &dto.OrderItemResponse{
			Name: item.Name,
			ImageURL: item.ImageURL,
			Price: item.Price,
			Qty: item.Qty,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		}
		orderItemDto.Buyer.ID = item.Order.UserID
		orderItemDto.Buyer.Name = item.Order.User.FirstName + " " + item.Order.User.LastName
		orderItemDto.Buyer.Email = item.Order.User.Email
		orderItemDto.Buyer.Phone = item.Order.User.Phone
		
		orderItemResp = append(orderItemResp, orderItemDto)
	}
	
	return orderItemResp, nil
}


func (s *TransactionService) GetOrderItemByID(id uint) (domain.OrderItem, error) {
	
	return s.Repo.FindOrderItemByID(id)
}