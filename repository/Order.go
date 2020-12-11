package repository

import (
	"go-practice/Models"

	"github.com/jinzhu/gorm"
)

type OrderRepository interface {
	OrderProduct(int, int, int) error
}

type orderRepository struct {
	connection *gorm.DB
}

func NewOrderRepository() OrderRepository {
	return &orderRepository{connection: DB()}
}

func (db *orderRepository) OrderProduct(userId int, productId int, quantity int) error {
	return db.connection.Create(&Models.Order{
		ProductID: uint(productId),
		Quantity:  quantity,
		UserID:    uint(userId),
	}).Error
}
