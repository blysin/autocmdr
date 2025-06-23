# API Documentation

This document describes the public API of the LangChain Chat App.

## Table of Contents

- [Configuration Package](#configuration-package)
- [Chat Package](#chat-package)
- [Prompts Package](#prompts-package)
- [Utils Package](#utils-package)

## Configuration Package

The `config` package provides configuration management functionality.

### Types

#### Config

```go
type Config struct {
    Model     string `mapstructure:"model" json:"model"`
    ServerURL string `mapstructure:"server_url" json:"server_url"`
    Token     string `mapstructure:"token" json:"token"`
    LogLevel  string `mapstructure:"log_level" json:"log_level"`
    ConfigDir string `mapstructure:"config_dir" json:"config_dir"`
}
```

Configuration structure for the application.

### Functions

#### DefaultConfig

```go
func DefaultConfig() *Config
```

Returns a new Config instance with default values.

#### Load

```go
func Load() (*Config, error)
```

Loads configuration from file and environment variables.

#### (c *Config) Save

```go
func (c *Config) Save() error
```

Saves the configuration to file.

#### (c *Config) Validate

```go
func (c *Config) Validate() error
```

Validates the configuration parameters.

#### (c *Config) GetConfigPath

```go
func (c *Config) GetConfigPath() string
```

Returns the path to the configuration file.

## Chat Package

The `chat` package provides chat functionality and interfaces.

### Types

#### AssistantResult

```go
type AssistantResult struct {
    Success       bool   `json:"success"`
    MultipleLines bool   `json:"multipleLines"`
    Script        string `json:"script"`
}
```

Represents the result from the AI assistant.

#### ChatOptions

```go
type ChatOptions struct {
    SystemPrompt   string
    MemorySize     int
    StreamResponse bool
}
```

Contains options for the chat session.

#### ExecutionResult

```go
type ExecutionResult struct {
    Success    bool   `json:"success"`
    Output     string `json:"output"`
    Error      string `json:"error,omitempty"`
    ExitCode   int    `json:"exit_code"`
    Duration   string `json:"duration"`
    Command    string `json:"command"`
}
```

Represents the result of script execution.

#### ChatSession

```go
type ChatSession struct {
    ID       string
    Memory   schema.Memory
    Options  *ChatOptions
    LastExec *ExecutionResult
}
```

Represents a chat session state.

### Interfaces

#### Assistant

```go
type Assistant interface {
    Run(ctx context.Context, llm llms.Model, memory schema.Memory) error
    ProcessInput(ctx context.Context, input string) (*AssistantResult, error)
    ExecuteScript(ctx context.Context, script string) (*ExecutionResult, error)
    LoadPrompt() string
    SetOptions(options *ChatOptions)
}
```

Defines the interface for chat assistants.

#### ScriptExecutor

```go
type ScriptExecutor interface {
    Execute(ctx context.Context, command string) (*ExecutionResult, error)
    CanExecute(command string) bool
    GetShell() string
}
```

Defines the interface for script execution.

#### PromptLoader

```go
type PromptLoader interface {
    LoadSystemPrompt() string
    LoadTemplate(name string) (string, error)
    GetAvailableTemplates() []string
}
```

Defines the interface for loading prompts.

### Functions

#### DefaultChatOptions

```go
func DefaultChatOptions() *ChatOptions
```

Returns default chat options.

#### NewChatSession

```go
func NewChatSession(id string, options *ChatOptions) *ChatSession
```

Creates a new chat session.

#### NewCliAssistant

```go
func NewCliAssistant(options *ChatOptions, logger *logrus.Logger) *CliAssistant
```

Creates a new CLI assistant.

## Prompts Package

The `prompts` package provides prompt template management.

### Types

#### Loader

```go
type Loader struct {
    // private fields
}
```

Handles loading and managing prompts.

### Functions

#### NewLoader

```go
func NewLoader() *Loader
```

Creates a new prompt loader.

#### (l *Loader) LoadSystemPrompt

```go
func (l *Loader) LoadSystemPrompt() string
```

Loads the appropriate system prompt based on the OS.

#### (l *Loader) LoadTemplate

```go
func (l *Loader) LoadTemplate(name string) (string, error)
```

Loads a specific prompt template by name.

#### (l *Loader) GetAvailableTemplates

```go
func (l *Loader) GetAvailableTemplates() []string
```

Returns a list of available template names.

#### (l *Loader) CreateConversationPrompt

```go
func (l *Loader) CreateConversationPrompt(systemPrompt string) string
```

Creates a conversation prompt template.

## Utils Package

The `utils` package provides utility functions.

### JSON Functions

#### ExtractFirstJSON

```go
func ExtractFirstJSON(input string) (string, error)
```

Extracts the first valid JSON object from a string.

#### PrettyPrintJSON

```go
func PrettyPrintJSON(jsonStr string) (string, error)
```

Formats JSON string with indentation.

#### ValidateJSON

```go
func ValidateJSON(jsonStr string) error
```

Checks if a string is valid JSON.

#### ParseJSONToMap

```go
func ParseJSONToMap(jsonStr string) (map[string]interface{}, error)
```

Parses JSON string to map[string]interface{}.

### Path Functions

#### GetUserConfigDir

```go
func GetUserConfigDir(appName string) (string, error)
```

Returns the user's configuration directory.

#### EnsureDir

```go
func EnsureDir(dir string) error
```

Creates a directory if it doesn't exist.

#### FileExists

```go
func FileExists(path string) bool
```

Checks if a file exists.

#### GetExecutablePath

```go
func GetExecutablePath() (string, error)
```

Returns the path of the current executable.

#### JoinPath

```go
func JoinPath(elements ...string) string
```

Safely joins path elements.

## Error Handling

All functions that can fail return an error as the last return value. Errors are wrapped with context using `fmt.Errorf` to provide meaningful error messages.

## Context Support

Functions that perform I/O operations or long-running tasks accept a `context.Context` parameter for cancellation and timeout support.

## Thread Safety

The API is designed to be thread-safe where applicable. However, individual instances of types like `Config` or `ChatSession` should not be accessed concurrently without proper synchronization.
