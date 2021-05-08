package spec

import (
	"context"
	"fmt"
	"regexp"

	"github.com/getkin/kin-openapi/jsoninfo"
	"github.com/getkin/kin-openapi/openapi3"
)

// Components is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#componentsObject
type Components struct {
	openapi3.ExtensionProps

	Schemas           openapi3.Schemas   `json:"schemas" yaml:"schemas"`
	Messages          Messages           `json:"messages" yaml:"messages"`
	SecuritySchemes   SecuritySchemes    `json:"securitySchemes" yaml:"securitySchemes"`
	Parameters        Parameters         `json:"parameters" yaml:"parameters"`
	CorrelationIds    CorrelationIDs     `json:"correlationIds" yaml:"correlationIds"`
	OperationTraits   OperationsTraits   `json:"operationTraits" yaml:"operationTraits"`
	MessageTraits     MessagesTraits     `json:"messageTraits" yaml:"messageTraits"`
	ServerBindings    ServersBindings    `json:"serverBindings" yaml:"serverBindings"`
	ChannelBindings   ChannelsBindings   `json:"channelBindings" yaml:"channelBindings"`
	OperationBindings OperationsBindings `json:"operationBindings" yaml:"operationBindings"`
	MessageBindings   MessagesBindings   `json:"messageBindings" yaml:"messageBindings"`
}

func (components *Components) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStrictStruct(components)
}

func (components *Components) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStrictStruct(data, components)
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

	for k, v := range components.Parameters {
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
