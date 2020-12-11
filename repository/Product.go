package repository

import (
	"go-practice/Models"

	"github.com/jinzhu/gorm"
)

type ProductRepository interface {
	GetAllProducts() ([]Models.Product, error)
	GetProductById(int) (Models.Product, error)
	CreateProduct(Models.Product) (Models.Product, error)
	UpdateProduct(Models.Product) (Models.Product, error)
	DeleteProduct(int) (Models.Product, error)
}

type productRepository struct {
	connection *gorm.DB
}

func NewProductRepository() ProductRepository {
	return &productRepository{
		connection: DB(),
	}
}

func (db *productRepository) GetAllProducts() (products []Models.Product, err error) {
	return products, db.connection.Find(&products).Error
}

func (db *productRepository) GetProductById(id int) (product Models.Product, err error) {
	return product, db.connection.First(&product, id).Error
}

func (db *productRepository) CreateProduct(product Models.Product) (Models.Product, error) {
	return product, db.connection.Create(&product).Error
}

func (db *productRepository) UpdateProduct(product Models.Product) (Models.Product, error) {
	if err := db.connection.First(&product, &product.ID).Error; err != nil {
		return product, err
	}
	return product, db.connection.Updates(&product).Error
}

func (db *productRepository) DeleteProduct(id int) (product Models.Product, err error) {
	if err := db.connection.Find(&product, &product.ID).Error; err != nil {
		return product, err
	}
	return product, db.connection.Delete(&product).Error
}
