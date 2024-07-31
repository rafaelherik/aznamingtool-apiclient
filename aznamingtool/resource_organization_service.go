package aznamingtool

import "github.com/rafaelherik/aznamingtool-apiclient/aznamingtool/models"

type ResourceOrganizationService struct {
	baseService *BaseService
}

func NewResourceOrganizationService(client *APIClient) *ResourceOrganizationService {
	return &ResourceOrganizationService{baseService: NewBaseService(client)}
}

func (s *ResourceOrganizationService) GetAllResourceOrganizations() (*[]models.ResourceOrganization, error) {
	var response []models.ResourceOrganization
	err := s.baseService.DoGet("RequestName", nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *ResourceOrganizationService) GetResourceOrganization(id string) (*models.ResourceOrganization, error) {
	var response models.ResourceOrganization
	err := s.baseService.DoGet("GetResourceOrganization", map[string]string{"id": id}, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *ResourceOrganizationService) CreateOrUpdateResourceOrganization(request models.ResourceOrganization) (*models.ResourceOrganization, error) {
	var response models.ResourceOrganization
	err := s.baseService.DoPost("CreateOrUpdateResourceOrganization", request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *ResourceOrganizationService) DeleteResourceOrganization(id string) (interface{}, error) {
	var response models.ResourceOrganization
	return s.baseService.DoDelete("DeleteResourceOrganization", map[string]string{"id": id}, &response)
}
