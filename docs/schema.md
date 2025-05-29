# Schema Generator 📊

## Overview

The `schema.go` file provides a powerful utility for generating JSON schemas from Go structs. This functionality is essential for creating properly typed tool definitions that Claude can understand and use correctly.

## Core Function

```go
func GenerateSchema[T any]() anthropic.ToolInputSchemaParam
```

This generic function takes a type parameter `T` and returns a JSON schema that describes the structure of that type in a format compatible with the Anthropic API.

## Implementation Details

### Using Go Generics

The function leverages Go's generics to provide a type-safe way to generate schemas for any struct type. This ensures that the schema accurately reflects the structure and constraints of the Go types.

### Reflector Configuration

```go
reflector := jsonschema.Reflector{
    AllowAdditionalProperties: false,
    DoNotReference:            true,
}
```

The schema generator is configured with specific settings:

| Setting | Value | Description |
|---------|-------|-------------|
| `AllowAdditionalProperties` | `false` | Enforces strict schema validation by disallowing properties not defined in the struct |
| `DoNotReference` | `true` | Ensures the schema is self-contained without external references |

### Schema Generation Process

1. Creates a zero value of the specified type
2. Uses reflection to analyze the structure of the type
3. Generates a JSON schema that describes the fields, types, and constraints
4. Returns the schema in the format expected by the Anthropic API

## Usage Example

```go
type MyToolInput struct {
    Name    string `json:"name" jsonschema_description:"The name parameter."`
    Count   int    `json:"count" jsonschema_description:"The count parameter."`
    Enabled bool   `json:"enabled,omitempty" jsonschema_description:"Optional enabled flag."`
}

var MyToolInputSchema = GenerateSchema[MyToolInput]()

var MyToolDefinition = ToolDefinition{
    Name:        "my_tool",
    Description: "A description of my tool.",
    InputSchema: MyToolInputSchema,
    Function:    MyToolFunction,
}
```

## Struct Tags

The schema generator recognizes several tags that can be used to customize the generated schema:

| Tag | Description |
|-----|-------------|
| `json:"name"` | Specifies the JSON field name |
| `json:",omitempty"` | Marks a field as optional |
| `jsonschema_description:"..."` | Provides a description for the field in the schema |

## Integration with Anthropic API

The schema is returned in the format expected by the Anthropic API's tool definition. This ensures that Claude has proper type information when using the tools, which enables:

1. More accurate tool usage
2. Better type checking on inputs
3. Appropriate handling of optional vs. required parameters

## Benefits

Using a schema generator rather than hardcoding JSON schemas provides several advantages:

1. **Type Safety**: The schema always matches the actual Go types
2. **DRY Principle**: Avoids duplication between type definitions and schemas
3. **Maintainability**: Changes to types automatically reflect in the schemas
4. **Consistency**: All tools use the same schema generation approach