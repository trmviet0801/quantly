package db

import (
	"fmt"
	"os"
	"sync"

	"github.com/trmviet0801/quantly/quantly-crawling-serivce/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlDB struct {
	DB *gorm.DB
}

var once sync.Once

func (m *MysqlDB) GetDatabase() *gorm.DB {
	once.Do(func() {
		dns := fmt.Sprintf(
			"%v:%v@tcp(%v:%v)/%v",
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)

		DB, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
		if err != nil {
			panic("can not connect do mysql database")
		}

		modelsToMigrate := []any{
			&models.SnapshotOverview{},
		}

		for _, model := range modelsToMigrate {
			if err = DB.AutoMigrate(model); err != nil {
				panic("can not migrate mysql database")
			}
		}
	})
	return m.DB
}
