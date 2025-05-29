package main

import (
    "encoding/json"
    "fmt"
    "os"
    "path"
    "strings"
)

var EditFileDefinition = ToolDefinition{
    Name:        "edit_file",
    Description: "Replace occurrences of old_str with new_str in a file, creating it if necessary.",
    InputSchema: EditFileInputSchema,
    Function:    EditFile,
}

type EditFileInput struct {
    Path   string `json:"path" jsonschema_description:"Path to the file."`
    OldStr string `json:"old_str" jsonschema_description:"Text to replace."`
    NewStr string `json:"new_str" jsonschema_description:"Replacement text."`
}

var EditFileInputSchema = GenerateSchema[EditFileInput]()

func EditFile(input json.RawMessage) (string, error) {
    var in EditFileInput
    if err := json.Unmarshal(input, &in); err != nil {
        return "", err
    }
    if in.Path == "" || in.OldStr == in.NewStr {
        return "", fmt.Errorf("invalid input parameters")
    }

    data, err := os.ReadFile(in.Path)
    if err != nil {
        if os.IsNotExist(err) && in.OldStr == "" {
            return createNewFile(in.Path, in.NewStr)
        }
        return "", err
    }

    old := string(data)
    updated := strings.ReplaceAll(old, in.OldStr, in.NewStr)
    if old == updated && in.OldStr != "" {
        return "", fmt.Errorf("old_str not found")
    }

    if err := os.WriteFile(in.Path, []byte(updated), 0644); err != nil {
        return "", err
    }
    return "OK", nil
}

func createNewFile(filePath, content string) (string, error) {
    dir := path.Dir(filePath)
    if dir != "." {
        if err := os.MkdirAll(dir, 0755); err != nil {
            return "", fmt.Errorf("mkdir failed: %w", err)
        }
    }
    if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
        return "", fmt.Errorf("write failed: %w", err)
    }
    return fmt.Sprintf("Created %s", filePath), nil
}
