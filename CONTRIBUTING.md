# Contributing to LangChain Chat App

Thank you for your interest in contributing to LangChain Chat App! This document provides guidelines and information for contributors.

## 🤝 How to Contribute

### Reporting Issues

Before creating an issue, please:

1. **Search existing issues** to avoid duplicates
2. **Use the issue templates** provided
3. **Provide detailed information** including:
   - Operating system and version
   - Go version
   - Steps to reproduce
   - Expected vs actual behavior
   - Relevant logs or error messages

### Suggesting Features

We welcome feature suggestions! Please:

1. **Check existing feature requests** first
2. **Use the feature request template**
3. **Explain the use case** and why it would be valuable
4. **Consider the scope** - smaller, focused features are easier to implement

### Code Contributions

#### Getting Started

1. **Fork the repository**
2. **Clone your fork**:
   ```bash
   git clone https://github.com/your-username/autocmdr-app.git
   cd autocmdr-app
   ```
3. **Add the upstream remote**:
   ```bash
   git remote add upstream https://github.com/blysin/autocmdr.git
   ```

#### Development Workflow

1. **Create a feature branch**:
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Install dependencies**:
   ```bash
   make deps
   ```

3. **Make your changes**:
   - Follow the coding standards (see below)
   - Add tests for new functionality
   - Update documentation as needed

4. **Test your changes**:
   ```bash
   make test
   make lint
   ```

5. **Commit your changes**:
   ```bash
   git add .
   git commit -m "feat: add your feature description"
   ```

6. **Push to your fork**:
   ```bash
   git push origin feature/your-feature-name
   ```

7. **Create a Pull Request**

#### Pull Request Guidelines

- **Use the PR template** provided
- **Write clear, descriptive titles**
- **Reference related issues** using `Fixes #123` or `Closes #123`
- **Keep PRs focused** - one feature/fix per PR
- **Update documentation** if needed
- **Add tests** for new functionality
- **Ensure CI passes** before requesting review

## 📝 Coding Standards

### Go Style Guide

We follow the standard Go style guide with some additional conventions:

#### Code Formatting

- Use `gofmt` and `goimports` for formatting
- Run `make fmt` before committing
- Line length should not exceed 120 characters

#### Naming Conventions

- Use descriptive names for variables, functions, and types
- Follow Go naming conventions (camelCase for private, PascalCase for public)
- Use meaningful package names

#### Error Handling

- Always handle errors explicitly
- Use structured logging for error context
- Wrap errors with meaningful context using `fmt.Errorf`

#### Testing

- Write unit tests for all new functionality
- Use table-driven tests where appropriate
- Aim for >80% test coverage
- Use meaningful test names that describe the scenario

#### Documentation

- Add godoc comments for all public functions and types
- Keep comments concise but informative
- Update README.md for user-facing changes

### Project Structure

Follow the established project structure:

```
autocmdr-app/
├── cmd/                    # Application entry points
├── pkg/                    # Public library code
├── internal/               # Private application code
├── examples/               # Usage examples
├── docs/                   # Documentation
├── scripts/                # Build and utility scripts
└── .github/                # GitHub workflows and templates
```

## 🧪 Testing

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run benchmarks
make bench

# Run linting
make lint
```

### Writing Tests

- Place test files alongside the code they test
- Use the `_test.go` suffix
- Follow the `TestFunctionName` pattern
- Use `testify` for assertions when helpful

Example:
```go
func TestConfigLoad(t *testing.T) {
    tests := []struct {
        name    string
        setup   func()
        want    *Config
        wantErr bool
    }{
        // test cases
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // test implementation
        })
    }
}
```

## 📋 Commit Message Guidelines

We use conventional commits for clear and consistent commit messages:

### Format

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

### Types

- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code refactoring
- `test`: Adding or updating tests
- `chore`: Maintenance tasks

### Examples

```
feat: add configuration validation
fix: resolve memory leak in chat session
docs: update installation instructions
test: add unit tests for config package
```

## 🔄 Release Process

Releases are automated through GitHub Actions:

1. **Version tags** trigger the release workflow
2. **Semantic versioning** is used (v1.2.3)
3. **Release notes** are generated from commit messages
4. **Binaries** are built for multiple platforms

## 📞 Getting Help

- **GitHub Discussions**: For questions and general discussion
- **GitHub Issues**: For bug reports and feature requests
- **Email**: For security-related concerns

## 🏆 Recognition

Contributors are recognized in:

- Release notes
- README.md contributors section
- GitHub contributors graph

## 📄 License

By contributing, you agree that your contributions will be licensed under the MIT License.

---

Thank you for contributing to LangChain Chat App! 🚀
