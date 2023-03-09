package spec

import (
	"context"
	"encoding/json"

	"github.com/getkin/kin-openapi/openapi3"
)

type Parameters map[string]*Parameter

// ParametersRefs is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#parameters-object
type ParametersRefs map[string]*ParameterRef

func (h ParametersRefs) Validate(ctx context.Context) error {
	for _, item := range h {
		if err := item.Validate(ctx); err != nil {
			return err
		}
	}

	return nil
}

// Parameter is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#parameterObject
type Parameter struct {
	Extensions map[string]interface{} `json:"-" yaml:"-"`

	Description string           `json:"description,omitempty" yaml:"description,omitempty"`
	Schema      *openapi3.Schema `json:"schema,omitempty" yaml:"schema,omitempty"`
	Location    string           `json:"location,omitempty" yaml:"location,omitempty"`
}

func (value *Parameter) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{}, 6+len(value.Extensions))
	for k, v := range value.Extensions {
		m[k] = v
	}

	if len(value.Description) != 0 {
		m["description"] = value.Description
	}
	if value.Schema != nil {
		m["schema"] = value.Schema
	}
	if len(value.Location) != 0 {
		m["location"] = value.Location
	}

	return json.Marshal(m)
}

func (value *Parameter) UnmarshalJSON(data []byte) error {
	type ParameterBis Parameter
	var x ParameterBis
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	_ = json.Unmarshal(data, &x.Extensions)

	delete(x.Extensions, "description")
	delete(x.Extensions, "schema")
	delete(x.Extensions, "location")

	*value = Parameter(x)

	return nil
}

func (value *Parameter) Validate(ctx context.Context) error {
	if v := value.Schema; v != nil {
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	return nil
}
