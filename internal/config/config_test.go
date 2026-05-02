package config

import (
	"strings"
	"testing"
)

func TestLoadFromEnv(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	t.Setenv(envAPIKey, "test-key")
	t.Setenv(envAPIToken, "test-token")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if cfg.APIKey != "test-key" {
		t.Errorf("APIKey = %q", cfg.APIKey)
	}
	if cfg.APIToken != "test-token" {
		t.Errorf("APIToken = %q", cfg.APIToken)
	}
	if cfg.BaseURL != defaultURL {
		t.Errorf("BaseURL = %q, want %q", cfg.BaseURL, defaultURL)
	}
}

func TestLoadMissingCreds(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	t.Setenv(envAPIKey, "")
	t.Setenv(envAPIToken, "")

	_, err := Load()
	if err == nil {
		t.Fatal("want error, got nil")
	}
	if !strings.Contains(err.Error(), envAPIKey) || !strings.Contains(err.Error(), envAPIToken) {
		t.Errorf("error missing env names: %v", err)
	}
}

func TestLoadKeyOnlyStillFails(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	t.Setenv(envAPIKey, "test-key")
	t.Setenv(envAPIToken, "")

	if _, err := Load(); err == nil {
		t.Fatal("want error when token missing")
	}
}
