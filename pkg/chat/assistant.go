// Package chat implements the core chat assistant functionality including conversation management and response generation.
package chat

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/blysin/autocmdr/pkg/prompts"
	"github.com/blysin/autocmdr/pkg/utils"
	"github.com/chzyer/readline"
	"github.com/sirupsen/logrus"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/outputparser"
	lcprompts "github.com/tmc/langchaingo/prompts"
	"github.com/tmc/langchaingo/schema"
)

// CliAssistant implements the Assistant interface for CLI interactions
type CliAssistant struct {
	options        *ChatOptions
	promptLoader   *prompts.Loader
	lastExecResult *ExecutionResult
	logger         *logrus.Logger
}

// NewCliAssistant creates a new CLI assistant
func NewCliAssistant(options *ChatOptions, logger *logrus.Logger) *CliAssistant {
	if options == nil {
		options = DefaultChatOptions()
	}
	if logger == nil {
		logger = logrus.New()
	}

	return &CliAssistant{
		options:      options,
		promptLoader: prompts.NewLoader(),
		logger:       logger,
	}
}

// Run starts the chat session
func (c *CliAssistant) Run(ctx context.Context, llm llms.Model, chatMemory schema.Memory) error {
	systemPrompt := c.LoadPrompt()
	c.logger.WithField("os", runtime.GOOS).Info("Starting chat session")

	chain := chains.LLMChain{
		Prompt: lcprompts.NewPromptTemplate(
			systemPrompt,
			[]string{"history", "input"},
		),
		LLM:          llm,
		Memory:       chatMemory,
		OutputParser: outputparser.NewSimple(),
		OutputKey:    "text",
	}

	reader := bufio.NewReader(os.Stdin)
	c.printWelcomeInfo()

	for {
		userInput, shouldContinue := c.handleUserInput(chatMemory, ctx)
		if !shouldContinue {
			break
		}
		if userInput == "" {
			continue
		}

		resp, err := c.processAIResponse(ctx, &chain, userInput)
		if err != nil {
			c.logger.WithError(err).Error("Failed to process AI response")
			fmt.Printf("Error: %v\n", err)
			continue
		}

		script, err := c.parseScript(resp)
		if err != nil {
			c.logger.WithError(err).Error("Failed to parse script")
			fmt.Printf("\nError: %v\n", err)
			continue
		}

		c.confirmAndExecute(reader, script, ctx)
	}

	return nil
}

func (c *CliAssistant) printWelcomeInfo() {
	fmt.Println("Welcome to the AutoCmdr App! Type 'exit' to exit, 'clear' to clear history, or 'help' for more info.")
}

// ProcessInput processes user input and returns AI response
func (c *CliAssistant) ProcessInput(ctx context.Context, input string) (*AssistantResult, error) {
	// This would be implemented for programmatic use
	return nil, fmt.Errorf("not implemented")
}

// ExecuteScript executes a script and returns the result
func (c *CliAssistant) ExecuteScript(ctx context.Context, script string) (*ExecutionResult, error) {
	startTime := time.Now()

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.CommandContext(ctx, "powershell", "-Command", script)
	default:
		cmd = exec.CommandContext(ctx, "bash", "-c", script)
	}

	output, err := cmd.CombinedOutput()
	duration := time.Since(startTime)

	result := &ExecutionResult{
		Command:  script,
		Duration: duration.String(),
		Output:   string(output),
	}

	if err != nil {
		result.Success = false
		result.Error = err.Error()
		if exitError, ok := err.(*exec.ExitError); ok {
			result.ExitCode = exitError.ExitCode()
		} else {
			result.ExitCode = -1
		}
	} else {
		result.Success = true
		result.ExitCode = 0
	}

	return result, nil
}

// LoadPrompt loads the system prompt
func (c *CliAssistant) LoadPrompt() string {
	systemPrompt := c.promptLoader.LoadSystemPrompt()
	return c.promptLoader.CreateConversationPrompt(systemPrompt)
}

// SetOptions sets chat options
func (c *CliAssistant) SetOptions(options *ChatOptions) {
	c.options = options
}

