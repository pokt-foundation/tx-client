package txclient

import (
	"errors"
	"time"
)

// Config struct to provide options
type Config struct {
	BaseURL, APIKey string
	Version         APIVersion
	Retries         int
	Timeout         time.Duration
}

var (
	ErrNoBaseURL      = errors.New("base url not provided")
	ErrNoAPIKey       = errors.New("api key not provided")
	ErrInvalidVersion = errors.New("invalid version provided")
)

func (c Config) validate() error {
	if c.BaseURL == "" {
		return ErrNoBaseURL
	}
	if c.APIKey == "" {
		return ErrNoAPIKey
	}

	if !validAPIVersions[c.Version] {
		return ErrInvalidVersion
	}

	return nil
}
