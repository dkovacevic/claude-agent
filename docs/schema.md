# JSON Schema Generation (schema.go)

![Component](https://img.shields.io/badge/Component-Utility-orange)

## Overview

`schema.go` provides a powerful utility function that automatically generates JSON Schema for tool parameters from Go struct types. This enables seamless integration with Claude's tool system while maintaining type safety and parameter validation.

## Implementation

### GenerateSchema Function

```go
func GenerateSchema[T any]() anthropic.ToolInputSchemaParam {
    reflector := jsonschema.Reflector{
        AllowAdditionalProperties: false,
        DoNotReference:            true,
    }
    var v T
    sch := reflector.Reflect(v)
    return anthropic.ToolInputSchemaParam{
        Properties: sch.Properties,
    }
}
```

This generic function:
1. Creates a jsonschema reflector with appropriate settings
2. Instantiates an empty value of type T
3. Uses reflection to generate a JSON schema from the type
4. Formats the schema to match Anthropic's API requirements

## Key Features

### Type Safety

By using Go's generics, the function ensures type safety between:
- The input type definition
- The JSON schema generation
- The runtime parameter parsing

### Clean API Integration

The function produces schema in exactly the format expected by the Anthropic API:
```go
anthropic.ToolInputSchemaParam{
    Properties: sch.Properties,
}
```

### Configuration

The reflector is configured with:
- `AllowAdditionalProperties: false` - Strictly enforce the schema
- `DoNotReference: true` - Inline all property definitions for clarity

## Usage Pattern

Throughout the codebase, this function is used consistently:

```go
// Define input structure with JSON tags and descriptions
type SomeToolInput struct {
    Param string `json:"param" jsonschema_description:"Description of param"`
}

// Generate schema automatically
var SomeToolInputSchema = GenerateSchema[SomeToolInput]()

// Use in tool definition
var SomeToolDefinition = ToolDefinition{
    // ...
    InputSchema: SomeToolInputSchema,
    // ...
}
```

## Benefits

1. **Consistency**: Ensures all tools follow the same schema pattern
2. **Maintainability**: Changes to input structures automatically update schemas
3. **Readability**: Keeps schema definitions close to the corresponding types
4. **Validation**: Leverages Go's type system for parameter validation

## Dependencies

- [github.com/invopop/jsonschema](https://github.com/invopop/jsonschema): Library for JSON Schema reflection
- [github.com/anthropics/anthropic-sdk-go](https://github.com/anthropics/anthropic-sdk-go): Anthropic Go SDK