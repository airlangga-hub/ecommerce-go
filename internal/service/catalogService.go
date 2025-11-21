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


func (s *CatalogService) GetCategories() ([]*domain.Category, error) {

	categories, err := s.Repo.FindCategories()

	if err != nil {
		return nil, err
	}

	return categories, err
}


func (s *CatalogService) GetCategoryByID(id uint) (*domain.Category, error) {

	category, err := s.Repo.FindCategoryByID(id)

	if err != nil {
		return nil, err
	}

	return category, err
}


func (s *CatalogService) EditCategory(id uint, input dto.CreateCategoryRequest) (*domain.Category, error) {

	category := &domain.Category{
		ID: id,
		Name: input.Name,
		ParentID: input.ParentID,
		ImageURL: input.ImageURL,
		DisplayOrder: input.DisplayOrder,
	}

	category, err := s.Repo.EditCategory(category)

	return category, err
}


func (s *CatalogService) DeleteCategory(id uint) error {

	if err := s.Repo.DeleteCategory(id); err != nil {
		return err
	}

	return nil
}


func (s *CatalogService) CreateProduct(userID uint, input dto.CreateProduct) error {

	product := domain.Product{
		Name: input.Name,
		Description: input.Description,
		UserID: userID,
		CategoryID: input.CategoryID,
		ImageURL: input.ImageURL,
		Price: input.Price,
		Stock: input.Stock,
	}

	return s.Repo.CreateProduct(product)
}


func (s *CatalogService) GetProducts() ([]*domain.Product, error) {

	products, err := s.Repo.FindProducts()

	return products, err
}


func (s *CatalogService) GetProductByID(id uint) (domain.Product, error) {

	product, err := s.Repo.FindProductByID(id)

	return product, err
}


func (s *CatalogService) EditProduct(id uint, input dto.CreateProduct) (domain.Product, error) {

	product := domain.Product{
		ID: id,
		Description: input.Description,
		CategoryID: input.CategoryID,
		ImageURL: input.ImageURL,
		Price: input.Price,
		Stock: input.Stock,
	}

	product, err := s.Repo.EditProduct(product)

	return product, err
}


func (s *CatalogService) UpdateStock(id uint, input dto.UpdateStock) (domain.Product, error) {

	product := domain.Product{
		ID: id,
		Stock: input.Stock,
	}

	product, err := s.Repo.EditProduct(product)

	return product, err
}


func (s *CatalogService) DeleteProduct(id uint) error {

	return s.Repo.DeleteProduct(domain.Product{ID: id})
}