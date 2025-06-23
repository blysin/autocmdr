# AutoCmdr 命令行辅助应用

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/blysin/autocmdr)](https://goreportcard.com/report/github.com/blysin/autocmdr)
[![CI](https://github.com/blysin/autocmdr/workflows/CI/badge.svg)](https://github.com/blysin/autocmdr/actions)

[English](README.md) | [简体中文](README_zh.md)

一个使用 Go 和 LangChain 构建的命令行辅助应用，旨在帮助用户安全地与 AI 模型交互以生成和执行系统命令。

## 🚀 特性

- **跨平台支持**：支持 Windows (PowerShell)、Linux 和 macOS (Bash)
- **交互式命令行**：具有 readline 支持的丰富命令行界面
- **智能命令生成**：具有安全检查的 AI 驱动命令生成
- **配置管理**：支持文件和环境变量的灵活配置
- **结构化日志**：具有可配置级别的全面日志记录
- **记忆管理**：具有可配置窗口大小的对话历史
- **安全优先**：内置安全检查和确认提示
- **可扩展架构**：遵循 Go 最佳实践的模块化设计

## 📦 安装

### 从源码安装

```bash
# 克隆仓库
git clone https://github.com/blysin/autocmdr.git
cd autocmdr-app

# 构建和安装
make install
```

### 使用 Go Install

```bash
go install github.com/blysin/autocmdr/cmd/autocmdr@latest
```

### 预编译二进制文件

从[发布页面](https://github.com/blysin/autocmdr/releases)下载最新版本。

## 🔧 配置

### 初始化配置

```bash
autocmdr -init -m "your-model" -u "http://your-ollama-server:11434"
```

### 配置选项

应用程序支持通过以下方式进行配置：

1. **配置文件**：`~/.autocmdr/config.json`
2. **环境变量**：以 `LANGCHAIN_CHAT_` 为前缀
3. **命令行参数**

#### 配置参数

| 参数 | 环境变量 | 默认值 | 描述 |
|-----------|---------------------|---------|-------------|
| `model` | `LANGCHAIN_CHAT_MODEL` | `qwen3:14b` | AI 模型名称 |
| `server_url` | `LANGCHAIN_CHAT_SERVER_URL` | `http://localhost:11434` | Ollama 服务器 URL |
| `token` | `LANGCHAIN_CHAT_TOKEN` | `""` | API 认证令牌 |
| `log_level` | `LANGCHAIN_CHAT_LOG_LEVEL` | `info` | 日志级别 (debug, info, warn, error) |

### 配置文件示例

```json
{
  "model": "qwen3:14b",
  "server_url": "http://localhost:11434",
  "token": "",
  "log_level": "info"
}
```

## 🎯 使用方法

### 基本用法

```bash
# 启动交互式聊天
autocmdr

# 查看当前配置
autocmdr -view

# 查看系统提示词
autocmdr -prompt

# 显示版本信息
autocmdr -version
```

### 交互式命令

在聊天会话中：

- 用自然语言输入您的请求
- AI 将生成适当的命令
- 使用 `y` 或 `n` 确认执行
- 使用 `clear` 清除对话历史
- 使用 `exit` 退出应用程序

### 会话示例

```
You: 列出当前目录中的所有文件
Bot: 我将帮助您列出当前目录中的所有文件。

{
  "success": "true",
  "multipleLines": "false",
  "script": "Get-ChildItem -Force"
}

是否直接执行脚本？(y/n)
You: y
✅ 脚本执行成功（退出代码：0）
输出：
Directory: C:\Users\example

Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2023/12/01     10:30                Documents
d-----         2023/12/01     10:30                Downloads
...
```

## 🏗️ 开发

### 前置条件

- Go 1.21 或更高版本
- Make（可选，用于使用 Makefile）

### 从源码构建

```bash
# 克隆仓库
git clone https://github.com/blysin/autocmdr.git
cd autocmdr-app

# 安装依赖
make deps

# 运行测试
make test

# 构建应用
make build

# 运行应用
make run
```

### 项目结构

```
autocmdr-app/
├── cmd/
│   └── autocmdr/        # 应用程序入口点
├── pkg/
│   ├── config/          # 配置管理
│   ├── chat/            # 聊天功能
│   ├── prompts/         # 提示词模板和加载
│   └── utils/           # 工具函数
├── internal/
│   └── version/         # 版本信息
├── examples/            # 使用示例
├── docs/               # 文档
├── scripts/            # 构建和工具脚本
└── .github/            # GitHub 工作流和模板
```

### 可用的 Make 目标

```bash
make help                     # 显示所有可用目标
make build                    # 构建应用
make test                     # 运行测试
make test-coverage           # 运行测试并生成覆盖率报告
make lint                    # 运行代码检查
make fmt                     # 格式化代码
make clean                   # 清理构建产物
```

## 🧪 测试

```bash
# 运行所有测试
make test

# 运行测试并生成覆盖率报告
make test-coverage

# 运行基准测试
make bench
```

## 📚 文档

- [API 文档](docs/api.md)
- [配置指南](docs/configuration.md)
- [贡献指南](CONTRIBUTING.md)

## 🤝 贡献

我们欢迎贡献！详情请参阅我们的[贡献指南](CONTRIBUTING.md)。

### 开发工作流程

1. Fork 仓库
2. 创建功能分支
3. 进行修改
4. 为新功能添加测试
5. 运行测试套件
6. 提交拉取请求

## 📄 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件。

## 🙏 致谢

- [LangChain Go](https://github.com/tmc/langchaingo) - LangChain 的 Go 实现
- [Ollama](https://ollama.ai/) - 本地 AI 模型服务
- [Cobra](https://github.com/spf13/cobra) - CLI 框架
- [Viper](https://github.com/spf13/viper) - 配置管理

## 📞 支持

- 📧 电子邮件：blysin@163.com
- 🐛 问题：[GitHub Issues](https://github.com/blysin/autocmdr/issues)
- 💬 讨论：[GitHub Discussions](https://github.com/blysin/autocmdr/discussions)

---

Made with ❤️ by the AutoCmdr App team
