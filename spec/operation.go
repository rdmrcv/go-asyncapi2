package spec

import (
	"context"
	"encoding/json"

	"github.com/getkin/kin-openapi/openapi3"
)

type OperationsTraits map[string]*OperationTrait

// OperationTrait is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#operationTraitObject
type OperationTrait struct {
	Extensions map[string]interface{} `json:"-" yaml:"-"`

	OperationID  string                 `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Summary      string                 `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description  string                 `json:"description,omitempty" yaml:"description,omitempty"`
	Tags         openapi3.Tags          `json:"tags,omitempty" yaml:"tags,omitempty"`
	ExternalDocs *openapi3.ExternalDocs `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	Bindings     *OperationBindings     `json:"bindings,omitempty" yaml:"bindings,omitempty"`
}

func (value *OperationTrait) produceM() map[string]interface{} {
	m := make(map[string]interface{}, 6+len(value.Extensions))
	for k, v := range value.Extensions {
		m[k] = v
	}

	if len(value.OperationID) != 0 {
		m["operationId"] = value.OperationID
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

	return m
}

func (value *OperationTrait) MarshalJSON() ([]byte, error) {
	return json.Marshal(value.produceM())
}

func (value *OperationTrait) UnmarshalJSON(data []byte) error {
	type OperationTraitBis OperationTrait
	var x OperationTraitBis
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	_ = json.Unmarshal(data, &x.Extensions)

	delete(x.Extensions, "operationId")
	delete(x.Extensions, "summary")
	delete(x.Extensions, "description")
	delete(x.Extensions, "tags")
	delete(x.Extensions, "externalDocs")
	delete(x.Extensions, "bindings")

	*value = OperationTrait(x)

	return nil
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

	Traits  []*OperationTraitRef `json:"traits,omitempty" yaml:"traits,omitempty"`
	Message *MessageOneOf        `json:"message,omitempty" yaml:"message,omitempty"`
}

func (value *Operation) MarshalJSON() ([]byte, error) {
	m := value.OperationTrait.produceM()

	if len(value.Traits) != 0 {
		m["traits"] = value.Traits
	}
	if value.Message != nil {
		m["message"] = value.Message
	}

	return json.Marshal(m)
}

func (value *Operation) UnmarshalJSON(data []byte) error {
	type OperationBis Operation
	var x OperationBis
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}

	if err := x.OperationTrait.UnmarshalJSON(data); err != nil {
		return err
	}

	*value = Operation(x)

	return nil
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
