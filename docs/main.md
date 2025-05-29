# Main Application (main.go)

![Component](https://img.shields.io/badge/Component-Entry%20Point-blue)

## Overview

`main.go` serves as the entry point for the Claude Agent CLI application. It initializes the core components, sets up the required dependencies, and launches the agent to begin processing user interactions.

## Implementation Details

### Imports

```go
import (
    "bufio"
    "context"
    "fmt"
    "os"

    "github.com/anthropics/anthropic-sdk-go"
    "agent/src/tools"
)
```

The main package imports:
- Standard library packages for I/O operations and context management
- The Anthropic SDK for Claude integration
- The application's custom tools package

### Main Function

The `main()` function performs several key tasks:

1. **Initializes the Anthropic Client**: Creates a new client instance that will handle API communication with Claude.
   ```go
   client := anthropic.NewClient()
   ```

2. **Sets up User Input Handling**: Creates a function that reads user input from the terminal.
   ```go
   scanner := bufio.NewScanner(os.Stdin)
   getUserMessage := func() (string, bool) {
       if !scanner.Scan() {
           return "", false
       }
       return scanner.Text(), true
   }
   ```

3. **Configures Available Tools**: Defines the set of tools that Claude can use during conversations.
   ```go
   tools := []tools.ToolDefinition{
       tools.ReadFileDefinition,
       tools.ListFilesDefinition,
       tools.EditFileDefinition,
       tools.GitCloneDefinition,
       tools.CreateDirDefinition,
       tools.CreateFileDefinition,
   }
   ```

4. **Creates and Runs the Agent**: Instantiates the agent with the necessary dependencies and starts its execution.
   ```go
   agent := NewAgent(&client, getUserMessage, tools)
   if err := agent.Run(context.TODO()); err != nil {
       fmt.Printf("Error: %s\n", err)
   }
   ```

## Key Features

- **Simple, Clean Design**: The main function is concise and focused solely on initialization and setup.
- **Dependency Injection**: Core dependencies (client, input handling, tools) are cleanly passed to the agent.
- **Error Handling**: Proper error handling for agent execution.

## Extensibility

Adding new tools to the application is straightforward:
1. Implement the tool in the tools package
2. Add the tool definition to the tools list in `main.go`

No other changes to the main application flow are required.

## Dependencies

- **Agent**: Requires the `NewAgent` function from the main package to create the agent instance.
- **Tools**: Depends on tool definitions from the tools package.
- **Anthropic SDK**: Requires the Anthropic API client for Claude communication.

## Environment Requirements

The application expects:
- The `ANTHROPIC_API_KEY` environment variable to be set with a valid API key
- Standard input/output for terminal interaction