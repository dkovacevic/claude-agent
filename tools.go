package main

import (
   "github.com/anthropics/anthropic-sdk-go"
   "encoding/json"
)

type ToolDefinition struct {
    Name        string
    Description string
    InputSchema anthropic.ToolInputSchemaParam
    Function    func(input json.RawMessage) (string, error)
}
