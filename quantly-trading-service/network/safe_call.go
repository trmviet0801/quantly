package network

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/trmviet0801/quantly/utils"
	"go.uber.org/zap"
)

// Do HTTP request
// Logging if facing error
func SafeCall(method string, url string, header map[string]string, body []byte) (*http.Response, error) {
	client := &http.Client{
		Timeout: 10 * time.Minute,
	}

	var requestBody io.Reader = nil
	if body != nil {
		requestBody = bytes.NewReader(body)
	}

	request, err := http.NewRequest(method, url, requestBody)
	if utils.IsError(err, "can not create new request") {
		zap.Error(fmt.Errorf("url: %v \n err: %s", url, err.Error()))
		return nil, err
	}

	auth := map[string]string{
		"username": os.Getenv("ALPACA_API_KEY"),
		"password": os.Getenv("ALPACA_API_SECRET"),
	}

	setHeaderForRequest(request, header)
	setAuthAttributesForRequest(request, auth)

	response, err := client.Do(request)
	if utils.IsError(err, "can not send request") {
		zap.Error(fmt.Errorf("url: %v \n err: %s", url, err.Error()))
		return nil, err
	}

	return response, nil
}

// utils func that sets header (MAP) for HTTP request
func setHeaderForRequest(request *http.Request, headers map[string]string) {
	if headers == nil {
		return
	}
	for key, value := range headers {
		request.Header.Set(key, value)
	}
}

// utils funcs that sets basic authentication for alpaca api gateway
func setAuthAttributesForRequest(request *http.Request, authCredentials map[string]string) {
	if authCredentials == nil {
		return
	}
	username, userOK := authCredentials["username"]
	password, passOK := authCredentials["password"]
	if userOK && passOK {
		request.SetBasicAuth(username, password)
	}
}
