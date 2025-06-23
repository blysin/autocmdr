// Package config provides functionality for loading, saving, and validating application configuration.
// It uses viper for configuration management and supports both file-based and environment variable configuration.
package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config holds the application configuration
type Config struct {
	Model     string `mapstructure:"model" json:"model"`
	ServerURL string `mapstructure:"server_url" json:"server_url"`
	Token     string `mapstructure:"token" json:"token"`
	LogLevel  string `mapstructure:"log_level" json:"log_level"`
	ConfigDir string `mapstructure:"config_dir" json:"config_dir"`
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logrus.WithError(err).Warn("Failed to get user home directory, using current directory")
		homeDir = "."
	}

	return &Config{
		Model:     "qwen3:14b",
		ServerURL: "http://localhost:11434",
		Token:     "",
		LogLevel:  "info",
		ConfigDir: filepath.Join(homeDir, ".autocmdr"),
	}
}

// Load loads configuration from file and environment variables
func Load() (*Config, error) {
	cfg := DefaultConfig()

	// Set up viper
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(cfg.ConfigDir)
	viper.AddConfigPath(".")

	// Set environment variable prefix
	viper.SetEnvPrefix("LANGCHAIN_CHAT")
	viper.AutomaticEnv()

	// Set default values
	viper.SetDefault("model", cfg.Model)
	viper.SetDefault("server_url", cfg.ServerURL)
	viper.SetDefault("token", cfg.Token)
	viper.SetDefault("log_level", cfg.LogLevel)
	viper.SetDefault("config_dir", cfg.ConfigDir)

	// Try to read config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
		logrus.Debug("Config file not found, using defaults and environment variables")
	} else {
		logrus.WithField("file", viper.ConfigFileUsed()).Debug("Config file loaded")
	}

	// Unmarshal config
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Ensure config directory exists
	if err := os.MkdirAll(cfg.ConfigDir, 0750); err != nil {
		return nil, fmt.Errorf("failed to create config directory: %w", err)
	}

	return cfg, nil
}

// Save saves the configuration to file
func (c *Config) Save() error {
	configPath := filepath.Join(c.ConfigDir, "config.json")

	// Ensure config directory exists
	if err := os.MkdirAll(c.ConfigDir, 0750); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	// Set viper values
	viper.Set("model", c.Model)
	viper.Set("server_url", c.ServerURL)
	viper.Set("token", c.Token)
	viper.Set("log_level", c.LogLevel)
	viper.Set("config_dir", c.ConfigDir)

	// Write config file
	if err := viper.WriteConfigAs(configPath); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	logrus.WithField("path", configPath).Info("Configuration saved")
	return nil
}

// Validate validates the configuration
func (c *Config) Validate() error {
	if c.Model == "" {
		return fmt.Errorf("model cannot be empty")
	}
	if c.ServerURL == "" {
		return fmt.Errorf("server_url cannot be empty")
	}
	return nil
}

// GetConfigPath returns the path to the config file
func (c *Config) GetConfigPath() string {
	return filepath.Join(c.ConfigDir, "config.json")
}
