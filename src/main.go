package main

import (
	"azure-naming-api-client/client"
	"azure-naming-api-client/config"
	"azure-naming-api-client/models"
	"fmt"
	"os"
)

func main() {
	os.Setenv("AZ_NAMING_TOOL_API_BASE_URL", "http://localhost:8081")
	os.Setenv("AZ_NAMING_TOOL_API_KEY", "603a01da-b170-4a0f-8d55-f809332faacd")
	cfg := config.LoadConfig(nil, nil)
	apiClient := client.NewAPIClient(cfg.BaseURL, cfg.APIKey)

	fmt.Println(apiClient.APIKey)
	fmt.Println(apiClient.BaseURL)

	restClient := client.NewAPIClient(cfg.BaseURL, cfg.APIKey)
	service := client.NewResourceNamingService(restClient)

	request := models.ResourceNameRequest{
		ResourceType:        "vm",
		ResourceId:          85,
		ResourceProjAppSvc:  "tnp",
		ResourceEnvironment: "dev",
		ResourceFunction:    "func",
		ResourceLocation:    "brs",
		ResourceInstance:    "1",
	}
	response, err := service.RequestName(request)
	if err != nil {
		fmt.Println(err)

	}

	fmt.Println(response.Message)

}
