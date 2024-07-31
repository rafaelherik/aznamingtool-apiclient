package aznamingtool

import "github.com/rafaelherik/aznamingtool-apiclient/aznamingtool/models"

type ResourceUnitService struct {
	baseService *BaseService
}

func NewResourceUnitService(client *APIClient) *ResourceUnitService {
	return &ResourceUnitService{baseService: NewBaseService(client)}
}

func (s *ResourceUnitService) GetAllResourceUnits() (*[]models.ResourceUnit, error) {
	var response []models.ResourceUnit
	err := s.baseService.DoGet("RequestName", nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *ResourceUnitService) GetResourceUnit(id string) (*models.ResourceUnit, error) {
	var response models.ResourceUnit
	err := s.baseService.DoGet("GetResourceUnit", map[string]string{"id": id}, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *ResourceUnitService) CreateOrUpdateResourceUnit(request models.ResourceUnit) (*models.ResourceUnit, error) {
	var response models.ResourceUnit
	err := s.baseService.DoPost("CreateOrUpdateResourceUnit", request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *ResourceUnitService) DeleteResourceUnit(id string) (interface{}, error) {
	var response models.ResourceUnit
	return s.baseService.DoDelete("DeleteResourceUnit", map[string]string{"id": id}, &response)
}
