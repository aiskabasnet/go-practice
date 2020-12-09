package Models

import (
	"github.com/jinzhu/gorm"
)

//Article => returns article struct
type Article struct {
	gorm.Model
	Title string `json:"title" gorm:"size:50"`
	Desc  string `json:"description"`
}

//TableName => returns tablename
func (aa *Article) TableName() string {
	return "article"
}
