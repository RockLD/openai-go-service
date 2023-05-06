package openai

import (
	"context"
	"net/http"
)

func (c *OpenAIClient) CreateChatCompletion(ctx context.Context, request ChatCompletionRequest) (response ChatCompletionResponse, err error) {
	if request.Stream {
		err = ErrChatCompletionStreamNotSupported
		return
	}
	urlSuffix := "/chat/completions"
	if !checkSupportsModel(urlSuffix, request.Model) {
		err = ErrChatCompletionInvalidModel
		return
	}
	req, err := c.requestBuilder.Build(ctx, http.MethodPost, c.generateUrl(urlSuffix), request)

}
