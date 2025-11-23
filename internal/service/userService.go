package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/airlangga-hub/ecommerce-go/config"
	"github.com/airlangga-hub/ecommerce-go/internal/domain"
	"github.com/airlangga-hub/ecommerce-go/internal/dto"
	"github.com/airlangga-hub/ecommerce-go/internal/helper"
	"github.com/airlangga-hub/ecommerce-go/internal/repository"
)


type UserService struct {
	Repo 	repository.UserRepository
	CRepo	repository.CatalogRepository
	Auth 	*helper.Auth
	Config 	*config.AppConfig
}


func (s *UserService) SignUp(input dto.UserSignUp) (string, error) {
	hashedPassword, err := s.Auth.CreateHashedPassword(input.Password)
	if err != nil {
		return "", err
	}

	user, err := s.Repo.CreateUser(
		domain.User{
			Email: input.Email,
			Password: hashedPassword,
			Phone: input.Phone,
		},
	)
	if err != nil {
		return "", nil
	}

	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
}


func (s *UserService) FindUserByEmail(email string) (domain.User, error) {
	user, err := s.Repo.FindUser(email)

	return user, err
}


func (s *UserService) UserLogin(email, password string) (string, error) {
	user, err := s.FindUserByEmail(email)
	if err != nil {
		return "", errors.New("user does not exist with the provided email")
	}

	err = s.Auth.VerifyPassword(password, user.Password)
	if err != nil {
		return "", err
	}

	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
}

func (s *UserService) isUserVerified(id uint) (domain.User, bool) {
	user, err := s.Repo.FindUserByID(id)

	return user, err == nil && user.Verified
}


func (s *UserService) CreateVerificationCode(user domain.User) (int, error) {
	// if user already verified
	if _, verified := s.isUserVerified(user.ID); verified {
		return 0, errors.New("user already verified")
	}

	// generate verification code
	code, err := s.Auth.GenerateCode()
	if err != nil {
		return 0, err
	}

	// update user
	u := domain.User{
		Expiry: time.Now().Add(time.Minute * 30),
		Code: code,
	}

	_, err = s.Repo.UpdateUser(user.ID, u)
	if err != nil {
		return 0, errors.New("failed to update user verification code")
	}

	return code, nil

	// send SMS
	// if user.Phone == "" {
	// 	return 0, errors.New("user does not have a phone number")
	// }

	// notificationClient := notification.NewNotificationClient(s.Config)

	// msg := fmt.Sprintf("Your verification code is %v", code)

	// if err := notificationClient.SendSMS(user.Phone, msg); err != nil {
	// 	return errors.New("error sending sms")
	// }

	// return nil
}


func (s *UserService) VerifyCode(id uint, code int) error {

	user, verified := s.isUserVerified(id)
	if verified {
		return errors.New("user already verified")
	}

	if user.Code != code {
		return errors.New("invalid verification code")
	}

	if user.Expiry.Before(time.Now()) {
		return errors.New("verification code expired")
	}

	u := domain.User{
		Verified: true,
	}

	_, err := s.Repo.UpdateUser(user.ID, u)

	if err != nil {
		return errors.New("failed to update user verified status")
	}

	return nil
}


func (s *UserService) CreateProfile(id uint, input dto.ProfileInput) error {
	
	user := domain.User{
		ID: id,
		FirstName: input.FirstName,
		LastName: input.LastName,
	}
	
	address := domain.Address{
		UserID: id,
		AddressLine1: input.Address.AddressLine1,
		AddressLine2: input.Address.AddressLine2,
		City: input.Address.City,
		Country: input.Address.Country,
	}
	
	if err := s.Repo.CreateProfile(user, address); err != nil {
		return err
	} 
	
	return nil
}


func (s *UserService) GetProfile(id uint) (domain.User, domain.Address, error) {
	
	user, address, err := s.Repo.GetProfile(id)
	
	if err != nil {
		return domain.User{}, domain.Address{}, err
	}
	
	return user, address, nil
}


func (s *UserService) UpdateProfile(id uint, input dto.ProfileInput) (domain.User, domain.Address, error) {
	user := domain.User{
		ID: id,
		FirstName: input.FirstName,
		LastName: input.LastName,
	}
	
	address := domain.Address{
		UserID: id,
		AddressLine1: input.Address.AddressLine1,
		AddressLine2: input.Address.AddressLine2,
		City: input.Address.City,
		Country: input.Address.Country,
	}
	
	user, address, err := s.Repo.UpdateProfile(user, address)
	
	if err != nil {
		return domain.User{}, domain.Address{}, err
	}
	
	return user, address, nil
}


func (s *UserService) UserBecomeSeller(id uint, input dto.SellerInput) (string, error) {
	// fetch user from db
	user, err := s.Repo.FindUserByID(id)
	if err != nil {
		return "", err
	}

	if user.UserType == domain.SELLER {
		return "", errors.New("user is already a seller")
	}

	// create bank account information in db
	if err := s.Repo.CreateBankAccount(domain.BankAccount{
		UserID: user.ID,
		BankAccountNumber: input.BankAccountNumber,
		SwiftCode: input.SwiftCode,
		PaymentType: input.PaymentType,
	}); err != nil {
		return "", fmt.Errorf("failed to create bank account in db: %v", err)
	}

	// update user
	user, err = s.Repo.UpdateUser(
		user.ID,
		domain.User{
			FirstName: input.FirstName,
			LastName: input.LastName,
			Phone: input.PhoneNumber,
			UserType: domain.SELLER,
		},
	)
	if err != nil {
		return "", fmt.Errorf("failed to update user type: %v", err)
	}

	// generate token
	token, err := s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
	if err != nil {
		return "", err
	}

	return token, nil
}


func (s *UserService) FindCart(userID uint) ([]*domain.CartItem, error) {
	return s.Repo.FindCartItems(userID)
}


func (s *UserService) CreateCart(input dto.CartRequest, userID uint) ([]*domain.CartItem, error) {
	
	product, err := s.CRepo.FindProductByID(input.ProductID)
	if err != nil {
		return nil, err
	}
	
	cartItem := domain.CartItem{
		ProductID: input.ProductID,
		Qty: input.Qty,
		Name: product.Name,
		ImageURL: product.ImageURL,
		Price: product.Price,
		UserID: userID,
		SellerID: product.UserID,
	}
		
	if err := s.Repo.CreateCartItem(cartItem); err != nil {
		return nil, err
	}
	
	return s.Repo.FindCartItems(userID)
}


func (s *UserService) CreateOrder(user domain.User) (int, error) {
	return 0, nil
}


func (s *UserService) GetOrders(user domain.User) ([]any, error) {
	return nil, nil
}


func (s *UserService) GetOrderByID(id, userID uint) ([]any, error) {
	return nil, nil
}