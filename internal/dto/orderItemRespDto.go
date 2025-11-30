package dto

import "time"


type OrderItemResponse struct {
	Name      string `json:"name"`
	ImageURL  string `json:"image_url"`
	Price     float64 `json:"price"`
	Qty       uint   `json:"qty"`

	Buyer struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
		Phone string `json:"phone"`
	} `json:"buyer"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}