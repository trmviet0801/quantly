package database

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/trmviet0801/quantly/models"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var once sync.Once

func GetDatabase() *gorm.DB {
	once.Do(func() {
		var err error
		err = godotenv.Load()
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

		modelsToMigrate := []any{
			&models.Contact{},
			&models.Account{},
			&models.BalanceSheet{},
			&models.CashFlow{},
			&models.Disclosure{},
			&models.ErrorResponse{},
			&models.Identity{},
			&models.Income{},
			&models.KycResult{},
			&models.Notification{},
			&models.Order{},
			&models.Portfolio{},
			&models.PortfolioHistory{},
			&models.Position{},
			&models.QuantModel{},
			&models.Stock{},
			&models.StockPrice{},
			&models.StopLoss{},
			&models.TakeProfit{},
			&models.TradeSignal{},
			&models.TradingConfiguration{},
			&models.TrustedContact{},
			&models.User{},
		}

		for _, model := range modelsToMigrate {
			if err := DB.AutoMigrate(model); err != nil {
				zap.L().Error("AutoMigrate failed", zap.String("model", fmt.Sprintf("%T", model)), zap.Error(err))
				panic("❌ Failed to migrate DB model: " + fmt.Sprintf("%T", model))
			}
		}
	})
	return DB
}
