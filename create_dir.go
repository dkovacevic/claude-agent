package main

import (
    "encoding/json"
    "fmt"
    "os"
)

// CreateDirDefinition defines a tool that makes a directory (and parents) at the given path.
var CreateDirDefinition = ToolDefinition{
    Name:        "create_dir",
    Description: "Create a directory (including any necessary parents) at the specified path.",
    InputSchema: CreateDirInputSchema,
    Function:    CreateDir,
}

type CreateDirInput struct {
    Path string `json:"path" jsonschema_description:"The directory path to create (can be nested)."`
}

var CreateDirInputSchema = GenerateSchema[CreateDirInput]()

// CreateDir creates the directory (and any parents) and returns a confirmation message.
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
