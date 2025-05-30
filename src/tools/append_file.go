package tools

import (
    "encoding/json"
    "fmt"
    "os"
    "path"
)

const MaxContentLength = 4000

// AppendFileDefinition defines a tool for appending content to a file.
var AppendFileDefinition = ToolDefinition{
	Name:        "append_file",
	Description: fmt.Sprintf(`Append content to a text file at the specified path. Content must not exceed %d characters and cannot be empty string.
	 If the file or its parent directories do not exist, they will be created.
	 Example: {"path": "tmp/example.txt", "content": "This is example content."}`, MaxContentLength),
	InputSchema: GenerateSchema[AppendFileInput](),
	Function:    AppendFile,
}

type AppendFileInput struct {
    Path    string `json:"path" jsonschema_description:"The file path to append to. Mandatory field"`
    Content string `json:"content" jsonschema_description:"The content to append. Mandatory field. Cannot be empty string"`
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

    if in.Content == "" {
        fmt.Println("\033[31mError: Content must be provided. JSON object must contain a non-empty 'content' field.\033[0m")
        return "", fmt.Errorf("content must be provided. JSON object must contain a non-empty 'content' field")
    }

    // Check content size against MaxContentLength
    if len(in.Content) > MaxContentLength {
        fmt.Printf("\033[31mError: Content exceeds %d characters. Got %d characters.\033[0m\n", MaxContentLength, len(in.Content))
        return "", fmt.Errorf("content exceeds %d character limit", MaxContentLength)
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
