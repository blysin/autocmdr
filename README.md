# AutoCmdr App

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/blysin/autocmdr)](https://goreportcard.com/report/github.com/blysin/autocmdr)
[![CI](https://github.com/blysin/autocmdr/workflows/CI/badge.svg)](https://github.com/blysin/autocmdr/actions)

[English](README.md) | [简体中文](README_zh.md)

A command-line assistant application built using Go and LangChain, designed to help users securely interact with AI models to generate and execute system commands.

## 🚀 Features

- **Cross-Platform Support**: Works on Windows (PowerShell), Linux, and macOS (Bash)
- **Interactive CLI**: Rich command-line interface with readline support
- **Smart Command Generation**: AI-powered command generation with safety checks
- **Configuration Management**: Flexible configuration with file and environment variable support
- **Structured Logging**: Comprehensive logging with configurable levels
- **Memory Management**: Conversation history with configurable window size
- **Safety First**: Built-in safety checks and confirmation prompts
- **Extensible Architecture**: Modular design following Go best practices

## 📦 Installation

### From Source

```bash
# Clone the repository
git clone https://github.com/blysin/autocmdr.git
cd autocmdr-app

# Build and install
make install
```

### Using Go Install

```bash
go install github.com/blysin/autocmdr/cmd/autocmdr@latest
```

### Pre-built Binaries

Download the latest release from the [releases page](https://github.com/blysin/autocmdr/releases).

## 🔧 Configuration

### Initialize Configuration

```bash
autocmdr -init -m "your-model" -u "http://your-ollama-server:11434"
```

### Configuration Options

The application supports configuration through:

1. **Configuration File**: `~/.autocmdr/config.json`
2. **Environment Variables**: Prefixed with `LANGCHAIN_CHAT_`
3. **Command Line Flags**

#### Configuration Parameters

| Parameter | Environment Variable | Default | Description |
|-----------|---------------------|---------|-------------|
| `model` | `LANGCHAIN_CHAT_MODEL` | `qwen3:14b` | AI model name |
| `server_url` | `LANGCHAIN_CHAT_SERVER_URL` | `http://localhost:11434` | Ollama server URL |
| `token` | `LANGCHAIN_CHAT_TOKEN` | `""` | API authentication token |
| `log_level` | `LANGCHAIN_CHAT_LOG_LEVEL` | `info` | Log level (debug, info, warn, error) |

### Example Configuration File

```json
{
  "model": "qwen3:14b",
  "server_url": "http://localhost:11434",
  "token": "",
  "log_level": "info"
}
```

## 🎯 Usage

### Basic Usage

```bash
# Start interactive chat
autocmdr

# View current configuration
autocmdr -view

# View system prompt
autocmdr -prompt

# Show version information
autocmdr -version
```

### Interactive Commands

Once in the chat session:

- Type your request in natural language
- The AI will generate appropriate commands
- Confirm execution with `y` or `n`
- Use `clear` to clear conversation history
- Use `exit` to quit the application

### Example Session

```
You: list all files in the current directory
Bot: I'll help you list all files in the current directory.

{
  "success": "true",
  "multipleLines": "false",
  "script": "Get-ChildItem -Force"
}

Execute script directly? (y/n)
You: y
✅ Script executed successfully (exit code: 0)
Output:
Directory: C:\Users\example

Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2023/12/01     10:30                Documents
d-----         2023/12/01     10:30                Downloads
...
```

## 🏗️ Development

### Prerequisites

- Go 1.21 or later
- Make (optional, for using Makefile)

### Building from Source

```bash
# Clone the repository
git clone https://github.com/blysin/autocmdr.git
cd autocmdr-app

# Install dependencies
make deps

# Run tests
make test

# Build the application
make build

# Run the application
make run
```

### Project Structure

```
autocmdr-app/
├── cmd/
│   └── autocmdr/        # Application entry point
├── pkg/
│   ├── config/               # Configuration management
│   ├── chat/                 # Chat functionality
│   ├── prompts/              # Prompt templates and loading
│   └── utils/                # Utility functions
├── internal/
│   └── version/              # Version information
├── examples/                 # Usage examples
├── docs/                     # Documentation
├── scripts/                  # Build and utility scripts
└── .github/                  # GitHub workflows and templates
```

### Available Make Targets

```bash
make help                     # Show all available targets
make build                    # Build the application
make test                     # Run tests
make test-coverage           # Run tests with coverage
make lint                    # Run linters
make fmt                     # Format code
make clean                   # Clean build artifacts
```

## 🧪 Testing

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run benchmarks
make bench
```

## 📚 Documentation

- [API Documentation](docs/api.md)
- [Configuration Guide](docs/configuration.md)
- [Contributing Guide](CONTRIBUTING.md)

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

### Development Workflow

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Run the test suite
6. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [LangChain Go](https://github.com/tmc/langchaingo) - Go implementation of LangChain
- [Ollama](https://ollama.ai/) - Local AI model serving
- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Viper](https://github.com/spf13/viper) - Configuration management

## 📞 Support

- 📧 Email: blysin@163.com
- 🐛 Issues: [GitHub Issues](https://github.com/blysin/autocmdr/issues)
- 💬 Discussions: [GitHub Discussions](https://github.com/blysin/autocmdr/discussions)

---

Made with ❤️ by the AutoCmdr App team
