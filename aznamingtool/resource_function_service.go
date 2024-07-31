package aznamingtool

import "github.com/rafaelherik/aznamingtool-apiclient/aznamingtool/models"

type ResourceFunctionService struct {
	baseService *BaseService
}

// NewResourceFunctionService creates a new instance of ResourceFunctionService with the provided API client.
//
// Parameters:
//   - client: A pointer to the APIClient instance.
//
// Returns:
//   - A pointer to the newly created ResourceFunctionService instance.
func NewResourceFunctionService(client *APIClient) *ResourceFunctionService {
	return &ResourceFunctionService{baseService: NewBaseService(client)}
}

func (s *ResourceFunctionService) GetAllResourceFunctions() (*[]models.ResourceFunction, error) {
	var response []models.ResourceFunction
	err := s.baseService.DoGet("RequestName", nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *ResourceFunctionService) GetResourceFunction(id string) (*models.ResourceFunction, error) {
	var response models.ResourceFunction
	err := s.baseService.DoGet("GetResourceFunction", map[string]string{"id": id}, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *ResourceFunctionService) CreateOrUpdateResourceFunction(request models.ResourceFunction) (*models.ResourceFunction, error) {
	var response models.ResourceFunction
	err := s.baseService.DoPost("CreateOrUpdateResourceFunction", request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *ResourceFunctionService) DeleteResourceFunction(id string) (interface{}, error) {
	var response models.ResourceFunction
	return s.baseService.DoDelete("DeleteResourceFunction", map[string]string{"id": id}, &response)
}
