package network

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/trmviet0801/quantly/models"
	"go.uber.org/zap"
)

func OnResult(response *http.Response, data interface{}) (bool, error) {
	if response == nil {
		zap.L().Error("response is empty")
		return false, fmt.Errorf("response is empty")
	}

	switch response.StatusCode {
	case 200, 201, 202, 203, 204, 205, 206, 207, 208, 226:
		parseHttpResponse(response, data)
		return true, nil
	default:
		return false, parseHttpErrorResponse(response)
	}
}

func parseHttpErrorResponse(response *http.Response) error {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		zap.L().Error("can not read error")
		return fmt.Errorf("can not read error")
	}

	var responseErr models.ErrorResponse
	err = json.Unmarshal(body, &responseErr)
	if err != nil {
		zap.L().Error("can not read error")
		return fmt.Errorf("can not parse error")
	}
	zap.L().Error("Error with request", zap.String("Code", fmt.Sprintf("%v", responseErr.Code)), zap.String("Msg", responseErr.Message))
	return fmt.Errorf("%v", responseErr)
}

func parseHttpResponse(response *http.Response, data interface{}) error {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		zap.L().Error(fmt.Sprintf("Can not read response body \n URL: %v", response.Request.URL))
		return fmt.Errorf("can not read response body \n URL: %v", response.Request.URL)
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		zap.L().Error(fmt.Sprintf("Can not parse response body \n URL: %v", response.Request.URL))
		return fmt.Errorf("can not parse response body \n URL: %v", response.Request.URL)
	}
	return nil
}
