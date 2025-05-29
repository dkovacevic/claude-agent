# Read File Tool (read_file.go)

![Tool](https://img.shields.io/badge/Tool-File%20System-green)

## Overview

The `read_file` tool provides Claude with the ability to read the contents of files from the local file system. This enables Claude to analyze code, view configuration files, or read any text file to provide more contextual assistance.

## Implementation

### Tool Definition

```go
var ReadFileDefinition = ToolDefinition{
    Name:        "read_file",
    Description: "Read the contents of a given relative file path. Use this when you want to see what's inside a file. Do not use this with directory names.",
    InputSchema: ReadFileInputSchema,
    Function:    ReadFile,
}
```

### Input Parameters

```go
type ReadFileInput struct {
    Path string `json:"path" jsonschema_description:"Relative path of a file."`
}
```

The tool accepts a single parameter:
- **path**: The relative path to the file that should be read

### Function Implementation

```go
func ReadFile(input json.RawMessage) (string, error) {
    var in ReadFileInput
    if err := json.Unmarshal(input, &in); err != nil {
        return "", err
    }
    data, err := os.ReadFile(in.Path)
    if err != nil {
        return "", err
    }
    return string(data), nil
}
```

The implementation:
1. Parses the input parameters from JSON
2. Uses `os.ReadFile` to read the file contents
3. Returns the file contents as a string or an error

## Usage Example

When Claude needs to read a file, it might use this tool as follows:

```
Claude: I'll check the contents of the README.md file for you.

tool: read_file({"path": "README.md"})
# Claude Agent CLI 🤖

![Claude Agent CLI](https://img.shields.io/badge/Claude-AI%20Assistant-5A67D8)
![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8)
![License](https://img.shields.io/badge/License-MIT-green)

## Overview

Claude Agent CLI is a powerful command-line interface...
```

## Security Considerations

This tool:
- Provides access to any file readable by the process
- Should be used with appropriate caution regarding sensitive files
- Returns clear error messages if the file cannot be read (e.g., due to permissions or non-existence)

## Error Handling

The tool handles several error conditions:
- JSON parsing errors for invalid input
- File not found errors
- Permission errors
- Other I/O errors that might occur during reading

All errors are passed through to Claude to allow appropriate response generation.

## Related Tools

- **list_files**: Complementary tool for discovering files that can be read
- **edit_file**: For modifying file contents after reading them