// handleUserInput handles user input with readline support
func (c *CliAssistant) handleUserInput(chatMemory schema.Memory, ctx context.Context) (string, bool) {
	rl, err := readline.New("You: ")
	if err != nil {
		c.logger.WithError(err).Fatal("Failed to create readline")
	}
	defer func() {
		if err := rl.Close(); err != nil {
			c.logger.WithError(err).Error("Failed to close readline")
		}
	}()

	line, err := rl.Readline()
	if err != nil {
		if errors.Is(err, readline.ErrInterrupt) {
			return "", false
		}
		c.logger.WithError(err).Fatal("Failed to read input")
	}

	userInput := strings.TrimSpace(line)

	if userInput == "" {
		return "", true
	}

	switch userInput {
	case "exit":
		c.logger.Info("Exiting the chat...")
		return "", false
	case "clear":
		if err := chatMemory.Clear(ctx); err != nil {
			c.logger.WithError(err).Error("Failed to clear memory")
			return "", false
		}
		c.logger.Info("Chat history cleared.")
		return "", true
	case "help":
		c.logger.Info("Available commands: exit, clear, help")
		return "", true
	default:
		return userInput, true
	}
}

// processAIResponse processes the AI response with streaming
func (c *CliAssistant) processAIResponse(ctx context.Context, chain *chains.LLMChain, userInput string) (string, error) {
	start := false

	// Append last execution result if available
	if c.lastExecResult != nil {
		userInput = fmt.Sprintf("Last execution result: %s\n%s", c.lastExecResult.Output, userInput)
	}

	resp, err := chains.Run(ctx, chain, userInput, chains.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		if !start {
			fmt.Print("Bot: ")
			start = true
		}
		fmt.Print(string(chunk))
		return nil
	}))

	if start {
		fmt.Println() // Add newline after streaming
	}

	return resp, err
}

// parseScript parses the AI response to extract script information
func (c *CliAssistant) parseScript(resp string) (*AssistantResult, error) {
	resp = strings.TrimSpace(resp)
	thinkEnd := "</think>"

	// Extract content after thinking block
	if idx := strings.LastIndex(resp, thinkEnd); idx != -1 {
		resp = resp[idx+len(thinkEnd):]
	}

	jsonStr, err := utils.ExtractFirstJSON(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to extract JSON from response: %w", err)
	}

	result := &AssistantResult{}
	if err := json.Unmarshal([]byte(jsonStr), result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return result, nil
}

// confirmAndExecute handles script confirmation and execution
func (c *CliAssistant) confirmAndExecute(reader *bufio.Reader, script *AssistantResult, ctx context.Context) {
	if !script.Success {
		fmt.Printf("\nAI did not provide a script: %s\n", script.Script)
		return
	}

	if script.MultipleLines {
		fmt.Println("\nAI: Please save the following content as a script file and execute:")
		fmt.Println(script.Script)
		return
	}

	scriptContent := strings.TrimSpace(script.Script)
	fmt.Println("\nExecute script directly? (y/n)")
	fmt.Print("You: ")

	confirm, err := reader.ReadString('\n')
	if err != nil {
		c.logger.WithError(err).Error("Failed to read confirmation")
		return
	}

	confirm = strings.TrimSpace(confirm)
	if confirm == "y" || confirm == "Y" {
		result, err := c.ExecuteScript(ctx, scriptContent)
		if err != nil {
			c.logger.WithError(err).Error("Failed to execute script")
			fmt.Printf("Execution error: %v\n", err)
			return
		}

		c.lastExecResult = result

		if result.Success {
			fmt.Printf("✅ Script executed successfully (exit code: %d)\n", result.ExitCode)
			if result.Output != "" {
				fmt.Printf("Output:\n%s\n", result.Output)
			}
		} else {
			fmt.Printf("❌ Script execution failed (exit code: %d)\n", result.ExitCode)
			if result.Error != "" {
				fmt.Printf("Error: %s\n", result.Error)
			}
			if result.Output != "" {
				fmt.Printf("Output:\n%s\n", result.Output)
			}
		}

		fmt.Printf("Duration: %s\n", result.Duration)
	}
}
