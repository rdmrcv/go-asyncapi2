package spec

import (
	"context"
	"fmt"

	"github.com/getkin/kin-openapi/jsoninfo"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-openapi/jsonpointer"
	"github.com/ligser/asyncapi/spec/validate"
)

type SecuritySchemes map[string]*SecurityScheme

var _ jsonpointer.JSONPointable = (*SecuritySchemes)(nil)

func (h SecuritySchemes) JSONLookup(token string) (interface{}, error) {
	value, ok := h[token]
	if value == nil || !ok {
		return nil, fmt.Errorf("object has no field %q", token)
	}

	return value, nil
}

// OAuthFlowObject is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#oauth-flow-object
type OAuthFlowObject struct {
	openapi3.ExtensionProps
	AuthorizationUrl string            `json:"authorizationUrl,omitempty" yaml:"authorizationUrl,omitempty"`
	TokenUrl         string            `json:"tokenUrl,omitempty" yaml:"tokenUrl,omitempty"`
	RefreshUrl       string            `json:"refreshUrl,omitempty" yaml:"refreshUrl,omitempty"`
	Scopes           map[string]string `json:"scopes" yaml:"scopes"`
}

func (value *OAuthFlowObject) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStrictStruct(value)
}

func (value *OAuthFlowObject) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStrictStruct(data, value)
}

func (value *OAuthFlowObject) Validate(ctx context.Context) error {
	if len(value.AuthorizationUrl) == 0 {
		return fmt.Errorf("field \"authorizationUrl\" is required: %w", validate.ErrWrongField)
	}

	if len(value.TokenUrl) == 0 {
		return fmt.Errorf("field \"tokenUrl\" is required: %w", validate.ErrWrongField)
	}

	if value.Scopes == nil {
		return fmt.Errorf("field \"scopes\" is required: %w", validate.ErrWrongField)
	}

	return nil
}

// OAuthFlows is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#oauthFlowsObject
type OAuthFlows struct {
	openapi3.ExtensionProps
	Implicit          *OAuthFlowObject `json:"implicit,omitempty" yaml:"implicit,omitempty"`
	Password          *OAuthFlowObject `json:"password,omitempty" yaml:"password,omitempty"`
	ClientCredentials *OAuthFlowObject `json:"clientCredentials,omitempty" yaml:"clientCredentials,omitempty"`
	AuthorizationCode *OAuthFlowObject `json:"authorizationCode,omitempty" yaml:"authorizationCode,omitempty"`
}

func (value *OAuthFlows) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStrictStruct(value)
}

func (value *OAuthFlows) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStrictStruct(data, value)
}

func (value *OAuthFlows) Validate(ctx context.Context) error {
	if v := value.Implicit; v != nil {
		return value.Implicit.Validate(ctx)
	}

	if v := value.Password; v != nil {
		return value.Implicit.Validate(ctx)
	}

	if v := value.ClientCredentials; v != nil {
		return value.Implicit.Validate(ctx)
	}

	if v := value.AuthorizationCode; v != nil {
		return value.Implicit.Validate(ctx)
	}

	return nil
}

// SecurityScheme is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#security-scheme-object
type SecurityScheme struct {
	openapi3.ExtensionProps
	Type             string      `json:"type" yaml:"type"`
	Description      string      `json:"description,omitempty" yaml:"description,omitempty"`
	Name             string      `json:"name,omitempty" yaml:"name,omitempty"`
	In               string      `json:"in,omitempty" yaml:"in,omitempty"`
	Scheme           string      `json:"scheme,omitempty" yaml:"scheme,omitempty"`
	BearerFormat     string      `json:"bearerFormat,omitempty" yaml:"bearerFormat,omitempty"`
	Flows            *OAuthFlows `json:"flows,omitempty" yaml:"flows,omitempty"`
	OpenIDConnectUrl string      `json:"openIdConnectUrl,omitempty" yaml:"openIdConnectUrl,omitempty"`
}

func (value *SecurityScheme) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStrictStruct(value)
}

func (value *SecurityScheme) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStrictStruct(data, value)
}

func (value *SecurityScheme) Validate(ctx context.Context) error {
	if len(value.Type) == 0 {
		return fmt.Errorf("field \"type\" is required: %w", validate.ErrWrongField)
	}

	if len(value.Name) == 0 {
		return fmt.Errorf("field \"name\" is required: %w", validate.ErrWrongField)
	}

	if len(value.In) == 0 {
		return fmt.Errorf("field \"in\" is required: %w", validate.ErrWrongField)
	}

	if v := value.Flows; v != nil {
		return v.Validate(ctx)
	} else {
		return fmt.Errorf("field \"flows\" is required: %w", validate.ErrWrongField)
	}

	if len(value.OpenIDConnectUrl) == 0 {
		return fmt.Errorf("field \"openIdConnectUrl\" is required: %w", validate.ErrWrongField)
	}

	return nil
}
