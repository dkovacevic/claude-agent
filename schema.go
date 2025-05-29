package main

import (
    "github.com/invopop/jsonschema"
    "github.com/anthropics/anthropic-sdk-go"
)

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
