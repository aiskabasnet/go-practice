package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Quantity    int    `json:"quantity"`
	Description string `json:"description"`
}

func (Product) TableName() string {
	return "product"
}
