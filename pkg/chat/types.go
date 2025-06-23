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

// ChatOptions contains options for the chat session
type ChatOptions struct {
	SystemPrompt   string
	MemorySize     int
	StreamResponse bool
}

// DefaultChatOptions returns default chat options
func DefaultChatOptions() *ChatOptions {
	return &ChatOptions{
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

// ChatSession represents a chat session state
type ChatSession struct {
	ID       string
	Memory   schema.Memory
	Options  *ChatOptions
	LastExec *ExecutionResult
}

// NewChatSession creates a new chat session
func NewChatSession(id string, options *ChatOptions) *ChatSession {
	if options == nil {
		options = DefaultChatOptions()
	}

	return &ChatSession{
		ID:      id,
		Options: options,
	}
}
