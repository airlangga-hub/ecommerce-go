package dto


type UserLogin struct {
	Email 		string 	`json:"email"`
	Password 	string 	`json:"password"`
}


type UserSignUp struct {
	UserLogin
	Phone		string	`json:"phone"`
}


type VerificationCode struct {
	Code		int		`json:"code"`
}


type SellerInput struct {
	FirstName				string		`json:"first_name"`
	LastName				string		`json:"last_name"`
	PhoneNumber				string		`json:"phone_number"`
	BankAccountNumber		uint		`json:"bank_account_number"`
	SwiftCode				string		`json:"swift_code"`
	PaymentType				string		`json:"payment_type"`
}


type AddressInput struct {
	ID				uint		`json:"id"`
	AddressLine1	string		`json:"address_line1"`
	AddressLine2	string		`json:"address_line2"`
	City			string		`json:"city"`
	Country			string		`json:"country"`
}


type ProfileInput struct {
	FirstName	string			`json:"first_name"`
	LastName	string			`json:"last_name"`
	Address		AddressInput	`json:"address"`
}