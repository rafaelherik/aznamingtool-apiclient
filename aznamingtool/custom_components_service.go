package aznamingtool

import (
	"github.com/rafaelherik/aznamingtool-apiclient/aznamingtool/models"
)

type CustomComponentService struct {
	baseService *BaseService
}

// NewCustomComponentService creates a new instance of CustomComponentService with the provided API client.
//
// Parameters:
//   - client: A pointer to the APIClient instance.
//
// Returns:
//   - A pointer to the newly created CustomComponentService instance.
func NewCustomComponentsService(client *APIClient) *CustomComponentService {
	return &CustomComponentService{baseService: NewBaseService(client)}
}

// GetAllCustomComponents requests all custom components.
//
// Returns:
//   - A pointer to models.ResourceNameResponse containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *CustomComponentService) GetAllCustomComponents() (*[]models.CustomComponent, error) {
	var response []models.CustomComponent
	err := s.baseService.DoGet("GetAllCustomComponents", nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// RequestName requests a new resource name based on the provided request data.
//
// Parameters:
//   - request: An instance of models.ResourceNameRequest containing the request data.
//
// Returns:
//   - A pointer to models.ResourceNameResponse containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *CustomComponentService) GetCustomComponent(id string) (*models.CustomComponent, error) {
	var response models.CustomComponent
	err := s.baseService.DoGet("GetCustomComponent", map[string]string{"id": id}, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// RequestName requests a new resource name based on the provided request data.
//
// Parameters:
//   - request: An instance of models.ResourceNameRequest containing the request data.
//
// Returns:
//   - A pointer to models.ResourceNameResponse containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *CustomComponentService) GetCustomComponentByParentId(parentComponentId string) (*models.CustomComponent, error) {
	var response models.CustomComponent
	err := s.baseService.DoGet("GetCustomComponentByParentId", map[string]string{"parentComponentId": parentComponentId}, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// RequestName requests a new resource name based on the provided request data.
//
// Parameters:
//   - request: An instance of models.ResourceNameRequest containing the request data.
//
// Returns:
//   - A pointer to models.ResourceNameResponse containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *CustomComponentService) GetCustomComponentByParentType(parenttype string) (*models.CustomComponent, error) {
	var response models.CustomComponent
	err := s.baseService.DoGet("GetCustomComponentByParentType", map[string]string{"parenttype": parenttype}, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// RequestName requests a new resource name based on the provided request data.
//
// Parameters:
//   - request: An instance of models.ResourceNameRequest containing the request data.
//
// Returns:
//   - A pointer to models.ResourceNameResponse containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *CustomComponentService) CreateOrUpdateCustomComponent(request models.CustomComponent) (*models.CustomComponent, error) {
	var response models.CustomComponent
	err := s.baseService.DoPost("CreateOrUpdateResourceComponent", request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// RequestName requests a new resource name based on the provided request data.
//
// Parameters:
//   - request: An instance of models.ResourceNameRequest containing the request data.
//
// Returns:
//   - A pointer to models.ResourceNameResponse containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *CustomComponentService) DeleteCustomComponent(id string) (interface{}, error) {
	var response models.ResourceOrganization
	return s.baseService.DoDelete("DeleteCustomComponent", map[string]string{"id": id}, &response)
}

// RequestName requests a new resource name based on the provided request data.
//
// Parameters:
//   - request: An instance of models.ResourceNameRequest containing the request data.
//
// Returns:
//   - A pointer to models.ResourceNameResponse containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *CustomComponentService) DeleteCustomComponentByParentId(parentcomponentid string) (interface{}, error) {
	var response models.ResourceOrganization
	return s.baseService.DoDelete("DeleteCustomComponentByParentId", map[string]string{"parentcomponentid": parentcomponentid}, &response)
}
