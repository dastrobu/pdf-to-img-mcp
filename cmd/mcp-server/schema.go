package main

import (
	"reflect"

	"github.com/google/jsonschema-go/jsonschema"
)

// GenerateSchema generates a JSON schema (draft 2020-12, as understood by the
// modelcontextprotocol/go-sdk) for the given Go type T. The go-sdk's AddTool
// uses it to validate and apply defaults to tool inputs.
//
// It mirrors the helper used by the mail-mcp project. Input structs use the
// `jsonschema:"..."` struct tag for field descriptions and `json:",omitempty"`
// for optional (non-required) fields.
func GenerateSchema[T any]() *jsonschema.Schema {
	schema, err := jsonschema.ForType(reflect.TypeFor[T](), &jsonschema.ForOptions{})
	if err != nil {
		// Should not happen for the valid Go types used in this package.
		panic(err)
	}
	return schema
}
