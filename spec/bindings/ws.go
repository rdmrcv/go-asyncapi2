package bindings

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/ligser/go-asyncapi2/spec/validate"
)

// WsServer is defined in AsyncAPI spec: https://github.com/asyncapi/bindings/tree/master/websockets#server-binding-object
type WsServer struct {
}

func (*WsServer) Validate(_ context.Context) error {
	return nil
}

// WsChannel is defined in AsyncAPI spec: https://github.com/asyncapi/bindings/tree/master/websockets#channel-binding-object
type WsChannel struct {
	Extensions map[string]interface{} `json:"-" yaml:"-"`

	Method         string           `json:"method,omitempty" yaml:"method,omitempty"`
	Query          *openapi3.Schema `json:"query,omitempty" yaml:"query,omitempty"`
	Headers        *openapi3.Schema `json:"headers,omitempty" yaml:"headers,omitempty"`
	BindingVersion string           `json:"bindingVersion,omitempty" yaml:"bindingVersion,omitempty"`
}

func (binding *WsChannel) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{}, 4+len(binding.Extensions))
	for k, v := range binding.Extensions {
		m[k] = v
	}

	if len(binding.Method) != 0 {
		m["method"] = binding.Method
	}
	if binding.Query != nil {
		m["query"] = binding.Query
	}
	if binding.Headers != nil {
		m["type"] = binding.Headers
	}
	if len(binding.BindingVersion) != 0 {
		m["bindingVersion"] = binding.BindingVersion
	}

	return json.Marshal(m)
}

func (binding *WsChannel) UnmarshalJSON(data []byte) error {
	type WsChannelBis WsChannel
	var x WsChannelBis
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	_ = json.Unmarshal(data, &x.Extensions)

	delete(x.Extensions, "method")
	delete(x.Extensions, "query")
	delete(x.Extensions, "type")
	delete(x.Extensions, "bindingVersion")

	*binding = WsChannel(x)

	return nil
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
