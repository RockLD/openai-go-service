package openai

import "net/http"

// 客户端配置
type ClientConfig struct {
	authToken  string
	BaseUrl    string
	ApiType    APIType
	Engine     string
	HTTPClient *http.Client
	APIVersion string
}

// 默认openai client配置
func DefaultOpenAIConfig(authToken string) ClientConfig {
	return ClientConfig{
		authToken:  authToken,
		BaseUrl:    openaiAPIURLv1,
		ApiType:    APITypeOpenAI,
		HTTPClient: &http.Client{},
	}
}

// 默认Azure配置
func DefaultAzureConfig() {

}
