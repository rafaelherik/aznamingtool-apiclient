package aznamingtool

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type BaseService struct {
	client *APIClient
}

func NewBaseService(client *APIClient) *BaseService {
	return &BaseService{client: client}
}

func (s *BaseService) DoPost(endpointKey string, requestData interface{}, response interface{}) error {
	endpoint, exists := s.client.ApiEndpoints[endpointKey]
	if !exists {
		return fmt.Errorf("endpoint not found for %s", endpointKey)
	}

	var req *http.Request
	var err error

	if requestData != nil {
		jsonData, err := json.Marshal(requestData)
		if err != nil {
			return err
		}
		req, err = http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return err
		}
	} else {
		req, err = http.NewRequest("POST", endpoint, nil)
		if err != nil {
			return err
		}
	}

	resp, err := s.client.DoRequest(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("received error status code: %d", resp.StatusCode)
	}

	return nil
}
