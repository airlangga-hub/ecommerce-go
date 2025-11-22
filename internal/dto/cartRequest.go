package dto


type CartRequest struct {
	ProductID	uint	`json:"product_id"`
	Qty			uint	`json:"qty"`
}