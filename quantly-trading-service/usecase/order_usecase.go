package usecase

import (
	"encoding/json"
	"io"
	"os"

	"github.com/joho/godotenv"
	"github.com/trmviet0801/quantly/dto"
	"github.com/trmviet0801/quantly/network"
	"github.com/trmviet0801/quantly/utils"
)

// POST request to ALPACE to place an order for account
func SubmitOrder(orderDto *dto.OrderDto, accountId string) *dto.OrderPostResponseDto {
	if orderDto.IsValid() {
		godotenv.Load()

		url := os.Getenv("ALPACE_ORDER_BASE_URL") + accountId + os.Getenv("ALPACE_ORDER_CREATE_ORDER")

		requestBodyJson, err := json.Marshal(orderDto)
		if utils.IsError(err, "can not parse order dto struct") {
			return nil
		}

		headers := map[string]string{
			"Content-Type": "application/json",
		}

		auth := map[string]string{
			"username": os.Getenv("ALPACA_API_KEY"),
			"password": os.Getenv("ALPACA_API_SECRET"),
		}

		response, err := network.SafeCall("POST", url, headers, requestBodyJson, auth)
		if utils.IsError(err, "can not send request") {
			return nil
		}
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if utils.IsError(err, "can not read response body") {
			return nil
		}

		var data dto.OrderPostResponseDto
		err = json.Unmarshal(body, &data)
		if utils.IsError(err, "can not parse response body") {
			return nil
		}
		return &data
	}
	return nil
}

// GET request to ALPACE to get all orders of specific account
func GetAllOrdersOfAccount(accountId string) *[]dto.OrderGetResponseDto {
	godotenv.Load()

	url := os.Getenv("ALPACE_ORDER_BASE_URL") + accountId + os.Getenv("ALPACE_ORDER_CREATE_ORDER")

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	auth := map[string]string{
		"username": os.Getenv("ALPACA_API_KEY"),
		"password": os.Getenv("ALPACA_API_SECRET"),
	}

	response, err := network.SafeCall("GET", url, headers, nil, auth)
	if err != nil {
		return nil
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if utils.IsError(err, "can not read response") {
		return nil
	}

	var data []dto.OrderGetResponseDto
	err = json.Unmarshal(body, &data)
	if utils.IsError(err, "can not parse response body") {
		return nil
	}
	return &data
}
