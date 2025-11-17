package api

import (
	"github.com/airlangga-hub/ecommerce-go/config"
	"github.com/airlangga-hub/ecommerce-go/internal/api/rest"
	"github.com/airlangga-hub/ecommerce-go/internal/api/rest/handlers"
	"github.com/gofiber/fiber/v2"
)


func StartServer(cfg config.AppConfig) {
	app := fiber.New()

	httpHandler := rest.HttpHandler{
		App: app,
	}

	setupRoutes(&httpHandler)

	app.Listen(cfg.ServerPort)
}


func setupRoutes(rh *rest.HttpHandler) {
	handlers.SetupUserRoutes(rh)
}