# Create File Tool (create_file.go)

![Tool](https://img.shields.io/badge/Tool-File%20System-green)

## Overview

The `create_file` tool allows Claude to create new files with specified content, automatically creating any necessary parent directories. This enables Claude to assist with code generation, configuration file creation, and documentation writing.

## Implementation

### Tool Definition

```go
var CreateFileDefinition = ToolDefinition{
    Name:        "create_file",
    Description: "Create a file at the specified path with provided content, creating parent directories if necessary.",
    InputSchema: CreateFileInputSchema,
    Function:    CreateFile,
}
```

### Input Parameters

```go
type CreateFileInput struct {
    Path    string `json:"path" jsonschema_description:"The file path to create. Mandatory"`
    Content string `json:"content" jsonschema_description:"The content to be written into the file. Mandatory, cannot be empty string"`
}
```

The tool accepts two parameters:
- **path**: The file path where the file should be created
- **content**: The text content to write into the file

### Function Implementation

```go
func CreateFile(input json.RawMessage) (string, error) {
    var in CreateFileInput
    if err := json.Unmarshal(input, &in); err != nil {
        return "", fmt.Errorf("invalid input: %w", err)
    }
    if in.Path == "" {
        return "", fmt.Errorf("path must be provided")
    }
    if in.Content == "" {
        return "", fmt.Errorf("content must be provided")
    }

    // Ensure parent directory exists
    dir := path.Dir(in.Path)
    if dir != "." {
        if err := os.MkdirAll(dir, 0755); err != nil {
            return "", fmt.Errorf("failed to create directory %q: %w", dir, err)
        }
    }

    // Write the file
    if err := os.WriteFile(in.Path, []byte(in.Content), 0644); err != nil {
        return "", fmt.Errorf("write failed: %w", err)
    }
    return fmt.Sprintf("Created %s", in.Path), nil
}
```

The implementation:
1. Parses and validates the input parameters
2. Creates any necessary parent directories
3. Writes the content to the specified file
4. Returns a success message or an error

## Key Features

- **Automatic Directory Creation**: Creates parent directories as needed
- **Input Validation**: Ensures both path and content are provided
- **Standard Permissions**: Uses 0644 permissions (rw-r--r--) for created files
- **Descriptive Error Messages**: Provides clear error context for troubleshooting

## Usage Example

When Claude needs to create a new file, it might use this tool as follows:

```
Claude: I'll create a simple Hello World program in Go for you.

tool: create_file({"path": "hello.go", "content": "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello, World!\")\n}\n"})
Created hello.go

Claude: I've created a Hello World program in `hello.go`. You can run it with `go run hello.go`.
```

## Error Handling

The tool handles several error cases:
- Invalid JSON input format
- Missing required parameters (path or content)
- Permission errors during directory creation
- Permission errors during file writing
- Other file system errors

## Related Tools

- **create_dir**: For creating directory structures separately
- **edit_file**: For modifying existing files
- **read_file**: For verifying the content of created files