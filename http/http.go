package http

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/waretool/go-common/logger"
)

const errorMessage = "error while sending http request %s %s due to: %s"

var client *resty.Client

func init() {
	client = resty.New()
	client.SetLogger(logger.GetLogger())
}

func NewRequest() *resty.Request {
	return client.R()
}

func Get[T any](url string, headers map[string]string) (T, error) {
	request := NewRequest()

	var bodyResponse T
	response, err := request.
		SetHeaders(headers).
		SetResult(&bodyResponse).
		Get(url)

	if err != nil {
		return bodyResponse, err
	}

	if response.IsError() {
		return bodyResponse, fmt.Errorf(errorMessage, request.Method, request.URL, response.Error())
	}

	return bodyResponse, nil
}
