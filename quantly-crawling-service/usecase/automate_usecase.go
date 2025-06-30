package usecase

import (
	"fmt"
	"time"

	"github.com/trmviet0801/quantly/quantly-crawling-serivce/utils"
)

// Auto-trigger crawling full stocks after 1 minute
// Retry if err
func CrawlManager() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for t := range ticker.C {
		fmt.Printf("Crawl Full Stocks At: %s", t)
		err := AutoCrawlAllStocks()
		if err != nil {
			RetryCrawlAllStocks(5)
		}
	}

	select {}
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
