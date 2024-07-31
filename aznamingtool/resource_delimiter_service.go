package aznamingtool

import "github.com/rafaelherik/aznamingtool-apiclient/aznamingtool/models"

type ResourceDelimiterService struct {
	baseService *BaseService
}

func NewResourceDelimiterService(client *APIClient) *ResourceDelimiterService {
	return &ResourceDelimiterService{baseService: NewBaseService(client)}
}

func (s *ResourceDelimiterService) GetAllResourceDelimiters() (*[]models.ResourceDelimiter, error) {
	var response []models.ResourceDelimiter
	err := s.baseService.DoGet("RequestName", nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *ResourceDelimiterService) GetResourceDelimiter(id string) (*models.ResourceDelimiter, error) {
	var response models.ResourceDelimiter
	err := s.baseService.DoGet("GetResourceDelimiter", map[string]string{"id": id}, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *ResourceDelimiterService) CreateOrUpdateResourceDelimiter(request models.ResourceDelimiter) (*models.ResourceDelimiter, error) {
	var response models.ResourceDelimiter
	err := s.baseService.DoPost("CreateOrUpdateResourceDelimiter", request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
