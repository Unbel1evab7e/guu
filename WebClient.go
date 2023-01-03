package guu

import (
	"encoding/json"
	"github.com/google/martian/log"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func ExecuteGet[T interface{}](urlToExecute string, params map[string]string, headers map[string]string) (*T, error) {
	client := http.Client{}

	requestUrl, err := url.Parse(urlToExecute)

	if err != nil {
		log.Errorf("Failed to convert url %s error: $2", urlToExecute, err)
		return nil, err
	}

	mapQueryParams(params, requestUrl)
	request, err := http.NewRequest(http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		log.Errorf("Failed to create request error: $1", err)
		return nil, err
	}

	mapHeaders(headers, request)

	response, err := client.Do(request)

	defer response.Body.Close()

	if err != nil {
		log.Errorf("Failed to execute GET to $1 error: $2", urlToExecute, err)
		return nil, err
	}

	buf, err := io.ReadAll(response.Body)

	if err != nil {
		log.Errorf("Failed to read response error: $1", err)
		return nil, err
	}

	var res T

	err = json.Unmarshal(buf, &res)

	if err != nil {
		log.Errorf("Failed to unmarshal response error: $1", err)
		return nil, err
	}

	return &res, nil
}

func ExecutePost[T interface{}](urlToExecute string, body any, params map[string]string, headers map[string]string) (*T, error) {
	client := http.Client{}

	requestUrl, err := url.Parse(urlToExecute)

	if err != nil {
		log.Errorf("Failed to convert url %s error: $2", urlToExecute, err)
		return nil, err
	}

	mapQueryParams(params, requestUrl)
	request, err := http.NewRequest(http.MethodPost, requestUrl.String(), nil)
	if err != nil {
		log.Errorf("Failed to create request error: $1", err)
		return nil, err
	}

	mapHeaders(headers, request)

	b, err := json.Marshal(body)

	if err != nil {
		log.Errorf("Failed to marshal body error: $1", err)
		return nil, err
	}

	request.Body = io.NopCloser(strings.NewReader(string(b)))

	response, err := client.Do(request)

	defer response.Body.Close()

	if err != nil {
		log.Errorf("Failed to execute POST to $1 error: $2", urlToExecute, err)
		return nil, err
	}

	buf, err := io.ReadAll(response.Body)

	if err != nil {
		log.Errorf("Failed to read response error: $1", err)
		return nil, err
	}

	var res T

	err = json.Unmarshal(buf, &res)

	if err != nil {
		log.Errorf("Failed to unmarshal response error: $1", err)
		return nil, err
	}

	return &res, nil
}

func mapHeaders(headers map[string]string, request *http.Request) {
	if headers != nil && len(headers) > 0 {
		for k, v := range headers {
			request.Header.Add(k, v)
		}
	}
}
func mapQueryParams(params map[string]string, startUrl *url.URL) {

	if params != nil && len(params) > 0 {
		for k, v := range params {
			startUrl.Query().Set(k, v)
		}
	}
}
