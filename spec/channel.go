package spec

import (
	"context"
	"fmt"

	"github.com/getkin/kin-openapi/jsoninfo"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-openapi/jsonpointer"
)

// Channels is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#channels-object
type Channels map[string]*Channel

func (h Channels) JSONLookup(token string) (interface{}, error) {
	value, ok := h[token]
	if value == nil || !ok {
		return nil, fmt.Errorf("object has no field %q", token)
	}

	return value, nil
}

var _ jsonpointer.JSONPointable = (Channels)(nil)

func (h Channels) Validate(ctx context.Context) error {
	for _, item := range h {
		if err := item.Validate(ctx); err != nil {
			return err
		}
	}

	return nil
}

// Channel is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#channel-item-object
type Channel struct {
	openapi3.ExtensionProps
	Description string              `json:"description,omitempty" yaml:"description,omitempty"`
	Subscribe   *OperationRef       `json:"subscribe,omitempty" yaml:"subscribe,omitempty"`
	Publish     *OperationRef       `json:"publish,omitempty" yaml:"publish,omitempty"`
	Parameters  ParametersRefs      `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Bindings    *ChannelBindingsRef `json:"bindings,omitempty" yaml:"bindings,omitempty"`
}

func (value *Channel) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStrictStruct(value)
}

func (value *Channel) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStrictStruct(data, value)
}

func (value *Channel) Validate(ctx context.Context) error {
	if v := value.Subscribe; v != nil {
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	if v := value.Publish; v != nil {
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	if v := value.Parameters; v != nil {
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
