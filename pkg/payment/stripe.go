package payment

import (
	"errors"
	"fmt"
	"log"

	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/paymentintent"
)


type PaymentClient interface {
	CreatePayment(amount float64, userID uint, orderRef string) (*stripe.PaymentIntent, error)
	GetPaymentStatus(paymentID string) (*stripe.PaymentIntent, error)
}


type payment struct {
	stripeSecretKey		string
	successURL			string
	cancelURL			string
}


func NewPaymentClient(stripeSecretKey, successURL, cancelURL string) PaymentClient {
	return &payment{
		stripeSecretKey: stripeSecretKey,
		successURL: successURL,
		cancelURL: cancelURL,
	}
}


func (p *payment) CreatePayment(amount float64, userID uint, orderRef string) (*stripe.PaymentIntent, error) {
	
	stripe.Key = p.stripeSecretKey
	
	amountInCents := amount * 100
	
	params := &stripe.PaymentIntentParams{
		Amount: stripe.Int64(int64(amountInCents)),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
	}
	
	params.AddMetadata("user_id", fmt.Sprintf("%d", userID))
	params.AddMetadata("order_ref", orderRef)
	
	pi, err := paymentintent.New(params)
	
	if err != nil {
		log.Println("--> stripe_err CreatePayment: ", err)
		return nil, errors.New("error creating payment intent")
	}
	
	return pi, nil
}


func (p *payment) GetPaymentStatus(paymentID string) (*stripe.PaymentIntent, error) {
	
	stripe.Key = p.stripeSecretKey
	
	params := &stripe.PaymentIntentParams{}
	
	pi, err := paymentintent.Get(paymentID, params)
	
	if err != nil {
		log.Println("--> stripe_err GetPaymentStatus: ", err)
		return nil, errors.New("error getting payment status")
	}
	
	return pi, nil
	
}