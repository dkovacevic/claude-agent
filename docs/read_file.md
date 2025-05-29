# Read File Tool 📖

## Overview

The `read_file.go` module implements a tool that allows Claude to read the contents of files from the user's file system. This gives Claude the ability to analyze, interpret, and work with existing files, providing more contextual and useful assistance.

## Tool Definition

```go
var ReadFileDefinition = ToolDefinition{
    Name:        "read_file",
    Description: "Read the contents of a given relative file path. Use this when you want to see what's inside a file. Do not use this with directory names.",
    InputSchema: ReadFileInputSchema,
    Function:    ReadFile,
}
```

## Input Schema

```go
type ReadFileInput struct {
    Path string `json:"path" jsonschema_description:"Relative path of a file."`
}
```

| Parameter | Description |
|-----------|-------------|
| `Path`    | The relative path to the file to be read |

## Functionality

The `ReadFile` function:

1. Takes a file path as input
2. Reads the entire contents of the file into memory
3. Returns the file contents as a string

## Implementation Details

### File Reading Logic

The implementation uses Go's standard library `os.ReadFile` function, which:
- Reads the entire file into memory at once
- Handles file opening and closing automatically
- Provides appropriate error messages for common failure cases

### Safety Considerations

The tool description explicitly warns against using it with directory paths to prevent misuse.

## Example Usage

From Claude's perspective:

```
You: Can you show me what's in the go.mod file?

Claude: I'll read the go.mod file for you to see its contents.

tool: read_file({"path":"go.mod"})
module agent

go 1.24.1

require (
	github.com/anthropics/anthropic-sdk-go v1.2.1
	github.com/invopop/jsonschema v0.13.0
)

require (
	cloud.google.com/go/auth v0.7.2 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.3 // indirect
	cloud.google.com/go/compute/metadata v0.5.0 // indirect
	github.com/aws/aws-sdk-go-v2 v1.30.3 // indirect
	...
)

Claude: Here's what's in the go.mod file:

The file defines a Go module named "agent" using Go version 1.24.1.

It has two direct dependencies:
1. github.com/anthropics/anthropic-sdk-go v1.2.1 - The Anthropic SDK for Go
2. github.com/invopop/jsonschema v0.13.0 - A library for JSON schema generation

The file also lists several indirect dependencies that are required by these primary dependencies, including various AWS SDK components, Google Cloud libraries, and other utility packages.

This confirms the project is using the Anthropic API to communicate with Claude and using jsonschema for handling the tool input schemas.
```

## Error Handling

The tool handles several potential error scenarios:

| Error Scenario | Handling |
|----------------|----------|
| Invalid input JSON | Returns a descriptive error message |
| File not found | Returns appropriate error |
| Permission issues | Propagates the system error message |
| File read errors | Returns the specific error |

## Performance Considerations

The current implementation reads the entire file into memory, which is efficient for small to medium-sized files. For very large files, this could potentially cause memory issues, but for typical usage in a CLI environment, this approach is appropriate.