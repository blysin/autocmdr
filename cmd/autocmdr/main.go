// Copyright © 2021 Blysin <blysin@163.com>
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/memory"

	"github.com/blysin/autocmdr/pkg/chat"
	"github.com/blysin/autocmdr/pkg/config"
	"github.com/blysin/autocmdr/pkg/prompts"
	"github.com/blysin/autocmdr/pkg/version"
)

// App represents the application.
type App struct {
	logger *logrus.Logger
	cfg    *config.Config
}

// NewApp creates a new App instance.
func NewApp() *App {
	return &App{}
}

// Args represents the command line arguments.
type Args struct {
	Init     bool
	View     bool
	Prompt   bool
	Version  bool
	Model    string
	Server   string
	Token    string
	LogLevel string
}

// ParseArgs parses command line arguments and returns them as a struct.
func (a *App) ParseArgs() *Args {
	var args Args
	flag.BoolVar(&args.Init, "init", false, "Initialize configuration")
	flag.BoolVar(&args.View, "view", false, "View current configuration")
	flag.BoolVar(&args.Prompt, "prompt", false, "View system prompt")
	flag.BoolVar(&args.Version, "version", false, "Show version information")
	flag.StringVar(&args.Model, "m", "", "Model name")
	flag.StringVar(&args.Server, "u", "", "Server URL")
	flag.StringVar(&args.Token, "t", "", "API token")
	flag.StringVar(&args.LogLevel, "log-level", "", "Log level (debug, info, warn, error)")
	flag.Parse()
	return &args
}

func (a *App) parseFlags() bool {
	args := a.ParseArgs()

	return a.handleFlags(args)
}

func (a *App) handleFlags(args *Args) (continueChat bool) {
	continueChat = false
	if args.Version {
		a.showVersion()
		return continueChat
	}

	var err error
	a.cfg, err = config.Load()
	if err != nil {
		fmt.Printf("Failed to load configuration: %v\n", err)
		os.Exit(1)
	}

	if args.LogLevel != "" {
		a.cfg.LogLevel = args.LogLevel
	}

	a.logger = setupLogger(a.cfg.LogLevel)

	if args.Init {
		a.handleInit(args.Model, args.Server, args.Token)
		return continueChat
	}

	if args.View {
		a.showConfig()
		return continueChat
	}

	if args.Prompt {
		a.showPrompt()
		return continueChat
	}

	if err = a.cfg.Validate(); err != nil {
		a.logger.WithError(err).Fatal("Invalid configuration")
	}
	continueChat = true
	return continueChat
}

func (a *App) showVersion() {
	versionInfo := version.Get()
	fmt.Println(versionInfo.String())
}

func (a *App) handleInit(model, server, token string) {
	if model != "" {
		a.cfg.Model = model
	}
	if server != "" {
		a.cfg.ServerURL = server
	}
	if token != "" {
		a.cfg.Token = token
	}

	if err := a.cfg.Save(); err != nil {
		a.logger.WithError(err).Fatal("Failed to save configuration")
	}
	fmt.Println("Configuration initialized successfully.")
}

func (a *App) showConfig() {
	fmt.Printf("Configuration:\n")
	fmt.Printf("  Model: %s\n", a.cfg.Model)
	fmt.Printf("  Server URL: %s\n", a.cfg.ServerURL)
	fmt.Printf("  Token: %s\n", maskToken(a.cfg.Token))
	fmt.Printf("  Log Level: %s\n", a.cfg.LogLevel)
	fmt.Printf("  Config Directory: %s\n", a.cfg.ConfigDir)
}

func (a *App) showPrompt() {
	loader := prompts.NewLoader()
	systemPrompt := loader.LoadSystemPrompt()
	fmt.Println(systemPrompt)
}

func (a *App) runChatSession() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	a.setupShutdownHandler(cancel)

	llm := a.initLLM()
	chatMemory := memory.NewConversationWindowBuffer(10)
	assistant := chat.NewCliAssistant(chat.DefaultChatOptions(), a.logger)

	a.logger.Info("Starting chat session")
	if err := assistant.Run(ctx, llm, chatMemory); err != nil {
		a.logger.WithError(err).Fatal("Chat session failed")
	}
	a.logger.Info("Chat session ended")
}

func (a *App) setupShutdownHandler(cancel context.CancelFunc) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		a.logger.Info("Received shutdown signal")
		cancel()
	}()
}

func (a *App) initLLM() *ollama.LLM {
	options := []ollama.Option{
		ollama.WithServerURL(a.cfg.ServerURL),
		ollama.WithModel(a.cfg.Model),
	}

	if a.cfg.Token != "" {
		a.logger.Debug("Using authentication token")
	}

	llm, err := ollama.New(options...)
	if err != nil {
		a.logger.WithError(err).Fatal("Failed to initialize LLM")
	}

	a.logger.WithFields(logrus.Fields{
		"model":      a.cfg.Model,
		"server_url": a.cfg.ServerURL,
	}).Info("LLM initialized successfully")

	return llm
}

func main() {
	app := NewApp()
	continueChat := app.parseFlags()
	if !continueChat {
		return
	}
	app.runChatSession()
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
