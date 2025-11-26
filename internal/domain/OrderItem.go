package domain

import "time"


type OrderItem struct {
	ID				uint			`json:"id" gorm:"primaryKey"`
	OrderID			uint			`json:"order_id"`
	Order			Order
	ProductID		uint			`json:"product_id"`
	Product			Product
	SellerID		uint			`json:"seller_id"`
	Seller			User			`gorm:"foreignKey:SellerID"`
	Name			string			`json:"name"`
	ImageURL		string			`json:"image_url"`
	Price			float64			`json:"price"`
	Qty				uint 			`json:"qty"`
	CreatedAt		time.Time		`json:"created_at"`
	UpdatedAt		time.Time		`json:"updated_at"`
}