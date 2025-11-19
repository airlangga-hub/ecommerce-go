package api

import (
	"log"

	"github.com/airlangga-hub/ecommerce-go/config"
	"github.com/airlangga-hub/ecommerce-go/internal/api/rest"
	"github.com/airlangga-hub/ecommerce-go/internal/api/rest/handlers"
	"github.com/airlangga-hub/ecommerce-go/internal/domain"
	"github.com/airlangga-hub/ecommerce-go/internal/helper"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func StartServer(cfg config.AppConfig) {
	app := fiber.New()

	db, err := gorm.Open(postgres.Open(cfg.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to db: ", err)
	}

	log.Println("database connected")

	// db migration
	db.AutoMigrate(&domain.User{})

	auth := helper.Auth{Secret: cfg.AppSecret}

	httpHandler := &rest.HttpHandler{
		App: app,
		DB: db,
		Auth: auth,
	}

	setupRoutes(httpHandler)

	app.Listen(cfg.ServerPort)
}


func setupRoutes(rh *rest.HttpHandler) {
	handlers.SetupUserRoutes(rh)
}