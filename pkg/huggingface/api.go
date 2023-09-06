package huggingface

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	getModelInfoTemplate = "%s/api/models/%s"
)

type HFAPIClient struct {
	BaseURL string
	Token   string
}

type GetModelInfoOptions struct {
	RepoID   string
	Revision string
}

type HFAPIClientOption func(*HFAPIClient)

func WithAPIToken(token string) HFAPIClientOption {
	return func(c *HFAPIClient) {
		c.Token = token
	}
}

func WithAPIBaseURL(baseURL string) HFAPIClientOption {
	return func(c *HFAPIClient) {
		if baseURL != "" {
			c.BaseURL = baseURL
		}
	}
}

func NewHFAPIClient(opts ...HFAPIClientOption) *HFAPIClient {
	client := &HFAPIClient{
		BaseURL: HuggingfaceURL,
	}
	for _, opt := range opts {
		opt(client)
	}
	return client
}

func (c *HFAPIClient) GetModelInfo(ctx context.Context, opts *GetModelInfoOptions) (*ModelInfo, error) {
	path := fmt.Sprintf(getModelInfoTemplate, c.BaseURL, opts.RepoID)
	if opts.Revision != "" {
		path = fmt.Sprintf("%s/revision/%s", path, url.QueryEscape(opts.Revision))
	}

	req, err := http.NewRequestWithContext(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}
	if c.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.Token)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get model info: %s", resp.Status)
	}

	var modelInfo ModelInfo
	err = json.NewDecoder(resp.Body).Decode(&modelInfo)
	if err != nil {
		return nil, err
	}

	return &modelInfo, nil
}
