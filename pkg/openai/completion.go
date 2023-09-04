package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const (
	completionsTemplate = "%s/engines/davinci/completions"
)

type CompletionsClient struct {
	baseURL string
	token   string
}

type CompletionsClientOption func(*CompletionsClient)

func WithBaseURL(baseURL string) CompletionsClientOption {
	return func(c *CompletionsClient) {
		if baseURL != "" {
			c.baseURL = baseURL
		}
	}
}

func WithToken(token string) CompletionsClientOption {
	return func(c *CompletionsClient) {
		if token != "" {
			c.token = token
		}
	}
}

func NewCompletionsClient(opts ...CompletionsClientOption) *CompletionsClient {
	client := &CompletionsClient{
		baseURL: "https://api.openai.com/v1",
		token:   os.Getenv("OPENAI_TOKEN"),
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

func (c *CompletionsClient) OpenAICompletionsURL() string {
	return fmt.Sprintf(completionsTemplate, c.baseURL)
}

func (c *CompletionsClient) Completions(ctx context.Context, prompt string) (string, error) {
	payload := map[string]interface{}{
		"prompt":     prompt,
		"max_tokens": 150,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.OpenAICompletionsURL(), bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var response struct {
		Choices []struct {
			Text string `json:"text"`
		} `json:"choices"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	return response.Choices[0].Text, nil
}
