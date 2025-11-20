package repository

import (
	"errors"
	"log"

	"github.com/airlangga-hub/ecommerce-go/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)


type UserRepository interface {
	CreateUser(user domain.User) (domain.User, error)
	FindUser(email string) (domain.User, error)
	FindUserByID(id uint) (domain.User, error)
	UpdateUser(id uint, user domain.User) (domain.User, error)
	CreateBankAccount(bank domain.BankAccount ) error
}


type userRepository struct {
	db *gorm.DB
}


func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}


func (ur *userRepository) CreateUser(user domain.User) (domain.User, error) {
	err := ur.db.Create(&user).Error

	if err != nil {
		log.Println("error creating user: ", err)
		return domain.User{}, errors.New("failed to create user")
	}

	return user, nil
}


func (ur *userRepository) FindUser(email string) (domain.User, error) {
	var user domain.User

	err := ur.db.First(&user, "email=?", email).Error

	if err != nil {
		log.Println("find user error: ", err)
		return domain.User{}, errors.New("user does not exist")
	}

	return user, nil
}


func (ur *userRepository) FindUserByID(id uint) (domain.User, error) {
	var user domain.User

	err := ur.db.First(&user, id).Error

	if err != nil {
		log.Println("find user error: ", err)
		return domain.User{}, errors.New("user does not exist")
	}
	return user, nil
}


func (ur *userRepository) UpdateUser(id uint, user domain.User) (domain.User, error) {
	u := domain.User{
		ID: id,
	}

	err := ur.db.Model(&u).Clauses(clause.Returning{}).Updates(user).Error

	if err != nil {
		log.Println("error on update: ", err)
		return domain.User{}, errors.New("failed to update user")
	}

	return u, nil
}


func (ur *userRepository) CreateBankAccount(bank domain.BankAccount ) error {
	return ur.db.Create(&bank).Error
}
