package infrastructure

import (
	"fmt"
	"go-practice/utils"
	"log"
	"os"
	"time"

	// go-sql-driver/mysql
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// SetupModels : initializing mysql database
func SetupModels() *gorm.DB {
	get := utils.GetEnvWithKey
	USER := get("DB_USERNAME")
	PASS := get("DB_PASSWORD")
	HOST := get("HOST_NAME")
	PORT := get("DB_PORT")
	DBNAME := get("DB_NAME")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // Disable color
		},
	)

	createDBDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	database, err := gorm.Open(mysql.Open(createDBDsn), &gorm.Config{Logger: newLogger})
	_ = database.Exec("CREATE DATABASE IF NOT EXISTS " + DBNAME + ";")

	if err != nil {
		panic(err.Error())
	}
	return database
}
