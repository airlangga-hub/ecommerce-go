package repository

import (
	"errors"
	"log"

	"github.com/airlangga-hub/ecommerce-go/internal/domain"
	"gorm.io/gorm"
)


type UserRepository interface {
	CreateUser(user *domain.User) (domain.User, error)
	FindUser(email string) (domain.User, error)
	FindUserByID(id uint) (domain.User, error)
	UpdateUser(id uint, user domain.User) (domain.User, error)
	CreateBankAccount(bank domain.BankAccount) error
}


type userRepository struct {
	db *gorm.DB
}


func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}


func (ur *userRepository) CreateUser(user *domain.User) (domain.User, error) {
	err := ur.db.Create(user).Error

	if err != nil {
		log.Println("error creating user: ", err)
		return domain.User{}, errors.New("failed to create user")
	}

	return *user, nil
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

	user.ID = id

	updated := domain.User{}

	tx := ur.db.Updates(user).Scan(&updated)

	if err := tx.Error; err != nil {
		log.Println("db_err UpdateUser: ", err)
		return domain.User{}, errors.New("failed to update user")
	}

	if tx.RowsAffected == 0 {
		return domain.User{}, errors.New("user not found, failed to update")
	}

	return updated, nil
}


func (ur *userRepository) CreateBankAccount(bank domain.BankAccount) error {
	return ur.db.Create(&bank).Error
}
