package openai

import (
	"fmt"
	"strings"
)

func checkSupportsModel(endpoint, model string) bool {
	return !disabledModels[endpoint][model]
}

func checkPromptType(prompt any) bool {
	_, isString := prompt.(string)
	_, isStringSlice := prompt.([]string)
	return isString || isStringSlice
}

func generateOpenAIUrl(suffix string, c OpenAIClient) string {
	return fmt.Sprintf("%s%s", c.config.BaseUrl, suffix)
}

func generateAzureUrl(suffix string, c AzureClient) string {
	baseURL := c.config.BaseUrl
	baseURL = strings.TrimRight(baseURL, "/")
	return fmt.Sprintf("%s/%s/%s/%s%s?api-version=%s",
		baseURL, azureAPIPrefix, azureDeploymentsPrefix, c.config.Engine, suffix, c.config.APIVersion)
}
