# Tools Interface 🧰

## Overview

The `tools.go` file defines the core interface for all tools that can be used by Claude within the application. It provides a standardized structure that allows for consistent handling of tool definitions, execution, and schema generation.

## ToolDefinition Struct

```go
type ToolDefinition struct {
    Name        string
    Description string
    InputSchema anthropic.ToolInputSchemaParam
    Function    func(input json.RawMessage) (string, error)
}
```

| Field | Description |
|-------|-------------|
| `Name` | The name of the tool (used by Claude to invoke it) |
| `Description` | Human-readable description of what the tool does |
| `InputSchema` | JSON schema describing the tool's input parameters |
| `Function` | The Go function that implements the tool's functionality |

## Design Philosophy

The `ToolDefinition` struct embodies several key design principles:

1. **Standardization**: All tools follow the same interface pattern
2. **Self-documentation**: Each tool carries its own description and schema
3. **Encapsulation**: Implementation details are hidden behind a consistent interface
4. **Simplicity**: The interface is minimal but complete

## Function Signature

```go
func(input json.RawMessage) (string, error)
```

Every tool function:
- Accepts a JSON input as a `json.RawMessage`
- Returns a string result or an error
- Is responsible for unmarshaling its own input
- Must handle its own error cases

## Integration with Anthropic API

The `ToolDefinition` structure is designed to work seamlessly with the Anthropic API:

- The `Name` field corresponds to the tool name Claude will use
- The `Description` field helps Claude understand when to use the tool
- The `InputSchema` provides type information for Claude's tool calls
- The `Function` implements the actual behavior invoked when Claude uses the tool

## Example Definition

```go
var ExampleToolDefinition = ToolDefinition{
    Name:        "example_tool",
    Description: "A simple example tool that does something useful.",
    InputSchema: ExampleToolInputSchema,
    Function:    ExampleToolFunction,
}
```

## How Tools Are Registered

Tools are registered in the `main.go` file by adding them to the tools slice passed to the Agent constructor:

```go
tools := []ToolDefinition{
    ReadFileDefinition,
    ListFilesDefinition,
    EditFileDefinition,
    // Add new tools here
}
```

## Creating New Tools

To create a new tool, you need to:

1. Define an input struct with appropriate JSON tags
2. Generate a schema for the input struct using `GenerateSchema`
3. Implement the tool function with the standard signature
4. Create a `ToolDefinition` instance with the appropriate fields
5. Add the tool to the list in `main.go`

## Benefits of the Design

This design provides several advantages:

1. **Extensibility**: New tools can be added without changing the core architecture
2. **Separation of Concerns**: Each tool is self-contained in its own file
3. **Consistency**: All tools follow the same pattern
4. **Ease of Testing**: Tools can be tested independently