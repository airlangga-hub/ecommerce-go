package domain


type Address struct {
	ID				uint		`json:"id" gorm:"primaryKey"`
	UserID			uint		`json:"user_id" gorm:"index;not null"`
	AddressLine1	string		`json:"address_line1"`
	AddressLine2	string		`json:"address_line2"`
	City			string		`json:"city"`
	Country			string		`json:"country"`
}