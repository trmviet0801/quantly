package automate

import (
	"fmt"
	"time"

	"github.com/trmviet0801/quantly/data"
	"github.com/trmviet0801/quantly/database"
	"github.com/trmviet0801/quantly/repos/model_repo"
	"go.uber.org/zap"
)

func AutomateController() {
	done := make(chan bool)
	go getNewData(done)
	<-done
}

// crawl stocks data
func getNewData(done chan bool) {
	count := 1
	zap.L().Info("Starting crawl stocks data")
	for 0 == 0 {
		url := "./res/us-stocks-" + fmt.Sprintf("%d", count) + ".csv"

		stocks := data.GetStocksFinancialIndexes(url)

		stockRepo := model_repo.StockRepo{
			DB: database.GetDatabase(),
		}

		for _, stock := range stocks {
			zap.L().Info("Crawling completed", zap.String("Symbol", stock.Symbol))

			time, _ := time.Parse("2006-01-02 15:04:05", "1970-01-01 00:00:00")
			stock.LatestTradeTime = time

			if stock.CurrentPrice != 0 {
				err := stockRepo.Update(stock)
				if err != nil {
					zap.L().Error("Can not save stock into DB",
						zap.String("Symbol", stock.Symbol),
					)
				} else {
					zap.L().Info("Save stock into DB successfully",
						zap.String("Symbol", stock.Symbol),
					)
				}
			}
		}

		if count+1 == 60 {
			count = 1
		} else {
			count++
		}

		time.Sleep(1 * time.Minute)
	}
	done <- true
}
