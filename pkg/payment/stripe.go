package payment

import "github.com/stripe/stripe-go/v78"


type PaymentClient interface {
	CreatePayment(amount float64, userID, orderID uint) (*stripe.CheckoutSession, error)
	GetPaymentStatus(paymentID string) (*stripe.CheckoutSession, error)
}


type payment struct {
	StripeSecretKey		string
	SuccessURL			string
	FailureURL			string
}


func NewPaymentClient(stripeSecretKey, successURL, failureURL string) PaymentClient {
	return &payment{
		StripeSecretKey: stripeSecretKey,
		SuccessURL: successURL,
		FailureURL: failureURL,
	}
}


func (p *payment) CreatePayment(amount float64, userID, orderID uint) (*stripe.CheckoutSession, error)


func (p *payment) GetPaymentStatus(paymentID string) (*stripe.CheckoutSession, error)