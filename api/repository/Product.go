package repository

import (
	"go-practice/models"

	"github.com/jinzhu/gorm"
)

type ProductRepository interface {
	GetAllProducts() ([]models.Product, error)
	GetProductById(int) (models.Product, error)
	CreateProduct(models.Product) (models.Product, error)
	UpdateProduct(models.Product) (models.Product, error)
	DeleteProduct(int) (models.Product, error)
}

type productRepository struct {
	connection *gorm.DB
}

func NewProductRepository() ProductRepository {
	return &productRepository{
		connection: DB(),
	}
}

func (db *productRepository) GetAllProducts() (products []models.Product, err error) {
	return products, db.connection.Find(&products).Error
}

func (db *productRepository) GetProductById(id int) (product models.Product, err error) {
	return product, db.connection.First(&product, id).Error
}

func (db *productRepository) CreateProduct(product models.Product) (models.Product, error) {
	return product, db.connection.Create(&product).Error
}

func (db *productRepository) UpdateProduct(product models.Product) (models.Product, error) {
	if err := db.connection.First(&product, &product.ID).Error; err != nil {
		return product, err
	}
	return product, db.connection.Updates(&product).Error
}

func (db *productRepository) DeleteProduct(id int) (product models.Product, err error) {
	if err := db.connection.Find(&product, &product.ID).Error; err != nil {
		return product, err
	}
	return product, db.connection.Delete(&product).Error
}
