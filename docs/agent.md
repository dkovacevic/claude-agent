# Agent Implementation (agent.go)

![Component](https://img.shields.io/badge/Component-Core-red)

## Overview

`agent.go` contains the core implementation of the Agent that manages the conversation flow between the user, Claude, and the execution of tools. It acts as the central orchestration component that ties together all aspects of the application.

## Agent Structure

```go
type Agent struct {
    client         *anthropic.Client
    getUserMessage func() (string, bool)
    tools          []tools.ToolDefinition
}
```

The Agent structure encapsulates:
- An Anthropic client for API communication
- A function to retrieve user input
- A list of available tool definitions

## Key Functions

### NewAgent

```go
func NewAgent(
    client *anthropic.Client,
    getUserMessage func() (string, bool),
    tools []tools.ToolDefinition,
) *Agent
```

Creates a new Agent instance with the provided dependencies.

### Run

```go
func (a *Agent) Run(ctx context.Context) error
```

The primary execution loop of the Agent that:
1. Manages the conversation state
2. Handles user input
3. Sends messages to Claude
4. Processes Claude's responses
5. Executes tools when requested
6. Returns tool results to Claude

### executeTool

```go
func (a *Agent) executeTool(
    id, name string,
    input json.RawMessage,
) anthropic.ContentBlockParamUnion
```

Executes a specific tool by:
1. Finding the matching tool definition by name
2. Unmarshaling and passing the input parameters
3. Executing the tool function
4. Formatting the result (or error) as a tool response

### runInference

```go
func (a *Agent) runInference(
    ctx context.Context,
    conversation []anthropic.MessageParam,
) (*anthropic.Message, error)
```

Sends the current conversation to Claude and retrieves a response by:
1. Configuring the available tools in the Claude API format
2. Setting up the API request with the conversation history
3. Calling the Anthropic API
4. Returning Claude's response

## Conversation Flow

The Agent implements a sophisticated conversation flow:

1. **User Input Phase**:
   - Prompt the user for input
   - Add the user's message to the conversation history

2. **Claude Response Phase**:
   - Send the conversation to Claude via the API
   - Process Claude's response
   - Display text content to the user

3. **Tool Execution Phase** (if tools are used):
   - Identify tool usage in Claude's response
   - Execute the requested tools
   - Format and collect tool results

4. **Result Processing Phase**:
   - Add tool results to the conversation
   - Send the updated conversation back to Claude (if tools were used)
   - Otherwise, return to the user input phase

## Key Features

- **Stateful Conversation Management**: Maintains the full conversation history
- **Elegant Tool Execution**: Seamlessly handles tool requests from Claude
- **Clean Terminal UI**: Uses ANSI color codes for a pleasant terminal experience
- **Proper Error Handling**: Comprehensive error handling throughout the process

## Technical Details

### Color Coding

The agent uses ANSI color codes to enhance the terminal experience:
- 🔵 Blue (\u001b[94m) for user messages
- 🟡 Yellow (\u001b[93m) for Claude's responses
- 🟢 Green (\u001b[92m) for tool executions

### Claude API Integration

The agent uses Claude 3 Sonnet (latest) as the default model:
```go
anthropic.ModelClaude3_7SonnetLatest
```

### Tool Interface

The agent handles tools through a standardized interface that:
1. Registers tools with Claude in the appropriate format
2. Dispatches tool calls to the correct implementation
3. Formats results in Claude's expected response format