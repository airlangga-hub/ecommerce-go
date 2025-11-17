package rest

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)


type HttpHandler struct {
	App *fiber.App
	DB *gorm.DB
}