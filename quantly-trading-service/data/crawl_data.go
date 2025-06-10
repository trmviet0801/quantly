package data

import (
	"sync"

	"github.com/trmviet0801/quantly/models"
)

func GetStocksFinancialIndexes() []*models.Stock {
	var stocks []*models.Stock = GetAllUsStock()
	var wg sync.WaitGroup

	for _, stock := range stocks {
		wg.Add(1)
		go stock.GetFinancialIndex(&wg)
	}

	wg.Wait()

	return stocks
}
