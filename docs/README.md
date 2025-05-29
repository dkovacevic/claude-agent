# Claude Agent CLI 🤖

![Claude Agent CLI](https://img.shields.io/badge/Claude-AI%20Assistant-5A67D8)
![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8)
![License](https://img.shields.io/badge/License-MIT-green)

## Overview

Claude Agent CLI is a powerful command-line interface that enables direct interaction with Anthropic's Claude AI assistant. This application leverages Claude's capabilities through a seamless terminal experience, allowing users to:

- Engage in natural language conversations with Claude
- Use specialized tools for file operations directly from the chat interface
- Build sophisticated workflows combining AI reasoning with file system access

## Key Features

- **Interactive Chat Interface**: Engage with Claude directly in your terminal
- **File System Integration**: Allow Claude to read, list, and edit files on your system
- **Tool-augmented AI**: Claude can use specialized tools to perform actions beyond conversation
- **Extensible Architecture**: Add new tools and capabilities with minimal code changes

## Architecture

The application follows a clean, modular architecture:

```
├── agent.go         # Core agent implementation
├── tools.go         # Tool definition interface
├── main.go          # Application entry point
├── schema.go        # JSON schema generator
├── read_file.go     # File reading tool
├── list_files.go    # File listing tool
├── edit_file.go     # File editing tool
```

## Getting Started

### Prerequisites

- Go 1.24 or higher
- Anthropic API key set in your environment as `ANTHROPIC_API_KEY`

### Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/claude-agent-cli.git
cd claude-agent-cli

# Build the application
go build -o claude-cli

# Run the application
./claude-cli
```

### Usage

Once running, you can interact with Claude directly in your terminal:

```
You: Hello Claude, can you help me with something?
Claude: Hello! I'm Claude, and I'd be happy to help you. What can I assist you with today?

You: Can you show me what files are in this directory?
Claude: I'll check what files are in the current directory for you.

tool: list_files({})
[".git/", ".idea/", "agent", "agent.go", "docs/", "edit_file.go", "go.mod", "go.sum", "list_files.go", "main.go", "read_file.go", "schema.go", "tools.go"]

Claude: Here are the files and directories in the current location:

1. .git/ (directory)
2. .idea/ (directory)
3. agent (executable)
4. agent.go
5. docs/ (directory)
6. edit_file.go
7. go.mod
8. go.sum
9. list_files.go
10. main.go
11. read_file.go
12. schema.go
13. tools.go

Is there a specific file you'd like to examine or modify?
```

## Available Tools

The Claude Agent CLI comes with three built-in tools that Claude can use:

1. **read_file**: Reads the contents of a file
2. **list_files**: Lists files and directories at a given path
3. **edit_file**: Edits file contents by replacing text

See the individual tool documentation for more details.

## Extending with New Tools

The application is designed to be easily extensible. To add a new tool:

1. Create a new file following the pattern of existing tool implementations
2. Define the tool's input schema, function logic, and description
3. Add the tool to the list in `main.go`

## License

This project is licensed under the MIT License - see the LICENSE file for details.

---

*For more detailed information about each component, please refer to the specific documentation files in this directory.*