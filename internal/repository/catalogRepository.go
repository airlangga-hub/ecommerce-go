package repository

import (
	"errors"
	"log"

	"github.com/airlangga-hub/ecommerce-go/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)


type CatalogRepository interface {
	CreateCategory(c *domain.Category) error
	FindCategories() ([]*domain.Category, error)
	FindCategoryByID(id uint) (*domain.Category, error)
	EditCategory(c *domain.Category) (*domain.Category, error)
	DeleteCategory(id uint) error

	CreateProduct(p domain.Product) error
	FindProducts() ([]*domain.Product, error)
	FindProductByID(id uint) (domain.Product, error)
	FindSellerProducts(id uint) ([]*domain.Product, error)
	EditProduct(p domain.Product) (domain.Product, error)
	DeleteProduct(p domain.Product) error
}


type catalogRepository struct {
	db *gorm.DB
}


func NewCatalogRepository(db *gorm.DB) CatalogRepository {
	return &catalogRepository{db: db}
}


func (cr *catalogRepository) CreateCategory(c *domain.Category) error {

	if err := cr.db.Create(c).Error; err != nil {
		log.Print(" --> db_err CreateCategory: ", err)
		return errors.New("error creating category")
	}

	return nil
}


func (cr *catalogRepository) FindCategories() ([]*domain.Category, error) {

	categories := []*domain.Category{}

	tx := cr.db.Find(&categories)

	if err := tx.Error; err != nil {
		log.Print(" --> db_err FindCategories: ", err)
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
		log.Print(" --> db_err FindCategoryByID: ", err)
		return nil, errors.New("category not found")
	}

	return category, nil
}


func (cr *catalogRepository) EditCategory(c *domain.Category) (*domain.Category, error) {

	tx := cr.db.Clauses(clause.Returning{}).Updates(c)

	if err := tx.Error; err != nil {
		log.Print(" --> db_err EditCategory: ", err)
		return nil, errors.New("failed to update category")
	}

	if tx.RowsAffected == 0 {
		return nil, errors.New("category not found")
	}

	return c, nil
}


func (cr *catalogRepository) DeleteCategory(id uint) error {

	tx := cr.db.Delete(&domain.Category{ID: id})

	if err := tx.Error; err != nil {
		log.Print(" --> db_err DeleteCategory: ", err)
		return errors.New("error deleting category")
	}

	if tx.RowsAffected == 0 {
		return errors.New("category not found, failed to delete")
	}

	return nil
}


func (cr *catalogRepository) CreateProduct(p domain.Product) error {

	if err := cr.db.Create(&p).Error; err != nil {
		log.Println(" --> db_err CreateProduct: ", err.Error())
		return errors.New("create product failed")
	}

	return nil
}


func (cr *catalogRepository) FindProducts() ([]*domain.Product, error) {

	products := []*domain.Product{}

	tx := cr.db.Find(&products)

	if err := tx.Error; err != nil {
		log.Print(" --> db_err FindProducts: ", err)
		return nil, errors.New("failed to find products")
	}

	if len(products) < 1 {
		return nil, errors.New("no products found")
	}

	return products, nil
}


func (cr *catalogRepository) FindProductByID(id uint) (domain.Product, error) {

	product := domain.Product{ID: id}

	if err := cr.db.First(&product).Error; err != nil {
		log.Print(" --> db_err FindProductByID: ", err)
		return domain.Product{}, errors.New("product not found")
	}

	return product, nil
}


func (cr *catalogRepository) FindSellerProducts(id uint) ([]*domain.Product, error) {

	products := []*domain.Product{}

	tx := cr.db.Where("user_id = ?", id).Find(&products)

	if err := tx.Error; err != nil {
		log.Print(" --> db_err FindSellerProducts: ", err)
		return nil, errors.New("failed to find seller products")
	}

	if len(products) == 0 {
		return nil, errors.New("no products found")
	}

	return products, nil
}


func (cr *catalogRepository) EditProduct(p domain.Product) (domain.Product, error) {

	tx := cr.db.Clauses(clause.Returning{}).Updates(&p)

	if err := tx.Error; err != nil {
		log.Print(" --> db_err EditProduct: ", err)
		return domain.Product{}, errors.New("error updating product")
	}

	if tx.RowsAffected == 0 {
		return domain.Product{}, errors.New("product not found, failed to update")
	}

	return p, nil
}


func (cr *catalogRepository) DeleteProduct(p domain.Product) error {

	tx := cr.db.Delete(&p)

	if err := tx.Error; err != nil {
		log.Print(" --> db_err DeleteProduct: ", err)
		return errors.New("error deleting product")
	}

	if tx.RowsAffected == 0 {
		return errors.New("product not found, failed to delete")
	}

	return nil
}