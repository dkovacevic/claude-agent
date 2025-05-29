# Tool Framework

![Component](https://img.shields.io/badge/Component-Framework-purple)

## Overview

The tool framework enables Claude to interact with the file system and perform operations beyond basic conversation. This framework provides a standardized way to define, register, and execute tools with proper parameter validation and error handling.

## Core Components

The tool system consists of two main components:

### 1. Tool Definition Interface (tools.go)

```go
type ToolDefinition struct {
    Name        string
    Description string
    InputSchema anthropic.ToolInputSchemaParam
    Function    func(input json.RawMessage) (string, error)
}
```

This structure encapsulates everything needed to define a tool:
- **Name**: The unique identifier for the tool
- **Description**: A human-readable description of the tool's purpose
- **InputSchema**: JSON schema defining the tool's parameters
- **Function**: The actual implementation that executes the tool's logic

### 2. Schema Generation (schema.go)

```go
func GenerateSchema[T any]() anthropic.ToolInputSchemaParam
```

This generic function automatically generates JSON schema for tool parameters from Go struct types, using:
- Struct tags for parameter descriptions
- Type information for validation constraints
- The jsonschema reflection library for schema generation

## Tool Implementation Pattern

All tools in the system follow a consistent implementation pattern:

```go
// 1. Define the tool's input structure
type SomeToolInput struct {
    Param1 string `json:"param1" jsonschema_description:"Description of param1"`
    Param2 int    `json:"param2" jsonschema_description:"Description of param2"`
}

// 2. Generate the input schema
var SomeToolInputSchema = GenerateSchema[SomeToolInput]()

// 3. Create the tool definition
var SomeToolDefinition = ToolDefinition{
    Name:        "some_tool",
    Description: "Description of what the tool does",
    InputSchema: SomeToolInputSchema,
    Function:    SomeTool,
}

// 4. Implement the tool function
func SomeTool(input json.RawMessage) (string, error) {
    // Parse input
    var in SomeToolInput
    if err := json.Unmarshal(input, &in); err != nil {
        return "", err
    }
    
    // Validate input (if needed)
    // ...
    
    // Execute tool logic
    // ...
    
    // Return result or error
    return result, nil
}
```

## Key Features

### Standardized Error Handling

All tools follow a consistent error handling pattern:
1. Return descriptive error messages with context
2. Use proper Go error wrapping for nested errors
3. Validate parameters before execution

### Automatic Parameter Validation

The JSON schema system provides automatic validation of:
- Required vs. optional parameters
- Parameter types and constraints
- Structured description of parameters

### Clean Extensibility

Adding new tools is straightforward:
1. Create a new file following the pattern above
2. Implement the tool's logic
3. Register the tool in main.go

No changes to the core agent or framework are required.

## Available Tools

The framework includes several built-in tools:

| Tool | File | Purpose |
|------|------|---------|
| `read_file` | read_file.go | Read file contents |
| `list_files` | list_files.go | List files in a directory |
| `edit_file` | edit_file.go | Modify file contents |
| `create_dir` | create_dir.go | Create directories |
| `create_file` | create_file.go | Create new files |
| `git_clone` | git_clone.go | Clone git repositories |

For detailed documentation on each tool, see their respective documentation files.