package spec

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ligser/go-asyncapi2/spec/validate"
)

type SecuritySchemes map[string]*SecurityScheme

// OAuthFlowObject is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#oauth-flow-object
type OAuthFlowObject struct {
	Extensions map[string]interface{} `json:"-" yaml:"-"`

	AuthorizationUrl string            `json:"authorizationUrl,omitempty" yaml:"authorizationUrl,omitempty"`
	TokenUrl         string            `json:"tokenUrl,omitempty" yaml:"tokenUrl,omitempty"`
	RefreshUrl       string            `json:"refreshUrl,omitempty" yaml:"refreshUrl,omitempty"`
	Scopes           map[string]string `json:"scopes" yaml:"scopes"`
}

func (value *OAuthFlowObject) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{}, 6+len(value.Extensions))
	for k, v := range value.Extensions {
		m[k] = v
	}

	if len(value.AuthorizationUrl) != 0 {
		m["authorizationUrl"] = value.AuthorizationUrl
	}
	if len(value.TokenUrl) != 0 {
		m["tokenUrl"] = value.TokenUrl
	}
	if len(value.RefreshUrl) != 0 {
		m["refreshUrl"] = value.RefreshUrl
	}
	if len(value.Scopes) != 0 {
		m["scopes"] = value.Scopes
	}

	return json.Marshal(m)
}

func (value *OAuthFlowObject) UnmarshalJSON(data []byte) error {
	type OAuthFlowObjectBis OAuthFlowObject
	var x OAuthFlowObjectBis
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	_ = json.Unmarshal(data, &x.Extensions)

	delete(x.Extensions, "authorizationUrl")
	delete(x.Extensions, "tokenUrl")
	delete(x.Extensions, "refreshUrl")
	delete(x.Extensions, "scopes")

	*value = OAuthFlowObject(x)

	return nil
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
	Extensions map[string]interface{} `json:"-" yaml:"-"`

	Implicit          *OAuthFlowObject `json:"implicit,omitempty" yaml:"implicit,omitempty"`
	Password          *OAuthFlowObject `json:"password,omitempty" yaml:"password,omitempty"`
	ClientCredentials *OAuthFlowObject `json:"clientCredentials,omitempty" yaml:"clientCredentials,omitempty"`
	AuthorizationCode *OAuthFlowObject `json:"authorizationCode,omitempty" yaml:"authorizationCode,omitempty"`
}

func (value *OAuthFlows) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{}, 4+len(value.Extensions))
	for k, v := range value.Extensions {
		m[k] = v
	}

	if value.Implicit != nil {
		m["implicit"] = value.Implicit
	}
	if value.Password != nil {
		m["password"] = value.Password
	}
	if value.ClientCredentials != nil {
		m["clientCredentials"] = value.ClientCredentials
	}
	if value.AuthorizationCode != nil {
		m["authorizationCode"] = value.AuthorizationCode
	}

	return json.Marshal(m)
}

func (value *OAuthFlows) UnmarshalJSON(data []byte) error {
	type OAuthFlowsBis OAuthFlows
	var x OAuthFlowsBis
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	_ = json.Unmarshal(data, &x.Extensions)

	delete(x.Extensions, "implicit")
	delete(x.Extensions, "password")
	delete(x.Extensions, "clientCredentials")
	delete(x.Extensions, "authorizationCode")

	*value = OAuthFlows(x)

	return nil
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

var validationsByType = map[string]func(value *SecurityScheme) error{
	"userPassword": func(value *SecurityScheme) error {
		return nil
	},
	"X509": func(value *SecurityScheme) error {
		return nil
	},
	"symmetricEncryption": func(value *SecurityScheme) error {
		return nil
	},
	"asymmetricEncryption": func(value *SecurityScheme) error {
		return nil
	},
	"plain": func(value *SecurityScheme) error {
		return nil
	},
	"scramSha256": func(value *SecurityScheme) error {
		return nil
	},
	"scramSha512": func(value *SecurityScheme) error {
		return nil
	},
	"gssapi": func(value *SecurityScheme) error {
		return nil
	},
	"httpApiKey": func(value *SecurityScheme) error {
		if len(value.Name) == 0 {
			return fmt.Errorf("field \"name\" is required: %w", validate.ErrWrongField)
		}
		if len(value.In) == 0 {
			return fmt.Errorf("field \"in\" is required: %w", validate.ErrWrongField)
		}

		return nil
	},
	"apiKey": func(value *SecurityScheme) error {
		if len(value.In) == 0 {
			return fmt.Errorf("field \"in\" is required: %w", validate.ErrWrongField)
		}

		return nil
	},
	"http": func(value *SecurityScheme) error {
		if len(value.Scheme) == 0 {
			return fmt.Errorf("field \"scheme\" is required: %w", validate.ErrWrongField)
		}
		if len(value.BearerFormat) == 0 {
			return fmt.Errorf("field \"bearerFormat\" is required: %w", validate.ErrWrongField)
		}

		return nil
	},
	"oauth2": func(value *SecurityScheme) error {
		if value.Flows == nil {
			return fmt.Errorf("field \"flows\" is required: %w", validate.ErrWrongField)
		}

		return nil
	},
	"openIdConnect": func(value *SecurityScheme) error {
		if len(value.OpenIDConnectUrl) != 0 {
			return fmt.Errorf("field \"openIdConnectUrl\" is required: %w", validate.ErrWrongField)
		}

		return nil
	},
}

// SecurityScheme is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#security-scheme-object
type SecurityScheme struct {
	Extensions map[string]interface{} `json:"-" yaml:"-"`

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
	m := make(map[string]interface{}, 8+len(value.Extensions))
	for k, v := range value.Extensions {
		m[k] = v
	}

	m["type"] = value.Type
	if len(value.Description) != 0 {
		m["description"] = value.Description
	}
	if len(value.Scheme) != 0 {
		m["scheme"] = value.Scheme
	}
	if len(value.Name) != 0 {
		m["name"] = value.Name
	}
	if len(value.In) != 0 {
		m["in"] = value.In
	}
	if len(value.BearerFormat) != 0 {
		m["bearerFormat"] = value.BearerFormat
	}
	if value.Flows != nil {
		m["flows"] = value.Flows
	}
	if len(value.OpenIDConnectUrl) != 0 {
		m["openIdConnectUrl"] = value.OpenIDConnectUrl
	}

	return json.Marshal(m)
}

func (value *SecurityScheme) UnmarshalJSON(data []byte) error {
	type SecuritySchemeBis SecurityScheme
	var x SecuritySchemeBis
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	_ = json.Unmarshal(data, &x.Extensions)

	delete(x.Extensions, "type")
	delete(x.Extensions, "description")
	delete(x.Extensions, "name")
	delete(x.Extensions, "in")
	delete(x.Extensions, "scheme")
	delete(x.Extensions, "bearerFormat")
	delete(x.Extensions, "flows")
	delete(x.Extensions, "openIdConnectUrl")

	*value = SecurityScheme(x)

	return nil
}

func (value *SecurityScheme) Validate(ctx context.Context) error {
	if len(value.Type) == 0 {
		return fmt.Errorf("field \"type\" is required: %w", validate.ErrWrongField)
	}

	validator, has := validationsByType[value.Type]
	if !has {
		return fmt.Errorf("field \"type\" is not expected: %w", validate.ErrWrongField)
	}

	return validator(value)
}
