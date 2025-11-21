package service

import (

	"github.com/airlangga-hub/ecommerce-go/config"
	"github.com/airlangga-hub/ecommerce-go/internal/domain"
	"github.com/airlangga-hub/ecommerce-go/internal/dto"
	"github.com/airlangga-hub/ecommerce-go/internal/helper"
	"github.com/airlangga-hub/ecommerce-go/internal/repository"
)


type CatalogService struct {
	Repo repository.CatalogRepository
	Auth *helper.Auth
	Config *config.AppConfig
}


func (s *CatalogService) CreateCategory(input dto.CreateCategoryRequest) error {

	category := domain.Category{
		Name: input.Name,
		ParentID: input.ParentID,
		ImageURL: input.ImageURL,
		DisplayOrder: input.DisplayOrder,
	}

	err := s.Repo.CreateCategory(&category)

	return err
}