package usecase

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/trmviet0801/quantly/dto"
	"github.com/trmviet0801/quantly/network"
	"github.com/trmviet0801/quantly/utils"
)

// POST request to ALPACE to place an order for account
func SubmitOrder(orderDto *dto.OrderDto, accountId string) (*dto.OrderPostResponseDto, error) {
	if orderDto.IsValid() {
		godotenv.Load()

		url := os.Getenv("ALPACE_ORDER_BASE_URL") + accountId + os.Getenv("ALPACE_ORDER_CREATE_ORDER")

		requestBodyJson, err := json.Marshal(orderDto)
		if utils.IsError(err, "can not parse order dto struct") {
			return nil, err
		}

		headers := map[string]string{
			"Content-Type": "application/json",
		}

		response, err := network.SafeCall("POST", url, headers, requestBodyJson)
		if utils.IsError(err, "can not send request") {
			return nil, err
		}
		defer response.Body.Close()

		var data dto.OrderPostResponseDto
		isOk, err := network.OnResult(response, &data)
		if isOk {
			return &data, nil
		}
		return nil, err
	}
	return nil, fmt.Errorf("invalid input")
}

// GET request to ALPACE to get all orders of specific account
func GetAllOrdersOfAccount(accountId string) (*[]dto.OrderGetResponseDto, error) {
	godotenv.Load()

	url := os.Getenv("ALPACE_ORDER_BASE_URL") + accountId + os.Getenv("ALPACE_ORDER_CREATE_ORDER")

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	response, err := network.SafeCall("GET", url, headers, nil)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var data []dto.OrderGetResponseDto
	isOk, err := network.OnResult(response, &data)
	if isOk {
		return &data, nil
	}
	return nil, err
}

func CancelOrder(accountId string, orderId string) {
	godotenv.Load()
	url := os.Getenv("ALPACE_ORDER_BASE_URL") + accountId + os.Getenv("ALPACE_ORDER_CREATE_ORDER") + "/" + orderId
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	network.SafeCall("DELETE", url, headers, nil)
}
