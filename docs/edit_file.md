# Edit File Tool 📝

## Overview

The `edit_file.go` module implements a tool that allows Claude to modify file contents by replacing text. This tool is particularly powerful as it enables Claude to create new files or update existing ones, making it possible for the AI to generate code, documentation, or configuration files directly on the user's system.

## Tool Definition

```go
var EditFileDefinition = ToolDefinition{
    Name:        "edit_file",
    Description: "Replace occurrences of old_str with new_str in a file, creating it if necessary.",
    InputSchema: EditFileInputSchema,
    Function:    EditFile,
}
```

## Input Schema

```go
type EditFileInput struct {
    Path   string `json:"path" jsonschema_description:"Path to the file."`
    OldStr string `json:"old_str" jsonschema_description:"Text to replace."`
    NewStr string `json:"new_str" jsonschema_description:"Replacement text."`
}
```

| Parameter | Description |
|-----------|-------------|
| `Path`    | The path to the file to edit or create |
| `OldStr`  | The text to find and replace in the file |
| `NewStr`  | The new text to insert in place of `OldStr` |

## Functionality

The `EditFile` function offers two primary modes of operation:

1. **Edit existing file**: Replace occurrences of `OldStr` with `NewStr` in an existing file
2. **Create new file**: When `OldStr` is empty and the file doesn't exist, create a new file with `NewStr` as content

## Implementation Details

### File Editing Logic

The function reads the entire file into memory, performs a string replacement, and writes the updated content back to the file. This approach is simple and effective for files of moderate size.

### File Creation Logic

When creating a new file, the function:
1. Ensures the parent directory exists (creating it if necessary)
2. Writes the provided content to the file
3. Returns a success message

### Safety Features

The tool includes several safety checks:
- Validates that parameters are valid before proceeding
- Verifies the `old_str` was actually found in the file (when not empty)
- Returns appropriate error messages for different failure scenarios

## Example Usage

From Claude's perspective:

```
You: Can you create a hello.go file that prints "Hello, World!"?

Claude: I'll create a simple Go file that prints "Hello, World!" for you.

tool: edit_file({"path":"hello.go","old_str":"","new_str":"package main\n\nimport \"fmt\"\n\nfunc main() {\n    fmt.Println(\"Hello, World!\")\n}\n"})
Created hello.go

Claude: I've created a file named `hello.go` with a basic "Hello, World!" program. Here's what I wrote:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

You can run this program with:

```bash
go run hello.go
```

This should output `Hello, World!` to your terminal.
```

## Error Handling

The tool handles several potential error conditions:

| Error Condition | Handling |
|-----------------|----------|
| Invalid parameters | Returns error message |
| File not found (when editing) | Returns appropriate error |
| Directory creation failure | Returns detailed error message |
| File write failure | Returns error with context |
| Text not found | Returns specific error message |

## Performance Considerations

The current implementation reads the entire file into memory, which is efficient for small to medium-sized files. For very large files, a streaming approach would be more appropriate but is not implemented in the current version.