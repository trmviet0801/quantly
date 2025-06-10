package data

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/trmviet0801/quantly/models"
	"github.com/trmviet0801/quantly/utils"
)

// Get basic information of all stocks in the us market
func GetAllUsStock() []*models.Stock {
	var rawData [][]string = loadUsStocksFromFile()
	var stocks = []*models.Stock{}

	for _, stock := range rawData {
		stocks = append(stocks, constructBasicInforForStock(stock))
	}
	return stocks
}

func constructBasicInforForStock(rawData []string) *models.Stock {
	ipoYears, err := strconv.ParseInt(rawData[3], 10, 16)
	utils.OnLogError(err, fmt.Sprintf("can not convert ipo year to int16: %v", rawData[3]))

	volume, err := strconv.ParseInt(rawData[4], 10, 64)
	utils.OnLogError(err, fmt.Sprintf("can not convert volume to int64: %v", rawData[4]))

	return &models.Stock{
		Symbol:   rawData[0],
		Name:     rawData[1],
		Country:  rawData[2],
		IpoYear:  int16(ipoYears),
		Volume:   volume,
		Sector:   rawData[5],
		Industry: rawData[6],
	}
}

func loadUsStocksFromFile() [][]string {
	err := godotenv.Load()
	utils.OnLogError(err, "can not load environment variable")

	file, err := os.Open(os.Getenv("US_STOCK_RAW_DATA"))
	utils.OnLogError(err, "Can not open csv file")
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	utils.OnLogError(err, "Can not read stock-file")

	return records
}
