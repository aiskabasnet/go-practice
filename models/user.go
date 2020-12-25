package models

import (
	"time"

	"github.com/gin-gonic/gin"
)

//User => Database Modal
type User struct {
	ID        string    `gorm:"primary_key;not null;unique" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Username  string    `json:"username`
	UserType  string    `json:"user_type`
	Password  string    `json:"password" binding:"required"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//TableName => Name of table
func (User) TableName() string {
	return "user"
}

//GetDesiredUser => get user
func (u User) GetDesiredUser() gin.H {
	return gin.H{
		"id":      u.ID,
		"name":    u.Name,
		"address": u.Address,
		"email":   u.Email,
		"phone":   u.Phone,
	}
}
