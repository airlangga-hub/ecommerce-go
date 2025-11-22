package domain

import "time"

type CartItem struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	ImageURL  string    `json:"image_url"`
	Price     float64   `json:"price"`
	Qty       uint      `json:"qty"`
	CreatedAt time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:current_timestamp"`
	UserID    uint      `json:"user_id" gorm:"uniqueIndex:idx_user_product"`
	ProductID uint      `json:"product_id" gorm:"uniqueIndex:idx_user_product"`
	SellerID  uint      `json:"seller_id"`

	User    User    `gorm:"foreignKey:UserID"`
	Product Product `gorm:"foreignKey:ProductID"`
	Seller  User    `gorm:"foreignKey:SellerID"`
}
