package repository

import (
	"go-practice/Models"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	AddUser(Models.User) (Models.User, error)
	GetAllUsers() ([]Models.User, error)
	GetUserByID(int) (Models.User, error)
	UpdateUser(Models.User) (Models.User, error)
	DeleteUser(int) (Models.User, error)
	GetProductOrdered(id int) ([]Models.Order, error)
}

type userRepository struct {
	connection *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{connection: DB()}
}

func (db *userRepository) AddUser(user Models.User) (Models.User, error) {
	return user, db.connection.Create(&user).Error
}

func (db *userRepository) GetAllUsers() (users []Models.User, err error) {
	return users, db.connection.Find(&users).Error
}

func (db *userRepository) GetUserByID(id int) (user Models.User, err error) {
	return user, db.connection.First(&user, id).Error
}

func (db *userRepository) UpdateUser(user Models.User) (Models.User, error) {
	if err := db.connection.First(&user, &user.ID).Error; err != nil {
		return user, err
	}
	return user, db.connection.Updates(&user).Error
}
func (db *userRepository) DeleteUser(id int) (user Models.User, err error) {
	if err := db.connection.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, db.connection.Delete(&user).Error
}

func (db *userRepository) GetProductOrdered(id int) (orders []Models.Order, err error) {
	return orders, db.connection.Where("id=?", id).Find(&orders).Error
}
