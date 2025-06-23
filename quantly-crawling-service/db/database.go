package db

import (
	"os"

	"github.com/redis/go-redis/v9"
)

type Database struct {
	DB *redis.Client
}

func (d *Database) GetDatabase() *redis.Client {
	if d.DB == nil {
		d.SetUpDatabaseClient()
	}
	return d.DB
}

func (d *Database) SetUpDatabaseClient() {
	d.DB = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Username: os.Getenv("REDIS_USERNAME"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})
}
