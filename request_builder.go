package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type requestBuilder struct {
}

func (b *requestBuilder) Build(ctx context.Context, method, url string, body any) (*http.Request, error) {
	if body == nil {
		return http.NewRequestWithContext(ctx, method, url, nil)
	}
	var reqBody []byte
	var err error
	reqBody, err = json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return http.NewRequestWithContext(
		ctx,
		method,
		url,
		bytes.NewBuffer(reqBody),
	)
}
