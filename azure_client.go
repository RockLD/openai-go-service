package openai

import (
	"net/http"
)

// OpenAIClient is Azure API client.
type AzureClient struct {
	config         ClientConfig
	requestBuilder requestBuilder
}

func NewAzureClient() {

}

func (c *AzureClient) sendRequest(req *http.Request, v interface{}) error {
	// Azure API Key authentication
	req.Header.Set(AzureAPIKeyHeader, c.config.authToken)
}
