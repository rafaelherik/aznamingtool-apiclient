package aznamingtool

import "github.com/rafaelherik/aznamingtool-apiclient/aznamingtool/models"

type ResourceLocationService struct {
	baseService *BaseService
}

func NewResourceLocationService(client *APIClient) *ResourceLocationService {
	return &ResourceLocationService{baseService: NewBaseService(client)}
}

func (s *ResourceLocationService) GetAllResourceLocations() (*[]models.ResourceLocation, error) {
	var response []models.ResourceLocation
	err := s.baseService.DoGet("RequestName", nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *ResourceLocationService) GetResourceLocation(id string) (*models.ResourceLocation, error) {
	var response models.ResourceLocation
	err := s.baseService.DoGet("GetResourceLocation", map[string]string{"id": id}, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *ResourceLocationService) CreateOrUpdateResourceLocation(request models.ResourceLocation) (*models.ResourceLocation, error) {
	var response models.ResourceLocation
	err := s.baseService.DoPost("CreateOrUpdateResourceLocation", request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *ResourceLocationService) DeleteResourceLocation(id string) (interface{}, error) {
	var response models.ResourceLocation
	return s.baseService.DoDelete("DeleteResourceLocation", map[string]string{"id": id}, &response)
}
