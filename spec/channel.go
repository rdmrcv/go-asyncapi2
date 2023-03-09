package spec

import (
	"context"
	"encoding/json"
)

// Channels is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#channels-object
type Channels map[string]*Channel

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
	Extensions map[string]interface{} `json:"-" yaml:"-"`

	Description string              `json:"description,omitempty" yaml:"description,omitempty"`
	Subscribe   *OperationRef       `json:"subscribe,omitempty" yaml:"subscribe,omitempty"`
	Publish     *OperationRef       `json:"publish,omitempty" yaml:"publish,omitempty"`
	Parameters  ParametersRefs      `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Bindings    *ChannelBindingsRef `json:"bindings,omitempty" yaml:"bindings,omitempty"`
}

func (value *Channel) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{}, 5+len(value.Extensions))
	for k, v := range value.Extensions {
		m[k] = v
	}

	m["description"] = value.Description
	if value.Subscribe != nil {
		m["subscribe"] = value.Subscribe
	}
	if value.Publish != nil {
		m["publish"] = value.Publish
	}
	if len(value.Parameters) > 0 {
		m["parameters"] = value.Parameters
	}
	if value.Bindings != nil {
		m["bindings"] = value.Bindings
	}

	return json.Marshal(m)
}

func (value *Channel) UnmarshalJSON(data []byte) error {
	type ChannelBis Channel
	var x ChannelBis
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	_ = json.Unmarshal(data, &x.Extensions)

	delete(x.Extensions, "description")
	delete(x.Extensions, "subscribe")
	delete(x.Extensions, "publish")
	delete(x.Extensions, "parameters")
	delete(x.Extensions, "bindings")

	*value = Channel(x)

	return nil
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
