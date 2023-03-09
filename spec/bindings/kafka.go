package bindings

import (
	"context"
	"encoding/json"

	"github.com/getkin/kin-openapi/openapi3"
)

// KafkaServer is defined in AsyncAPI spec: https://github.com/asyncapi/bindings/tree/master/kafka#server-binding-object
type KafkaServer struct {
}

func (binding *KafkaServer) Validate(ctx context.Context) error {
	return nil
}

// KafkaChannel is defined in AsyncAPI spec: https://github.com/asyncapi/bindings/tree/master/kafka#channel-binding-object
type KafkaChannel struct {
}

func (*KafkaChannel) Validate(context.Context) error {
	return nil
}

// KafkaOperation is defined in AsyncAPI spec: https://github.com/asyncapi/bindings/tree/master/kafka#operation-binding-object
type KafkaOperation struct {
	Extensions map[string]interface{} `json:"-" yaml:"-"`

	GroupID        *openapi3.Schema `json:"groupId" yaml:"groupId"`
	ClientID       *openapi3.Schema `json:"clientId" yaml:"clientId"`
	BindingVersion string           `json:"bindingVersion" yaml:"bindingVersion"`
}

func (binding *KafkaOperation) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{}, 3+len(binding.Extensions))
	for k, v := range binding.Extensions {
		m[k] = v
	}

	if binding.GroupID != nil {
		m["groupId"] = binding.GroupID
	}
	if binding.ClientID != nil {
		m["clientId"] = binding.ClientID
	}
	if len(binding.BindingVersion) != 0 {
		m["bindingVersion"] = binding.BindingVersion
	}

	return json.Marshal(m)
}

func (binding *KafkaOperation) UnmarshalJSON(data []byte) error {
	type KafkaOperationBis KafkaOperation
	var x KafkaOperationBis
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	_ = json.Unmarshal(data, &x.Extensions)

	delete(x.Extensions, "groupId")
	delete(x.Extensions, "clientId")
	delete(x.Extensions, "bindingVersion")

	*binding = KafkaOperation(x)

	return nil
}

func (binding *KafkaOperation) Validate(ctx context.Context) error {
	if v := binding.GroupID; v != nil {
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	if v := binding.ClientID; v != nil {
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	return nil
}

// KafkaMessage is defined in AsyncAPI spec: https://github.com/asyncapi/bindings/tree/master/kafka#message-binding-object
type KafkaMessage struct {
	Extensions map[string]interface{} `json:"-" yaml:"-"`

	Key            *openapi3.Schema `json:"key" yaml:"key"`
	BindingVersion string           `json:"bindingVersion" yaml:"bindingVersion"`
}

func (value *KafkaMessage) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{}, 2+len(value.Extensions))
	for k, v := range value.Extensions {
		m[k] = v
	}

	if value.Key != nil {
		m["key"] = value.Key
	}
	if len(value.BindingVersion) != 0 {
		m["bindingVersion"] = value.BindingVersion
	}

	return json.Marshal(m)
}

func (value *KafkaMessage) UnmarshalJSON(data []byte) error {
	type KafkaMessageBis KafkaMessage
	var x KafkaMessageBis
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	_ = json.Unmarshal(data, &x.Extensions)

	delete(x.Extensions, "key")
	delete(x.Extensions, "bindingVersion")

	*value = KafkaMessage(x)

	return nil
}

func (value *KafkaMessage) Validate(ctx context.Context) error {
	if v := value.Key; v != nil {
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	return nil
}
