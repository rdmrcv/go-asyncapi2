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

// MessagePayload tries to guess if a scheme serialized in the message payload
// or its some generic object.
type MessagePayload struct {
	Any    map[string]interface{}
	Schema *openapi3.Schema
}

func (m *MessagePayload) MarshalJSON() ([]byte, error) {
	if m.Schema != nil {
		return jsoninfo.MarshalStrictStruct(m.Schema)
	}

	return json.Marshal(m.Any)
}

func (m *MessagePayload) UnmarshalJSON(data []byte) error {
	schema := openapi3.Schema{}

	err := jsoninfo.UnmarshalStrictStruct(data, &schema)

	if !schema.IsEmpty() {
		m.Schema = &schema
		return err
	}

	m.Any = map[string]interface{}{}

	return json.Unmarshal(data, &m.Any)
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
	Headers       *openapi3.SchemaRef    `json:"headers" yaml:"headers"`
	CorrelationID *CorrelationIDRef      `json:"correlationId" yaml:"correlationId"`
	SchemaFormat  SchemaFormat           `json:"schemaFormat" yaml:"schemaFormat"`
	ContentType   string                 `json:"contentType" yaml:"contentType"`
	Name          string                 `json:"name" yaml:"name"`
	Title         string                 `json:"title" yaml:"title"`
	Summary       string                 `json:"summary" yaml:"summary"`
	Description   string                 `json:"description" yaml:"description"`
	Tags          openapi3.Tags          `json:"tags" yaml:"tags"`
	ExternalDocs  *openapi3.ExternalDocs `json:"externalDocs" yaml:"externalDocs"`
	Bindings      *MessageBindings       `json:"bindings" yaml:"bindings"`
	Examples      []interface{}          `json:"examples" yaml:"examples"`
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

	Payload MessagePayload     `json:"payload" yaml:"payload"`
	Traits  []*MessageTraitRef `json:"traits" yaml:"traits"`
}

func (value *Message) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStrictStruct(value)
}

func (value *Message) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStrictStruct(data, value)
}

func (value *Message) Validate(ctx context.Context) error {
	if v := value.Payload.Schema; v != nil {
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	return value.MessageTrait.Validate(ctx)
}
