package repository

import (
	"go-practice/models"
	"log"

	"gorm.io/gorm"
)

type UserRepository interface {
	AddUser(models.User) (models.User, error)
	GetAllUsers() ([]models.User, error)
	GetUserByID(int) (models.User, error)
	GetUserByEmail(string) (models.User, error)
	UpdateUser(models.User) (models.User, error)
	DeleteUser(int) (models.User, error)
	GetProductOrdered(id int) ([]models.Order, error)
	Migrate() error
}

type userRepository struct {
	connection *gorm.DB
}

// NewUserRepository => new user repo
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{connection: db}
}

func (db *userRepository) Migrate() error {
	log.Print("[UserRepository]...Migrate")
	return db.connection.AutoMigrate(&models.User{})
}

func (db *userRepository) AddUser(user models.User) (models.User, error) {
	return user, db.connection.Create(&user).Error
}

func (db *userRepository) GetAllUsers() (users []models.User, err error) {
	return users, db.connection.Find(&users).Error
}

func (db *userRepository) GetUserByID(id int) (user models.User, err error) {
	return user, db.connection.First(&user, id).Error
}
func (db *userRepository) GetUserByEmail(email string) (user models.User, err error) {
	return user, db.connection.First(&user, "email=?", email).Error
}
func (db *userRepository) UpdateUser(user models.User) (models.User, error) {
	if err := db.connection.First(&user, &user.ID).Error; err != nil {
		return user, err
	}
	return user, db.connection.Updates(&user).Error
}
func (db *userRepository) DeleteUser(id int) (user models.User, err error) {
	if err := db.connection.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, db.connection.Delete(&user).Error
}

func (db *userRepository) GetProductOrdered(id int) (orders []models.Order, err error) {
	return orders, db.connection.Where("id=?", id).Find(&orders).Error
}
