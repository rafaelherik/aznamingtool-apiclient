package aznamingtool

import "github.com/rafaelherik/aznamingtool-apiclient/aznamingtool/models"

type ResourceComponentService struct {
	baseService *BaseService
}

func NewResourceComponentService(client *APIClient) *ResourceComponentService {
	return &ResourceComponentService{baseService: NewBaseService(client)}
}

func (s *ResourceComponentService) GetAllResourceComponents() (*[]models.ResourceComponent, error) {
	var response []models.ResourceComponent
	err := s.baseService.DoGet("RequestName", nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *ResourceComponentService) GetResourceComponent(id string) (*models.ResourceComponent, error) {
	var response models.ResourceComponent
	err := s.baseService.DoGet("GetResourceComponent", map[string]string{"id": id}, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *ResourceComponentService) CreateOrUpdateResourceComponent(request models.ResourceComponent) (*models.ResourceComponent, error) {
	var response models.ResourceComponent
	err := s.baseService.DoPost("CreateOrUpdateResourceComponent", request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
