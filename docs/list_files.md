# List Files Tool 📂

## Overview

The `list_files.go` module implements a tool that allows Claude to enumerate files and directories within a specified path. This tool provides Claude with "eyes" into the file system, enabling it to make informed decisions about file operations and provide contextual assistance to users.

## Tool Definition

```go
var ListFilesDefinition = ToolDefinition{
    Name:        "list_files",
    Description: "List files and directories at a given path. Defaults to current directory.",
    InputSchema: ListFilesInputSchema,
    Function:    ListFiles,
}
```

## Input Schema

```go
type ListFilesInput struct {
    Path string `json:"path,omitempty" jsonschema_description:"Optional path to list files from."`
}
```

| Parameter | Description |
|-----------|-------------|
| `Path`    | The directory path to list files from (optional, defaults to current directory) |

## Functionality

The `ListFiles` function:

1. Accepts an optional path parameter (defaults to current directory if not specified)
2. Recursively walks through the directory structure
3. Builds a comprehensive list of all files and directories
4. Returns the list as a JSON array string

## Implementation Details

### Recursive Directory Traversal

The tool uses Go's `filepath.Walk` function to recursively traverse the directory structure, which offers several advantages:

- Handles nested directories automatically
- Follows symbolic links correctly
- Manages error cases gracefully

### Path Formatting

The function formats the output to:
- Make paths relative to the requested directory
- Add trailing slashes to directories for easy identification
- Exclude the root directory itself from the results

### JSON Output

The results are returned as a JSON array of strings, making it easy for Claude to parse and work with the data.

## Example Usage

From Claude's perspective:

```
You: What files are in the current directory?

Claude: Let me check what files are in the current directory.

tool: list_files({})
[".git/", ".idea/", "agent", "agent.go", "docs/", "edit_file.go", "go.mod", "go.sum", "list_files.go", "main.go", "read_file.go", "schema.go", "tools.go"]

Claude: Here are the files and directories in the current location:

Files:
- agent.go
- edit_file.go
- go.mod
- go.sum
- list_files.go
- main.go
- read_file.go
- schema.go
- tools.go
- agent (executable)

Directories:
- .git/
- .idea/
- docs/

Is there anything specific you'd like to know about these files or directories?
```

## Error Handling

The tool handles several potential error scenarios:

| Error Scenario | Handling |
|----------------|----------|
| Invalid input JSON | Returns a descriptive error message |
| Directory not found | Returns appropriate error |
| Permission issues | Propagates the system error message |
| Path traversal errors | Forwards the specific error |

## Performance Considerations

For very large directories with many files, this tool will return all entries, which could potentially be a large response. In most typical usage scenarios, this is not a concern, but it's something to be aware of when working with directories containing thousands of files.