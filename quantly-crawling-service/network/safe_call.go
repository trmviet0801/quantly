package network

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/trmviet0801/quantly/quantly-crawling-serivce/utils"
)

// Make http request
// Using Bearer token for authentication
// Caller is responsible for closing the response body.
func SafeCall(url, method string, headers map[string]string, body []byte) (*http.Response, error) {
	client := http.Client{
		Timeout: 15 * time.Second,
	}

	var bodyReader io.Reader
	if body != nil {
		bodyReader = bytes.NewReader(body)
	}

	request, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		utils.OnError(err)
		return nil, err
	}

	setHeadersForRequest(request, headers)

	response, err := client.Do(request)
	if err != nil {
		utils.OnError(err)
		return nil, err
	}

	return response, nil
}

// Sets the provided headers on the HTTP request (Authorization: Bearer <token>)
func setHeadersForRequest(request *http.Request, headers map[string]string) {
	for k, v := range headers {
		request.Header.Set(k, v)
	}
}
