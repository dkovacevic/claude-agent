package main

import (
    "bufio"
    "context"
    "fmt"
    "os"

    "github.com/anthropics/anthropic-sdk-go"
    "agent/src/tools"
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

    tools := []tools.ToolDefinition{
        tools.ReadFileDefinition,
        tools.ListFilesDefinition,
        tools.EditFileDefinition,
        tools.CreateDirDefinition,
        //tools.CreateFileDefinition,
        tools.AppendFileDefinition,
        tools.GitCloneDefinition,
        tools.GitPatchDefinition,
    }
    agent := NewAgent(&client, getUserMessage, tools)

    if err := agent.Run(context.TODO()); err != nil {
        fmt.Printf("Error: %s\n", err)
    }
}
