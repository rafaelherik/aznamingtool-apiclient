package config

import (
	"os"
)

type Config struct {
	BaseURL      string
	APIKey       string
	ApiEndpoints map[string]string
}

func LoadConfig(baseUrl *string, apiKey *string) *Config {
	var resolvedBaseURL string
	if baseUrl != nil {
		resolvedBaseURL = *baseUrl
	} else {
		resolvedBaseURL = os.Getenv("AZ_NAMING_TOOL_API_BASE_URL")
	}

	var resolvedApiKey string
	if apiKey != nil {
		resolvedApiKey = *apiKey
	} else {
		resolvedApiKey = os.Getenv("AZ_NAMING_TOOL_API_KEY")
	}

	return &Config{
		BaseURL: resolvedBaseURL,
		APIKey:  resolvedApiKey,
		ApiEndpoints: map[string]string{
			"RequestName":               resolvedBaseURL + "/api/ResourceNamingRequests/RequestName",
			"RequestNameWithComponents": resolvedBaseURL + "/api/ResourceNamingRequests/RequestNameWithComponents",
			"ValidateName":              resolvedBaseURL + "/api/ResourceNamingRequests/ValidateName",
		},
	}
}
