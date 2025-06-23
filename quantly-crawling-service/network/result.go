package network

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/trmviet0801/quantly/quantly-crawling-serivce/models"
	"github.com/trmviet0801/quantly/quantly-crawling-serivce/utils"
)

// Pre-Handle response from http request
// Returns error if request fails
// Converts Json to object if request successed
func Result[T any](response *http.Response) (T, error) {
	var result T
	switch response.StatusCode {
	case 200:
		result, err := extractData[T](response)
		if err != nil {
			utils.OnError(err)
			return result, err
		}

		return result, nil
	default:
		fmt.Println(response.Request.URL)
		body, _ := io.ReadAll(response.Body)
		err := fmt.Errorf("request failed with status code %d msg : %s", response.StatusCode, string(body))
		utils.OnError(err)
		return result, err
	}
}

// Convert json in http response (from BrightData) to object
func extractData[T any](response *http.Response) (T, error) {
	var result T

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		utils.OnError(err)
		return result, err
	}

	switch interface{}(result).(type) {
	case []models.Stock:
		jsonString, isOk := utils.WrapJSONObjectAsArray(body)

		if !isOk {
			err := fmt.Errorf("can not format response body")
			return result, err
		}
		err := json.Unmarshal([]byte(jsonString), &result)
		if err != nil {
			return result, err
		}
	default:
		err = json.Unmarshal(body, &result)
		if err != nil {
			return result, err
		}
	}

	return result, nil
}
