package database

import (
	"database/sql"
	"fmt"
	"os"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var Database *sql.DB
var once sync.Once

func GetDatabase() *sql.DB {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			panic("Can not get DB information")
		}

		databaseInformation := fmt.Sprintf(
			"%v:%v@tcp(%v:%v)/%v",
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)

		Database, err = sql.Open("mysql", databaseInformation)
		if err != nil {
			panic("Cannot open connection to database")
		}

		Database.SetConnMaxLifetime(time.Minute * 5)
		Database.SetConnMaxIdleTime(time.Minute * 3)
		Database.SetMaxOpenConns(10)
		Database.SetMaxIdleConns(5)

		fmt.Println("DATABASE CONNECTED SUCCESSFULLY")
	})
	return Database
}
