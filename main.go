package main

import (
	"fmt"
	"go-practice/Config"
	"go-practice/Models"
	"go-practice/Routes"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var err error

func main() {
	err1 := godotenv.Load()
	if err1 != nil {
		log.Fatal("Error loading")
	}

	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{}, &Models.Article{}, &Models.Order{}, &Models.Product{})
	r := Routes.SetupRouter()
	//running
	r.Run(os.Getenv("PORT"))
}
