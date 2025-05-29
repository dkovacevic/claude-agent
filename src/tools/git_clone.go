package tools

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

// GitCloneDefinition defines a tool that clones a public Git repo.
var GitCloneDefinition = ToolDefinition{
	Name:        "git_clone",
	Description: "Clone a public Git repository. Provide `repo_url` and `dest_path`.",
	InputSchema: GitCloneInputSchema,
	Function:    GitClone,
}

type GitCloneInput struct {
	RepoURL  string `json:"repo_url" jsonschema_description:"URL of the public Git repository to clone."`
	DestPath string `json:"dest_path" jsonschema_description:"Local directory path where the repo will be cloned."`
}

var GitCloneInputSchema = GenerateSchema[GitCloneInput]()

// GitClone runs `git clone <repo_url> <dest_path>` and returns the command output.
func GitClone(input json.RawMessage) (string, error) {
	var in GitCloneInput
	if err := json.Unmarshal(input, &in); err != nil {
		return "", fmt.Errorf("invalid input: %w", err)
	}
	if in.RepoURL == "" || in.DestPath == "" {
		return "", fmt.Errorf("both repo_url and dest_path must be provided")
	}

	cmd := exec.Command("git", "clone", in.RepoURL, in.DestPath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("git clone failed: %s\n%s", err, string(out))
	}
	return string(out), nil
}
