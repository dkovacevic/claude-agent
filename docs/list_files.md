# List Files Tool (list_files.go)

![Tool](https://img.shields.io/badge/Tool-File%20System-green)

## Overview

The `list_files` tool enables Claude to discover files and directories within the file system. This provides Claude with awareness of the available files and directories, allowing it to navigate the file structure and recommend relevant files for reading or editing.

## Implementation

### Tool Definition

```go
var ListFilesDefinition = ToolDefinition{
    Name:        "list_files",
    Description: "List files and directories at a given path. If no path is provided, lists files in the current directory.",
    InputSchema: ListFilesInputSchema,
    Function:    ListFiles,
}
```

### Input Parameters

```go
type ListFilesInput struct {
    Path string `json:"path,omitempty" jsonschema_description:"Optional path to list files from."`
}
```

The tool accepts a single optional parameter:
- **path**: The directory path to list (defaults to current directory if omitted)

### Function Implementation

```go
func ListFiles(input json.RawMessage) (string, error) {
    var in ListFilesInput
    if err := json.Unmarshal(input, &in); err != nil {
        return "", err
    }
    dir := in.Path
    if dir == "" {
        dir = "."
    }

    var files []string
    err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        rel, err := filepath.Rel(dir, path)
        if err != nil {
            return err
        }
        if rel != "." {
            if info.IsDir() {
                files = append(files, rel+"/")
            } else {
                files = append(files, rel)
            }
        }
        return nil
    })
    if err != nil {
        return "", err
    }
    out, err := json.Marshal(files)
    return string(out), err
}
```

The implementation:
1. Parses the input parameters from JSON
2. Defaults to the current directory if no path is provided
3. Uses `filepath.Walk` to recursively traverse the directory structure
4. Builds a list of files and directories, marking directories with a trailing slash
5. Returns the list as a JSON array

## Key Features

- **Directory Indication**: Directories are marked with a trailing slash for easy identification
- **Relative Paths**: All paths are relative to the requested directory
- **Recursive Listing**: Shows the complete directory structure, not just the top level
- **Error Handling**: Properly handles permission errors and other file system issues

## Usage Example

When Claude needs to discover files, it might use this tool as follows:

```
Claude: Let me check what files are in the current directory.

tool: list_files({})
["README.md", "go.mod", "go.sum", "src/", "docs/"]

Claude: I can see the following files and directories:
- README.md
- go.mod
- go.sum
- src/ (directory)
- docs/ (directory)

Would you like me to examine any of these files or list the contents of one of the directories?
```

## Related Tools

- **read_file**: For reading the contents of files discovered via list_files
- **edit_file**: For modifying files discovered via list_files