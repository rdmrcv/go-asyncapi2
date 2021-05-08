package spec

import (
	"context"
	"fmt"

	"github.com/getkin/kin-openapi/jsoninfo"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-openapi/jsonpointer"
	"github.com/ligser/asyncapi/spec/bindings"
)

type ServersBindings map[string]*ServerBindings

var _ jsonpointer.JSONPointable = (*ServersBindings)(nil)

func (h ServersBindings) JSONLookup(token string) (interface{}, error) {
	value, ok := h[token]
	if value == nil || !ok {
		return nil, fmt.Errorf("object has no field %q", token)
	}

	return value, nil
}

// ServerBindings is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#server-bindings-object
type ServerBindings struct {
	openapi3.ExtensionProps
	Http  *bindings.HttpServer `json:"http" yaml:"http"`
	Ws    *bindings.WsServer   `json:"ws" yaml:"ws"`
	Kafka interface{}          `json:"kafka" yaml:"kafka"`
	Amqp  interface{}          `json:"amqp" yaml:"amqp"`
	Amqp1 interface{}          `json:"amqp1" yaml:"amqp1"`
	Mqtt  interface{}          `json:"mqtt" yaml:"mqtt"`
	Mqtt5 interface{}          `json:"mqtt5" yaml:"mqtt5"`
	Nats  interface{}          `json:"nats" yaml:"nats"`
	Jms   interface{}          `json:"jms" yaml:"jms"`
	Sns   interface{}          `json:"sns" yaml:"sns"`
	Sqs   interface{}          `json:"sqs" yaml:"sqs"`
	Stomp interface{}          `json:"stomp" yaml:"stomp"`
	Redis interface{}          `json:"redis" yaml:"redis"`
}

var _ jsonpointer.JSONPointable = (*ServerBindings)(nil)

func (value *ServerBindings) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStrictStruct(value)
}

func (value *ServerBindings) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStrictStruct(data, value)
}

