package service

import (
	"errors"
	"time"

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
	hashedPassword, err := s.CreateHashedPassword(input.Password)
	if err != nil {
		return "", err
	}

	user, err := s.CreateUser(
		domain.User{
			Email: input.Email,
			Password: hashedPassword,
			Phone: input.Phone,
		},
	)
	if err != nil {
		return "", nil
	}

	return s.GenerateToken(user.ID, user.Email, user.UserType)
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

	err = s.VerifyPassword(password, user.Password)
	if err != nil {
		return "", err
	}

	return s.GenerateToken(user.ID, user.Email, user.UserType)
}

func (s UserService) isUserVerified(id uint) bool {
	user, err := s.FindUserByID(id)

	return err == nil && user.Verified
}


func (s UserService) CreateVerificationCode(user domain.User) (int, error) {
	// if user already verified
	if s.isUserVerified(user.ID) {
		return 0, errors.New("user already verified")
	}

	// generate verification code
	code, err := s.GenerateCode()
	if err != nil {
		return 0, err
	}

	// update user
	u := domain.User{
		Expiry: time.Now().Add(time.Minute * 30),
		Code: code,
	}

	_, err = s.UpdateUser(user.ID, u)
	if err != nil {
		return 0, errors.New("failed to update user verification code")
	}

	// send SMS

	return code, nil
}


func (s UserService) VerifyCode(id uint, code int) error {
	if s.isUserVerified(id) {
		return errors.New("user already verified")
	}

	user, err := s.FindUserByID(id)

	if user.Code != code {
		return errors.New("invalid verification code")
	}
	if err != nil {
		return err
	}
	if user.Expiry.Before(time.Now()) {
		return errors.New("verification code expired")
	}

	u := domain.User{
		Verified: true,
	}

	_, err = s.UpdateUser(user.ID, u)

	if err != nil {
		return errors.New("failed to update user verified status")
	}

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