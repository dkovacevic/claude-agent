package main

import (
    "bufio"
    "context"
    "fmt"
    "os"

    "github.com/anthropics/anthropic-sdk-go"
)

func main() {
    client := anthropic.NewClient()

    scanner := bufio.NewScanner(os.Stdin)
    getUserMessage := func() (string, bool) {
        if !scanner.Scan() {
            return "", false
        }
        return scanner.Text(), true
    }

    tools := []ToolDefinition{
        ReadFileDefinition,
        ListFilesDefinition,
        EditFileDefinition,
    }
    agent := NewAgent(&client, getUserMessage, tools)

    if err := agent.Run(context.TODO()); err != nil {
        fmt.Printf("Error: %s\n", err)
    }
}
