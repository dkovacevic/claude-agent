# Claude Agent CLI: Go Project Documentation

![Claude Agent CLI](https://img.shields.io/badge/Claude-AI%20Assistant-5A67D8)
![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8)
![License](https://img.shields.io/badge/License-MIT-green)

## Project Overview

Claude Agent CLI is an elegant, powerful command-line application that facilitates direct interaction with Anthropic's Claude AI assistant. This project seamlessly integrates Claude's advanced AI capabilities with a terminal interface, allowing users to engage in natural language conversations while providing Claude with the ability to interact with the local file system through specialized tools.

## Core Architecture

The project follows a clean, modular architecture built around several key components:

```
src/
├── main.go             # Application entry point and initialization
├── agent.go            # Core agent implementation handling Claude interaction
└── tools/              # Directory containing all tool implementations
    ├── tools.go        # Core tool definition interface
    ├── schema.go       # JSON schema generator for tool parameters
    ├── read_file.go    # Tool to read file contents
    ├── list_files.go   # Tool to list files and directories
    ├── edit_file.go    # Tool to modify file contents
    ├── create_dir.go   # Tool to create directories
    ├── create_file.go  # Tool to create new files
    └── git_clone.go    # Tool to clone git repositories
```

## Key Components

### Agent Framework

The heart of the application is the agent framework which:
- Maintains a conversation with the user through the terminal
- Sends user messages to Claude via the Anthropic API
- Processes Claude's responses, particularly handling tool use
- Executes tools on behalf of Claude and returns results
- Manages the conversation state and flow

### Tool System

The tool system provides Claude with the ability to perform actions beyond conversation:
- Standardized tool definition interface
- Automatic JSON schema generation for tool parameters
- Clean separation between tool definition and implementation
- Easily extensible to add new capabilities

## Available Tools

The project includes several powerful tools that Claude can use:

| Tool | Description |
|------|-------------|
| `read_file` | Reads the contents of a file at a specified path |
| `list_files` | Lists all files and directories at a given path |
| `edit_file` | Modifies file contents by replacing text |
| `create_dir` | Creates a directory including any necessary parent directories |
| `create_file` | Creates a new file with specified content |
| `git_clone` | Clones a public Git repository to a local directory |

## Technical Implementation

The project is implemented in Go and leverages several key technologies:

- **Go 1.24+**: Modern Go features for clean, efficient code
- **Anthropic API**: Direct integration with Claude via the official Go SDK
- **JSON Schema**: Structured parameter definitions for tools
- **Terminal I/O**: Clean terminal interface for user interaction

## Getting Started

To run the project:

1. Ensure you have Go 1.24+ installed
2. Set your Anthropic API key in the environment as `ANTHROPIC_API_KEY`
3. Clone the repository
4. Run with `go run ./src` or build with `go build -o agent ./src`

## Further Reading

For detailed information about specific components, please refer to the following documentation:

- [Main Application](docs/main.md) - Entry point and initialization
- [Agent Implementation](docs/agent.md) - Core conversation handling
- [Tool Framework](docs/tools.md) - Tool system architecture
