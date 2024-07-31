package aznamingtool

import "github.com/rafaelherik/aznamingtool-apiclient/aznamingtool/models"

type ResourceProjectService struct {
	baseService *BaseService
}

func NewResourceProjectService(client *APIClient) *ResourceProjectService {
	return &ResourceProjectService{baseService: NewBaseService(client)}
}

func (s *ResourceProjectService) GetAllResourceProjects() (*[]models.ResourceProject, error) {
	var response []models.ResourceProject
	err := s.baseService.DoGet("RequestName", nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *ResourceProjectService) GetResourceProject(id string) (*models.ResourceProject, error) {
	var response models.ResourceProject
	err := s.baseService.DoGet("GetResourceProject", map[string]string{"id": id}, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *ResourceProjectService) CreateOrUpdateResourceProject(request models.ResourceProject) (*models.ResourceProject, error) {
	var response models.ResourceProject
	err := s.baseService.DoPost("CreateOrUpdateResourceProject", request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *ResourceProjectService) DeleteResourceProject(id string) (interface{}, error) {
	var response models.ResourceProject
	return s.baseService.DoDelete("DeleteResourceProject", map[string]string{"id": id}, &response)
}
