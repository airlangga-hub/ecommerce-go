package repository

import (
	"errors"
	"log"

	"github.com/airlangga-hub/ecommerce-go/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)


type UserRepository interface {
	CreateUser(user 	domain.User) (domain.User, error)
	FindUser(email string) (domain.User, error)
	FindUserByID(id uint) (domain.User, error)
	UpdateUser(id uint, user domain.User) (domain.User, error)
	CreateBankAccount(bank domain.BankAccount) error
	
	CreateProfile(user domain.User, address domain.Address) error
	UpdateProfile(user domain.User, address *domain.Address) (domain.User, error) 
	
	FindCartItems(userID uint) ([]*domain.CartItem, error)
	FindCartItemByID(userID, productID uint) (domain.CartItem, error)
	CreateCartItem(c domain.CartItem) error
	UpdateCartItem(c domain.CartItem) (domain.CartItem, error)
	DeleteCartItem(id uint) error
	DeleteCartItems(userID uint) error
	
	CreateOrder(order domain.Order) error
	FindOrders(userID uint) ([]*domain.Order, error)
	FindOrderByID(id, userID uint) (domain.Order, error)
}


type userRepository struct {
	db *gorm.DB
}


func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}


func (ur *userRepository) CreateUser(user domain.User) (domain.User, error) {

	if err := ur.db.Create(&user).Error; err != nil {
		log.Println(" --> db_err CreateUser: ", err)
		return domain.User{}, errors.New("failed to create user")
	}

	return user, nil
}


func (ur *userRepository) FindUser(email string) (domain.User, error) {
	
	var user domain.User

	err := ur.db.First(&user, "email=?", email).Error

	if err != nil {
		log.Println(" --> db_err FindUser: ", err)
		return domain.User{}, errors.New("user does not exist")
	}

	return user, nil
}


func (ur *userRepository) FindUserByID(id uint) (domain.User, error) {
	
	user := domain.User{ID: id}

	err := ur.db.Preload("BankAccounts").
		Preload("Addresses").
		Preload("Orders").
		First(&user).Error

	if err != nil {
		log.Println(" --> db_err FindUserByID: ", err)
		return domain.User{}, errors.New("user does not exist")
	}
	
	return user, nil
}


func (ur *userRepository) UpdateUser(id uint, user domain.User) (domain.User, error) {

	user.ID = id

	tx := ur.db.Clauses(clause.Returning{}).Updates(&user)

	if err := tx.Error; err != nil {
		log.Println(" --> db_err UpdateUser: ", err)
		return domain.User{}, errors.New("failed to update user")
	}

	if tx.RowsAffected == 0 {
		return domain.User{}, errors.New("user not found, failed to update")
	}

	return user, nil
}


func (ur *userRepository) CreateBankAccount(bank domain.BankAccount) error {
	
	if err := ur.db.Create(&bank).Error; err != nil {
		log.Println(" --> db_err CreateBankAccount: ", err)
		return errors.New("failed to create bank account")
	}
	
	return nil
}


func (ur *userRepository) CreateProfile(user domain.User, address domain.Address) error {
	
	if err := ur.db.Create(&address).Error; err != nil {
		log.Println(" --> db_err CreateProfile (create address): ", err)
		return errors.New("error creating profile")
	}
	
	tx := ur.db.Updates(user)
	
	if err := tx.Error; err != nil {
		log.Println(" --> db_err CreateProfile (update user): ", err)
		return errors.New("error creating profile")		
	}
	
	if tx.RowsAffected == 0 {
		return errors.New("user not found, failed to create profile")
	}
	
	return nil
}


func (ur *userRepository) UpdateProfile(user domain.User, address *domain.Address) (domain.User, error) {
	
	tx := ur.db.Updates(user)
	
	if err := tx.Error; err != nil {
		log.Println(" --> db_err UpdateProfile (update user): ", err)
		return domain.User{}, errors.New("error updating profile")
	}
	
	if tx.RowsAffected == 0 {
		return domain.User{}, errors.New("user not found, failed to update profile")
	}

	if address != nil {
		
		tx = ur.db.Updates(address)
		
		if err := tx.Error; err != nil {
			log.Println(" --> db_err UpdateProfile (update address): ", err)
			return domain.User{}, errors.New("error updating profile")
		}
		
		if tx.RowsAffected == 0 {
			return domain.User{}, errors.New("user not found, failed to update profile")
		}
	}
	
	if err := ur.db.Preload("Addresses").First(&user).Error; err != nil {
		log.Println(" --> db_err UpdateProfile (find user): ", err)
		return domain.User{}, errors.New("error finding user")
	}
		
	return user, nil
}


