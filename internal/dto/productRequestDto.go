package dto


type CreateProduct struct {
	Name			string		`json:"name"`
	Description		string		`json:"description"`
	CategoryID		uint		`json:"category_id"`
	ImageURL		string		`json:"image_url"`
	Price			float64		`json:"price"`
	Stock			uint		`json:"stock"`
}


type UpdateStock struct {
	Stock			uint		`json:"stock"`
}