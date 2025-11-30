package domain

import "time"


type Order struct {
	ID					uint			`json:"id" gorm:"primaryKey"`
	UserID				uint			`json:"user_id" gorm:"index;not null"`
	User				User			`json:"user"`
	Status				string			`json:"status"`
	Amount				float64			`json:"amount"`
	TransactionID		string			`json:"transaction_id"`
	OrderRef			string			`json:"order_ref"`
	PaymentID			string			`json:"payment_id"`
	OrderItems			[]*OrderItem	`json:"-"`
	CreatedAt			time.Time		`json:"created_at"`
	UpdatedAt			time.Time		`json:"updated_at"`
}