func (ur *userRepository) CreateCartItem(c domain.CartItem) error {
	
	if err := ur.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "user_id"}, {Name: "product_id"}},
		DoUpdates: clause.Assignments(map[string]any{
			"qty": gorm.Expr("excluded.qty"),
			"updated_at": gorm.Expr("now()"),
		}),
	}).Create(&c).Error; err != nil {
		log.Println(" --> db_err CreateCartItem: ", err)
		return errors.New("error creating cart item")
	}
	
	return nil
}


func (ur *userRepository) FindCartItems(userID uint) ([]*domain.CartItem, error) {
	
	cartItems := []*domain.CartItem{}
	
	tx := ur.db.Find(&cartItems, "user_id = ?", userID)
	
	if err := tx.Error; err != nil {
		log.Println(" --> db_err FindCartItems: ", err)
		return nil, errors.New("error finding cart items")
	}
	
	if len(cartItems) == 0 {
		return nil, errors.New("user id not found, failed to find cart items")
	}
	
	return cartItems, nil
}


func (ur *userRepository) FindCartItemByID(userID, productID uint) (domain.CartItem, error) {
	
	cartItem := domain.CartItem{}
	
	if err := ur.db.First(&cartItem, "user_id = ? AND product_id = ?", userID, productID).Error; err != nil {
		log.Println(" --> db_err FindCartItemByID: ", err)
		return domain.CartItem{}, errors.New("cart item not found")
	}
	
	return cartItem, nil
}


func (ur *userRepository) UpdateCartItem(c domain.CartItem) (domain.CartItem, error) {
	
	tx := ur.db.Clauses(clause.Returning{}).Updates(&c)
	
	if err := tx.Error; err != nil {
		log.Println(" --> db_err UpdateCartItem: ", err)
		return domain.CartItem{}, errors.New("error updating cart item")
	}
	
	if tx.RowsAffected == 0 {
		return domain.CartItem{}, errors.New("cart item not found, failed to update")
	}
	
	return c, nil
}


func (ur *userRepository) DeleteCartItem(id uint) error {
	
	tx := ur.db.Delete(&domain.CartItem{ID: id})
	
	if err := tx.Error; err != nil {
		log.Println(" --> db_err DeleteCartItem: ", err)
		return errors.New("error deleting cart item")
	}
	
	if tx.RowsAffected == 0 {
		return errors.New("cart item not found, failed to delete")
	}
	
	return nil
}


 func (ur *userRepository) DeleteCartItems(userID uint) error {
 
	tx := ur.db.Delete(&domain.CartItem{}, "user_id = ?", userID)
	
	if err := tx.Error; err != nil {
		log.Println(" --> db_err DeleteCartItems: ", err)
		return errors.New("error deleting cart items")
	}
	
	if tx.RowsAffected == 0 {
		return errors.New("user id not found, failed to delete cart items")
	}
	
	return nil
}


func (ur *userRepository) CreateOrder(order domain.Order) error {
	
	if err := ur.db.Create(&order).Error; err != nil {
		log.Println("--> db_err CreateOrder: ", err)
		return errors.New("failed to create order")
	}
	
	return nil
}


func (ur *userRepository) FindOrders(userID uint) ([]*domain.Order, error){
	
	orders := []*domain.Order{}
	
	tx := ur.db.Where("user_id=?", userID).Find(&orders)
	
	if err := tx.Error; err != nil {
		log.Println("--> db_err FindOrders: ", err)
		return nil, errors.New("failed to find orders")
	}
	
	if len(orders) == 0 {
		return nil, errors.New("user has no orders")
	}
	
	return orders, nil
}


func (ur *userRepository) FindOrderByID(id, userID uint) (domain.Order, error) {
	
	order := domain.Order{ID: id}
	
	if err := ur.db.Preload("OrderItems").Where("user_id=?", userID).First(&order).Error; err != nil {
		log.Println("--> db_err FindOrderByID: ", err)
		return domain.Order{}, errors.New("failed to find order by id")
	}
	
	return order, nil
}