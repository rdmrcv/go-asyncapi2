package spec

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/getkin/kin-openapi/openapi3"
)

// Components scheme is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#componentsObject
type Components struct {
	Extensions map[string]interface{} `json:"-" yaml:"-"`

	Schemas           openapi3.Schemas   `json:"schemas,omitempty" yaml:"schemas,omitempty"`
	Messages          Messages           `json:"messages,omitempty" yaml:"messages,omitempty"`
	SecuritySchemes   SecuritySchemes    `json:"securitySchemes,omitempty" yaml:"securitySchemes,omitempty"`
	Parameters        Parameters         `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	CorrelationIds    CorrelationIDs     `json:"correlationIds,omitempty" yaml:"correlationIds,omitempty"`
	OperationTraits   OperationsTraits   `json:"operationTraits,omitempty" yaml:"operationTraits,omitempty"`
	MessageTraits     MessagesTraits     `json:"messageTraits,omitempty" yaml:"messageTraits,omitempty"`
	ServerBindings    ServersBindings    `json:"serverBindings,omitempty" yaml:"serverBindings,omitempty"`
	ChannelBindings   ChannelsBindings   `json:"channelBindings,omitempty" yaml:"channelBindings,omitempty"`
	OperationBindings OperationsBindings `json:"operationBindings,omitempty" yaml:"operationBindings,omitempty"`
	MessageBindings   MessagesBindings   `json:"messageBindings,omitempty" yaml:"messageBindings,omitempty"`
}

func (components *Components) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{}, 11+len(components.Extensions))
	for k, v := range components.Extensions {
		m[k] = v
	}

	if len(components.Schemas) != 0 {
		m["schemas"] = components.Schemas
	}
	if len(components.Messages) != 0 {
		m["messages"] = components.Messages
	}
	if len(components.SecuritySchemes) != 0 {
		m["securitySchemes"] = components.SecuritySchemes
	}
	if len(components.Parameters) != 0 {
		m["parameters"] = components.Parameters
	}
	if len(components.CorrelationIds) != 0 {
		m["correlationIds"] = components.CorrelationIds
	}
	if len(components.OperationTraits) != 0 {
		m["operationTraits"] = components.OperationTraits
	}
	if len(components.MessageTraits) != 0 {
		m["messageTraits"] = components.MessageTraits
	}
	if len(components.ServerBindings) != 0 {
		m["serverBindings"] = components.ServerBindings
	}
	if len(components.ChannelBindings) != 0 {
		m["channelBindings"] = components.ChannelBindings
	}
	if len(components.OperationBindings) != 0 {
		m["operationBindings"] = components.OperationBindings
	}
	if len(components.MessageBindings) != 0 {
		m["messageBindings"] = components.MessageBindings
	}

	return json.Marshal(m)
}

func (components *Components) UnmarshalJSON(data []byte) error {
	type ComponentsBis Components
	var x ComponentsBis
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	_ = json.Unmarshal(data, &x.Extensions)

	delete(x.Extensions, "schemas")
	delete(x.Extensions, "messages")
	delete(x.Extensions, "securitySchemes")
	delete(x.Extensions, "parameters")
	delete(x.Extensions, "correlationIds")
	delete(x.Extensions, "operationTraits")
	delete(x.Extensions, "messageTraits")
	delete(x.Extensions, "serverBindings")
	delete(x.Extensions, "channelBindings")
	delete(x.Extensions, "operationBindings")
	delete(x.Extensions, "messageBindings")

	*components = Components(x)

	return nil
}

func (components *Components) Validate(ctx context.Context) (err error) {
	for k, v := range components.Schemas {
		if err = ValidateIdentifier(k); err != nil {
			return
		}
		if err = v.Validate(ctx); err != nil {
			return
		}
	}

	for k, v := range components.Messages {
		if err = ValidateIdentifier(k); err != nil {
			return
		}
		if err = v.Validate(ctx); err != nil {
			return
		}
	}

	for k, v := range components.SecuritySchemes {
		if err = ValidateIdentifier(k); err != nil {
			return
		}
		if err = v.Validate(ctx); err != nil {
			return
		}
	}

	for k, v := range components.Parameters {
		if err = ValidateIdentifier(k); err != nil {
			return
		}
		if err = v.Validate(ctx); err != nil {
			return
		}
	}

	for k, v := range components.CorrelationIds {
		if err = ValidateIdentifier(k); err != nil {
			return
		}
		if err = v.Validate(ctx); err != nil {
			return
		}
	}

	for k, v := range components.OperationTraits {
		if err = ValidateIdentifier(k); err != nil {
			return
		}
		if err = v.Validate(ctx); err != nil {
			return
		}
	}

	for k, v := range components.MessageTraits {
		if err = ValidateIdentifier(k); err != nil {
			return
		}
		if err = v.Validate(ctx); err != nil {
			return
		}
	}

	for k, v := range components.ServerBindings {
		if err = ValidateIdentifier(k); err != nil {
			return
		}
		if err = v.Validate(ctx); err != nil {
			return
		}
	}

	for k, v := range components.ChannelBindings {
		if err = ValidateIdentifier(k); err != nil {
			return
		}
		if err = v.Validate(ctx); err != nil {
			return
		}
	}

	for k, v := range components.OperationBindings {
		if err = ValidateIdentifier(k); err != nil {
			return
		}
		if err = v.Validate(ctx); err != nil {
			return
		}
	}

	for k, v := range components.MessageBindings {
		if err = ValidateIdentifier(k); err != nil {
			return
		}
		if err = v.Validate(ctx); err != nil {
			return
		}
	}

	return
}

const identifierPattern = `^[a-zA-Z0-9._-]+$`

// IdentifierRegExp verifies whether Component object key matches 'identifierPattern' pattern, according to OpenAPI v3.x.0.
// However, to be able supporting legacy OpenAPI v2.x, there is a need to customize above pattern in order not to fail
// converted v2-v3 validation
var IdentifierRegExp = regexp.MustCompile(identifierPattern)

func ValidateIdentifier(value string) error {
	if IdentifierRegExp.MatchString(value) {
		return nil
	}

	return fmt.Errorf("identifier %q is not supported by OpenAPIv3 standard (regexp: %q)", value, identifierPattern)
}
