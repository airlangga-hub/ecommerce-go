package handlers

import (
	"github.com/airlangga-hub/ecommerce-go/internal/api/rest"
	"github.com/airlangga-hub/ecommerce-go/internal/repository"
	"github.com/airlangga-hub/ecommerce-go/internal/service"
	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	Svc service.TransactionService
}

func SetupTransactionHandler(rh *rest.HttpHandler) {

	app := rh.App

	transactionService := service.TransactionService{
		Repo: repository.NewTransactionRepository(rh.DB),
		Auth: rh.Auth,
	}

	handler := TransactionHandler{Svc: transactionService}

	// buyer endpoints
	buyerRoutes := app.Group("/", handler.Svc.Auth.Authorize)
	buyerRoutes.Post("/payment", handler.MakePayment)

	// seller endpoints
	sellerRoutes := app.Group("/seller", handler.Svc.Auth.AuthorizeSeller)
	sellerRoutes.Get("/order", handler.GetOrders)
	sellerRoutes.Get("/order/:id", handler.GetOrderDetails)
}


func (h *TransactionHandler) MakePayment(ctx *fiber.Ctx) error {
	
	return ctx.Status(200).JSON(fiber.Map{
		"message": "payment success",
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