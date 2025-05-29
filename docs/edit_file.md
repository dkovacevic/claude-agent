# Edit File Tool (edit_file.go)

![Tool](https://img.shields.io/badge/Tool-File%20System-green)

## Overview

The `edit_file` tool gives Claude the ability to modify file contents by replacing specific text with new content. This enables Claude to assist with code editing, configuration updates, and text file modifications.

## Implementation

### Tool Definition

```go
var EditFileDefinition = ToolDefinition{
    Name: "edit_file",
    Description: `Make edits to a text file.

Replaces 'old_str' with 'new_str' in the given file. 'old_str' and 'new_str' MUST be different from each other.

If the file specified with path doesn't exist, it will be created.
`,
    InputSchema: EditFileInputSchema,
    Function:    EditFile,
}
```

### Input Parameters

```go
type EditFileInput struct {
    Path   string `json:"path" jsonschema_description:"The path to the file"`
    OldStr string `json:"old_str" jsonschema_description:"Text to search for - must match exactly and must only have one match exactly"`
    NewStr string `json:"new_str" jsonschema_description:"Text to replace old_str with"`
}
```

The tool accepts three parameters:
- **path**: The path to the file to edit
- **old_str**: The text to be replaced
- **new_str**: The replacement text

### Function Implementation

The implementation:
1. Parses the input parameters from JSON
2. Validates that the parameters are valid
3. Reads the existing file content
4. Replaces the old string with the new string
5. Writes the updated content back to the file

Special handling is included for:
- Creating new files if the file doesn't exist and old_str is empty
- Validating that old_str and new_str are different
- Ensuring that old_str exists in the file content

## Key Features

- **String Replacement**: Performs exact string replacement within files
- **File Creation**: Can create new files if they don't exist
- **Parent Directory Creation**: Automatically creates parent directories as needed
- **Validation**: Ensures the replacement operation is valid before execution

## Usage Example

When Claude needs to modify a file, it might use this tool as follows:

```
Claude: I'll update the error message in main.go for you.

tool: edit_file({"path": "main.go", "old_str": "fmt.Printf(\"Error: %s\\n\", err)", "new_str": "fmt.Printf(\"Error occurred: %s\\n\", err)"})
OK

Claude: I've updated the error message in main.go. The message now reads "Error occurred: %s" instead of "Error: %s".
```

## Error Handling

The tool handles several error cases:
- Invalid parameters (empty path, identical old_str and new_str)
- File not found errors (unless creating a new file)
- Write permission errors
- Parent directory creation errors
- String not found errors

## Related Tools

- **read_file**: Often used before edit_file to examine content before modifying
- **list_files**: For discovering files that might need editing
- **create_file**: An alternative for creating new files with content