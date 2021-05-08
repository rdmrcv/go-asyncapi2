package spec

import (
	"context"
	"fmt"

	"github.com/getkin/kin-openapi/jsoninfo"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/ligser/asyncapi/spec/validate"
)

// T is the root of an OpenAPI v3 document
// T is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#A2SObject
type T struct {
	openapi3.ExtensionProps
	AsyncAPI     string                 `json:"asyncapi" yaml:"asyncapi"`
	ID           string                 `json:"id" yaml:"id"`
	Info         *openapi3.Info         `json:"info" yaml:"info"`
	Servers      Servers                `json:"servers,omitempty" yaml:"servers,omitempty"`
	Channels     Channels               `json:"channels" yaml:"channels"`
	Components   Components             `json:"components,omitempty" yaml:"components,omitempty"`
	Tags         openapi3.Tags          `json:"tags,omitempty" yaml:"tags,omitempty"`
	ExternalDocs *openapi3.ExternalDocs `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}

func (doc *T) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStrictStruct(doc)
}

func (doc *T) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStrictStruct(data, doc)
}

func (doc *T) Validate(ctx context.Context) error {
	if len(doc.AsyncAPI) == 0 {
		return fmt.Errorf("field asyncapi is required: %w", validate.ErrWrongField)
	}

	if v := doc.Info; v != nil {
		if err := doc.Info.Validate(ctx); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("field info is required: %w", validate.ErrWrongField)
	}

	if err := doc.Components.Validate(ctx); err != nil {
		return err
	}

	if v := doc.Servers; len(v) != 0 {
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	return nil
}
