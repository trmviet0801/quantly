package automate

import (
	"fmt"
	"time"

	"github.com/trmviet0801/quantly/data"
	"github.com/trmviet0801/quantly/database"
	"github.com/trmviet0801/quantly/models"
	"github.com/trmviet0801/quantly/repos/model_repo"
	"go.uber.org/zap"
)

func AutomateController() {
	go getNewData()
	select {}
}

// crawling stocks data
// saving new data [Stock + StockPrice] into DB
func getNewData() {
	count := 1
	zap.L().Info("Starting crawl stocks data")
	for {
		url := "./res/us-stocks-" + fmt.Sprintf("%d", count) + ".csv"

		stocks := data.GetStocksFinancialIndexes(url)

		stockRepo := model_repo.StockRepo{
			DB: database.GetDatabase(),
		}

		stockPriceRepo := model_repo.StockPriceRepo{
			DB: database.GetDatabase(),
		}

		for _, stock := range stocks {
			zap.L().Info("Crawling completed", zap.String("Symbol", stock.Symbol))

			time, _ := time.Parse("2006-01-02 15:04:05", "1970-01-01 00:00:00")
			stock.LatestTradeTime = time

			// CurrentPrice == 0 -> error when crawling
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
					saveStockPrice(stock, &stockPriceRepo)
				}
			}
		}

		increaseIndex(&count)

		time.Sleep(1 * time.Minute)
	}
}

func increaseIndex(count *int) {
	*count++
	if *count == 60 {
		*count = 1
	}
}

func saveStockPrice(stock *models.Stock, stockPriceRepo *model_repo.StockPriceRepo) {
	stockPrice := &models.StockPrice{
		Symbol:    stock.Symbol,
		Price:     stock.CurrentPrice,
		Timestamp: time.Now(),
	}

	err := stockPriceRepo.Update(stockPrice)
	if err != nil {
		zap.L().Info("Save stock price into DB successfully",
			zap.String("Symbol", stockPrice.Symbol),
		)
	} else {
		zap.L().Error("Can not save Stock Price",
			zap.String("Symbol", stock.Symbol),
			zap.String("Price", fmt.Sprintf("%v", stock.CurrentPrice)),
		)
	}
}
