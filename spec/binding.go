package spec

import (
	"context"
	"encoding/json"

	"github.com/ligser/go-asyncapi2/spec/bindings"
)

type ServersBindings map[string]*ServerBindings

// ServerBindings is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#server-bindings-object
type ServerBindings struct {
	Extensions map[string]interface{} `json:"-" yaml:"-"`

	Http  *bindings.HttpServer  `json:"http,omitempty" yaml:"http,omitempty"`
	Ws    *bindings.WsServer    `json:"ws,omitempty" yaml:"ws,omitempty"`
	Kafka *bindings.KafkaServer `json:"kafka,omitempty" yaml:"kafka,omitempty"`

	Amqp  interface{} `json:"amqp,omitempty" yaml:"amqp,omitempty"`
	Amqp1 interface{} `json:"amqp1,omitempty" yaml:"amqp1,omitempty"`
	Mqtt  interface{} `json:"mqtt,omitempty" yaml:"mqtt,omitempty"`
	Mqtt5 interface{} `json:"mqtt5,omitempty" yaml:"mqtt5,omitempty"`
	Nats  interface{} `json:"nats,omitempty" yaml:"nats,omitempty"`
	Jms   interface{} `json:"jms,omitempty" yaml:"jms,omitempty"`
	Sns   interface{} `json:"sns,omitempty" yaml:"sns,omitempty"`
	Sqs   interface{} `json:"sqs,omitempty" yaml:"sqs,omitempty"`
	Stomp interface{} `json:"stomp,omitempty" yaml:"stomp,omitempty"`
	Redis interface{} `json:"redis,omitempty" yaml:"redis,omitempty"`
}

func (value *ServerBindings) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{}, 13+len(value.Extensions))
	for k, v := range value.Extensions {
		m[k] = v
	}

	if value.Http != nil {
		m["http"] = value.Http
	}
	if value.Ws != nil {
		m["ws"] = value.Ws
	}
	if value.Kafka != nil {
		m["kafka"] = value.Kafka
	}
	if value.Amqp != nil {
		m["amqp"] = value.Amqp
	}
	if value.Amqp1 != nil {
		m["amqp1"] = value.Amqp1
	}
	if value.Mqtt != nil {
		m["mqtt"] = value.Mqtt
	}
	if value.Mqtt5 != nil {
		m["mqtt5"] = value.Mqtt5
	}
	if value.Nats != nil {
		m["nats"] = value.Nats
	}
	if value.Jms != nil {
		m["jms"] = value.Jms
	}
	if value.Sns != nil {
		m["sns"] = value.Sns
	}
	if value.Sqs != nil {
		m["sqs"] = value.Sqs
	}
	if value.Stomp != nil {
		m["stomp"] = value.Stomp
	}
	if value.Redis != nil {
		m["redis"] = value.Redis
	}

	return json.Marshal(m)
}

func (value *ServerBindings) UnmarshalJSON(data []byte) error {
	type ServerBindingsBis ServerBindings
	var x ServerBindingsBis
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	_ = json.Unmarshal(data, &x.Extensions)

	delete(x.Extensions, "http")
	delete(x.Extensions, "ws")
	delete(x.Extensions, "kafka")
	delete(x.Extensions, "amqp")
	delete(x.Extensions, "amqp1")
	delete(x.Extensions, "mqtt")
	delete(x.Extensions, "mqtt5")
	delete(x.Extensions, "nats")
	delete(x.Extensions, "jms")
	delete(x.Extensions, "sns")
	delete(x.Extensions, "sqs")
	delete(x.Extensions, "stomp")
	delete(x.Extensions, "redis")

	*value = ServerBindings(x)

	return nil
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

	if v := value.Kafka; v != nil {
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	return nil
}

type ChannelsBindings map[string]*ChannelBindings

// ChannelBindings is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#channel-bindings-object
type ChannelBindings struct {
	Extensions map[string]interface{} `json:"-" yaml:"-"`

	Http  *bindings.HttpChannel  `json:"http,omitempty" yaml:"http,omitempty"`
	Ws    *bindings.WsChannel    `json:"ws,omitempty" yaml:"ws,omitempty"`
	Kafka *bindings.KafkaChannel `json:"kafka,omitempty" yaml:"kafka,omitempty"`
	Amqp  interface{}            `json:"amqp,omitempty" yaml:"amqp,omitempty"`
	Amqp1 interface{}            `json:"amqp1,omitempty" yaml:"amqp1,omitempty"`
	Mqtt  interface{}            `json:"mqtt,omitempty" yaml:"mqtt,omitempty"`
	Mqtt5 interface{}            `json:"mqtt5,omitempty" yaml:"mqtt5,omitempty"`
	Nats  interface{}            `json:"nats,omitempty" yaml:"nats,omitempty"`
	Jms   interface{}            `json:"jms,omitempty" yaml:"jms,omitempty"`
	Sns   interface{}            `json:"sns,omitempty" yaml:"sns,omitempty"`
	Sqs   interface{}            `json:"sqs,omitempty" yaml:"sqs,omitempty"`
	Stomp interface{}            `json:"stomp,omitempty" yaml:"stomp,omitempty"`
	Redis interface{}            `json:"redis,omitempty" yaml:"redis,omitempty"`
}

func (value *ChannelBindings) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{}, 13+len(value.Extensions))
	for k, v := range value.Extensions {
		m[k] = v
	}

	if value.Http != nil {
		m["http"] = value.Http
	}
	if value.Ws != nil {
		m["ws"] = value.Ws
	}
	if value.Kafka != nil {
		m["kafka"] = value.Kafka
	}
	if value.Amqp != nil {
		m["amqp"] = value.Amqp
	}
	if value.Amqp1 != nil {
		m["amqp1"] = value.Amqp1
	}
	if value.Mqtt != nil {
		m["mqtt"] = value.Mqtt
	}
	if value.Mqtt5 != nil {
		m["mqtt5"] = value.Mqtt5
	}
	if value.Nats != nil {
		m["nats"] = value.Nats
	}
	if value.Jms != nil {
		m["jms"] = value.Jms
	}
	if value.Sns != nil {
		m["sns"] = value.Sns
	}
	if value.Sqs != nil {
		m["sqs"] = value.Sqs
	}
	if value.Stomp != nil {
		m["stomp"] = value.Stomp
	}
	if value.Redis != nil {
		m["redis"] = value.Redis
	}

	return json.Marshal(m)
}

