package network

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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
		err := fmt.Errorf("request failed with status code %d", response.StatusCode)
		utils.OnError(err)
		return result, err
	}
}

// Convert json in http response to object
func extractData[T any](response *http.Response) (T, error) {
	var result T

	body, err := io.ReadAll(response.Body)
	if err != nil {
		utils.OnError(err)
		return result, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		utils.OnError(err)
		return result, err
	}
	return result, nil
}
