package aznamingtool

import (
	"fmt"

	"github.com/rafaelherik/aznamingtool-apiclient/pkg/aznamingtool/models"
)

type ResourceNamingService struct {
	baseService *BaseService
}

func NewResourceNamingService(client *APIClient) *ResourceNamingService {
	return &ResourceNamingService{baseService: NewBaseService(client)}
}

func (s *ResourceNamingService) RequestName(request models.ResourceNameRequest) (*models.ResourceNameResponse, error) {
	var response models.ResourceNameResponse
	err := s.baseService.DoPost("RequestName", request, &response)
	if err != nil {
		return nil, err
	}

	if !response.Success {
		return &response, fmt.Errorf("request failed: %v", response)
	}

	return &response, nil
}

func (s *ResourceNamingService) RequestNameWithComponents(request models.ResourceNameRequestWithComponents) (*models.ResourceNameResponse, error) {
	var response models.ResourceNameResponse
	err := s.baseService.DoPost("RequestName", request, &response)
	if err != nil {
		return nil, err
	}

	if !response.Success {
		return &response, fmt.Errorf("request failed: %v", response)
	}

	return &response, nil
}

func (s *ResourceNamingService) ValidatetName(request models.ValidateNameRequest) (*models.ValidateNameResponse, error) {
	var response models.ValidateNameResponse
	err := s.baseService.DoPost("ValidateName", request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
