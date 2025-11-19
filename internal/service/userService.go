package service

import (
	"errors"
	"fmt"
	"log"

	"github.com/airlangga-hub/ecommerce-go/internal/domain"
	"github.com/airlangga-hub/ecommerce-go/internal/dto"
	"github.com/airlangga-hub/ecommerce-go/internal/helper"
	"github.com/airlangga-hub/ecommerce-go/internal/repository"
)


type UserService struct {
	repository.UserRepository
	helper.Auth
}


func (s UserService) SignUp(input dto.UserSignUp) (string, error) {
	log.Println(input)

	user, err := s.CreateUser(
		domain.User{
			Email: input.Email,
			Password: input.Password,
			Phone: input.Phone,
		},
	)

	userInfo := fmt.Sprintf("%v, %v, %v", user.ID, user.Email, user.UserType)

	return userInfo, err
}


func (s UserService) FindUserByEmail(email string) (*domain.User, error) {
	user, err := s.FindUser(email)

	return &user, err
}


func (s UserService) UserLogin(email, password string) (string, error) {
	user, err := s.FindUserByEmail(email)
	if err != nil {
		return "", errors.New("user does not exist with the provided email")
	}

	return user.Email, nil
}


func (s UserService) GetVerificationCode(user domain.User) (int, error) {
	return 0, nil
}


func (s UserService) VerifyCode(id uint, code int) error {
	return nil
}


func (s UserService) CreateProfile(id uint, input any) error {
	return nil
}


func (s UserService) GetProfile(id uint) (*domain.User, error) {
	return nil, nil
}


func (s UserService) UpdateProfile(id uint, input any) error {
	return nil
}


func (s UserService) BecomeSeller(id uint, input any) (string, error) {
	return "", nil
}


func (s UserService) FindCart(id uint) ([]any, error) {
	return nil, nil
}


func (s UserService) CreateCart(input any, user domain.User) ([]any, error) {
	return nil, nil
}


func (s UserService) CreateOrder(user domain.User) (int, error) {
	return 0, nil
}


func (s UserService) GetOrders(user domain.User) ([]any, error) {
	return nil, nil
}


func (s UserService) GetOrderByID(id, userID uint) ([]any, error) {
	return nil, nil
}