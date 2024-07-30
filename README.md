# Azure Naming tool API client using go

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

	"github.com/rafaelherik/aznamingtool-apiclient/models"
	"github.com/rafaelherik/aznamingtool-apiclient/client"
)

func main() {
	// Create a new client
	apiClient := client.NewApiClient("http://your-naming-tool-url","your-api-key")

	// Initialize the service
	resourceNamingService := client.NewResourceNamingService(apiClient)

	// Create a request
	request := models.ResourceNameRequest{
		Name: "example-resource",
		// Add other required fields
	}

	// Request a name
	response, err := resourceNamingService.RequestName(request)
	if err != nil {
		log.Fatalf("Error requesting name: %v", err)
	}

	fmt.Printf("Resource Name: %v\n", response.Name)
}

```