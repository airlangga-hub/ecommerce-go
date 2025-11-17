package api

import (
	"github.com/airlangga-hub/ecommerce-go/config"
	"github.com/gofiber/fiber/v2"
)


func StartServer(cfg config.AppConfig) {
	app := fiber.New()

	app.Listen(cfg.ServerPort)
}