func (value *ChannelBindings) UnmarshalJSON(data []byte) error {
	type ChannelBindingsBis ChannelBindings
	var x ChannelBindingsBis
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	_ = json.Unmarshal(data, &x.Extensions)

	delete(x.Extensions, "http")
	delete(x.Extensions, "ws")
	delete(x.Extensions, "kafka")
	delete(x.Extensions, "amqp")
	delete(x.Extensions, "amqp1")
	delete(x.Extensions, "mqtt")
	delete(x.Extensions, "mqtt5")
	delete(x.Extensions, "nats")
	delete(x.Extensions, "jms")
	delete(x.Extensions, "sns")
	delete(x.Extensions, "sqs")
	delete(x.Extensions, "stomp")
	delete(x.Extensions, "redis")

	*value = ChannelBindings(x)

	return nil
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

type OperationsBindings map[string]*OperationBindings

// OperationBindings is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#operation-bindings-object
type OperationBindings struct {
	Extensions map[string]interface{} `json:"-" yaml:"-"`

	Http  *bindings.HttpOperation  `json:"http,omitempty" yaml:"http,omitempty"`
	Ws    *bindings.WsOperation    `json:"ws,omitempty" yaml:"ws,omitempty"`
	Kafka *bindings.KafkaOperation `json:"kafka,omitempty" yaml:"kafka,omitempty"`

	Amqp  interface{} `json:"amqp,omitempty" yaml:"amqp,omitempty"`
	Amqp1 interface{} `json:"amqp1,omitempty" yaml:"amqp1,omitempty"`
	Mqtt  interface{} `json:"mqtt,omitempty" yaml:"mqtt,omitempty"`
	Mqtt5 interface{} `json:"mqtt5,omitempty" yaml:"mqtt5,omitempty"`
	Nats  interface{} `json:"nats,omitempty" yaml:"nats,omitempty"`
	Jms   interface{} `json:"jms,omitempty" yaml:"jms,omitempty"`
	Sns   interface{} `json:"sns,omitempty" yaml:"sns,omitempty"`
	Sqs   interface{} `json:"sqs,omitempty" yaml:"sqs,omitempty"`
	Stomp interface{} `json:"stomp,omitempty" yaml:"stomp,omitempty"`
	Redis interface{} `json:"redis,omitempty" yaml:"redis,omitempty"`
}

func (value *OperationBindings) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{}, 13+len(value.Extensions))
	for k, v := range value.Extensions {
		m[k] = v
	}

	if value.Http != nil {
		m["http"] = value.Http
	}
	if value.Ws != nil {
		m["ws"] = value.Ws
	}
	if value.Kafka != nil {
		m["kafka"] = value.Kafka
	}
	if value.Amqp != nil {
		m["amqp"] = value.Amqp
	}
	if value.Amqp1 != nil {
		m["amqp1"] = value.Amqp1
	}
	if value.Mqtt != nil {
		m["mqtt"] = value.Mqtt
	}
	if value.Mqtt5 != nil {
		m["mqtt5"] = value.Mqtt5
	}
	if value.Nats != nil {
		m["nats"] = value.Nats
	}
	if value.Jms != nil {
		m["jms"] = value.Jms
	}
	if value.Sns != nil {
		m["sns"] = value.Sns
	}
	if value.Sqs != nil {
		m["sqs"] = value.Sqs
	}
	if value.Stomp != nil {
		m["stomp"] = value.Stomp
	}
	if value.Redis != nil {
		m["redis"] = value.Redis
	}

	return json.Marshal(m)
}

