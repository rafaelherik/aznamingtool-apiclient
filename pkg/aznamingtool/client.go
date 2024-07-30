package aznamingtool

import (
	"net/http"
)

type APIClient struct {
	BaseURL      string
	APIKey       string
	ApiEndpoints map[string]string
	HttpClient   *http.Client
}

func NewAPIClient(baseURL string, apiKey string) *APIClient {

	return &APIClient{
		BaseURL: baseURL,
		APIKey:  apiKey,
		ApiEndpoints: map[string]string{
			"RequestName":               baseURL + "/api/ResourceNamingRequests/RequestName",
			"RequestNameWithComponents": baseURL + "/api/ResourceNamingRequests/RequestNameWithComponents",
			"ValidateName":              baseURL + "/api/ResourceNamingRequests/ValidateName",
		},
		HttpClient: &http.Client{},
	}
}

func (c *APIClient) DoRequest(req *http.Request) (*http.Response, error) {
	req.Header.Set("APIKey", c.APIKey)
	return c.HttpClient.Do(req)
}
