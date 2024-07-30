package tests

import (
	"azure-naming-api-client/client"
	"azure-naming-api-client/config"
	"azure-naming-api-client/models"
	"fmt"
	"os"
	"testing"
)

func TestResourceNamingService_RequestName(t *testing.T) {

	// ARRANGE
	os.Setenv("AZ_NAMING_TOOL_API_BASE_URL", "http://localhost:8081")
	os.Setenv("AZ_NAMING_TOOL_API_KEY", "603a01da-b170-4a0f-8d55-f809332faacd")
	cfg := config.LoadConfig(nil, nil)
	restClient := client.NewAPIClient(cfg.BaseURL, cfg.APIKey)
	service := client.NewResourceNamingService(restClient)

	// ACT
	t.Run("RequestName Success", func(t *testing.T) {
		request := models.ResourceNameRequest{
			ResourceType:        "vm",
			ResourceId:          85,
			ResourceProjAppSvc:  "tnp",
			ResourceEnvironment: "dev",
			ResourceFunction:    "func",
			ResourceLocation:    "weu",
			ResourceInstance:    "1",
		}
		response, err := service.RequestName(request)
		if err != nil {
			fmt.Println(err)
			t.Fatalf("expected no error, got %v", err)
		}

		// ASSERT
		if response.Success != true {
			t.Errorf("The response was not successful")
		}
	})

	t.Run("RequestName Failure", func(t *testing.T) {

		request := models.ResourceNameRequest{
			ResourceType:        "ERROR",
			ResourceId:          85,
			ResourceProjAppSvc:  "tnp",
			ResourceEnvironment: "dev",
			ResourceFunction:    "func",
		}
		response, err := service.RequestName(request)
		if err == nil {
			t.Fatalf("expected error, got none")
		}

		if response.Success {
			t.Errorf("expected success to be false, got true")
		}
	})
}
