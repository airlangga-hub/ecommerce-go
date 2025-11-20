package main

import (
	"log"

	"github.com/airlangga-hub/ecommerce-go/config"
	"github.com/airlangga-hub/ecommerce-go/internal/api"
)


func main() {
	cfg, err := config.SetupEnv()

	if err != nil {
		log.Fatalf("config is not setup properly: %v\n", err)
	}

	api.StartServer(&cfg)
}