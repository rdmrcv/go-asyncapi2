package bindings

import (
	"context"
	"fmt"
	"net/http"

	"github.com/getkin/kin-openapi/jsoninfo"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/ligser/asyncapi/spec/validate"
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

var _ jsoninfo.StrictStruct = &HttpOperation{}

// HttpOperation is defined in AsyncAPI spec: https://github.com/asyncapi/bindings/tree/master/http#operation-binding-object
type HttpOperation struct {
	openapi3.ExtensionProps
	Type           HttpOperationBindingType `json:"type,omitempty" yaml:"type,omitempty"`
	Method         string                   `json:"method,omitempty" yaml:"method,omitempty"`
	Query          *openapi3.Schema         `json:"query,omitempty" yaml:"query,omitempty"`
	BindingVersion string                   `json:"bindingVersion,omitempty" yaml:"bindingVersion,omitempty"`
}

func (binding *HttpOperation) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStrictStruct(binding)
}

func (binding *HttpOperation) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStrictStruct(data, binding)
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
		if v.Type != "object" || len(v.Properties) == 0 {
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

var _ jsoninfo.StrictStruct = &HttpMessage{}

// HttpMessage is defined in AsyncAPI spec: https://github.com/asyncapi/bindings/tree/master/http#message-binding-object
type HttpMessage struct {
	openapi3.ExtensionProps
	Headers        *openapi3.Schema `json:"headers,omitempty" yaml:"headers,omitempty"`
	BindingVersion string           `json:"bindingVersion,omitempty" yaml:"bindingVersion,omitempty"`
}

func (value *HttpMessage) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStrictStruct(value)
}

func (value *HttpMessage) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStrictStruct(data, value)
}

func (value *HttpMessage) Validate(ctx context.Context) error {
	if v := value.Headers; v != nil {
		if v.Type != "object" || len(v.Properties) == 0 {
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
