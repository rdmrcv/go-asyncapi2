package spec

import (
	"context"
	"fmt"

	"github.com/getkin/kin-openapi/jsoninfo"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-openapi/jsonpointer"
	"github.com/ligser/asyncapi/spec/validate"
)

type CorrelationIDs map[string]*CorrelationID

var _ jsonpointer.JSONPointable = (*CorrelationIDs)(nil)

func (h CorrelationIDs) JSONLookup(token string) (interface{}, error) {
	value, ok := h[token]
	if value == nil || !ok {
		return nil, fmt.Errorf("object has no field %q", token)
	}

	return value, nil
}

// CorrelationID is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#correlationIdObject
type CorrelationID struct {
	openapi3.ExtensionProps
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Location    string `json:"location" yaml:"location"`
}

func (value *CorrelationID) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStrictStruct(value)
}

func (value *CorrelationID) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStrictStruct(data, value)
}

func (value *CorrelationID) Validate(context.Context) error {
	if value.Location == "" {
		return fmt.Errorf("location field is required: %w", validate.ErrWrongField)
	}

	return nil
}
