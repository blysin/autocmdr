package prompts

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

// Loader handles loading and managing prompts
type Loader struct {
	osVersion string
}

// NewLoader creates a new prompt loader
func NewLoader() *Loader {
	return &Loader{
		osVersion: getOSVersion(),
	}
}

// LoadSystemPrompt loads the appropriate system prompt based on the OS
func (l *Loader) LoadSystemPrompt() string {
	var prompt string
	switch runtime.GOOS {
	case "windows":
		prompt = PowershellAssistant
	default:
		prompt = ShellAssistant
	}

	// Replace template placeholders
	prompt = strings.ReplaceAll(prompt, "<'>", "`")
	prompt = strings.ReplaceAll(prompt, "{{.osVersion}}", l.osVersion)

	return prompt
}

// LoadTemplate loads a specific prompt template by name
func (l *Loader) LoadTemplate(name string) (string, error) {
	switch name {
	case "powershell":
		return PowershellAssistant, nil
	case "shell":
		return ShellAssistant, nil
	default:
		return "", fmt.Errorf("unknown template: %s", name)
	}
}

// GetAvailableTemplates returns a list of available template names
func (l *Loader) GetAvailableTemplates() []string {
	return []string{"powershell", "shell"}
}

// GetOSVersion returns the operating system version
func (l *Loader) GetOSVersion() string {
	return l.osVersion
}

// getOSVersion retrieves the OS version information
func getOSVersion() string {
	switch runtime.GOOS {
	case "windows":
		return getWindowsVersion()
	case "linux":
		return getLinuxVersion()
	case "darwin":
		return getMacOSVersion()
	default:
		return fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)
	}
}

// getWindowsVersion gets Windows version information
func getWindowsVersion() string {
	cmd := exec.Command("cmd", "/c", "ver")
	output, err := cmd.Output()
	if err != nil {
		return "Windows (version unknown)"
	}
	return strings.TrimSpace(string(output))
}

// getLinuxVersion gets Linux version information
func getLinuxVersion() string {
	// Try to get version from /etc/os-release
	cmd := exec.Command("sh", "-c", "cat /etc/os-release | grep PRETTY_NAME | cut -d'=' -f2 | tr -d '\"'")
	output, err := cmd.Output()
	if err == nil && len(output) > 0 {
		return strings.TrimSpace(string(output))
	}

	// Fallback to uname
	cmd = exec.Command("uname", "-a")
	output, err = cmd.Output()
	if err != nil {
		return "Linux (version unknown)"
	}
	return strings.TrimSpace(string(output))
}

// getMacOSVersion gets macOS version information
func getMacOSVersion() string {
	cmd := exec.Command("sw_vers", "-productName", "-productVersion")
	output, err := cmd.Output()
	if err != nil {
		return "macOS (version unknown)"
	}
	return strings.TrimSpace(string(output))
}

// CreateConversationPrompt creates a conversation prompt template
func (l *Loader) CreateConversationPrompt(systemPrompt string) string {
	return systemPrompt + `
Current conversation:
{{.history}}
Human: {{.input}}
AI:`
}
