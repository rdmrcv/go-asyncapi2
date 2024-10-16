package bindings

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/rdmrcv/go-asyncapi2/spec/validate"
)

// HttpServer is defined in AsyncAPI spec: https://github.com/asyncapi/bindings/tree/master/http#server-binding-object
type HttpServer struct {
}

func (*HttpServer) Validate(context.Context) error {
	return nil
}

// HttpChannel is defined in AsyncAPI spec: https://github.com/asyncapi/bindings/tree/master/http#channel-binding-object
type HttpChannel struct {
}

func (*HttpChannel) Validate(context.Context) error {
	return nil
}

type HttpOperationBindingType string

const (
	HttpOperationBindingRequest  = "request"
	HttpOperationBindingResponse = "response"
)

var httpValidMethodsSet = map[string]struct{}{
	http.MethodGet:     {},
	http.MethodPost:    {},
	http.MethodPut:     {},
	http.MethodPatch:   {},
	http.MethodDelete:  {},
	http.MethodHead:    {},
	http.MethodOptions: {},
	http.MethodConnect: {},
	http.MethodTrace:   {},
}

// HttpOperation is defined in AsyncAPI spec: https://github.com/asyncapi/bindings/tree/master/http#operation-binding-object
type HttpOperation struct {
	Extensions map[string]interface{} `json:"-" yaml:"-"`

	Type           HttpOperationBindingType `json:"type,omitempty" yaml:"type,omitempty"`
	Method         string                   `json:"method,omitempty" yaml:"method,omitempty"`
	Query          *openapi3.Schema         `json:"query,omitempty" yaml:"query,omitempty"`
	BindingVersion string                   `json:"bindingVersion,omitempty" yaml:"bindingVersion,omitempty"`
}

func (binding *HttpOperation) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{}, 4+len(binding.Extensions))
	for k, v := range binding.Extensions {
		m[k] = v
	}

	if len(binding.Type) != 0 {
		m["type"] = binding.Type
	}
	if len(binding.Method) != 0 {
		m["method"] = binding.Method
	}
	if binding.Query != nil {
		m["query"] = binding.Query
	}
	if len(binding.BindingVersion) != 0 {
		m["bindingVersion"] = binding.BindingVersion
	}

	return json.Marshal(m)
}

func (binding *HttpOperation) UnmarshalJSON(data []byte) error {
	type HttpOperationBis HttpOperation
	var x HttpOperationBis
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	_ = json.Unmarshal(data, &x.Extensions)

	delete(x.Extensions, "type")
	delete(x.Extensions, "method")
	delete(x.Extensions, "query")
	delete(x.Extensions, "bindingVersion")

	*binding = HttpOperation(x)

	return nil
}

func (binding *HttpOperation) Validate(ctx context.Context) error {
	switch binding.Type {
	case HttpOperationBindingRequest:
		if _, ok := httpValidMethodsSet[binding.Method]; !ok {
			return fmt.Errorf(
				"when the type field is request the mehtod should be set to a valid value: %w",
				validate.ErrWrongField,
			)
		}
	case HttpOperationBindingResponse:
	default:
		return fmt.Errorf(
			"type should be set and must be either request or response: %w",
			validate.ErrWrongField,
		)
	}

	if v := binding.Query; v != nil {
		if !v.Type.Is(openapi3.TypeObject) || len(v.Properties) == 0 {
			return fmt.Errorf(
				"the schema in the query field MUST be of type object and have a properties key: %w",
				validate.ErrWrongField,
			)
		}

		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	return nil
}

// HttpMessage is defined in AsyncAPI spec: https://github.com/asyncapi/bindings/tree/master/http#message-binding-object
type HttpMessage struct {
	Extensions map[string]interface{} `json:"-" yaml:"-"`

	Headers        *openapi3.Schema `json:"headers,omitempty" yaml:"headers,omitempty"`
	BindingVersion string           `json:"bindingVersion,omitempty" yaml:"bindingVersion,omitempty"`
}

func (value *HttpMessage) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{}, 2+len(value.Extensions))
	for k, v := range value.Extensions {
		m[k] = v
	}

	if value.Headers != nil {
		m["headers"] = value.Headers
	}
	if len(value.BindingVersion) != 0 {
		m["bindingVersion"] = value.BindingVersion
	}

	return json.Marshal(m)
}

func (value *HttpMessage) UnmarshalJSON(data []byte) error {
	type HttpMessageBis HttpMessage
	var x HttpMessageBis
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	_ = json.Unmarshal(data, &x.Extensions)

	delete(x.Extensions, "headers")
	delete(x.Extensions, "bindingVersion")

	*value = HttpMessage(x)

	return nil
}

func (value *HttpMessage) Validate(ctx context.Context) error {
	if v := value.Headers; v != nil {
		if !v.Type.Is(openapi3.TypeObject) || len(v.Properties) == 0 {
			return fmt.Errorf(
				"the schema in the headers field MUST be of type object and have a properties key: %w",
				validate.ErrWrongField,
			)
		}

		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	return nil
}
