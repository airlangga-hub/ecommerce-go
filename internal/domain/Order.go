package domain

import "time"


type Order struct {
	ID					uint			`json:"id" gorm:"primaryKey"`
	UserID				uint			`json:"user_id" gorm:"index;not null"`
	User				User			`json:"-"`
	Status				string			`json:"status"`
	Amount				float64			`json:"amount"`
	TransactionID		string			`json:"transaction_id"`
	OrderRefNumber		uint			`json:"order_ref_number"`
	PaymentID			string			`json:"payment_id"`
	OrderItems			[]*OrderItem	`json:"order_items"`
	CreatedAt			time.Time		`json:"created_at"`
	UpdatedAt			time.Time		`json:"updated_at"`
}