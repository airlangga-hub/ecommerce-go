package handlers

import (
	"github.com/airlangga-hub/ecommerce-go/internal/api/rest"
	"github.com/airlangga-hub/ecommerce-go/internal/dto"
	"github.com/airlangga-hub/ecommerce-go/internal/repository"
	"github.com/airlangga-hub/ecommerce-go/internal/service"
	"github.com/gofiber/fiber/v2"
)


type UserHandler struct {
	service.UserService
}


func SetupUserRoutes(rh *rest.HttpHandler) {
	app := rh.App

	userService := service.UserService{
		UserRepository: repository.NewUserRepository(rh.DB),
		Auth: rh.Auth,
	}
	handler := &UserHandler{userService}

	// Public endpoints
	publicRoutes := app.Group("/users")

	publicRoutes.Post("/register", handler.Register)
	publicRoutes.Post("/login", handler.Login)

	// Private endpoints
	privateRoutes := publicRoutes.Group("/", handler.Authorize)

	privateRoutes.Post("/verify", handler.Verify)
	privateRoutes.Get("/verify", handler.GetVerificationCode)

	privateRoutes.Post("/profile", handler.CreateProfile)
	privateRoutes.Get("/profile", handler.GetProfile)

	privateRoutes.Post("/cart", handler.AddToCart)
	privateRoutes.Get("/cart", handler.GetCart)

	privateRoutes.Get("/order", handler.GetOrders)
	privateRoutes.Get("/order/:id", handler.GetOrder)

	privateRoutes.Post("/become-seller", handler.BecomeSeller)
}


func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	user := dto.UserSignUp{}
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(400).JSON(&fiber.Map{"message": "please provide valid inputs for signup", "error": err.Error()})
	}

	token, err := h.SignUp(user)
	if err != nil {
		return ctx.Status(500).JSON(&fiber.Map{"message": "error on signup", "error": err.Error()})
	}

	return ctx.Status(200).JSON(&fiber.Map{"message": "register", "token": token})
}


func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	userLogin := dto.UserLogin{}
	err := ctx.BodyParser(&userLogin)
	if err != nil {
		return ctx.Status(401).JSON(&fiber.Map{"message": "please provide valid user_id & password", "error": err.Error()})
	}

	token, err := h.UserLogin(userLogin.Email, userLogin.Password)

	if err != nil {
		return ctx.Status(401).JSON(&fiber.Map{"message": "unauthorized user", "error": err.Error()})
	}

	return ctx.Status(200).JSON(&fiber.Map{"message": "login", "token": token})
}


func (h *UserHandler) Verify(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(&fiber.Map{"message": "verify"})
}


func (h *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {
	user := h.GetCurrentUser(ctx)

	code, err := h.CreateVerificationCode(user)

	if err != nil {
	return ctx.Status(500).JSON(&fiber.Map{"message": "failed to generate verification code", "error": err.Error()})
	}

	return ctx.Status(200).JSON(&fiber.Map{"message": "get verification code", "data": code})
}


func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(&fiber.Map{"message": "create profile"})
}


func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	user := h.GetCurrentUser(ctx)

	return ctx.Status(200).JSON(&fiber.Map{"message": "get profile", "user": user})
}


func (h *UserHandler) AddToCart(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(&fiber.Map{"message": "add to cart"})
}


func (h *UserHandler) GetCart(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(&fiber.Map{"message": "get cart"})
}


func (h *UserHandler) GetOrders(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(&fiber.Map{"message": "get orders"})
}


func (h *UserHandler) GetOrder(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(&fiber.Map{"message": "get order by id"})
}


func (h *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(&fiber.Map{"message": "become seller"})
}