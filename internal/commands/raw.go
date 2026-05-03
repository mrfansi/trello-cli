package commands

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/mrfansi/trecli/internal/client"
	"github.com/mrfansi/trecli/internal/cmdutil"
	"github.com/mrfansi/trecli/internal/config"
	"github.com/spf13/cobra"
)

func rawCmd() *cobra.Command {
	var (
		pathVars    []string
		queryVars   []string
		headerVars  []string
		dataArg     string
		methodFlag  string
	)

	cmd := &cobra.Command{
		Use:   "raw <METHOD> <PATH>",
		Short: "Send a raw request to the Trello API (escape hatch for endpoints not yet wrapped)",
		Long: `Send any HTTP request to the Trello API with auth auto-injected.

Path templating: use {name} placeholders and supply --path name=value.
Query params: --query key=value (repeatable).
Headers: --header key=value (repeatable).
Body: --data '{"key":"v"}' or --data @file.json.

Examples:
  trecli raw GET /members/me
  trecli raw GET /boards/{id} --path id=abc --query fields=name,url
  trecli raw POST /cards --query idList=xyz --query name="New"
  trecli raw PUT /cards/{id} --path id=abc --data @update.json`,
		Args: cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			method, path, err := parseMethodPath(args, methodFlag)
			if err != nil {
				return err
			}

			cfg, err := config.Load()
			if err != nil {
				return err
			}

			path, err = applyPathVars(path, pathVars)
			if err != nil {
				return err
			}

			body, err := loadBody(dataArg)
			if err != nil {
				return err
			}

			httpClient, baseURL, auth := client.Raw(cfg)
			fullURL := strings.TrimRight(baseURL, "/") + "/" + strings.TrimLeft(path, "/")

			req, err := buildRequest(fullURL, method, body, queryVars, headerVars)
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
		},
	}

	cmd.Flags().StringArrayVar(&pathVars, "path", nil, "Path variable: key=value (repeatable)")
	cmd.Flags().StringArrayVar(&queryVars, "query", nil, "Query parameter: key=value (repeatable)")
	cmd.Flags().StringArrayVar(&headerVars, "header", nil, "Header: key=value (repeatable)")
	cmd.Flags().StringVar(&dataArg, "data", "", "Request body: literal JSON or @path/to/file.json")
	cmd.Flags().StringVarP(&methodFlag, "method", "X", "", "HTTP method (alternative to positional arg)")
	return cmd
}

func parseMethodPath(args []string, methodFlag string) (method, path string, err error) {
	switch len(args) {
	case 1:
		if methodFlag == "" {
			return "", "", fmt.Errorf("method required: pass as first arg or --method")
		}
		return strings.ToUpper(methodFlag), args[0], nil
	case 2:
		return strings.ToUpper(args[0]), args[1], nil
	default:
		return "", "", fmt.Errorf("usage: raw <METHOD> <PATH>")
	}
}

func applyPathVars(path string, vars []string) (string, error) {
	for _, kv := range vars {
		k, v, err := splitKV(kv, "--path")
		if err != nil {
			return "", err
		}
		placeholder := "{" + k + "}"
		if !strings.Contains(path, placeholder) {
			return "", fmt.Errorf("--path %s: placeholder %s not found in %q", k, placeholder, path)
		}
		path = strings.ReplaceAll(path, placeholder, url.PathEscape(v))
	}
	if i := strings.Index(path, "{"); i >= 0 {
		if j := strings.Index(path[i:], "}"); j >= 0 {
			return "", fmt.Errorf("unresolved path placeholder %s", path[i:i+j+1])
		}
	}
	return path, nil
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

func buildRequest(rawURL, method string, body []byte, queryVars, headerVars []string) (*http.Request, error) {
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

	var bodyReader io.Reader
	if body != nil {
		bodyReader = bytes.NewReader(body)
	}
	req, err := http.NewRequestWithContext(context.Background(), method, u.String(), bodyReader)
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	for _, kv := range headerVars {
		k, v, err := splitKV(kv, "--header")
		if err != nil {
			return nil, err
		}
		req.Header.Set(k, v)
	}
	return req, nil
}
