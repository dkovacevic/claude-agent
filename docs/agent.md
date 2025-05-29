# Agent Module 🧠

## Overview

The `agent.go` file contains the core logic for the Claude Agent CLI. It manages the conversation between the user and Claude, handles tool invocation, and orchestrates the application's main loop.

## Key Components

### Agent Struct

```go
type Agent struct {
    client         *anthropic.Client
    getUserMessage func() (string, bool)
    tools          []ToolDefinition
}
```

| Field | Description |
|-------|-------------|
| `client` | Anthropic API client for communicating with Claude |
| `getUserMessage` | Function that retrieves input from the user |
| `tools` | Slice of available tools that Claude can use |

### Constructor

```go
func NewAgent(
    client *anthropic.Client,
    getUserMessage func() (string, bool),
    tools []ToolDefinition,
) *Agent
```

Creates a new Agent instance with the specified Anthropic client, user input function, and available tools.

### Main Loop

The `Run` method implements the main conversation loop:

1. Get user input
2. Send the conversation to Claude
3. Process Claude's response:
   - Display text content to the user
   - Execute tools when requested
4. Add tool results to the conversation
5. Repeat

## Key Methods

### Run

```go
func (a *Agent) Run(ctx context.Context) error
```

The main loop of the agent that manages the conversation flow and tool execution.

### executeTool

```go
func (a *Agent) executeTool(id, name string, input json.RawMessage) anthropic.ContentBlockParamUnion
```

Executes a requested tool and returns its results in the format expected by the Anthropic API.

### runInference

```go
func (a *Agent) runInference(ctx context.Context, conversation []anthropic.MessageParam) (*anthropic.Message, error)
```

Makes an API call to Claude with the current conversation history and available tools.

## Terminal UI

The agent implements a simple but effective terminal UI with color-coded output:

- **Blue**: User messages
- **Yellow**: Claude's text responses
- **Green**: Tool executions

## Error Handling

The agent properly handles errors from:
- Tool execution
- API communication
- User input processing

## Design Patterns

This module follows several design patterns:

1. **Dependency Injection**: Dependencies are passed to the agent constructor
2. **Strategy Pattern**: The user input function is an interchangeable strategy
3. **Command Pattern**: Tools are encapsulated commands with standardized interfaces

## Usage Example

```go
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
```