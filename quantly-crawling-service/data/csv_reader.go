package data

import (
	"encoding/csv"
	"os"

	"github.com/trmviet0801/quantly/quantly-crawling-serivce/utils"
)

// Get basic information of all stocks in the us market
func GetAllUsStockId(url string) ([]string, error) {
	rawData, err := loadUsStocksFromFile(url)
	if err != nil {
		return nil, err
	}

	var stocksId = []string{}

	for _, stock := range rawData {
		stocksId = append(stocksId, stock[0])
	}
	return stocksId, nil
}

func loadUsStocksFromFile(url string) ([][]string, error) {
	file, err := os.Open(url)
	if err != nil {
		utils.OnError(err)
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		utils.OnError(err)
		return nil, err
	}

	return records, nil
}
