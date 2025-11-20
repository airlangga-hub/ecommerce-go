package repository

import (
	"errors"
	"log"

	"github.com/airlangga-hub/ecommerce-go/internal/domain"
	"gorm.io/gorm"
)


type CatalogRepository interface {
	CreateCategory(c *domain.Category) error
	FindCategories() ([]*domain.Category, error)
	FindCategoryByID(id uint) (*domain.Category, error)
	EditCategory(c *domain.Category) (*domain.Category, error)
	DeleteCategory(id uint) error
}


type catalogRepository struct {
	db *gorm.DB
}


func NewCatalogRepository(db *gorm.DB) CatalogRepository {
	return &catalogRepository{db: db}
}


func (cr *catalogRepository) CreateCategory(c *domain.Category) error {

	if err := cr.db.Create(c).Error; err != nil {
		log.Print("db_err CreateCategory: ", err)
		return errors.New("error creating category")
	}

	return nil
}


func (cr *catalogRepository) FindCategories() ([]*domain.Category, error) {

	categories := []*domain.Category{}

	if err := cr.db.Find(&categories).Error; err != nil {
		log.Print("db_err FindCategories: ", err)
		return nil, errors.New("error finding categories")
	}

	return categories, nil
}


func (cr *catalogRepository) FindCategoryByID(id uint) (*domain.Category, error) {

	category := &domain.Category{ID: id}

	if err := cr.db.First(category).Error; err != nil {
		log.Print("db_err FindCategoryByID: ", err)
		return nil, errors.New("category not found")
	}

	return category, nil
}


func (cr *catalogRepository) EditCategory(c *domain.Category) (*domain.Category, error) {

	if err := cr.db.Save(c).Error; err != nil {
		log.Print("db_err EditCategory: ", err)
		return nil, errors.New("failed to update category")
	}

	return c, nil
}


func (cr *catalogRepository) DeleteCategory(id uint) error {
	
	if err := cr.db.Delete(&domain.Category{ID: id}).Error; err != nil {
		log.Print("db_err DeleteCategory: ", err)
		return errors.New("error deleting category")
	}

	return nil
}