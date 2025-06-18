package network

import (
	"bytes"
	"net/http"
	"time"

	"github.com/trmviet0801/quantly/utils"
)

func SafeCall(method string, url string, header map[string]string, body []byte, auth map[string]string) (*http.Response, error) {
	client := &http.Client{
		Timeout: 10 * time.Minute,
	}

	request, err := http.NewRequest(method, url, bytes.NewReader(body))
	if utils.IsError(err, "can not create new request") {
		return nil, err
	}

	setHeaderForRequest(request, header)
	setAuthAttributesForRequest(request, auth)

	response, err := client.Do(request)
	if utils.IsError(err, "can not send request") {
		return nil, err
	}

	return response, nil
}

func setHeaderForRequest(request *http.Request, headers map[string]string) {
	for key, value := range headers {
		request.Header.Set(key, value)
	}
}

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
