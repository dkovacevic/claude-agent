package tools

import (
    "encoding/json"
    "fmt"
    "os/exec"
    "strings"
)

var GitPatchDefinition = ToolDefinition{
    Name:        "git_patch",
    Description: "Apply a unified diff patch to a Git repository. Provide `repo_path` and `patch` (the diff text).",
    InputSchema: GenerateSchema[GitPatchInput](),
    Function:    GitPatch,
}

type GitPatchInput struct {
    RepoPath string `json:"repo_path" jsonschema_description:"Local path of the Git repository to which the patch will be applied."`
    Patch    string `json:"patch" jsonschema_description:"Unified diff text to apply via git apply."`
}

var GitPatchInputSchema = GenerateSchema[GitPatchInput]()

func GitPatch(input json.RawMessage) (string, error) {
    var in GitPatchInput
    if err := json.Unmarshal(input, &in); err != nil {
        return "", fmt.Errorf("invalid input: %w", err)
    }
    if in.RepoPath == "" {
        return "", fmt.Errorf("repo_path must be provided")
    }
    if in.Patch == "" {
        return "", fmt.Errorf("patch content must be provided")
    }

    // Run: git -C <repo_path> apply -
    cmd := exec.Command("git", "-C", in.RepoPath, "apply", "-")
    cmd.Stdin = strings.NewReader(in.Patch)
    out, err := cmd.CombinedOutput()
    if err != nil {
        return "", fmt.Errorf("git apply failed: %s\n%s", err, string(out))
    }
    return string(out), nil
}