func (value *OperationBindings) UnmarshalJSON(data []byte) error {
	type OperationBindingsBis OperationBindings
	var x OperationBindingsBis
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	_ = json.Unmarshal(data, &x.Extensions)

	delete(x.Extensions, "http")
	delete(x.Extensions, "ws")
	delete(x.Extensions, "kafka")
	delete(x.Extensions, "amqp")
	delete(x.Extensions, "amqp1")
	delete(x.Extensions, "mqtt")
	delete(x.Extensions, "mqtt5")
	delete(x.Extensions, "nats")
	delete(x.Extensions, "jms")
	delete(x.Extensions, "sns")
	delete(x.Extensions, "sqs")
	delete(x.Extensions, "stomp")
	delete(x.Extensions, "redis")

	*value = OperationBindings(x)

	return nil
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

	if v := value.Kafka; v != nil {
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	return nil
}

type MessagesBindings map[string]*MessageBindings

// MessageBindings is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#message-bindings-object
type MessageBindings struct {
	Extensions map[string]interface{} `json:"-" yaml:"-"`

	Http  *bindings.HttpMessage  `json:"http,omitempty" yaml:"http,omitempty"`
	Ws    *bindings.WsMessage    `json:"ws,omitempty" yaml:"ws,omitempty"`
	Kafka *bindings.KafkaMessage `json:"kafka,omitempty" yaml:"kafka,omitempty"`

	Amqp  interface{} `json:"amqp,omitempty" yaml:"amqp,omitempty"`
	Amqp1 interface{} `json:"amqp1,omitempty" yaml:"amqp1,omitempty"`
	Mqtt  interface{} `json:"mqtt,omitempty" yaml:"mqtt,omitempty"`
	Mqtt5 interface{} `json:"mqtt5,omitempty" yaml:"mqtt5,omitempty"`
	Nats  interface{} `json:"nats,omitempty" yaml:"nats,omitempty"`
	Jms   interface{} `json:"jms,omitempty" yaml:"jms,omitempty"`
	Sns   interface{} `json:"sns,omitempty" yaml:"sns,omitempty"`
	Sqs   interface{} `json:"sqs,omitempty" yaml:"sqs,omitempty"`
	Stomp interface{} `json:"stomp,omitempty" yaml:"stomp,omitempty"`
	Redis interface{} `json:"redis,omitempty" yaml:"redis,omitempty"`
}

func (value *MessageBindings) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{}, 13+len(value.Extensions))
	for k, v := range value.Extensions {
		m[k] = v
	}

	if value.Http != nil {
		m["http"] = value.Http
	}
	if value.Ws != nil {
		m["ws"] = value.Ws
	}
	if value.Kafka != nil {
		m["kafka"] = value.Kafka
	}
	if value.Amqp != nil {
		m["amqp"] = value.Amqp
	}
	if value.Amqp1 != nil {
		m["amqp1"] = value.Amqp1
	}
	if value.Mqtt != nil {
		m["mqtt"] = value.Mqtt
	}
	if value.Mqtt5 != nil {
		m["mqtt5"] = value.Mqtt5
	}
	if value.Nats != nil {
		m["nats"] = value.Nats
	}
	if value.Jms != nil {
		m["jms"] = value.Jms
	}
	if value.Sns != nil {
		m["sns"] = value.Sns
	}
	if value.Sqs != nil {
		m["sqs"] = value.Sqs
	}
	if value.Stomp != nil {
		m["stomp"] = value.Stomp
	}
	if value.Redis != nil {
		m["redis"] = value.Redis
	}

	return json.Marshal(m)
}

func (value *MessageBindings) UnmarshalJSON(data []byte) error {
	type MessageBindingsBis MessageBindings
	var x MessageBindingsBis
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	_ = json.Unmarshal(data, &x.Extensions)

	delete(x.Extensions, "http")
	delete(x.Extensions, "ws")
	delete(x.Extensions, "kafka")
	delete(x.Extensions, "amqp")
	delete(x.Extensions, "amqp1")
	delete(x.Extensions, "mqtt")
	delete(x.Extensions, "mqtt5")
	delete(x.Extensions, "nats")
	delete(x.Extensions, "jms")
	delete(x.Extensions, "sns")
	delete(x.Extensions, "sqs")
	delete(x.Extensions, "stomp")
	delete(x.Extensions, "redis")

	*value = MessageBindings(x)

	return nil
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

	if v := value.Kafka; v != nil {
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	return nil
}
