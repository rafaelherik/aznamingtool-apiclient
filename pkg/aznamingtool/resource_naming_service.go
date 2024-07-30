package aznamingtool

import (
	"aznamingtool/pkg/aznamingtool/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ResourceNamingService struct {
	client *APIClient
}

func NewResourceNamingService(client *APIClient) *ResourceNamingService {
	return &ResourceNamingService{client: client}
}

func (s *ResourceNamingService) RequestName(request models.ResourceNameRequest) (*models.ResourceNameResponse, error) {
	endpoint := s.client.ApiEndpoints["RequestName"]

	if endpoint == "" {
		return nil, fmt.Errorf("endpoint not found for Resource Name Request")
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.DoRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var resourceName models.ResourceNameResponse
	if err := json.NewDecoder(resp.Body).Decode(&resourceName); err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return &resourceName, fmt.Errorf("received error status code: %d", resp.StatusCode)
	}

	if !resourceName.Success {
		return &resourceName, fmt.Errorf("request failed: %v", resourceName)
	}

	return &resourceName, nil
}

func (s *ResourceNamingService) RequestNameWithComponents(request models.ResourceNameRequestWithComponents) (*models.ResourceNameResponse, error) {
	endpoint, exists := s.client.ApiEndpoints["RequestNameWithComponents"]
	if !exists {
		return nil, fmt.Errorf("endpoint not found for Resource Name With Components Request")
	}

	req, err := http.NewRequest("POST", endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.DoRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var resourceName models.ResourceNameResponse
	if err := json.NewDecoder(resp.Body).Decode(&resourceName); err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return &resourceName, fmt.Errorf("received error status code: %d", resp.StatusCode)
	}

	if !resourceName.Success {
		return &resourceName, fmt.Errorf("request failed: %v", resourceName)
	}

	return &resourceName, nil
}

func (s *ResourceNamingService) ValidatetName(request models.ValidateNameRequest) (*models.ValidateNameResponse, error) {
	endpoint, exists := s.client.ApiEndpoints["ValidateName"]
	if !exists {
		return nil, fmt.Errorf("endpoint not found for Validate Name Request")
	}

	req, err := http.NewRequest("POST", endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.DoRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result models.ValidateNameResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return &result, fmt.Errorf("received error status code: %d", resp.StatusCode)
	}

	return &result, nil
}
