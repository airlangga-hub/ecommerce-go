package handlers

import (
	"github.com/airlangga-hub/ecommerce-go/internal/api/rest"
	"github.com/gofiber/fiber/v2"
)


type UserHandler struct {

}


func SetupUserRoutes(rh *rest.HttpHandler) {
	app := rh.App

	handler := &UserHandler{}

	// Public endpoints
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

	// Private endpoints
}


func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(&fiber.Map{"message": "register"})
}


func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(&fiber.Map{"message": "login"})
}