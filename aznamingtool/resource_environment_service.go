package aznamingtool

import "github.com/rafaelherik/aznamingtool-apiclient/aznamingtool/models"

type ResourceEnvironmentService struct {
	baseService *BaseService
}

// NewResourceEnvironmentService creates a new instance of ResourceEnvironmentService with the provided API client.
//
// Parameters:
//   - client: A pointer to the APIClient instance.
//
// Returns:
//   - A pointer to the newly created ResourceEnvironmentService instance.
func NewResourceEnvironmentService(client *APIClient) *ResourceEnvironmentService {
	return &ResourceEnvironmentService{baseService: NewBaseService(client)}
}

func (s *ResourceEnvironmentService) GetAllResourceEnvironments() (*[]models.ResourceEnvironment, error) {
	var response []models.ResourceEnvironment
	err := s.baseService.DoGet("RequestName", nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *ResourceEnvironmentService) GetResourceEnvironment(id string) (*models.ResourceEnvironment, error) {
	var response models.ResourceEnvironment
	err := s.baseService.DoGet("GetResourceEnvironment", map[string]string{"id": id}, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *ResourceEnvironmentService) CreateOrUpdateResourceEnvironment(request models.ResourceEnvironment) (*models.ResourceEnvironment, error) {
	var response models.ResourceEnvironment
	err := s.baseService.DoPost("CreateOrUpdateResourceEnvironment", request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *ResourceEnvironmentService) DeleteResourceEnvironment(id string) (interface{}, error) {
	var response models.ResourceEnvironment
	return s.baseService.DoDelete("DeleteResourceEnvironment", map[string]string{"id": id}, &response)
}
