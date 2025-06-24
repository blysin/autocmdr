package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()

	if cfg.Model == "" {
		t.Error("expected Model to be non-empty")
	}
	if cfg.ServerURL == "" {
		t.Error("expected ServerURL to be non-empty")
	}
	if cfg.LogLevel != "info" {
		t.Errorf("expected LogLevel to be 'info', got '%s'", cfg.LogLevel)
	}
	if !strings.Contains(cfg.ConfigDir, ".autocmdr") {
		t.Errorf("expected ConfigDir to contain '.autocmdr', got '%s'", cfg.ConfigDir)
	}
}

func TestConfigSaveAndLoad(t *testing.T) {
	// Create temporary directory
	tempDir, err := os.MkdirTemp("", "autocmdr-test")
	if err != nil {
		t.Fatalf("failed to create temp directory: %v", err)
	}
	defer func(path string) {
		removeErr := os.RemoveAll(path)
		if removeErr != nil {
			t.Errorf("failed to remove temp directory: %v", removeErr)
		}
	}(tempDir)

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
	if err != nil {
		t.Fatalf("failed to save config: %v", err)
	}

	// Check if config file exists
	configPath := filepath.Join(tempDir, "config.json")
	if _, statErr := os.Stat(configPath); os.IsNotExist(statErr) {
		t.Errorf("config file does not exist at %s", configPath)
	}

	// Set environment variable to use temp directory
	if err := os.Setenv("LANGCHAIN_CHAT_CONFIG_DIR", tempDir); err != nil {
		t.Fatalf("failed to set LANGCHAIN_CHAT_CONFIG_DIR: %v", err)
	}

	// 使用匿名函数包装 Unsetenv 以捕获错误
	defer func() {
		if err := os.Unsetenv("LANGCHAIN_CHAT_CONFIG_DIR"); err != nil {
			t.Errorf("failed to unset LANGCHAIN_CHAT_CONFIG_DIR: %v", err)
		}
	}()

	// Load config
	loadedCfg, err := Load()
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	// Compare configs
	if cfg.Model != loadedCfg.Model {
		t.Errorf("expected Model %s, got %s", cfg.Model, loadedCfg.Model)
	}
	if cfg.ServerURL != loadedCfg.ServerURL {
		t.Errorf("expected ServerURL %s, got %s", cfg.ServerURL, loadedCfg.ServerURL)
	}
	if cfg.Token != loadedCfg.Token {
		t.Errorf("expected Token %s, got %s", cfg.Token, loadedCfg.Token)
	}
	if cfg.LogLevel != loadedCfg.LogLevel {
		t.Errorf("expected LogLevel %s, got %s", cfg.LogLevel, loadedCfg.LogLevel)
	}
}

func TestGetConfigPath(t *testing.T) {
	cfg := &Config{
		ConfigDir: filepath.Join("test", "dir"),
	}

	expected := filepath.Join("test", "dir", "config.json")
	if expected != cfg.GetConfigPath() {
		t.Errorf("expected %s, got %s", expected, cfg.GetConfigPath())
	}
}
