package domain

import "time"


type Product struct {
	ID 				uint 		`json:"id" gorm:"PrimaryKey"`
	Name 			string 		`json:"name" gorm:"index"`
	Description 	string 		`json:"description"`
	Price			float64		`json:"price"`
	Stock			uint		`json:"stock"`
	ImageURL 		string 		`json:"image_url"`
	UserID 			uint 		`json:"user_id"`
	CategoryID	 	uint 		`json:"category_id"`
	CreatedAt		time.Time	`json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt		time.Time	`json:"updated_at" gorm:"default:current_timestamp"`
}