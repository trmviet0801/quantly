package usecase

import (
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/trmviet0801/quantly/quantly-crawling-serivce/db"
	"github.com/trmviet0801/quantly/quantly-crawling-serivce/utils"
	"go.uber.org/zap"
)

// Auto-trigger crawling full stocks after 1 minute
// Retry if err
func CrawlManager() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for t := range ticker.C {
		fmt.Printf("Crawl Full Stocks At: %s", t)
		err := AutoCrawlAllStocks()
		if err != nil {
			RetryCrawlAllStocks(5)
		}

		SyncLatestCrawlingVersion()
	}

	select {}
}

// Check for the latest snapshot version
// Post latest snapshot version to redis (it redis-snapshot-version is outdate)
func SyncLatestCrawlingVersion() error {
	database := db.Database{}
	conn, err := database.GetDatabase()
	if err != nil {
		err = fmt.Errorf("can not create connection to redis")
		utils.OnError(err)
		return err
	}

	// Snapshot_Overviews in Redis
	snapshotOverviews, err := GetLatestSnapshot(conn)
	if err != nil {
		// No existed data in Redis
		if err == redis.Nil {

			err = PostLatestSnapshotOverview(conn)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}

	latestSnapshotOverview, err := GetLatestReadySnapshotOverviewInBD(conn)
	if err != nil {
		return err
	}

	if snapshotOverviews[0].Id != latestSnapshotOverview.Id {
		err = PostLatestSnapshotOverview(conn)
		if err != nil {
			return err
		}
		return nil
	}
	zap.L().Info("Sync data successfully", zap.String("Data", "Latest Snapshot Overview"))
	return nil
}

// Get latest snapshot_overview from BrightData -> Post to Redis
func PostLatestSnapshotOverview(conn *redis.Client) error {
	latestSnapshotOverview, err := GetLatestReadySnapshotOverviewInBD(conn)
	if err != nil {
		return err
	}

	err = PostSnapshotOverviewToRedisDB(latestSnapshotOverview, conn)
	if err != nil {
		return err
	}
	zap.L().Info("Redis Insertion Successfully", zap.String("Data", "Latest Snapshot Overview"))
	return nil
}

func RetryCrawlAllStocks(maxTimes int) {
	if maxTimes == 1 {
		err := AutoCrawlAllStocks()
		if err != nil {
			utils.OnError(err)
		}
		return
	}
	err := AutoCrawlAllStocks()
	if err != nil {
		RetryCrawlAllStocks(maxTimes - 1)
	}
}

func AutoCrawlAllStocks() error {
	snapshotOverviews, err := GetSnapshotOverviews()
	if err != nil {
		return err
	}

	runningSnapshots, err := GetRunningSnapshot(snapshotOverviews)
	if err != nil {
		return err
	}

	// Trigger new crawling process
	if len(runningSnapshots) == 0 {
		_, err := TriggerFullStockCrawl()
		if err != nil {
			return err
		}
		return nil
	}

	err = fmt.Errorf("there are running processes")
	utils.OnError(err)
	return err
}
