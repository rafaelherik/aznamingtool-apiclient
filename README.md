# Azure Naming tool API client using go

[![Build Status](https://github.com/rafaelherik/aznamingtool-apiclient/actions/workflows/build.yaml/badge.svg)](https://github.com/rafaelherik/aznamingtool-apiclient/actions/workflows/build.yaml)

This repository provides a simplified API client in Go for interacting with the AzureNamingTool API. The AzureNamingTool Go Client enables developers to easily integrate Azure naming services into their Go applications with minimal setup and effort.

### Features
- Simplified API interactions with AzureNamingTool
- Request name generation and validation
- Easy integration with Go projects
- Error handling and response parsing

## Setup Environment


## Usage

```go
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


```