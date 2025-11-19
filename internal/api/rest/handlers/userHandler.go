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
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

	// Private endpoints
	app.Post("/verify", handler.Verify)
	app.Get("/verify", handler.GetVerificationCode)

	app.Post("/profile", handler.CreateProfile)
	app.Get("/profile", handler.GetProfile)

	app.Post("/cart", handler.AddToCart)
	app.Get("/cart", handler.GetCart)

	app.Get("/order", handler.GetOrders)
	app.Get("/order/:id", handler.GetOrder)

	app.Post("/become-seller", handler.BecomeSeller)
}


func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	user := dto.UserSignUp{}
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(400).JSON(&fiber.Map{"message": "please provide valid inputs for signup"})
	}

	token, err := h.SignUp(user)
	if err != nil {
		return ctx.Status(500).JSON(&fiber.Map{"message": "error on signup", "reason": err.Error()})
	}

	return ctx.Status(200).JSON(&fiber.Map{"message": token})
}


func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	userLogin := dto.UserLogin{}
	err := ctx.BodyParser(&userLogin)
	if err != nil {
		return ctx.Status(401).JSON(&fiber.Map{"message": "please provide valid user_id & password"})
	}

	token, err := h.UserLogin(userLogin.Email, userLogin.Password)

	if err != nil {
		return ctx.Status(401).JSON(&fiber.Map{"message": "unauthorized user"})
	}

	return ctx.Status(200).JSON(&fiber.Map{"message": "login", "token": token})
}


func (h *UserHandler) Verify(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(&fiber.Map{"message": "verify"})
}


func (h *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(&fiber.Map{"message": "get verification code"})
}


func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(&fiber.Map{"message": "create profile"})
}


func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(&fiber.Map{"message": "get profile"})
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