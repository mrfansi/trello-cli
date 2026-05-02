package cmdutil

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mrfansi/trello-cli/internal/client"
	"github.com/mrfansi/trello-cli/internal/config"
	"github.com/mrfansi/trello-cli/internal/output"
	"github.com/mrfansi/trello-cli/internal/trello"
)

type Ctx struct {
	Client   *trello.Client
	Renderer *output.Renderer
}

func Build(jsonMode bool) (*Ctx, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}
	c, err := client.New(cfg)
	if err != nil {
		return nil, fmt.Errorf("init client: %w", err)
	}
	return &Ctx{Client: c, Renderer: output.New(jsonMode)}, nil
}

func Context() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	sigCtx, sigCancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	return sigCtx, func() { sigCancel(); cancel() }
}

// Decode reads HTTP response into v, returning error on non-2xx.
func Decode(resp *http.Response, v any) error {
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("trello api %s: %s", resp.Status, string(body))
	}
	if v == nil {
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(v)
}

