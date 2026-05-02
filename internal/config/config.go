package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	APIKey   string
	APIToken string
	BaseURL  string
}

const (
	envAPIKey   = "TRELLO_API_KEY"
	envAPIToken = "TRELLO_TOKEN"
	defaultURL  = "https://api.trello.com/1"
)

func Load() (*Config, error) {
	v := viper.New()
	v.SetEnvPrefix("TRELLO")
	v.AutomaticEnv()
	v.BindEnv("api_key", envAPIKey)
	v.BindEnv("token", envAPIToken)
	v.SetDefault("base_url", defaultURL)

	home, err := os.UserHomeDir()
	if err == nil {
		cfgDir := filepath.Join(home, ".trello-cli")
		v.SetConfigName("config")
		v.SetConfigType("yaml")
		v.AddConfigPath(cfgDir)
		_ = v.ReadInConfig()
	}

	cfg := &Config{
		APIKey:   v.GetString("api_key"),
		APIToken: v.GetString("token"),
		BaseURL:  v.GetString("base_url"),
	}
	if cfg.APIKey == "" || cfg.APIToken == "" {
		return nil, fmt.Errorf("missing credentials: set %s and %s env vars or write ~/.trello-cli/config.yaml (api_key, token)", envAPIKey, envAPIToken)
	}
	return cfg, nil
}
