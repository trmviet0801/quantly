package usecase

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/trmviet0801/quantly/dto"
	"github.com/trmviet0801/quantly/utils"
	"go.uber.org/zap"
)

// POST request to ALPACE to place an order for account
func SubmitOrder(orderDto *dto.OrderDto, accountId string) {
	if orderDto.IsValid() {
		godotenv.Load()

		client := &http.Client{
			Timeout: 10 * time.Minute,
		}

		url := os.Getenv("ALPACE_ORDER_BASE_URL") + accountId + os.Getenv("ALPACE_ORDER_CREATE_ORDER")

		requestBodyJson, err := json.Marshal(orderDto)
		if utils.IsError(err, "can not parse order dto struct") {
			return
		}

		requestBody := bytes.NewReader(requestBodyJson)

		request, err := http.NewRequest("POST", url, requestBody)
		if utils.IsError(err, "can not create new request") {
			return
		}
		request.SetBasicAuth(os.Getenv("ALPACA_API_KEY"), os.Getenv("ALPACA_API_SECRET"))
		request.Header.Set("Content-Type", "application/json")

		response, err := client.Do(request)
		if utils.IsError(err, "can not send request") {
			return
		}
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if utils.IsError(err, "can not read response body") {
			return
		}

		var data dto.OrderResponseDto
		err = json.Unmarshal(body, &data)
		if utils.IsError(err, "can not parse response body") {
			return
		}

		zap.L().Info("Response", zap.String("value", string(data.String())))
	}
}
