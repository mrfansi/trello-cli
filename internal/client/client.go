package client

import (
	"context"
	"net/http"
	"time"

	"github.com/mrfansi/trecli/internal/config"
	"github.com/mrfansi/trecli/internal/trello"
)

func New(cfg *config.Config) (*trello.Client, error) {
	httpClient := &http.Client{Timeout: 30 * time.Second}
	return trello.NewClient(cfg.BaseURL,
		trello.WithHTTPClient(httpClient),
		trello.WithRequestEditorFn(AuthEditor(cfg)),
	)
}

// AuthEditor returns a request editor that injects key+token query params.
func AuthEditor(cfg *config.Config) func(context.Context, *http.Request) error {
	return func(_ context.Context, req *http.Request) error {
		q := req.URL.Query()
		q.Set("key", cfg.APIKey)
		q.Set("token", cfg.APIToken)
		req.URL.RawQuery = q.Encode()
		return nil
	}
}

// Raw creates a base HTTP client + base URL + auth editor for ad-hoc requests.
func Raw(cfg *config.Config) (*http.Client, string, func(context.Context, *http.Request) error) {
	return &http.Client{Timeout: 30 * time.Second}, cfg.BaseURL, AuthEditor(cfg)
}
