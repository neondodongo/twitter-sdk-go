package twitter

import (
	"fmt"
	"net/http"
	"strings"
)

// ControllerV2 exposes methods for interfacing with the Twitter API v2 web service layer
type ControllerV2 interface {
	GetTweetByID(tweetID string, queryParams map[string]string) (GetTweetByIDResponse, error)
	GetTweets(tweetIDs []string, queryParams map[string]string) (GetTweetsResponse, error)
	GetUserByUsername(username string, queryParams map[string]string) (GetUserByUsernameResponse, error)
	SearchTweetsRecent(queryParams map[string]string) (SearchTweetsResponse, error)
}

// operator implements the Twitter API v2 client contract via HTTP client
type operatorV2 struct {
	cfgV2 ControllerConfigV2
	http  *http.Client
}

// ControllerConfigV2 the Twitter API v2 client's available HTTP connection, security and third-party options
type ControllerConfigV2 struct {
	APIKey       string `json:"api_key" yaml:"api_key"`
	APISecretKey string `json:"api_secret_key" yaml:"api_secret_key"`
	BaseURL      string `json:"base_url" yaml:"base_url"`
	BearerToken  string `json:"bearer_token" yaml:"bearer_token"`
	// TODO: Additional config options i.e. HTTP client options, Monitoring enabled, ...
}

// NewControllerV2 construct a new Twitter API v2 client operator
func NewControllerV2(cfg ControllerConfigV2) (ControllerV2, error) {
	if err := sanitizeAndValidateConfigV2(&cfg); err != nil {
		return nil, fmt.Errorf("failed to sanitize and validate API client v2 configuration; %w", err)
	}

	return &operatorV2{
		cfgV2: cfg,
		http:  http.DefaultClient, // TODO: create an HTTP client with custom config options
	}, nil
}

// check for empty, whitespace-only or zero configuration values and apply default settings when relevant
func sanitizeAndValidateConfigV2(cfg *ControllerConfigV2) error {
	cfg.APIKey = strings.TrimSpace(cfg.APIKey)
	if cfg.APIKey == "" {
		return fmt.Errorf("cannot construct twitter API client v2 with an empty or whitespace-only API key")
	}

	cfg.APISecretKey = strings.TrimSpace(cfg.APISecretKey)
	if cfg.APISecretKey == "" {
		return fmt.Errorf("cannot construct twitter API client v2 with an empty or whitespace-only API secret key")
	}

	cfg.BaseURL = strings.TrimSpace(cfg.BaseURL)
	if cfg.BaseURL == "" {
		cfg.BaseURL = defaultBaseURL
	}

	cfg.BearerToken = strings.TrimSpace(cfg.BearerToken)
	if cfg.BearerToken == "" {
		return fmt.Errorf("cannot construct twitter API client v2 with an empty or whitespace-only oauth 2.0 bearer token")
	}

	return nil
}
