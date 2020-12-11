package repository

import (
	"go-practice/Models"

	"log"
	"os"

	"github.com/jinzhu/gorm"
)

func DB() *gorm.DB {
	db, err := gorm.Open("mysql", os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@/"+os.Getenv("DB_NAME")+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("Couldnot connect to database")
		return nil
	}
	db.AutoMigrate(&Models.User{}, &Models.Article{}, &Models.Order{}, &Models.Product{})
	return db
}
