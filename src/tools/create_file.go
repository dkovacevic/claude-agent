package tools

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

// CreateFileDefinition defines a tool that creates a file with given content.
var CreateFileDefinition = ToolDefinition{
	Name:        "create_file",
	Description: "Create a file at the specified path with provided content, creating parent directories if necessary.",
	InputSchema: CreateFileInputSchema,
	Function:    CreateFile,
}

type CreateFileInput struct {
	Path    string `json:"path" jsonschema_description:"The file path to create. Mandatory"`
	Content string `json:"content" jsonschema_description:"The content to be written into the file. Mandatory, cannot be empty string"`
}

var CreateFileInputSchema = GenerateSchema[CreateFileInput]()

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
