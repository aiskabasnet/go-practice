package Models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	ProductID uint
	UserID    uint
	Quantity  int     `json:"quantity"`
	User      User    `gorm:"foreignkey:UserID"`
	Product   Product `gorm:"foreignkey:ProductID"`
}

func (Order) TableName() string {
	return "order"
}
