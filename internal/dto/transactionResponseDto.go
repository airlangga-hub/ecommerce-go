package dto

type SellerOrderDetails struct {
	OrderRefNumber  int    `json:"order_ref_number"`
	OrderStatus     int    `json:"order_status"`
	CreatedAt       string `json:"created_at"`
	OrderItemID     uint   `json:"order_item_id"`
	ProductID       uint   `json:"product_id"`
	Name            string `json:"name"`
	ImageURL        string `json:"image_url"`
	Price           string `json:"price"`
	Qty             uint   `json:"qty"`
	CustomerName    string `json:"customer_name"`
	CustomerEmail   string `json:"customer_email"`
	CustomerPhone   string `json:"customer_phone"`
	CustomerAddress string `json:"customer_address"`
}