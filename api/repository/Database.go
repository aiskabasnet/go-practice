package repository

import (
	"go-practice/models"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func DB() *gorm.DB {
	db, err := gorm.Open("mysql", os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@/"+os.Getenv("DB_NAME")+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("Couldnot connect to database:" + err.Error())
		return nil
	}
	db.AutoMigrate(&models.User{}, &models.Order{}, &models.Product{})
	return db
}
