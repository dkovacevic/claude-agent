package tools

import (
    "encoding/json"
    "fmt"
    "os"
    "path"
)

var AppendFileDefinition = ToolDefinition{
    Name:        "append_file",
    Description: "Append content to a text file at the specified path. Content must not exceed 1000 characters. If the file or its parent directories do not exist, they will be created.",
    InputSchema: GenerateSchema[AppendFileInput](),
    Function:    AppendFile,
}

type AppendFileInput struct {
    Path    string `json:"path" jsonschema_description:"The file path to append to."`
    Content string `json:"content" jsonschema_description:"The content to append (max 1000 characters)."`
}

var AppendFileInputSchema = GenerateSchema[AppendFileInput]()

func AppendFile(input json.RawMessage) (string, error) {
    var in AppendFileInput
    if err := json.Unmarshal(input, &in); err != nil {
        return "", fmt.Errorf("invalid input: %w", err)
    }
    if in.Path == "" {
        return "", fmt.Errorf("path must be provided")
    }
    if len(in.Content) > 1000 {
        return "", fmt.Errorf("content exceeds 1000 character limit")
    }

    // Ensure parent directory exists
    dir := path.Dir(in.Path)
    if dir != "." {
        if err := os.MkdirAll(dir, 0755); err != nil {
            return "", fmt.Errorf("failed to create directories: %w", err)
        }
    }

    // Open file for appending (create if not exists)
    f, err := os.OpenFile(in.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        return "", fmt.Errorf("failed to open file: %w", err)
    }
    defer f.Close()

    // Append the content
    if _, err := f.WriteString(in.Content); err != nil {
        return "", fmt.Errorf("failed to append content: %w", err)
    }

    return fmt.Sprintf("Successfully appended to %s", in.Path), nil
}
