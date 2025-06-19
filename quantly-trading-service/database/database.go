package database

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/trmviet0801/quantly/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var once sync.Once

func GetDatabase() *gorm.DB {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			panic("Cannot load environment variables")
		}
		dns := fmt.Sprintf(
			"%v:%v@tcp(%v:%v)/%v",
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)

		DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
		if err != nil {
			panic("Cannot connect to database: " + err.Error())
		}

		DB.AutoMigrate(&models.StockPrice{}, &models.Position{})
	})
	return DB
}
