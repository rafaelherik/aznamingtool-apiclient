package main

import (
	"fmt"
	"log"

	"github.com/rafaelherik/aznamingtool-apiclient/pkg/aznamingtool"
	"github.com/rafaelherik/aznamingtool-apiclient/pkg/aznamingtool/models"
)

func main() {

	restClient := aznamingtool.NewAPIClient("http://localhost:8081", "603a01da-b170-4a0f-8d55-f809332faacd")
	resourceNamingService := aznamingtool.NewResourceNamingService(restClient)
	request := models.ResourceNameRequest{
		ResourceType:        "vm",
		ResourceId:          85,
		ResourceProjAppSvc:  "tnp",
		ResourceEnvironment: "dev",
		ResourceFunction:    "func",
		ResourceLocation:    "weu",
		ResourceInstance:    "1",
	}
	response, err := resourceNamingService.RequestName(request)
	if err != nil {
		log.Fatalf("Error requesting name: %v", err)
	}

	fmt.Printf("Resource Name: %v\n", response.ResourceName)

}
