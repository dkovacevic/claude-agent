# Create Directory Tool (create_dir.go)

![Tool](https://img.shields.io/badge/Tool-File%20System-green)

## Overview

The `create_dir` tool enables Claude to create new directories in the file system, including any necessary parent directories. This tool is essential for organizing file structures and preparing directory paths before creating files.

## Implementation

### Tool Definition

```go
var CreateDirDefinition = ToolDefinition{
    Name:        "create_dir",
    Description: "Create a directory (including any necessary parents) at the specified path.",
    InputSchema: CreateDirInputSchema,
    Function:    CreateDir,
}
```

### Input Parameters

```go
type CreateDirInput struct {
    Path string `json:"path" jsonschema_description:"The directory path to create (can be nested)."`
}
```

The tool accepts a single parameter:
- **path**: The directory path to create, which can include multiple levels of nesting

### Function Implementation

```go
func CreateDir(input json.RawMessage) (string, error) {
    var in CreateDirInput
    if err := json.Unmarshal(input, &in); err != nil {
        return "", fmt.Errorf("invalid input: %w", err)
    }
    if in.Path == "" {
        return "", fmt.Errorf("path must be provided")
    }
    if err := os.MkdirAll(in.Path, 0755); err != nil {
        return "", fmt.Errorf("failed to create directory: %w", err)
    }
    return fmt.Sprintf("Successfully created directory %s", in.Path), nil
}
```

The implementation:
1. Parses the input parameters from JSON
2. Validates that a path is provided
3. Uses `os.MkdirAll` to create the directory and any necessary parent directories
4. Returns a success message or an error

## Key Features

- **Recursive Creation**: Creates all necessary parent directories in a single operation
- **Standard Permissions**: Uses 0755 permissions (rwxr-xr-x) for created directories
- **Idempotent Operation**: Succeeds even if the directory already exists
- **Descriptive Error Messages**: Provides clear error messages for troubleshooting

## Usage Example

When Claude needs to create a directory structure, it might use this tool as follows:

```
Claude: I'll create the necessary directory structure for your project.

tool: create_dir({"path": "src/utils/helpers"})
Successfully created directory src/utils/helpers

Claude: I've created the directory structure at `src/utils/helpers`. Now we can proceed with creating the utility files in this location.
```

## Error Handling

The tool handles several error conditions:
- Invalid input JSON format
- Empty path parameter
- Permission errors when creating directories
- Other file system errors that might occur

## Related Tools

- **create_file**: Often used after create_dir to populate the newly created directories
- **list_files**: For verifying the directory structure after creation