# Main Application Entry Point 🚀

## Overview

The `main.go` file serves as the entry point for the Claude Agent CLI application. It initializes all necessary components, configures the agent, and starts the main interaction loop.

## Key Components

### Imports

```go
import (
    "bufio"
    "context"
    "fmt"
    "os"

    "github.com/anthropics/anthropic-sdk-go"
)
```

The application uses standard Go libraries for I/O operations and the official Anthropic SDK for Go to communicate with the Claude API.

### Main Function

```go
func main() {
    client := anthropic.NewClient()

    scanner := bufio.NewScanner(os.Stdin)
    getUserMessage := func() (string, bool) {
        if !scanner.Scan() {
            return "", false
        }
        return scanner.Text(), true
    }

    tools := []ToolDefinition{
        ReadFileDefinition,
        ListFilesDefinition,
        EditFileDefinition,
    }
    agent := NewAgent(&client, getUserMessage, tools)

    if err := agent.Run(context.TODO()); err != nil {
        fmt.Printf("Error: %s\n", err)
    }
}
```

## Initialization Process

The main function performs several important initialization steps:

1. **API Client**: Creates a new Anthropic API client using default configuration
2. **Input Handler**: Sets up a function to read user input from stdin
3. **Tools Registration**: Registers all available tools with the agent
4. **Agent Creation**: Instantiates the agent with the configured components
5. **Execution**: Starts the agent's main loop with a background context

## User Input Handling

The application uses a standard bufio Scanner to read user input from the command line:

```go
scanner := bufio.NewScanner(os.Stdin)
getUserMessage := func() (string, bool) {
    if !scanner.Scan() {
        return "", false
    }
    return scanner.Text(), true
}
```

This function returns the input text and a boolean indicating success. If reading fails (e.g., due to EOF), it returns `false` to signal termination.

## Available Tools

The main function registers three tools with the agent:

1. **ReadFileDefinition**: Allows reading file contents
2. **ListFilesDefinition**: Enables listing files in directories
3. **EditFileDefinition**: Permits creating or editing files

Additional tools can be added to this list to extend the agent's capabilities.

## Error Handling

The application implements simple but effective error handling:

```go
if err := agent.Run(context.TODO()); err != nil {
    fmt.Printf("Error: %s\n", err)
}
```

Any errors that occur during the agent's execution are caught and printed to the console before the application exits.

## Context Management

The application uses `context.TODO()` as a placeholder context. In a more sophisticated implementation, this could be replaced with a cancelable context to allow for graceful shutdown.

## Authentication

The Anthropic client is initialized with default settings, which means it will look for the `ANTHROPIC_API_KEY` environment variable to authenticate API requests.

## Extension Points

The main function can be extended in several ways:

1. **Additional Tools**: Register more tools in the tools slice
2. **Custom Input**: Replace the getUserMessage function with a custom implementation
3. **Configuration**: Add command-line flags or config file parsing
4. **Context Management**: Implement proper context handling with cancellation