package handlers

import (
	"encoding/json"

	"github.com/airlangga-hub/ecommerce-go/internal/api/rest"
	"github.com/airlangga-hub/ecommerce-go/internal/domain"
	"github.com/airlangga-hub/ecommerce-go/internal/helper"
	"github.com/airlangga-hub/ecommerce-go/internal/repository"
	"github.com/airlangga-hub/ecommerce-go/internal/service"
	"github.com/airlangga-hub/ecommerce-go/pkg/payment"
	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	Svc 			service.TransactionService
	UserSvc			service.UserService
	PaymentClient	payment.PaymentClient
}

func SetupTransactionRoutes(rh *rest.HttpHandler) {

	app := rh.App

	transactionService := service.TransactionService{
		Repo: repository.NewTransactionRepository(rh.DB),
		Auth: rh.Auth,
	}
	
	userService := service.UserService{
		Repo: repository.NewUserRepository(rh.DB),
		CRepo: repository.NewCatalogRepository(rh.DB),
		Config: rh.Config,
	}

	handler := &TransactionHandler{
		Svc: transactionService,
		UserSvc: userService,
		PaymentClient: rh.PaymentClient,
	}

	// buyer endpoints
	buyerRoutes := app.Group("/buyer", handler.Svc.Auth.Authorize)
	buyerRoutes.Post("/payment", handler.MakePayment)
	buyerRoutes.Get("/verify", handler.VerifyPayment)

	// seller endpoints
	sellerRoutes := app.Group("/seller", handler.Svc.Auth.AuthorizeSeller)
	sellerRoutes.Get("/order", handler.GetOrders)
	sellerRoutes.Get("/order/:id", handler.GetOrderDetails)
}


func (h *TransactionHandler) MakePayment(ctx *fiber.Ctx) error {
	
	user := h.Svc.Auth.GetCurrentUser(ctx)
	stripePubKey := h.UserSvc.Config.StripePubKey
	
	activePayment, err := h.Svc.GetActivePayment(user.ID)
	if activePayment != nil {
		return ctx.Status(200).JSON(fiber.Map{
			"message": "create payment",
			"stripe_pub_key": stripePubKey,
			"client_secret": activePayment.ClientSecret,
		})
	}
	if err != nil {
		return rest.ErrorResponse(ctx, 500, err)
	}
	
	_, amount, err := h.UserSvc.FindCart(user.ID)
	if err != nil {
		return rest.ErrorResponse(ctx, 404, err)
	}
	
	orderRef, err := helper.RandomNumbers(8)
	if err != nil {
		return rest.ErrorResponse(ctx, 500, err)
	}
	
	paymentIntent, err := h.PaymentClient.CreatePayment(amount, user.ID, orderRef)	
	if err != nil {
		return rest.ErrorResponse(ctx, 500, err)
	}
	
	if err := h.Svc.SavePayment(&domain.Payment{
		UserID: user.ID,
		OrderRef: orderRef,
		Amount: amount,
		Status: domain.PaymentStatusInitial, 
		PaymentID: paymentIntent.ID,
		ClientSecret: paymentIntent.ClientSecret,
	}); err != nil {
		return rest.ErrorResponse(ctx, 500, err)
	}
	
	return ctx.Status(200).JSON(fiber.Map{
		"message": "create payment",
		"stripe_pub_key": stripePubKey,
		"client_secret": paymentIntent.ClientSecret,
	})
	
	// after making payment, run this cURL to make the payment status "succeeded"
	
	// curl https://api.stripe.com/v1/payment_intents/pi_xxx/confirm \
	// -u sk_test_xxx: \
	// -d payment_method=pm_card_visa \
	// -d return_url="https://example.com"
	
	// Replace pi_xxx with your actual PaymentIntent ID and sk_test_xxx with your Stripe Secret Key.
}


func (h *TransactionHandler) VerifyPayment(ctx *fiber.Ctx) error {
	
	// grab authorized user
	user := h.Svc.Auth.GetCurrentUser(ctx)
	
	// active payment exist?
	activePayment, err := h.Svc.GetActivePayment(user.ID)
	if err != nil || activePayment == nil {
		return ctx.Status(404).JSON(fiber.Map{
			"error": "no active payment exists",
		})
	}
	
	// fetch payment status from stripe
	paymentIntent, err := h.PaymentClient.GetPaymentStatus(activePayment.PaymentID)
	if err != nil {
		return rest.ErrorResponse(ctx, 404, err)
	}
	
	paymentJson, err := json.Marshal(paymentIntent)
	if err != nil {
		return rest.ErrorResponse(ctx, 500, err)
	}
	
	paymentLog := string(paymentJson)
	paymentStatus := "failed"
	
	// if payment succeeded then create order
	if paymentIntent.Status == "succeeded" {
		paymentStatus = "success"
		
		// create order
		if err := h.UserSvc.CreateOrder(user.ID, activePayment.Amount, activePayment.OrderRef, activePayment.PaymentID); err != nil {
			return rest.ErrorResponse(ctx, 500, err)
		}
	}
	
	// update payment status
	activePayment.Status = domain.PaymentStatus(paymentStatus)
	activePayment.Response = paymentLog
	if err := h.Svc.UpdatePayment(activePayment); err != nil {
		return rest.ErrorResponse(ctx, 500, err)
	}
	
	return rest.OkResponse(ctx, "payment verified successfully", paymentIntent) 
}


func (h *TransactionHandler) GetOrders(ctx *fiber.Ctx) error {
	
	return ctx.Status(200).JSON(fiber.Map{
		"message": "get orders success",
	})
}


func (h *TransactionHandler) GetOrderDetails(ctx *fiber.Ctx) error {
	
	return ctx.Status(200).JSON(fiber.Map{
		"message": "get order details success",
	})
}