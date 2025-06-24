package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// GetUserConfigDir returns the user's configuration directory
func GetUserConfigDir(appName string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}

	var configDir string
	switch runtime.GOOS {
	case "windows":
		configDir = filepath.Join(homeDir, "AppData", "Roaming", appName)
	case "darwin":
		configDir = filepath.Join(homeDir, "Library", "Application Support", appName)
	default:
		configDir = filepath.Join(homeDir, "."+appName)
	}

	return configDir, nil
}

// EnsureDir creates a directory if it doesn't exist
func EnsureDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0o750); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}
	return nil
}

// FileExists checks if a file exists
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// GetExecutablePath returns the path of the current executable
func GetExecutablePath() (string, error) {
	execPath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("failed to get executable path: %w", err)
	}
	return filepath.Dir(execPath), nil
}

// GetWorkingDir returns the current working directory
func GetWorkingDir() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get working directory: %w", err)
	}
	return wd, nil
}

// JoinPath safely joins path elements
func JoinPath(elements ...string) string {
	return filepath.Join(elements...)
}

// CleanPath cleans and normalizes a path
func CleanPath(path string) string {
	return filepath.Clean(path)
}

// IsAbsolutePath checks if a path is absolute
func IsAbsolutePath(path string) bool {
	return filepath.IsAbs(path)
}

// GetTempDir returns the system temporary directory
func GetTempDir() string {
	return os.TempDir()
}

// CreateTempFile creates a temporary file with the given prefix and suffix
func CreateTempFile(prefix, suffix string) (*os.File, error) {
	return os.CreateTemp("", prefix+"*"+suffix)
}
