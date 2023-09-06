package huggingface

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

var (
	repoTypes = map[string]bool{
		"model":   true,
		"dataset": true,
		"space":   true,
	}
	repoTypesURLPrefixes = map[string]string{
		"dataset": "datasets/",
		"space":   "spaces/",
	}
)

type HFHubClient struct {
	defaultRevision string
	baseURL         string
	token           string
}

type HFHubClientOption func(*HFHubClient)

func WithDefaultRevision(defaultRevision string) HFHubClientOption {
	return func(c *HFHubClient) {
		c.defaultRevision = defaultRevision
	}
}

func WithBaseURL(baseURL string) HFHubClientOption {
	return func(c *HFHubClient) {
		if baseURL != "" {
			c.baseURL = baseURL
		}
	}
}

func WithToken(token string) HFHubClientOption {
	return func(c *HFHubClient) {
		if token != "" {
			c.token = token
		}
	}
}

func NewHFHubClient(opts ...HFHubClientOption) *HFHubClient {
	client := &HFHubClient{
		defaultRevision: defaultRevision,
		baseURL:         HuggingfaceURL,
		token:           os.Getenv("HF_TOKEN"),
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

func (c *HFHubClient) hfHubURL(repoID string, filename string, subfolder string, repoType string, revision string) (string, error) {
	if subfolder == "" {
		subfolder = ""
	}
	if subfolder != "" {
		filename = filepath.Join(subfolder, filename)
	}

	if !repoTypes[repoType] {
		return "", errors.New("invalid repo type")
	}

	if prefix, ok := repoTypesURLPrefixes[repoType]; ok {
		repoID = fmt.Sprintf("%s%s", prefix, repoID)
	}

	if revision == "" {
		revision = c.defaultRevision
	}

	u, err := url.Parse(fmt.Sprintf(huggingfaceDownloadFileTemplate, c.baseURL, repoID, revision, filename))
	if err != nil {
		return "", err
	}

	if c.baseURL != "" && strings.HasPrefix(u.String(), huggingfaceDownloadFileTemplate) {
		u.Scheme = "https"
		u.Host = c.baseURL
	}

	return u.String(), nil
}

func (c *HFHubClient) GetModel(ctx context.Context, repoID string, revision string) (io.ReadCloser, error) {
	return c.StreamFile(ctx, repoID, "pytorch_model.bin", "", "model", revision)
}

func (c *HFHubClient) StreamFile(
	ctx context.Context,
	repoID string,
	filename string,
	subfolder string,
	repoType string,
	revision string,
) (io.ReadCloser, error) {
	url, err := c.hfHubURL(repoID, filename, subfolder, repoType, revision)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func (c *HFHubClient) DownloadFile(ctx context.Context,
	repoID string,
	filename string,
	subfolder string,
	repoType string,
	revision string,
	outputFilePath string,
) error {
	file, err := c.StreamFile(ctx, repoID, filename, subfolder, repoType, revision)
	if err != nil {
		return err
	}
	defer file.Close()

	out, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}

	return nil
}
