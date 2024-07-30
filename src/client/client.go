package client

import (
	"azure-naming-api-client/config"
	"net/http"
	"time"
)

type APIClient struct {
	BaseURL      string
	APIKey       string
	ApiEndpoints map[string]string
	Client       *http.Client
}

func NewAPIClient(baseURL, apiKey string) *APIClient {

	cfg := config.LoadConfig(nil, nil)

	return &APIClient{
		BaseURL:      cfg.BaseURL,
		APIKey:       cfg.APIKey,
		ApiEndpoints: cfg.ApiEndpoints,
		Client: &http.Client{
			Timeout: time.Second * 30,
		},
	}
}

func (c *APIClient) DoRequest(req *http.Request) (*http.Response, error) {
	req.Header.Set("APIKey", c.APIKey)
	return c.Client.Do(req)
}
