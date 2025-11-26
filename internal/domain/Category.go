package domain

import "time"


type Category struct {
	ID 				uint 		`json:"id" gorm:"PrimaryKey"`
	Name 			string 		`json:"name" gorm:"index;unique;not null"`
	ParentID 		*uint 		`json:"parent_id"`
	Parent			*Category	`gorm:"foreignKey:ParentID"`
	Products 		[]*Product 	`json:"products"`
	ImageURL 		string	 	`json:"image_url"`
	DisplayOrder 	int 		`json:"display_order"`
	CreatedAt		time.Time	`json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt		time.Time	`json:"updated_at" gorm:"default:current_timestamp"`
}