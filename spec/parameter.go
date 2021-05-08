package spec

import (
	"context"
	"fmt"

	"github.com/getkin/kin-openapi/jsoninfo"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-openapi/jsonpointer"
)

type Parameters map[string]*Parameter

var _ jsonpointer.JSONPointable = (*Parameters)(nil)

func (h Parameters) JSONLookup(token string) (interface{}, error) {
	value, ok := h[token]
	if value == nil || !ok {
		return nil, fmt.Errorf("object has no field %q", token)
	}

	return value, nil
}

// ParametersRefs is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#parameters-object
type ParametersRefs map[string]*ParameterRef

var _ jsonpointer.JSONPointable = (*ParametersRefs)(nil)

func (h ParametersRefs) Validate(ctx context.Context) error {
	for _, item := range h {
		if err := item.Validate(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (h ParametersRefs) JSONLookup(token string) (interface{}, error) {
	value, ok := h[token]
	if value == nil || !ok {
		return nil, fmt.Errorf("object has no field %q", token)
	}

	return value, nil
}

// Parameter is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#parameterObject
type Parameter struct {
	openapi3.ExtensionProps
	Description string           `json:"description,omitempty" yaml:"description,omitempty"`
	Schema      *openapi3.Schema `json:"schema,omitempty" yaml:"schema,omitempty"`
	Location    string           `json:"location,omitempty" yaml:"location,omitempty"`
}

func (value *Parameter) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStrictStruct(value)
}

func (value *Parameter) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStrictStruct(data, value)
}

func (value *Parameter) Validate(ctx context.Context) error {
	if v := value.Schema; v != nil {
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	return nil
}
