package service

import (
	"github.com/airlangga-hub/ecommerce-go/config"
	"github.com/airlangga-hub/ecommerce-go/internal/helper"
	"github.com/airlangga-hub/ecommerce-go/internal/repository"
)


type CatalogService struct {
	CatalogRepo repository.CatalogRepository
	Auth *helper.Auth
	Config *config.AppConfig
}