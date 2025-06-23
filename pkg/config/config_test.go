package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()

	assert.NotEmpty(t, cfg.Model)
	assert.NotEmpty(t, cfg.ServerURL)
	assert.Equal(t, "info", cfg.LogLevel)
	assert.Contains(t, cfg.ConfigDir, ".autocmdr")
}

func TestConfigValidate(t *testing.T) {
	tests := []struct {
		name    string
		config  *Config
		wantErr bool
	}{
		{
			name:    "valid config",
			config:  DefaultConfig(),
			wantErr: false,
		},
		{
			name: "empty model",
			config: &Config{
				Model:     "",
				ServerURL: "http://localhost:11434",
			},
			wantErr: true,
		},
		{
			name: "empty server URL",
			config: &Config{
				Model:     "test-model",
				ServerURL: "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestConfigSaveAndLoad(t *testing.T) {
	// Create temporary directory
	tempDir, err := os.MkdirTemp("", "autocmdr-test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Create test config
	cfg := &Config{
		Model:     "test-model",
		ServerURL: "http://test.example.com:11434",
		Token:     "test-token",
		LogLevel:  "debug",
		ConfigDir: tempDir,
	}

	// Save config
	err = cfg.Save()
	require.NoError(t, err)

	// Check if config file exists
	configPath := filepath.Join(tempDir, "config.json")
	assert.FileExists(t, configPath)

	// Set environment variable to use temp directory
	os.Setenv("LANGCHAIN_CHAT_CONFIG_DIR", tempDir)
	defer os.Unsetenv("LANGCHAIN_CHAT_CONFIG_DIR")

	// Load config
	loadedCfg, err := Load()
	require.NoError(t, err)

	// Compare configs
	assert.Equal(t, cfg.Model, loadedCfg.Model)
	assert.Equal(t, cfg.ServerURL, loadedCfg.ServerURL)
	assert.Equal(t, cfg.Token, loadedCfg.Token)
	assert.Equal(t, cfg.LogLevel, loadedCfg.LogLevel)
}

func TestGetConfigPath(t *testing.T) {
	cfg := &Config{
		ConfigDir: "/test/dir",
	}

	expected := filepath.Join("/test/dir", "config.json")
	assert.Equal(t, expected, cfg.GetConfigPath())
}
