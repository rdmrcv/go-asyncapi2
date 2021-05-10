package bindings

import (
	"context"
	"fmt"
	"net/http"

	"github.com/getkin/kin-openapi/jsoninfo"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/ligser/go-asyncapi2/spec/validate"
)

// WsServer is defined in AsyncAPI spec: https://github.com/asyncapi/bindings/tree/master/websockets#server-binding-object
type WsServer struct {
}

func (*WsServer) Validate(_ context.Context) error {
	return nil
}

var _ jsoninfo.StrictStruct = &WsChannel{}

// WsChannel is defined in AsyncAPI spec: https://github.com/asyncapi/bindings/tree/master/websockets#channel-binding-object
type WsChannel struct {
	openapi3.ExtensionProps
	Method         string           `json:"method,omitempty" yaml:"method,omitempty"`
	Query          *openapi3.Schema `json:"query,omitempty" yaml:"query,omitempty"`
	Headers        *openapi3.Schema `json:"headers,omitempty" yaml:"headers,omitempty"`
	BindingVersion string           `json:"bindingVersion,omitempty" yaml:"bindingVersion,omitempty"`
}

func (binding *WsChannel) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStrictStruct(binding)
}

func (binding *WsChannel) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStrictStruct(data, binding)
}

func (binding *WsChannel) Validate(ctx context.Context) error {
	if binding.Method != http.MethodGet && binding.Method != http.MethodPost {
		return fmt.Errorf("method value MUST be either GET or POST: %w", validate.ErrWrongField)
	}

	if v := binding.Query; v != nil {
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	if v := binding.Headers; v != nil {
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	return nil
}

// WsOperation is defined in AsyncAPI spec: https://github.com/asyncapi/bindings/tree/master/websockets#operation-binding-object
type WsOperation struct {
}

func (*WsOperation) Validate(context.Context) error {
	return nil
}

// WsMessage is defined in AsyncAPI spec: https://github.com/asyncapi/bindings/tree/master/websockets#message-binding-object
type WsMessage struct {
}

func (*WsMessage) Validate(context.Context) error {
	return nil
}
