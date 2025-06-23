package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/blysin/autocmdr/pkg/chat"
	"github.com/blysin/autocmdr/pkg/config"
	"github.com/blysin/autocmdr/pkg/prompts"
	"github.com/sirupsen/logrus"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/memory"
)

func main() {
	// This example shows advanced usage with custom configuration
	fmt.Println("LangChain Chat App - Advanced Example")

	// Create custom configuration
	cfg := &config.Config{
		Model:     "custom-model",
		ServerURL: "http://localhost:11434",
		Token:     "",
		LogLevel:  "debug",
		ConfigDir: "./config",
	}

	// Setup advanced logger
	logger := setupAdvancedLogger(cfg.LogLevel)

	// Setup graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		logger.Info("Received shutdown signal")
		cancel()
	}()

	// Initialize LLM with custom options
	options := []ollama.Option{
		ollama.WithServerURL(cfg.ServerURL),
		ollama.WithModel(cfg.Model),
	}

	llm, err := ollama.New(options...)
	if err != nil {
		logger.WithError(err).Fatal("Failed to initialize LLM")
	}

	// Create custom chat memory with larger window
	chatMemory := memory.NewConversationWindowBuffer(20)

	// Create custom chat options
	chatOptions := &chat.ChatOptions{
		MemorySize:     20,
		StreamResponse: true,
	}

	// Load custom prompt
	promptLoader := prompts.NewLoader()
	systemPrompt := promptLoader.LoadSystemPrompt()
	logger.WithField("prompt_length", len(systemPrompt)).Debug("Loaded system prompt")

	// Create chat assistant with custom options
	assistant := chat.NewCliAssistant(chatOptions, logger)

	// Demonstrate programmatic script execution
	logger.Info("Demonstrating programmatic script execution")
	result, err := assistant.ExecuteScript(ctx, "echo 'Hello from advanced example'")
	if err != nil {
		logger.WithError(err).Error("Failed to execute test script")
	} else {
		logger.WithFields(logrus.Fields{
			"success":   result.Success,
			"exit_code": result.ExitCode,
			"duration":  result.Duration,
		}).Info("Test script executed")
		fmt.Printf("Output: %s\n", result.Output)
	}

	// Start interactive chat session
	logger.Info("Starting interactive chat session")
	if err := assistant.Run(ctx, llm, chatMemory); err != nil {
		logger.WithError(err).Fatal("Chat session failed")
	}

	logger.Info("Advanced example completed")
}

func setupAdvancedLogger(logLevel string) *logrus.Logger {
	logger := logrus.New()

	// Set custom formatter
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000Z07:00",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "message",
		},
	})

	// Set log level
	switch logLevel {
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	case "info":
		logger.SetLevel(logrus.InfoLevel)
	case "warn":
		logger.SetLevel(logrus.WarnLevel)
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	default:
		logger.SetLevel(logrus.InfoLevel)
	}

	// Add hooks for advanced logging
	logger.AddHook(&contextHook{})

	return logger
}

// contextHook adds context information to log entries
type contextHook struct{}

func (h *contextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *contextHook) Fire(entry *logrus.Entry) error {
	entry.Data["component"] = "autocmdr-app"
	entry.Data["version"] = "1.0.0"
	return nil
}
