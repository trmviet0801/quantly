package usecase

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/trmviet0801/quantly/quantly-crawling-serivce/data"
	"github.com/trmviet0801/quantly/quantly-crawling-serivce/dto"
	"github.com/trmviet0801/quantly/quantly-crawling-serivce/models"
	"github.com/trmviet0801/quantly/quantly-crawling-serivce/network"
	"github.com/trmviet0801/quantly/quantly-crawling-serivce/utils"
)

// Get dataset that ready from BrightData
// If multiple stocks -> BrightData does not wrap it in a list -> have to re-format response body
func GetCrawledData(snapshotId string) ([]*models.Stock, error) {
	var stocksInfo []models.Stock

	url := fmt.Sprintf("%s%s%s", os.Getenv("BRIGHT_DATA_BASE_URL"), os.Getenv("BRIGHT_DATA_GET_DATASET_URL_SUB_URL"), snapshotId)

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

// Trigger crawling process for all stocks on BrightData
// Load stocks from local files -> Create POST request to BrightData -> Response snapshotId
func TriggerFullStockCrawl() (*dto.CrawlRequestResponseDto, error) {
	stockIds, err := data.GetAllUsStockId(os.Getenv("SnP500_URL"))
	if err != nil {
		return nil, err
	}

	requestBody := dto.CrawlRequestPostDto{}
	requestBody.ConstructBody(stockIds)

	url := fmt.Sprintf("%s%s", os.Getenv("BRIGHT_DATA_BASE_URL"), os.Getenv("BRIGHT_DATA_TRIGGER_CRAWLING_SUB_URL"))

	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", os.Getenv("BRIGHT_DATA_BEARER_TOKEN")),
		"Content-Type":  "application/json",
	}

	jsonRequestBody, err := json.Marshal(requestBody)
	if err != nil {
		err = fmt.Errorf("can not marshal request body for url: %s", url)
		utils.OnError(err)
		return nil, err
	}

	response, err := network.SafeCall(url, "POST", headers, []byte(jsonRequestBody))
	if err != nil {
		return nil, err
	}

	result, err := network.Result[dto.CrawlRequestResponseDto](response)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
