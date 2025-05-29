package main

import (
    "context"
    "encoding/json"
    "fmt"

    "github.com/anthropics/anthropic-sdk-go"
)

type Agent struct {
    client         *anthropic.Client
    getUserMessage func() (string, bool)
    tools          []ToolDefinition
}

func NewAgent(
    client *anthropic.Client,
    getUserMessage func() (string, bool),
    tools []ToolDefinition,
) *Agent {
    return &Agent{client, getUserMessage, tools}
}

func (a *Agent) Run(ctx context.Context) error {
    conversation := []anthropic.MessageParam{}
    fmt.Println("Chat with Claude (use 'ctrl-c' to quit)")

    readUserInput := true
    for {
        if readUserInput {
            fmt.Print("\u001b[94mYou\u001b[0m: ")
            input, ok := a.getUserMessage()
            if !ok {
                break
            }
            conversation = append(conversation, anthropic.NewUserMessage(
                anthropic.NewTextBlock(input),
            ))
        }

        msg, err := a.runInference(ctx, conversation)
        if err != nil {
            return err
        }
        conversation = append(conversation, msg.ToParam())

        var toolResults []anthropic.ContentBlockParamUnion
        for _, c := range msg.Content {
            switch c.Type {
            case "text":
                fmt.Printf("\u001b[93mClaude\u001b[0m: %s\n", c.Text)
            case "tool_use":
                result := a.executeTool(c.ID, c.Name, c.Input)
                toolResults = append(toolResults, result)
            }
        }

        if len(toolResults) == 0 {
            readUserInput = true
        } else {
            readUserInput = false
            conversation = append(conversation, anthropic.NewUserMessage(toolResults...))
        }
    }
    return nil
}

func (a *Agent) executeTool(
    id, name string,
    input json.RawMessage,
) anthropic.ContentBlockParamUnion {
    for _, t := range a.tools {
        if t.Name == name {
            fmt.Printf("\u001b[92mtool\u001b[0m: %s(%s)\n", name, input)
            res, err := t.Function(input)
            if err != nil {
                return anthropic.NewToolResultBlock(id, err.Error(), true)
            }
            return anthropic.NewToolResultBlock(id, res, false)
        }
    }
    return anthropic.NewToolResultBlock(id, "tool not found", true)
}

func (a *Agent) runInference(
    ctx context.Context,
    conversation []anthropic.MessageParam,
) (*anthropic.Message, error) {
    var anthTools []anthropic.ToolUnionParam
    for _, t := range a.tools {
        anthTools = append(anthTools, anthropic.ToolUnionParam{
            OfTool: &anthropic.ToolParam{
                Name:        t.Name,
                Description: anthropic.String(t.Description),
                InputSchema: t.InputSchema,
            },
        })
    }
    return a.client.Messages.New(ctx, anthropic.MessageNewParams{
        Model:     anthropic.ModelClaude3_7SonnetLatest,
        MaxTokens: 1024,
        Messages:  conversation,
        Tools:     anthTools,
    })
}
