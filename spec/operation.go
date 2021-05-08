package spec

import (
	"context"
	"fmt"

	"github.com/getkin/kin-openapi/jsoninfo"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-openapi/jsonpointer"
)

type OperationsTraits map[string]*OperationTrait

var _ jsonpointer.JSONPointable = (*OperationsTraits)(nil)

func (h OperationsTraits) JSONLookup(token string) (interface{}, error) {
	value, ok := h[token]
	if value == nil || !ok {
		return nil, fmt.Errorf("object has no field %q", token)
	}

	return value, nil
}

// OperationTrait is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#operationTraitObject
type OperationTrait struct {
	openapi3.ExtensionProps
	OperationID  string                 `json:"operationId" yaml:"operationId"`
	Summary      string                 `json:"summary" yaml:"summary"`
	Description  string                 `json:"description" yaml:"description"`
	Tags         openapi3.Tags          `json:"tags" yaml:"tags"`
	ExternalDocs *openapi3.ExternalDocs `json:"externalDocs" yaml:"externalDocs"`
	Bindings     *OperationBindings     `json:"bindings" yaml:"bindings"`
}

func (value *OperationTrait) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStrictStruct(value)
}

func (value *OperationTrait) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStrictStruct(data, value)
}

func (value *OperationTrait) Validate(ctx context.Context) error {
	if v := value.Bindings; v != nil {
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	return nil
}

// Operation is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#operation-object
type Operation struct {
	OperationTrait

	Traits  []*OperationTraitRef `json:"traits" yaml:"traits"`
	Message *MessageOneOf        `json:"message" yaml:"message"`
}

func (value *Operation) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStrictStruct(value)
}

func (value *Operation) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStrictStruct(data, value)
}

func (value *Operation) Validate(ctx context.Context) error {
	if v := value.Traits; len(v) > 0 {
		for _, item := range v {
			if err := item.Validate(ctx); err != nil {
				return err
			}
		}
	}

	if v := value.Message; v != nil {
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	return value.OperationTrait.Validate(ctx)
}
