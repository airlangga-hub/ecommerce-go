package api

import (
	"log"

	"github.com/airlangga-hub/ecommerce-go/config"
	"github.com/airlangga-hub/ecommerce-go/internal/api/rest"
	"github.com/airlangga-hub/ecommerce-go/internal/api/rest/handlers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func StartServer(cfg config.AppConfig) {
	app := fiber.New()

	httpHandler := rest.HttpHandler{
		App: app,
	}

	db, err := gorm.Open(postgres.Open(cfg.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to db: ", err)
	}

	log.Println("database connected")
	log.Println(db)

	setupRoutes(&httpHandler)

	app.Listen(cfg.ServerPort)
}


func setupRoutes(rh *rest.HttpHandler) {
	handlers.SetupUserRoutes(rh)
}