func (value *ServerBindings) Validate(ctx context.Context) error {
	if v := value.Http; v != nil {
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	if v := value.Ws; v != nil {
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (value ServerBindings) JSONLookup(token string) (interface{}, error) {
	switch token {
	case "http":
		return value.Http, nil
	case "ws":
		return value.Ws, nil
	case "kafka":
		return value.Kafka, nil
	case "amqp":
		return value.Amqp, nil
	case "amqp1":
		return value.Amqp1, nil
	case "mqtt":
		return value.Mqtt, nil
	case "mqtt5":
		return value.Mqtt5, nil
	case "nats":
		return value.Nats, nil
	case "jms":
		return value.Jms, nil
	case "sns":
		return value.Sns, nil
	case "sqs":
		return value.Sqs, nil
	case "stomp":
		return value.Stomp, nil
	case "redis":
		return value.Redis, nil
	}

	v, _, err := jsonpointer.GetForToken(value.ExtensionProps, token)
	return v, err
}

type ChannelsBindings map[string]*ChannelBindings

var _ jsonpointer.JSONPointable = (*ChannelsBindings)(nil)

func (h ChannelsBindings) JSONLookup(token string) (interface{}, error) {
	value, ok := h[token]
	if value == nil || !ok {
		return nil, fmt.Errorf("object has no field %q", token)
	}

	return value, nil
}

// ChannelBindings is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#channel-bindings-object
type ChannelBindings struct {
	openapi3.ExtensionProps
	Http  *bindings.HttpChannel `json:"http" yaml:"http"`
	Ws    *bindings.WsChannel   `json:"ws" yaml:"ws"`
	Kafka interface{}           `json:"kafka" yaml:"kafka"`
	Amqp  interface{}           `json:"amqp" yaml:"amqp"`
	Amqp1 interface{}           `json:"amqp1" yaml:"amqp1"`
	Mqtt  interface{}           `json:"mqtt" yaml:"mqtt"`
	Mqtt5 interface{}           `json:"mqtt5" yaml:"mqtt5"`
	Nats  interface{}           `json:"nats" yaml:"nats"`
	Jms   interface{}           `json:"jms" yaml:"jms"`
	Sns   interface{}           `json:"sns" yaml:"sns"`
	Sqs   interface{}           `json:"sqs" yaml:"sqs"`
	Stomp interface{}           `json:"stomp" yaml:"stomp"`
	Redis interface{}           `json:"redis" yaml:"redis"`
}

var _ jsonpointer.JSONPointable = (*ChannelBindings)(nil)

func (value *ChannelBindings) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStrictStruct(value)
}

func (value *ChannelBindings) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStrictStruct(data, value)
}

func (value *ChannelBindings) Validate(ctx context.Context) error {
	if v := value.Http; v != nil {
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	if v := value.Ws; v != nil {
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (value ChannelBindings) JSONLookup(token string) (interface{}, error) {
	switch token {
	case "http":
		return value.Http, nil
	case "ws":
		return value.Ws, nil
	case "kafka":
		return value.Kafka, nil
	case "amqp":
		return value.Amqp, nil
	case "amqp1":
		return value.Amqp1, nil
	case "mqtt":
		return value.Mqtt, nil
	case "mqtt5":
		return value.Mqtt5, nil
	case "nats":
		return value.Nats, nil
	case "jms":
		return value.Jms, nil
	case "sns":
		return value.Sns, nil
	case "sqs":
		return value.Sqs, nil
	case "stomp":
		return value.Stomp, nil
	case "redis":
		return value.Redis, nil
	}

	v, _, err := jsonpointer.GetForToken(value.ExtensionProps, token)
	return v, err
}

type OperationsBindings map[string]*OperationBindings

var _ jsonpointer.JSONPointable = (*OperationsBindings)(nil)

func (h OperationsBindings) JSONLookup(token string) (interface{}, error) {
	value, ok := h[token]
	if value == nil || !ok {
		return nil, fmt.Errorf("object has no field %q", token)
	}

	return value, nil
}

// OperationBindings is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#operation-bindings-object
type OperationBindings struct {
	openapi3.ExtensionProps
	Http  *bindings.HttpOperation `json:"http" yaml:"http"`
	Ws    *bindings.WsOperation   `json:"ws" yaml:"ws"`
	Kafka interface{}             `json:"kafka" yaml:"kafka"`
	Amqp  interface{}             `json:"amqp" yaml:"amqp"`
	Amqp1 interface{}             `json:"amqp1" yaml:"amqp1"`
	Mqtt  interface{}             `json:"mqtt" yaml:"mqtt"`
	Mqtt5 interface{}             `json:"mqtt5" yaml:"mqtt5"`
	Nats  interface{}             `json:"nats" yaml:"nats"`
	Jms   interface{}             `json:"jms" yaml:"jms"`
	Sns   interface{}             `json:"sns" yaml:"sns"`
	Sqs   interface{}             `json:"sqs" yaml:"sqs"`
	Stomp interface{}             `json:"stomp" yaml:"stomp"`
	Redis interface{}             `json:"redis" yaml:"redis"`
}

var _ jsonpointer.JSONPointable = (*OperationBindings)(nil)

func (value *OperationBindings) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStrictStruct(value)
}

func (value *OperationBindings) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStrictStruct(data, value)
}

func (value *OperationBindings) Validate(ctx context.Context) error {
	if v := value.Http; v != nil {
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	if v := value.Ws; v != nil {
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (value OperationBindings) JSONLookup(token string) (interface{}, error) {
	switch token {
	case "http":
		return value.Http, nil
	case "ws":
		return value.Ws, nil
	case "kafka":
		return value.Kafka, nil
	case "amqp":
		return value.Amqp, nil
	case "amqp1":
		return value.Amqp1, nil
	case "mqtt":
		return value.Mqtt, nil
	case "mqtt5":
		return value.Mqtt5, nil
	case "nats":
		return value.Nats, nil
	case "jms":
		return value.Jms, nil
	case "sns":
		return value.Sns, nil
	case "sqs":
		return value.Sqs, nil
	case "stomp":
		return value.Stomp, nil
	case "redis":
		return value.Redis, nil
	}

	v, _, err := jsonpointer.GetForToken(value.ExtensionProps, token)
	return v, err
}

type MessagesBindings map[string]*MessageBindings

var _ jsonpointer.JSONPointable = (*MessagesBindings)(nil)

func (h MessagesBindings) JSONLookup(token string) (interface{}, error) {
	value, ok := h[token]
	if value == nil || !ok {
		return nil, fmt.Errorf("object has no field %q", token)
	}

	return value, nil
}

// MessageBindings is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#message-bindings-object
type MessageBindings struct {
	openapi3.ExtensionProps
	Http  *bindings.HttpMessage `json:"http" yaml:"http"`
	Ws    *bindings.WsMessage   `json:"ws" yaml:"ws"`
	Kafka interface{}           `json:"kafka" yaml:"kafka"`
	Amqp  interface{}           `json:"amqp" yaml:"amqp"`
	Amqp1 interface{}           `json:"amqp1" yaml:"amqp1"`
	Mqtt  interface{}           `json:"mqtt" yaml:"mqtt"`
	Mqtt5 interface{}           `json:"mqtt5" yaml:"mqtt5"`
	Nats  interface{}           `json:"nats" yaml:"nats"`
	Jms   interface{}           `json:"jms" yaml:"jms"`
	Sns   interface{}           `json:"sns" yaml:"sns"`
	Sqs   interface{}           `json:"sqs" yaml:"sqs"`
	Stomp interface{}           `json:"stomp" yaml:"stomp"`
	Redis interface{}           `json:"redis" yaml:"redis"`
}

var _ jsonpointer.JSONPointable = (*MessageBindings)(nil)

func (value *MessageBindings) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStrictStruct(value)
}

func (value *MessageBindings) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStrictStruct(data, value)
}

func (value *MessageBindings) Validate(ctx context.Context) error {
	if v := value.Http; v != nil {
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	if v := value.Ws; v != nil {
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (value MessageBindings) JSONLookup(token string) (interface{}, error) {
	switch token {
	case "http":
		return value.Http, nil
	case "ws":
		return value.Ws, nil
	case "kafka":
		return value.Kafka, nil
	case "amqp":
		return value.Amqp, nil
	case "amqp1":
		return value.Amqp1, nil
	case "mqtt":
		return value.Mqtt, nil
	case "mqtt5":
		return value.Mqtt5, nil
	case "nats":
		return value.Nats, nil
	case "jms":
		return value.Jms, nil
	case "sns":
		return value.Sns, nil
	case "sqs":
		return value.Sqs, nil
	case "stomp":
		return value.Stomp, nil
	case "redis":
		return value.Redis, nil
	}

	v, _, err := jsonpointer.GetForToken(value.ExtensionProps, token)
	return v, err
}
