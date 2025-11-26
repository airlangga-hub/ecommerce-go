package handlers

import (
	"strconv"

	"github.com/airlangga-hub/ecommerce-go/internal/api/rest"
	"github.com/airlangga-hub/ecommerce-go/internal/dto"
	"github.com/airlangga-hub/ecommerce-go/internal/repository"
	"github.com/airlangga-hub/ecommerce-go/internal/service"
	"github.com/gofiber/fiber/v2"
)


type UserHandler struct {
	Svc *service.UserService
}


func SetupUserRoutes(rh *rest.HttpHandler) {
	app := rh.App

	userService := &service.UserService{
		Repo: repository.NewUserRepository(rh.DB),
		CRepo: repository.NewCatalogRepository(rh.DB),
		Auth: rh.Auth,
		Config: rh.Config,
	}
	handler := &UserHandler{Svc: userService}

	// Public endpoints
	publicRoutes := app.Group("/users")

	publicRoutes.Post("/register", handler.Register)
	publicRoutes.Post("/login", handler.Login)

	// Private endpoints
	privateRoutes := publicRoutes.Group("/", handler.Svc.Auth.Authorize)

	privateRoutes.Post("/verify", handler.Verify)
	privateRoutes.Get("/verify", handler.GetVerificationCode)

	privateRoutes.Post("/profile", handler.CreateProfile)
	privateRoutes.Get("/profile", handler.GetProfile)
	privateRoutes.Patch("/profile/", handler.UpdateProfile)

	privateRoutes.Post("/cart", handler.AddToCart)
	privateRoutes.Get("/cart", handler.GetCart)

	privateRoutes.Post("/order", handler.CreateOrder)
	privateRoutes.Get("/order", handler.GetOrders)
	privateRoutes.Get("/order/:id", handler.GetOrderByID)

	privateRoutes.Post("/become-seller", handler.BecomeSeller)
}


func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	user := dto.UserSignUp{}
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "invalid request body for signup",
			"error": err.Error(),
		})
	}

	token, err := h.Svc.SignUp(user)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "error on signup",
			"error": err.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "register success",
		"token": token,
	})
}


func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	userLogin := dto.UserLogin{}
	err := ctx.BodyParser(&userLogin)
	if err != nil {
		return ctx.Status(401).JSON(fiber.Map{
			"message": "please provide valid user_id & password",
			"error": err.Error(),
		})
	}

	token, err := h.Svc.UserLogin(userLogin.Email, userLogin.Password)

	if err != nil {
		return ctx.Status(401).JSON(fiber.Map{
			"message": "unauthorized user",
			"error": err.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "login success",
		"token": token,
	})
}


func (h *UserHandler) Verify(ctx *fiber.Ctx) error {
	user := h.Svc.Auth.GetCurrentUser(ctx)

	var verificationCode dto.VerificationCode

	if err := ctx.BodyParser(&verificationCode); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "invalid verification code",
			"error": err.Error(),
		})
	}

	if err := h.Svc.VerifyCode(user.ID, verificationCode.Code); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to verify code",
			"error": err.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "verified successfully",
	})
}


func (h *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {
	user := h.Svc.Auth.GetCurrentUser(ctx)

	code, err := h.Svc.CreateVerificationCode(user.ID)

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
				"message": "failed to generate verification code",
				"error": err.Error(),
			})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "get verification code success",
		"code": code,
	})
}


func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	
	user := h.Svc.Auth.GetCurrentUser(ctx)
	
	profileInput := dto.ProfileInput{}
	
	if err := ctx.BodyParser(&profileInput); err != nil {
		return rest.BadRequest(ctx, "invalid request body for create profile")
	}
	
	if err := h.Svc.CreateProfile(user.ID, profileInput); err != nil {
		return rest.ErrorResponse(ctx, 500, err)
	}
	
	return ctx.Status(200).JSON(fiber.Map{
		"message": "create profile success",
	})
}


func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	user := h.Svc.Auth.GetCurrentUser(ctx)
	
	u, address, err := h.Svc.GetProfile(user.ID)
	
	if err != nil {
		return rest.ErrorResponse(ctx, 500, err)
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "get profile success",
		"user": u,
		"address": address,
	})
}


func (h *UserHandler) UpdateProfile(ctx *fiber.Ctx) error {
	
	user := h.Svc.Auth.GetCurrentUser(ctx)
	
	profileInput := dto.ProfileInput{}
	
	if err := ctx.BodyParser(&profileInput); err != nil {
		return rest.BadRequest(ctx, "invalid request body for update profile")
	}
	
	u, address, err := h.Svc.UpdateProfile(user.ID, profileInput)
	
	if err != nil {
		return rest.ErrorResponse(ctx, 500, err)
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "update profile success",
		"user": u,
		"address": address,
	})
}


func (h *UserHandler) AddToCart(ctx *fiber.Ctx) error {
	
	cartReq := dto.CartRequest{}
	
	if err := ctx.BodyParser(&cartReq); err != nil {
		return rest.BadRequest(ctx, "invalid cart request body")
	}
	
	user := h.Svc.Auth.GetCurrentUser(ctx)
	
	cartItems, err := h.Svc.CreateCart(cartReq, user.ID)
	
	if err != nil {
		return rest.ErrorResponse(ctx, 500, err)
	}
	
	return ctx.Status(200).JSON(fiber.Map{
		"message": "add to cart success",
		"data": cartItems,
	})
}


func (h *UserHandler) GetCart(ctx *fiber.Ctx) error {
	
	user := h.Svc.Auth.GetCurrentUser(ctx)
	
	cartItems, err := h.Svc.FindCart(user.ID)
	if err != nil {
		return rest.ErrorResponse(ctx, 404, err)
	}
	
	return ctx.Status(200).JSON(fiber.Map{
		"message": "get cart success",
		"data": cartItems,
	})
}


func (h *UserHandler) CreateOrder(ctx *fiber.Ctx) error {
	
	user := h.Svc.Auth.GetCurrentUser(ctx)
	
	orderRef, err := h.Svc.CreateOrder(user.ID)
	
	if err != nil {
		return rest.ErrorResponse(ctx, 500, err)
	}
	
	return ctx.Status(200).JSON(fiber.Map{
		"message": "create order success",
		"order": orderRef,
	})
}


func (h *UserHandler) GetOrders(ctx *fiber.Ctx) error {
	
	user := h.Svc.Auth.GetCurrentUser(ctx)
	
	orders, err := h.Svc.GetOrders(user.ID)
	
	if err != nil {
		return rest.ErrorResponse(ctx, 404, err)
	}
	
	return ctx.Status(200).JSON(fiber.Map{
		"message": "get orders success",
		"orders": orders,
	})
}


func (h *UserHandler) GetOrderByID(ctx *fiber.Ctx) error {
	
	user := h.Svc.Auth.GetCurrentUser(ctx)
	
	id, _ := strconv.Atoi(ctx.Params("id"))
	
	order, err := h.Svc.GetOrderByID(uint(id), user.ID)
	
	if err != nil {
		return rest.ErrorResponse(ctx, 404, err)
	}
	
	return ctx.Status(200).JSON(fiber.Map{
		"message": "get order by id success",
		"order": order,
	})
}


func (h *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {
	user := h.Svc.Auth.GetCurrentUser(ctx)

	var sellerInput dto.SellerInput

	if err := ctx.BodyParser(&sellerInput); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "invalid seller input",
			"error": err.Error() ,
		})
	}

	token, err := h.Svc.UserBecomeSeller(user.ID, sellerInput)

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to become seller",
			"error": err.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "become seller success",
		"token": token,
	})
}