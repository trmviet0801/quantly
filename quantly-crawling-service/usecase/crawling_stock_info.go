package usecase

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/trmviet0801/quantly/quantly-crawling-serivce/models"
	"github.com/trmviet0801/quantly/quantly-crawling-serivce/network"
	"github.com/trmviet0801/quantly/quantly-crawling-serivce/utils"
)

func CrawlStockInfo(snapshotId string) ([]*models.Stock, error) {
	var stocksInfo []models.Stock

	err := godotenv.Load()
	if err != nil {
		utils.OnError(err)
		return nil, err
	}

	url := fmt.Sprintf("%s%s%s", os.Getenv("BRIGHT_DATA_BASE_URL"), os.Getenv("BRIGHT_DATA_GET_DATASET_URL_PREFIX"), snapshotId)

	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", os.Getenv("BRIGHT_DATA_BEARER_TOKEN")),
		"Content-Type":  "application/json",
	}

	response, err := network.SafeCall(url, "GET", headers, nil)
	if err != nil {
		return nil, err
	}

	stocksInfo, err = network.Result[[]models.Stock](response)
	if err != nil {
		return nil, err
	}

	return utils.ToPointerArray(stocksInfo), nil
}
