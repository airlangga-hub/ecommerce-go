package payment

import (
	"errors"
	"fmt"
	"log"

	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/checkout/session"
)


type PaymentClient interface {
	CreatePayment(amount float64, userID, orderID uint) (*stripe.CheckoutSession, error)
	GetPaymentStatus(paymentID string) (*stripe.CheckoutSession, error)
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


func (p *payment) CreatePayment(amount float64, userID, orderID uint) (*stripe.CheckoutSession, error) {
	
	amountInCents := amount * 100
	
	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					UnitAmount: stripe.Int64(int64(amountInCents)),
					Currency: stripe.String("usd"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("Electronic"),
					},
				},
				Quantity: stripe.Int64(1),
			}, 
		},
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(p.successURL),
		CancelURL: stripe.String(p.cancelURL),
	}
	
	params.AddMetadata("user_id", fmt.Sprintf("%d", userID))
	params.AddMetadata("order_id", fmt.Sprintf("%d", orderID))
	
	session, err := session.New(params)
	
	if err != nil {
		log.Println("--> stripe_err CreatePayment: ", err)
		return nil, errors.New("error creating payment session")
	}
	
	return session, nil
}


func (p *payment) GetPaymentStatus(paymentID string) (*stripe.CheckoutSession, error)