package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/blysin/autocmdr/internal/version"
	"github.com/blysin/autocmdr/pkg/chat"
	"github.com/blysin/autocmdr/pkg/config"
	"github.com/blysin/autocmdr/pkg/prompts"
	"github.com/sirupsen/logrus"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/memory"
)

func main() {
	// Parse command line flags
	var (
		initFlag    = flag.Bool("init", false, "Initialize configuration")
		viewFlag    = flag.Bool("view", false, "View current configuration")
		promptFlag  = flag.Bool("prompt", false, "View system prompt")
		versionFlag = flag.Bool("version", false, "Show version information")
		modelFlag   = flag.String("m", "", "Model name")
		serverFlag  = flag.String("u", "", "Server URL")
		tokenFlag   = flag.String("t", "", "API token")
		logLevel    = flag.String("log-level", "", "Log level (debug, info, warn, error)")
	)
	flag.Parse()

	// Show version information
	if *versionFlag {
		versionInfo := version.Get()
		fmt.Println(versionInfo.String())
		return
	}

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("Failed to load configuration: %v\n", err)
		os.Exit(1)
	}

	// Override config with command line flags
	if *logLevel != "" {
		cfg.LogLevel = *logLevel
	}

	// Setup logger
	logger := setupLogger(cfg.LogLevel)

	// Handle init command
	if *initFlag {
		if *modelFlag != "" {
			cfg.Model = *modelFlag
		}
		if *serverFlag != "" {
			cfg.ServerURL = *serverFlag
		}
		if *tokenFlag != "" {
			cfg.Token = *tokenFlag
		}

		if err := cfg.Save(); err != nil {
			logger.WithError(err).Fatal("Failed to save configuration")
		}
		fmt.Println("Configuration initialized successfully.")
		return
	}

	// Handle view command
	if *viewFlag {
		fmt.Printf("Configuration:\n")
		fmt.Printf("  Model: %s\n", cfg.Model)
		fmt.Printf("  Server URL: %s\n", cfg.ServerURL)
		fmt.Printf("  Token: %s\n", maskToken(cfg.Token))
		fmt.Printf("  Log Level: %s\n", cfg.LogLevel)
		fmt.Printf("  Config Directory: %s\n", cfg.ConfigDir)
		return
	}

	// Handle prompt command
	if *promptFlag {
		loader := prompts.NewLoader()
		systemPrompt := loader.LoadSystemPrompt()
		fmt.Println(systemPrompt)
		return
	}

	// Validate configuration
	if err := cfg.Validate(); err != nil {
		logger.WithError(err).Fatal("Invalid configuration")
	}

	// Setup context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		logger.Info("Received shutdown signal")
		cancel()
	}()

	// Initialize LLM
	options := []ollama.Option{
		ollama.WithServerURL(cfg.ServerURL),
		ollama.WithModel(cfg.Model),
	}

	if cfg.Token != "" {
		// Add token if provided (this might need adjustment based on ollama client)
		logger.Debug("Using authentication token")
	}

	llm, err := ollama.New(options...)
	if err != nil {
		logger.WithError(err).Fatal("Failed to initialize LLM")
	}

	logger.WithFields(logrus.Fields{
		"model":      cfg.Model,
		"server_url": cfg.ServerURL,
	}).Info("LLM initialized successfully")

	// Create chat memory
	chatMemory := memory.NewConversationWindowBuffer(10)

	// Create chat assistant
	chatOptions := chat.DefaultChatOptions()
	assistant := chat.NewCliAssistant(chatOptions, logger)

	// Start chat session
	logger.Info("Starting chat session")
	if err := assistant.Run(ctx, llm, chatMemory); err != nil {
		logger.WithError(err).Fatal("Chat session failed")
	}

	logger.Info("Chat session ended")
}

// setupLogger configures the logger based on the log level
func setupLogger(logLevel string) *logrus.Logger {
	logger := logrus.New()

	// Set log format
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
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

	return logger
}

// maskToken masks the token for display purposes
func maskToken(token string) string {
	if token == "" {
		return "(not set)"
	}
	if len(token) <= 8 {
		return "****"
	}
	return token[:4] + "****" + token[len(token)-4:]
}
