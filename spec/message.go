package spec

import (
	"context"
	"encoding/json"

	"github.com/getkin/kin-openapi/openapi3"
)

type Messages map[string]*Message

type MessagesTraits map[string]*MessageTrait

// MessageTrait is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#message-trait-object
type MessageTrait struct {
	Extensions map[string]interface{} `json:"-" yaml:"-"`

	Headers       *openapi3.SchemaRef      `json:"headers,omitempty" yaml:"headers,omitempty"`
	CorrelationID *CorrelationIDRef        `json:"correlationId,omitempty" yaml:"correlationId,omitempty"`
	SchemaFormat  string                   `json:"schemaFormat,omitempty" yaml:"schemaFormat,omitempty"`
	ContentType   string                   `json:"contentType,omitempty" yaml:"contentType,omitempty"`
	Name          string                   `json:"name,omitempty" yaml:"name,omitempty"`
	Title         string                   `json:"title,omitempty" yaml:"title,omitempty"`
	Summary       string                   `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description   string                   `json:"description,omitempty" yaml:"description,omitempty"`
	Tags          openapi3.Tags            `json:"tags,omitempty" yaml:"tags,omitempty"`
	ExternalDocs  *openapi3.ExternalDocs   `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	Bindings      *MessageBindings         `json:"bindings,omitempty" yaml:"bindings,omitempty"`
	Examples      []map[string]interface{} `json:"examples,omitempty" yaml:"examples,omitempty"`
}

func (value *MessageTrait) MarshalJSON() ([]byte, error) {
	return json.Marshal(value.produceM())
}

func (value *MessageTrait) produceM() map[string]interface{} {
	m := make(map[string]interface{}, 12+len(value.Extensions))
	for k, v := range value.Extensions {
		m[k] = v
	}

	if value.Headers != nil {
		m["headers"] = value.Headers
	}
	if value.CorrelationID != nil {
		m["correlationId"] = value.CorrelationID
	}
	if len(value.SchemaFormat) != 0 {
		m["schemaFormat"] = value.SchemaFormat
	}
	if len(value.ContentType) != 0 {
		m["contentType"] = value.ContentType
	}
	if len(value.Name) != 0 {
		m["name"] = value.Name
	}
	if len(value.Title) != 0 {
		m["title"] = value.Title
	}
	if len(value.Summary) != 0 {
		m["summary"] = value.Summary
	}
	if len(value.Description) != 0 {
		m["description"] = value.Description
	}
	if len(value.Tags) != 0 {
		m["tags"] = value.Tags
	}
	if value.ExternalDocs != nil {
		m["externalDocs"] = value.ExternalDocs
	}
	if value.Bindings != nil {
		m["bindings"] = value.Bindings
	}
	if len(value.Examples) != 0 {
		m["examples"] = value.Examples
	}
	return m
}

func (value *MessageTrait) UnmarshalJSON(data []byte) error {
	type MessageTraitBis MessageTrait
	var x MessageTraitBis
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	_ = json.Unmarshal(data, &x.Extensions)

	delete(x.Extensions, "headers")
	delete(x.Extensions, "correlationId")
	delete(x.Extensions, "schemaFormat")
	delete(x.Extensions, "contentType")
	delete(x.Extensions, "name")
	delete(x.Extensions, "title")
	delete(x.Extensions, "summary")
	delete(x.Extensions, "description")
	delete(x.Extensions, "tags")
	delete(x.Extensions, "externalDocs")
	delete(x.Extensions, "bindings")
	delete(x.Extensions, "examples")

	*value = MessageTrait(x)

	return nil
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
	m := value.MessageTrait.produceM()

	if value.Payload != nil {
		m["payload"] = value.Payload
	}
	if len(value.Traits) != 0 {
		m["traits"] = value.Traits
	}

	return json.Marshal(m)
}

func (value *Message) UnmarshalJSON(data []byte) error {
	type MessageBis Message
	var x MessageBis
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}

	if err := x.MessageTrait.UnmarshalJSON(data); err != nil {
		return err
	}

	*value = Message(x)

	return nil
}

func (value *Message) Validate(ctx context.Context) error {
	if v := value.Payload; v != nil {
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	for _, v := range value.Traits {
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	return value.MessageTrait.Validate(ctx)
}
