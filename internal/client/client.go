package client

import (
	"context"
	"net/http"
	"time"

	"github.com/mrfansi/trello-cli/internal/config"
	"github.com/mrfansi/trello-cli/internal/trello"
)

func New(cfg *config.Config) (*trello.Client, error) {
	httpClient := &http.Client{Timeout: 30 * time.Second}
	auth := func(_ context.Context, req *http.Request) error {
		q := req.URL.Query()
		q.Set("key", cfg.APIKey)
		q.Set("token", cfg.APIToken)
		req.URL.RawQuery = q.Encode()
		return nil
	}
	return trello.NewClient(cfg.BaseURL,
		trello.WithHTTPClient(httpClient),
		trello.WithRequestEditorFn(auth),
	)
}
