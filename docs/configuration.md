# Configuration Guide

This guide explains how to configure the LangChain Chat App.

## Configuration Methods

The application supports multiple configuration methods, in order of precedence:

1. **Command Line Flags** (highest priority)
2. **Environment Variables**
3. **Configuration File**
4. **Default Values** (lowest priority)

## Configuration Parameters

### Core Parameters

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `model` | string | `qwen3:14b` | AI model name to use |
| `server_url` | string | `http://localhost:11434` | Ollama server URL |
| `token` | string | `""` | API authentication token |
| `log_level` | string | `info` | Log level (debug, info, warn, error) |
| `config_dir` | string | `~/.autocmdr` | Configuration directory path |

## Command Line Flags

```bash
# Initialize configuration
autocmdr -init -m "model-name" -u "http://server:11434" -t "token"

# Override configuration temporarily
autocmdr -m "different-model" --log-level debug

# View current configuration
autocmdr -view

# Show version information
autocmdr -version

# View system prompt
autocmdr -prompt
```

### Available Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--init` | | Initialize configuration |
| `--view` | | View current configuration |
| `--prompt` | | View system prompt |
| `--version` | | Show version information |
| `--model` | `-m` | Model name |
| `--server-url` | `-u` | Server URL |
| `--token` | `-t` | API token |
| `--log-level` | | Log level |

## Environment Variables

All environment variables are prefixed with `LANGCHAIN_CHAT_`:

```bash
export LANGCHAIN_CHAT_MODEL="custom-model"
export LANGCHAIN_CHAT_SERVER_URL="http://localhost:11434"
export LANGCHAIN_CHAT_TOKEN="your-token"
export LANGCHAIN_CHAT_LOG_LEVEL="debug"
export LANGCHAIN_CHAT_CONFIG_DIR="/custom/config/path"
```

## Configuration File

The configuration file is stored in JSON format at `~/.autocmdr/config.json` by default.

### Example Configuration File

```json
{
  "model": "qwen3:14b",
  "server_url": "http://localhost:11434",
  "token": "",
  "log_level": "info",
  "config_dir": "/home/user/.autocmdr"
}
```

### Creating Configuration File

You can create the configuration file manually or use the init command:

```bash
# Initialize with default values
autocmdr -init

# Initialize with custom values
autocmdr -init -m "custom-model" -u "http://custom-server:11434"
```

## Configuration Directory

The configuration directory contains:

- `config.json` - Main configuration file
- Log files (if file logging is enabled)
- Cache files
- Temporary files

### Default Locations

| OS | Default Path |
|----|--------------|
| Linux | `~/.autocmdr` |
| macOS | `~/Library/Application Support/autocmdr` |
| Windows | `%APPDATA%\autocmdr` |

## Model Configuration

### Supported Models

The application works with any Ollama-compatible model. Popular choices include:

- `qwen3:14b` - Default model, good balance of performance and quality
- `llama2:7b` - Lightweight option
- `codellama:13b` - Optimized for code generation
- `mistral:7b` - Fast and efficient

### Model Selection

Choose a model based on your needs:

- **Performance**: Smaller models (7B parameters) are faster
- **Quality**: Larger models (13B+ parameters) provide better responses
- **Specialization**: Code-specific models for programming tasks

## Server Configuration

### Ollama Server Setup

1. **Install Ollama**:
   ```bash
   # Linux/macOS
   curl -fsSL https://ollama.ai/install.sh | sh
   
   # Windows
   # Download from https://ollama.ai/download
   ```

2. **Start Ollama Server**:
   ```bash
   ollama serve
   ```

3. **Pull a Model**:
   ```bash
   ollama pull qwen3:14b
   ```

### Remote Server Configuration

For remote Ollama servers:

```json
{
  "server_url": "http://remote-server:11434",
  "token": "optional-auth-token"
}
```

## Authentication

### Token-based Authentication

If your Ollama server requires authentication:

```bash
# Set token via environment variable
export LANGCHAIN_CHAT_TOKEN="your-auth-token"

# Or via configuration file
autocmdr -init -t "your-auth-token"
```

### Network Security

For production deployments:

- Use HTTPS URLs when possible
- Store tokens securely (environment variables preferred)
- Restrict network access to Ollama server
- Consider using reverse proxy with authentication

## Logging Configuration

### Log Levels

| Level | Description |
|-------|-------------|
| `debug` | Detailed debugging information |
| `info` | General information messages |
| `warn` | Warning messages |
| `error` | Error messages only |

### Log Format

The application uses structured logging with the following fields:

- `timestamp` - ISO 8601 timestamp
- `level` - Log level
- `message` - Log message
- `component` - Component name
- Additional context fields

### Example Log Output

```json
{
  "timestamp": "2024-12-23T10:30:45.123Z",
  "level": "info",
  "message": "LLM initialized successfully",
  "component": "autocmdr-app",
  "model": "qwen3:14b",
  "server_url": "http://localhost:11434"
}
```

## Validation

The application validates configuration on startup:

### Required Fields

- `model` - Cannot be empty
- `server_url` - Must be a valid URL

### Optional Fields

- `token` - Can be empty for servers without authentication
- `log_level` - Defaults to "info" if invalid
- `config_dir` - Defaults to standard location if invalid

## Troubleshooting

### Common Issues

1. **Configuration file not found**:
   - Run `autocmdr -init` to create default configuration
   - Check file permissions

2. **Invalid server URL**:
   - Verify Ollama server is running
   - Check network connectivity
   - Validate URL format

3. **Model not available**:
   - Pull the model: `ollama pull model-name`
   - Check model name spelling
   - Verify server has the model

4. **Permission denied**:
   - Check configuration directory permissions
   - Ensure user has write access

### Debug Mode

Enable debug logging for troubleshooting:

```bash
autocmdr --log-level debug
```

This will show detailed information about:
- Configuration loading process
- Network requests
- Model interactions
- Error details

## Best Practices

1. **Security**:
   - Use environment variables for sensitive data
   - Don't commit configuration files with tokens
   - Use HTTPS in production

2. **Performance**:
   - Choose appropriate model size for your hardware
   - Configure memory limits if needed
   - Monitor resource usage

3. **Maintenance**:
   - Regularly update models
   - Monitor log files
   - Backup configuration

4. **Development**:
   - Use separate configurations for development/production
   - Test configuration changes in safe environment
   - Document custom configurations
