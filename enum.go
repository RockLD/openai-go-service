package openai

import "errors"

// API类型，当下有OPENAI和AZURE
type APIType string

const (
	APITypeOpenAI          APIType = "OPEN_AI"
	APITypeAzure           APIType = "AZURE"
	openaiAPIURLv1                 = "https://api.openai.com/v1"
	azureAPIPrefix                 = "openai"
	azureDeploymentsPrefix         = "deployments"

	GPT432K0314             = "gpt-4-32k-0314"
	GPT432K                 = "gpt-4-32k"
	GPT40314                = "gpt-4-0314"
	GPT4                    = "gpt-4"
	GPT3Dot5Turbo0301       = "gpt-3.5-turbo-0301"
	GPT3Dot5Turbo           = "gpt-3.5-turbo"
	GPT3TextDavinci003      = "text-davinci-003"
	GPT3TextDavinci002      = "text-davinci-002"
	GPT3TextCurie001        = "text-curie-001"
	GPT3TextBabbage001      = "text-babbage-001"
	GPT3TextAda001          = "text-ada-001"
	GPT3TextDavinci001      = "text-davinci-001"
	GPT3DavinciInstructBeta = "davinci-instruct-beta"
	GPT3Davinci             = "davinci"
	GPT3CurieInstructBeta   = "curie-instruct-beta"
	GPT3Curie               = "curie"
	GPT3Ada                 = "ada"
	GPT3Babbage             = "babbage"
	CodexCodeDavinci002     = "code-davinci-002"
	CodexCodeCushman001     = "code-cushman-001"
	CodexCodeDavinci001     = "code-davinci-001"
)

var (
	ErrChatCompletionInvalidModel       = errors.New("this model is not supported with this method, please use CreateCompletion client method instead") //nolint:lll
	ErrChatCompletionStreamNotSupported = errors.New("streaming is not supported with this method, please use CreateChatCompletionStream")              //nolint:lll
)

var disabledModels = map[string]map[string]bool{
	"/completions": {
		GPT3Dot5Turbo:     true,
		GPT3Dot5Turbo0301: true,
		GPT4:              true,
		GPT40314:          true,
		GPT432K:           true,
		GPT432K0314:       true,
	},
	"/chat/completions": {
		CodexCodeDavinci002:     true,
		CodexCodeCushman001:     true,
		CodexCodeDavinci001:     true,
		GPT3TextDavinci003:      true,
		GPT3TextDavinci002:      true,
		GPT3TextCurie001:        true,
		GPT3TextBabbage001:      true,
		GPT3TextAda001:          true,
		GPT3TextDavinci001:      true,
		GPT3DavinciInstructBeta: true,
		GPT3Davinci:             true,
		GPT3CurieInstructBeta:   true,
		GPT3Curie:               true,
		GPT3Ada:                 true,
		GPT3Babbage:             true,
	},
}
