package main

import (
    "encoding/json"
    "os"
)

var ReadFileDefinition = ToolDefinition{
    Name:        "read_file",
    Description: "Read the contents of a given relative file path. Use this when you want to see what's inside a file. Do not use this with directory names.",
    InputSchema: ReadFileInputSchema,
    Function:    ReadFile,
}

type ReadFileInput struct {
    Path string `json:"path" jsonschema_description:"Relative path of a file."`
}

var ReadFileInputSchema = GenerateSchema[ReadFileInput]()

func ReadFile(input json.RawMessage) (string, error) {
    var in ReadFileInput
    if err := json.Unmarshal(input, &in); err != nil {
        return "", err
    }
    data, err := os.ReadFile(in.Path)
    if err != nil {
        return "", err
    }
    return string(data), nil
}
