package bindings

import (
	"context"

	"github.com/getkin/kin-openapi/jsoninfo"
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

var _ jsoninfo.StrictStruct = &KafkaOperation{}

// KafkaOperation is defined in AsyncAPI spec: https://github.com/asyncapi/bindings/tree/master/kafka#operation-binding-object
type KafkaOperation struct {
	openapi3.ExtensionProps
	GroupID        *openapi3.Schema `json:"groupId" yaml:"groupId"`
	ClientID       *openapi3.Schema `json:"clientId" yaml:"clientId"`
	bindingVersion string           `json:"bindingVersion" yaml:"bindingVersion"`
}

func (binding *KafkaOperation) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStrictStruct(binding)
}

func (binding *KafkaOperation) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStrictStruct(data, binding)
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

var _ jsoninfo.StrictStruct = &KafkaMessage{}

// KafkaMessage is defined in AsyncAPI spec: https://github.com/asyncapi/bindings/tree/master/kafka#message-binding-object
type KafkaMessage struct {
	openapi3.ExtensionProps
	Key            *openapi3.Schema `json:"key" yaml:"key"`
	bindingVersion string           `json:"bindingVersion" yaml:"bindingVersion"`
}

func (value *KafkaMessage) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStrictStruct(value)
}

func (value *KafkaMessage) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStrictStruct(data, value)
}

func (value *KafkaMessage) Validate(ctx context.Context) error {
	if v := value.Key; v != nil {
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	return nil
}
