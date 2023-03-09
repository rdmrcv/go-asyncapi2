package spec

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/ligser/go-asyncapi2/spec/validate"
)

const asyncAPIVersion = "2.0.0"

// T is the root of an OpenAPI v3 document
// T is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#A2SObject
type T struct {
	Extensions map[string]interface{} `json:"-" yaml:"-"`

	AsyncAPI           string                 `json:"asyncapi" yaml:"asyncapi"`
	ID                 string                 `json:"id,omitempty" yaml:"id,omitempty"`
	Info               *openapi3.Info         `json:"info" yaml:"info"`
	DefaultContentType string                 `json:"defaultContentType" yaml:"defaultContentType"`
	Servers            Servers                `json:"servers,omitempty" yaml:"servers,omitempty"`
	Channels           Channels               `json:"channels,omitempty" yaml:"channels,omitempty"`
	Components         *Components            `json:"components,omitempty" yaml:"components,omitempty"`
	Tags               openapi3.Tags          `json:"tags,omitempty" yaml:"tags,omitempty"`
	ExternalDocs       *openapi3.ExternalDocs `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}

func (doc *T) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{}, 4+len(doc.Extensions))
	for k, v := range doc.Extensions {
		m[k] = v
	}
	m["asyncapi"] = doc.AsyncAPI
	if len(doc.ID) != 0 {
		m["id"] = doc.ID
	}

	m["info"] = doc.Info
	m["defaultContentType"] = doc.DefaultContentType
	if x := doc.Servers; len(x) != 0 {
		m["servers"] = x
	}
	if x := doc.Channels; len(x) != 0 {
		m["channels"] = x
	}
	if x := doc.Components; x != nil {
		m["components"] = x
	}
	if x := doc.Tags; len(x) != 0 {
		m["tags"] = x
	}
	if x := doc.ExternalDocs; x != nil {
		m["externalDocs"] = x
	}

	return json.Marshal(m)
}

func (doc *T) UnmarshalJSON(data []byte) error {
	type TBis T
	var x TBis
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	_ = json.Unmarshal(data, &x.Extensions)
	delete(x.Extensions, "asyncapi")
	delete(x.Extensions, "id")
	delete(x.Extensions, "info")
	delete(x.Extensions, "defaultContentType")
	delete(x.Extensions, "servers")
	delete(x.Extensions, "channels")
	delete(x.Extensions, "components")
	delete(x.Extensions, "tags")
	delete(x.Extensions, "externalDocs")
	*doc = T(x)
	return nil
}

func (doc *T) Validate(ctx context.Context) error {
	if doc.AsyncAPI != asyncAPIVersion {
		return fmt.Errorf("field asyncapi is required and should be equal %q: %w", asyncAPIVersion, validate.ErrWrongField)
	}

	if v := doc.Info; v != nil {
		if err := doc.Info.Validate(ctx); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("field info is required: %w", validate.ErrWrongField)
	}

	if v := doc.Servers; len(v) != 0 {
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	if v := doc.Channels; len(v) != 0 {
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	if err := doc.Components.Validate(ctx); err != nil {
		return err
	}

	return nil
}
