package spec

import (
	"fmt"

	"github.com/getkin/kin-openapi/jsoninfo"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-openapi/jsonpointer"
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
	AuthorizationUrl string            `json:"authorizationUrl" yaml:"authorizationUrl"`
	TokenUrl         string            `json:"tokenUrl" yaml:"tokenUrl"`
	RefreshUrl       string            `json:"refreshUrl" yaml:"refreshUrl"`
	Scopes           map[string]string `json:"scopes" yaml:"scopes"`
}

func (value *OAuthFlowObject) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStrictStruct(value)
}

func (value *OAuthFlowObject) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStrictStruct(data, value)
}

// OAuthFlows is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#oauthFlowsObject
type OAuthFlows struct {
	openapi3.ExtensionProps
	Implicit          *OAuthFlowObject `json:"implicit" yaml:"implicit"`
	Password          *OAuthFlowObject `json:"password" yaml:"password"`
	ClientCredentials *OAuthFlowObject `json:"clientCredentials" yaml:"clientCredentials"`
	AuthorizationCode *OAuthFlowObject `json:"authorizationCode" yaml:"authorizationCode"`
}

func (value *OAuthFlows) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStrictStruct(value)
}

func (value *OAuthFlows) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStrictStruct(data, value)
}

// SecurityScheme is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#security-scheme-object
type SecurityScheme struct {
	openapi3.ExtensionProps
	Type             string      `json:"type" yaml:"type"`
	Description      string      `json:"description" yaml:"description"`
	Name             string      `json:"name" yaml:"name"`
	In               string      `json:"in" yaml:"in"`
	Scheme           string      `json:"scheme" yaml:"scheme"`
	BearerFormat     string      `json:"bearerFormat" yaml:"bearerFormat"`
	Flows            *OAuthFlows `json:"flows" yaml:"flows"`
	OpenIDConnectUrl string      `json:"openIdConnectUrl" yaml:"openIdConnectUrl"`
}

func (value *SecurityScheme) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStrictStruct(value)
}

func (value *SecurityScheme) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStrictStruct(data, value)
}
