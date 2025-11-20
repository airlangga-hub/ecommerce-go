package service

import (
	"github.com/airlangga-hub/ecommerce-go/config"
	"github.com/airlangga-hub/ecommerce-go/internal/helper"
	"github.com/airlangga-hub/ecommerce-go/internal/repository"
)


type CatalogService struct {
	repository.CatalogRepository
	*helper.Auth
	*config.AppConfig
}