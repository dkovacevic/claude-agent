package tools

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
)

var EditFileDefinition = ToolDefinition{
	Name: "edit_file",
	Description: `Make edits to a text file.

Replaces 'old_str' with 'new_str' in the given file. 'old_str' and 'new_str' MUST be different from each other.

If the file specified with path doesn't exist, it will be created.
`,
	InputSchema: EditFileInputSchema,
	Function:    EditFile,
}

type EditFileInput struct {
	Path   string `json:"path" jsonschema_description:"The path to the file"`
	OldStr string `json:"old_str" jsonschema_description:"Text to search for - must match exactly and must only have one match exactly"`
	NewStr string `json:"new_str" jsonschema_description:"Text to replace old_str with"`
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
