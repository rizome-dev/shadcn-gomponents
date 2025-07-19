// Package example provides a sample library structure for the tmpl project.
package example

import (
	"errors"
	"fmt"
)

// ErrInvalidInput is returned when the input is invalid.
var ErrInvalidInput = errors.New("invalid input")

// Config holds the configuration for the example.
type Config struct {
	Name    string
	Verbose bool
}

// DefaultConfig returns a Config with default values.
func DefaultConfig() *Config {
	return &Config{
		Name:    "default",
		Verbose: false,
	}
}

// Client represents an example client.
type Client struct {
	config *Config
}

// NewClient creates a new Client with the given configuration.
func NewClient(cfg *Config) (*Client, error) {
	if cfg == nil {
		return nil, ErrInvalidInput
	}

	if cfg.Name == "" {
		cfg.Name = "default"
	}

	return &Client{
		config: cfg,
	}, nil
}

// Process performs some example processing.
func (c *Client) Process(input string) (string, error) {
	if input == "" {
		return "", ErrInvalidInput
	}

	result := fmt.Sprintf("Processed by %s: %s", c.config.Name, input)

	if c.config.Verbose {
		fmt.Printf("[DEBUG] %s\n", result)
	}

	return result, nil
}
