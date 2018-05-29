package luas_http_client

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	luasBaseURL         = "http://luasforecasts.rpa.ie"
	luasForecastPath    = "xml/get.ashx"
	luasBaseQueryParams = map[string]string{
		"action":  "forecast",
		"encrypt": "false",
	}
)

type HTTPClient struct {
	BaseURL   *url.URL
	ClientImp *http.Client
}

type Response struct {
	Body []byte
}

func NewClient(baseURL string) (*HTTPClient, error) {
	resolvedBaseURL := luasBaseURL
	if baseURL != "" {
		resolvedBaseURL = baseURL
	}
	parsedURL, err := url.Parse(resolvedBaseURL)
	if err != nil {
		return nil, err
	}
	return &HTTPClient{
		BaseURL:   parsedURL,
		ClientImp: http.DefaultClient, // Maybe support different client in the future
	}, nil
}

func (httpClient *HTTPClient) GetForecast(stop string) ([]byte, error) {
	luasBaseQueryParams["stop"] = stop
	resp, err := httpClient.Get(luasForecastPath, luasBaseQueryParams)
	delete(luasBaseQueryParams, "stop")
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (httpClient *HTTPClient) Get(path string, queryParams map[string]string) (*Response, error) {
	url := httpClient.prepareURL(path, queryParams)
	resp, err := httpClient.ClientImp.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return &Response{Body: body}, nil
}

func (httpClient *HTTPClient) prepareURL(path string, queryParams map[string]string) string {
	url := *httpClient.BaseURL // dereference for copying
	url.Path = path
	values := url.Query()
	for key, value := range queryParams {
		values.Add(key, value)
	}
	url.RawQuery = values.Encode()
	return url.String()
}
