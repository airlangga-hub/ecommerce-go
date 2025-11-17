package repository

import (
	"errors"
	"log"

	"github.com/airlangga-hub/ecommerce-go/internal/domain"
	"gorm.io/gorm"
)


type UserRepository interface {
	CreateUser(user domain.User) (domain.User, error)
	FindUser(email string) (domain.User, error)
	FindUserByID(id uint) (domain.User, error)
	UpdateUser(id uint, user domain.User) (domain.User, error)
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
	return domain.User{}, nil
}


func (ur *userRepository) FindUserByID(id uint) (domain.User, error) {
	return domain.User{}, nil
}


func (ur *userRepository) UpdateUser(id uint, user domain.User) (domain.User, error) {
	return domain.User{}, nil
}