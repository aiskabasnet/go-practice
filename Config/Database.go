package Config

import (
	"fmt"

	"os"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

//BuildDBConfig hhhn
func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     os.Getenv("HOST"),
		Port:     3306,
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	}
	return &dbConfig
}

//DbURL -> dd
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Database,
	)
}
