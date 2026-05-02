package auto

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/mrfansi/trello-cli/internal/client"
	"github.com/mrfansi/trello-cli/internal/cmdutil"
	"github.com/mrfansi/trello-cli/internal/config"
	"github.com/spf13/cobra"
)

// RunOp is exported so handcrafted aliases (e.g., `me`) can dispatch
// through the same path as generated commands.
func RunOp(cmd *cobra.Command, method, path string, pathVars, queryVars []string, data string) error {
	return execRaw(cmd, method, path, pathVars, queryVars, data)
}

// execRaw is the runtime entry point for every generated command.
// It substitutes path placeholders, appends query params, sends the
// request via the same auth path as the curated raw command, and
// streams the response body to stdout. Non-2xx responses print to
// stderr and return an error.
func execRaw(cmd *cobra.Command, method, path string, pathVars, queryVars []string, data string) error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	for _, kv := range pathVars {
		k, v, err := splitKV(kv, "--path")
		if err != nil {
			return err
		}
		ph := "{" + k + "}"
		if !strings.Contains(path, ph) {
			return fmt.Errorf("path placeholder %s not found in %q", ph, path)
		}
		path = strings.ReplaceAll(path, ph, url.PathEscape(v))
	}

	body, err := loadBody(data)
	if err != nil {
		return err
	}

	httpClient, baseURL, auth := client.Raw(cfg)
	fullURL := strings.TrimRight(baseURL, "/") + "/" + strings.TrimLeft(path, "/")

	req, err := buildRequest(fullURL, method, body, queryVars)
	if err != nil {
		return err
	}

	ctx, cancel := cmdutil.Context()
	defer cancel()
	req = req.WithContext(ctx)

	if err := auth(ctx, req); err != nil {
		return fmt.Errorf("auth: %w", err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 300 {
		fmt.Fprintf(os.Stderr, "trello api %s: %s\n", resp.Status, string(respBody))
		return fmt.Errorf("non-2xx response")
	}
	fmt.Fprintln(cmd.OutOrStdout(), string(respBody))
	return nil
}

func splitKV(s, label string) (string, string, error) {
	i := strings.Index(s, "=")
	if i <= 0 {
		return "", "", fmt.Errorf("%s expects key=value, got %q", label, s)
	}
	return s[:i], s[i+1:], nil
}

func loadBody(data string) ([]byte, error) {
	if data == "" {
		return nil, nil
	}
	if strings.HasPrefix(data, "@") {
		return os.ReadFile(data[1:])
	}
	return []byte(data), nil
}

func buildRequest(rawURL, method string, body []byte, queryVars []string) (*http.Request, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("parse url: %w", err)
	}
	q := u.Query()
	for _, kv := range queryVars {
		k, v, err := splitKV(kv, "--query")
		if err != nil {
			return nil, err
		}
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()

	var reader io.Reader
	if body != nil {
		reader = bytes.NewReader(body)
	}
	req, err := http.NewRequestWithContext(context.Background(), method, u.String(), reader)
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, nil
}
