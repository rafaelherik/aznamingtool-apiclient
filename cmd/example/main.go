package tests

import (
	"fmt"
	"testing"

	"github.com/rafaelherik/aznamingtool-apiclient/pkg/aznamingtool"
	"github.com/rafaelherik/aznamingtool-apiclient/pkg/aznamingtool/models"
)

func TestResourceNamingService_RequestName(t *testing.T) {

	restClient := aznamingtool.NewAPIClient("http://localhost:8081", "603a01da-b170-4a0f-8d55-f809332faacd")
	service := aznamingtool.NewResourceNamingService(restClient)

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
