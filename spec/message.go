package spec

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/getkin/kin-openapi/jsoninfo"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-openapi/jsonpointer"
)

type Messages map[string]*Message

var _ jsonpointer.JSONPointable = (*Messages)(nil)

func (h Messages) JSONLookup(token string) (interface{}, error) {
	value, ok := h[token]
	if value == nil || !ok {
		return nil, fmt.Errorf("object has no field %q", token)
	}

	return value, nil
}

type MessagesTraits map[string]*MessageTrait

var _ jsonpointer.JSONPointable = (*MessagesTraits)(nil)

func (h MessagesTraits) JSONLookup(token string) (interface{}, error) {
	value, ok := h[token]
	if value == nil || !ok {
		return nil, fmt.Errorf("object has no field %q", token)
	}

	return value, nil
}

// SchemaFormat tries to guess if a scheme serialized in the schemaFormat field
// or its just schema name.
type SchemaFormat struct {
	SchemaFormat string
	Schema       *openapi3.Schema
}

func (m *SchemaFormat) MarshalJSON() ([]byte, error) {
	if m.Schema != nil {
		return jsoninfo.MarshalStrictStruct(m.Schema)
	}

	return json.Marshal(m.SchemaFormat)
}

func (m *SchemaFormat) UnmarshalJSON(data []byte) error {
	schema := openapi3.Schema{}

	err := jsoninfo.UnmarshalStrictStruct(data, &schema)

	if !schema.IsEmpty() {
		m.Schema = &schema
		return err
	}

	return json.Unmarshal(data, &m.SchemaFormat)
}

// MessageTrait is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#message-trait-object
type MessageTrait struct {
	openapi3.ExtensionProps
	Headers       *openapi3.SchemaRef    `json:"headers,omitempty" yaml:"headers,omitempty"`
	CorrelationID *CorrelationIDRef      `json:"correlationId,omitempty" yaml:"correlationId,omitempty"`
	SchemaFormat  *SchemaFormat          `json:"schemaFormat,omitempty" yaml:"schemaFormat,omitempty"`
	ContentType   string                 `json:"contentType,omitempty" yaml:"contentType,omitempty"`
	Name          string                 `json:"name,omitempty" yaml:"name,omitempty"`
	Title         string                 `json:"title,omitempty" yaml:"title,omitempty"`
	Summary       string                 `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description   string                 `json:"description,omitempty" yaml:"description,omitempty"`
	Tags          openapi3.Tags          `json:"tags,omitempty" yaml:"tags,omitempty"`
	ExternalDocs  *openapi3.ExternalDocs `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	Bindings      *MessageBindings       `json:"bindings,omitempty" yaml:"bindings,omitempty"`
	Examples      []interface{}          `json:"examples,omitempty" yaml:"examples,omitempty"`
}

func (value *MessageTrait) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStrictStruct(value)
}

func (value *MessageTrait) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStrictStruct(data, value)
}

func (value *MessageTrait) Validate(ctx context.Context) error {
	if v := value.Headers; v != nil {
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	if v := value.CorrelationID; v != nil {
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	if v := value.Bindings; v != nil {
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	return nil
}

// Message is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#message-object
type Message struct {
	MessageTrait

	Payload *openapi3.SchemaRef `json:"payload,omitempty" yaml:"payload,omitempty"`
	Traits  []*MessageTraitRef  `json:"traits,omitempty" yaml:"traits,omitempty"`
}

func (value *Message) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStrictStruct(value)
}

func (value *Message) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStrictStruct(data, value)
}

func (value *Message) Validate(ctx context.Context) error {
	if v := value.Payload; v != nil {
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	return value.MessageTrait.Validate(ctx)
}
