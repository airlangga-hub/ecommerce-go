package api

import (
	"log"

	"github.com/airlangga-hub/ecommerce-go/config"
	"github.com/airlangga-hub/ecommerce-go/internal/api/rest"
	"github.com/airlangga-hub/ecommerce-go/internal/api/rest/handlers"
	"github.com/airlangga-hub/ecommerce-go/internal/domain"
	"github.com/airlangga-hub/ecommerce-go/internal/helper"
	"github.com/airlangga-hub/ecommerce-go/pkg/payment"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func StartServer(cfg *config.AppConfig) {

	db, err := gorm.Open(postgres.Open(cfg.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to db: ", err)
	}

	log.Println("database connected")

	// db migration
	err = db.AutoMigrate(
		&domain.User{},
		&domain.BankAccount{},
		&domain.Category{},
		&domain.Product{},
		&domain.CartItem{},
		&domain.Address{},
		&domain.Order{},
		&domain.OrderItem{},
		&domain.Payment{},
	)

	if err != nil {
		log.Fatal("DB Migration failed")
	}
	log.Println("DB Migration successful")

	auth := &helper.Auth{Secret: cfg.AppSecret}
	
	paymentClient := payment.NewPaymentClient(cfg.StripeSecret, cfg.SuccessURL, cfg.CancelURL)
	
	app := fiber.New()

	httpHandler := &rest.HttpHandler{
		App: app,
		DB: db,
		Auth: auth,
		Config: cfg,
		PaymentClient: paymentClient,
	}

	setupRoutes(httpHandler)

	app.Listen(cfg.ServerPort)
}


func setupRoutes(rh *rest.HttpHandler) {
	handlers.SetupCatalogRoutes(rh)
	handlers.SetupUserRoutes(rh)
	handlers.SetupTransactionRoutes(rh)
}