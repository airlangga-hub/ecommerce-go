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

	tx := cr.db.Find(&categories)

	if tx.Error != nil {
		log.Print("db_err FindCategories: ", tx.Error)
		return nil, errors.New("error finding categories")
	}

	if len(categories) < 1 {
		return nil, errors.New("no categories found")
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

	tx := cr.db.Updates(c)

	if tx.Error != nil {
		log.Print("db_err EditCategory: ", tx.Error)
		return nil, errors.New("failed to update category")
	}

	if tx.RowsAffected == 0 {
		return nil, errors.New("category not found")
	}

	return c, nil
}


func (cr *catalogRepository) DeleteCategory(id uint) error {

	tx := cr.db.Delete(&domain.Category{ID: id})

	if tx.Error != nil {
		log.Print("db_err DeleteCategory: ", tx.Error)
		return errors.New("error deleting category")
	}

	if tx.RowsAffected == 0 {
		return errors.New("category not found, failed to delete")
	}

	return nil
}