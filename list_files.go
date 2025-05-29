package main

import (
    "encoding/json"
    "os"
    "path/filepath"
)

var ListFilesDefinition = ToolDefinition{
    Name:        "list_files",
	Description: "List files and directories at a given path. If no path is provided, lists files in the current directory.",
    InputSchema: ListFilesInputSchema,
    Function:    ListFiles,
}

type ListFilesInput struct {
    Path string `json:"path,omitempty" jsonschema_description:"Optional path to list files from."`
}

var ListFilesInputSchema = GenerateSchema[ListFilesInput]()

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
