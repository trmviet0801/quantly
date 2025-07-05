package db

import (
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
	"github.com/trmviet0801/quantly/quantly-crawling-serivce/utils"
)

type Database struct {
	DB *redis.Client
}

func (d *Database) GetDatabase() (*redis.Client, error) {
	if d.DB == nil {
		d.Init()
	}
	return d.DB, nil
}

// Create connection to redis db
// Create indexes for redis db
func (d *Database) Init() error {
	if d.DB != nil {
		return nil
	}
	d.SetUpDatabaseClient()
	return d.SetUpIndexes()
}

// Create all pre-definded indexes
func (d *Database) SetUpIndexes() error {
	if d.DB != nil {
		err := SetUpSnapshotOverviewIndex(d.DB)
		if err != nil {
			return err
		}

		err = SetUpStockIndex(d.DB)
		if err != nil {
			return err
		}

		err = SetUpSnapshotIndex(d.DB)
		if err != nil {
			return err
		}
		return nil
	}
	err := fmt.Errorf("can not create indexes: DB is nil")
	utils.OnError(err)
	return err
}

func (d *Database) SetUpDatabaseClient() {
	d.DB = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Username: os.Getenv("REDIS_USERNAME"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})
}
