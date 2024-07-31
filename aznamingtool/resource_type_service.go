package aznamingtool

import "github.com/rafaelherik/aznamingtool-apiclient/aznamingtool/models"

type ResourceTypeService struct {
	baseService *BaseService
}

func NewResourceTypeService(client *APIClient) *ResourceTypeService {
	return &ResourceTypeService{baseService: NewBaseService(client)}
}

func (s *ResourceTypeService) GetAllResourceTypes() (*[]models.ResourceType, error) {
	var response []models.ResourceType
	err := s.baseService.DoGet("RequestName", nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *ResourceTypeService) GetResourceType(id string) (*models.ResourceType, error) {
	var response models.ResourceType
	err := s.baseService.DoGet("GetResourceType", map[string]string{"id": id}, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
