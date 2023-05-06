package openai

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// OpenAIClient is OpenAI API client.
type OpenAIClient struct {
	config         ClientConfig
	requestBuilder requestBuilder
}

var ctx context.Context

func NewOpenAIClient(authToken string) *OpenAIClient {
	config := DefaultOpenAIConfig(authToken)
	return NewOpenAIClientWithConfig(config)
}

func NewOpenAIClientWithConfig(config ClientConfig) *OpenAIClient {
	return &OpenAIClient{
		config: config,
	}
}

func (c *OpenAIClient) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Accept", "application/json; charset=utf-8")
	// OpenAI authentication
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.config.authToken))

	// Check whether Content-Type is already set, Upload Files API requires
	// Content-Type == multipart/form-data
	contentType := req.Header.Get("Content-Type")
	if contentType == "" {
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	}

	// if len(c.config.OrgID) > 0 {
	// 	req.Header.Set("OpenAI-Organization", c.config.OrgID)
	// }

	res, err := c.config.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes ErrorResponse
		err = json.NewDecoder(res.Body).Decode(&errRes)
		if err != nil || errRes.Error == nil {
			reqErr := RequestError{
				StatusCode: res.StatusCode,
				Err:        err,
			}
			return fmt.Errorf("error, %w", &reqErr)
		}
		errRes.Error.StatusCode = res.StatusCode
		return fmt.Errorf("error, status code: %d, message: %w", res.StatusCode, errRes.Error)
	}

	if v != nil {
		if err = json.NewDecoder(res.Body).Decode(v); err != nil {
			return err
		}
	}

	return nil
}
