package rest

import (
	"github.com/airlangga-hub/ecommerce-go/config"
	"github.com/airlangga-hub/ecommerce-go/internal/helper"
	"github.com/airlangga-hub/ecommerce-go/pkg/payment"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)


type HttpHandler struct {
	App 			*fiber.App
	DB 				*gorm.DB
	Auth 			*helper.Auth
	Config 			*config.AppConfig
	PaymentClient	payment.PaymentClient
}