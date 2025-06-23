package main

import (
	"context"
	"fmt"
	"log"

	"github.com/blysin/autocmdr/pkg/chat"
	"github.com/blysin/autocmdr/pkg/config"
	"github.com/sirupsen/logrus"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/memory"
)

func main() {
	// This example shows basic usage of the LangChain Chat App
	fmt.Println("LangChain Chat App - Basic Example")

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Validate configuration
	if err := cfg.Validate(); err != nil {
		log.Fatalf("Invalid configuration: %v", err)
	}

	// Setup logger
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	// Initialize LLM
	options := []ollama.Option{
		ollama.WithServerURL(cfg.ServerURL),
		ollama.WithModel(cfg.Model),
	}

	llm, err := ollama.New(options...)
	if err != nil {
		log.Fatalf("Failed to initialize LLM: %v", err)
	}

	// Create chat memory
	chatMemory := memory.NewConversationWindowBuffer(5)

	// Create chat assistant
	chatOptions := chat.DefaultChatOptions()
	assistant := chat.NewCliAssistant(chatOptions, logger)

	// Start chat session
	ctx := context.Background()
	fmt.Println("Starting chat session...")

	if err := assistant.Run(ctx, llm, chatMemory); err != nil {
		log.Fatalf("Chat session failed: %v", err)
	}

	fmt.Println("Chat session ended.")
}
