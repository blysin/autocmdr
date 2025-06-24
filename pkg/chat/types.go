package chat

import (
	"context"

	"github.com/tmc/langchaingo/schema"
)

// AssistantResult represents the result from the AI assistant
type AssistantResult struct {
	Success       bool   `json:"success"`
	MultipleLines bool   `json:"multipleLines"`
	Script        string `json:"script"`
}

// Options contains options for the chat session
type Options struct {
	SystemPrompt   string
	MemorySize     int
	StreamResponse bool
}

// DefaultChatOptions returns default chat options
func DefaultChatOptions() *Options {
	return &Options{
		MemorySize:     10,
		StreamResponse: true,
	}
}

// StreamFunc is a function type for handling streaming responses
type StreamFunc func(ctx context.Context, chunk []byte) error

// ExecutionResult represents the result of script execution
type ExecutionResult struct {
	Success  bool   `json:"success"`
	Output   string `json:"output"`
	Error    string `json:"error,omitempty"`
	ExitCode int    `json:"exit_code"`
	Duration string `json:"duration"`
	Command  string `json:"command"`
}

// Session represents a chat session state
type Session struct {
	ID       string
	Memory   schema.Memory
	Options  *Options
	LastExec *ExecutionResult
}
