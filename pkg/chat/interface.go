package chat

import (
	"context"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/schema"
)

// Assistant defines the interface for chat assistants
type Assistant interface {
	// Run starts the chat session
	Run(ctx context.Context, llm llms.Model, memory schema.Memory) error

	// ProcessInput processes user input and returns AI response
	ProcessInput(ctx context.Context, input string) (*AssistantResult, error)

	// ExecuteScript executes a script and returns the result
	ExecuteScript(ctx context.Context, script string) (*ExecutionResult, error)

	// LoadPrompt loads the system prompt
	LoadPrompt() string

	// SetOptions sets chat options
	SetOptions(options *Options)
}

// ScriptExecutor defines the interface for script execution
type ScriptExecutor interface {
	// Execute executes a script command
	Execute(ctx context.Context, command string) (*ExecutionResult, error)

	// CanExecute checks if a command can be executed
	CanExecute(command string) bool

	// GetShell returns the shell type (powershell, bash, etc.)
	GetShell() string
}

// PromptLoader defines the interface for loading prompts
type PromptLoader interface {
	// LoadSystemPrompt loads the system prompt based on OS
	LoadSystemPrompt() string

	// LoadTemplate loads a prompt template by name
	LoadTemplate(name string) (string, error)

	// GetAvailableTemplates returns available template names
	GetAvailableTemplates() []string
}
