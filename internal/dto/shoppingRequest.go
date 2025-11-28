package dto


type CartRequest struct {
	ProductID	uint	`json:"product_id"`
	Qty			uint	`json:"qty"`
}


type CreatePayment struct {
	UserID			uint		`json:"user_id"`
	Amount			float64		`json:"amount"`
	ClientSecret	string		`json:"client_secret"`
	PaymentID		string		`json:"payment_id"`
	OrderRef		string		`json:"order_ref"`
}