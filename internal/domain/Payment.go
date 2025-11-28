package domain

import "time"


type Payment struct {
	ID				uint				`json:"id" gorm:"primaryKey"`
	UserID			uint				`json:"user_id"`
	User			User				`json:"-"`
	OrderRef		string				`json:"order_ref"`
	CaptureMethod	string				`json:"capture_method"`
	Amount			float64				`json:"amount"`
	CustomerID		string				`json:"customer_id"`
	PaymentID		string				`json:"payment_id"`
	ClientSecret	string				`json:"client_secret"`
	Status			PaymentStatus		`json:"status" gorm:"default:initial"`
	Response		string				`json:"response"`
	CreatedAt		time.Time			`json:"created_at"`
	UpdatedAt		time.Time			`json:"updated_at"`
}


type PaymentStatus string


const (
	PaymentStatusInitial PaymentStatus = "initial"
	PaymentStatusSuccess PaymentStatus = "success"
	PaymentStatusFailed PaymentStatus = "failed"
	PaymentStatusPending PaymentStatus = "pending"
)