package handlers

import (
	"github.com/airlangga-hub/ecommerce-go/internal/api/rest"
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
	buyerRoutes := app.Group("/", handler.Svc.Auth.Authorize)
	buyerRoutes.Post("/payment", handler.MakePayment)

	// seller endpoints
	sellerRoutes := app.Group("/seller", handler.Svc.Auth.AuthorizeSeller)
	sellerRoutes.Get("/order", handler.GetOrders)
	sellerRoutes.Get("/order/:id", handler.GetOrderDetails)
}


func (h *TransactionHandler) MakePayment(ctx *fiber.Ctx) error {
	
	user := h.Svc.Auth.GetCurrentUser(ctx)
	
	_, amount, err := h.UserSvc.FindCart(user.ID)
	if err != nil {
		return rest.ErrorResponse(ctx, 404, err)
	}
	
	orderID, err := helper.RandomNumbers(8)
	if err != nil {
		return rest.ErrorResponse(ctx, 500, err)
	}
	
	stripeCheckout, err := h.PaymentClient.CreatePayment(amount, user.ID, orderID)	
	if err != nil {
		return rest.ErrorResponse(ctx, 500, err)
	}
	
	return ctx.Status(200).JSON(fiber.Map{
		"message": "payment success",
		"result": stripeCheckout,
		"success_url": stripeCheckout.URL,
	})
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