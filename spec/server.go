package spec

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/url"
	"regexp"
	"strings"
)

var (
	serverKeyRegexp = regexp.MustCompile("^[A-Za-z0-9_\\-]+$")

	ErrServerKeyInvalid = fmt.Errorf("server should match pattern %q", serverKeyRegexp.String())
)

// Servers object is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#serversObject
type Servers map[string]*Server

func (servers Servers) Validate(ctx context.Context) error {
	for k, v := range servers {
		if !serverKeyRegexp.MatchString(k) {
			return ErrServerKeyInvalid
		}

		if err := v.Validate(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (servers Servers) MatchURL(parsedURL *url.URL) (*Server, []string, string) {
	rawURL := parsedURL.String()
	if i := strings.IndexByte(rawURL, '?'); i >= 0 {
		rawURL = rawURL[:i]
	}

	for _, server := range servers {
		pathParams, remaining, ok := server.MatchRawURL(rawURL)
		if ok {
			return server, pathParams, remaining
		}
	}

	return nil, nil, ""
}

// SecurityRequirements is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#securityRequirementObject
type SecurityRequirements map[string][]string

// Server is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#serverObject
type Server struct {
	Extensions map[string]interface{} `json:"-" yaml:"-"`

	URL             string                     `json:"url" yaml:"url"`
	Protocol        string                     `json:"protocol" yaml:"protocol"`
	ProtocolVersion string                     `json:"protocolVersion,omitempty" yaml:"protocolVersion,omitempty"`
	Description     string                     `json:"description,omitempty" yaml:"description,omitempty"`
	Security        []SecurityRequirements     `json:"security,omitempty" yaml:"security,omitempty"`
	Bindings        *ServerBindings            `json:"bindings,omitempty" yaml:"bindings,omitempty"`
	Variables       map[string]*ServerVariable `json:"variables,omitempty" yaml:"variables,omitempty"`
}

func (value *Server) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{}, 7+len(value.Extensions))
	for k, v := range value.Extensions {
		m[k] = v
	}

	m["url"] = value.URL
	m["protocol"] = value.Protocol
	if len(value.ProtocolVersion) != 0 {
		m["protocolVersion"] = value.ProtocolVersion
	}
	if len(value.Description) != 0 {
		m["description"] = value.Description
	}
	if len(value.Security) != 0 {
		m["security"] = value.Security
	}
	if value.Bindings != nil {
		m["bindings"] = value.Bindings
	}
	if len(value.Variables) != 0 {
		m["variables"] = value.Variables
	}

	return json.Marshal(m)
}

func (value *Server) UnmarshalJSON(data []byte) error {
	type ServerBis Server
	var x ServerBis
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	_ = json.Unmarshal(data, &x.Extensions)

	delete(x.Extensions, "url")
	delete(x.Extensions, "protocol")
	delete(x.Extensions, "protocolVersion")
	delete(x.Extensions, "description")
	delete(x.Extensions, "security")
	delete(x.Extensions, "bindings")
	delete(x.Extensions, "variables")

	*value = Server(x)

	return nil
}

func (value *Server) ParameterNames() ([]string, error) {
	pattern := value.URL
	var params []string
	for len(pattern) > 0 {
		i := strings.IndexByte(pattern, '{')
		if i < 0 {
			break
		}
		pattern = pattern[i+1:]
		i = strings.IndexByte(pattern, '}')
		if i < 0 {
			return nil, errors.New("missing '}'")
		}
		params = append(params, strings.TrimSpace(pattern[:i]))
		pattern = pattern[i+1:]
	}
	return params, nil
}

func (value *Server) MatchRawURL(input string) ([]string, string, bool) {
	pattern := value.URL
	var params []string
	for len(pattern) > 0 {
		c := pattern[0]
		if len(pattern) == 1 && c == '/' {
			break
		}
		if c == '{' {
			// Find end of pattern
			i := strings.IndexByte(pattern, '}')
			if i < 0 {
				return nil, "", false
			}
			pattern = pattern[i+1:]

			// Find next matching pattern character or next '/' whichever comes first
			np := -1
			if len(pattern) > 0 {
				np = strings.IndexByte(input, pattern[0])
			}
			ns := strings.IndexByte(input, '/')

			if np < 0 {
				i = ns
			} else if ns < 0 {
				i = np
			} else {
				i = int(math.Min(float64(np), float64(ns)))
			}
			if i < 0 {
				i = len(input)
			}
			params = append(params, input[:i])
			input = input[i:]
			continue
		}
		if len(input) == 0 || input[0] != c {
			return nil, "", false
		}
		pattern = pattern[1:]
		input = input[1:]
	}
	if input == "" {
		input = "/"
	}
	if input[0] != '/' {
		return nil, "", false
	}
	return params, input, true
}

func (value *Server) Validate(ctx context.Context) (err error) {
	if value.URL == "" {
		return errors.New("value of url must be a non-empty string")
	}

	opening, closing := strings.Count(value.URL, "{"), strings.Count(value.URL, "}")
	if opening != closing {
		return errors.New("server URL has mismatched { and }")
	}

	if opening != len(value.Variables) {
		return errors.New("server has undeclared variables")
	}

	for name, v := range value.Variables {
		if !strings.Contains(value.URL, fmt.Sprintf("{%s}", name)) {
			return errors.New("server has undeclared variables")
		}

		if err = v.Validate(ctx); err != nil {
			return
		}
	}

	if v := value.Bindings; v != nil {
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}

	return
}

// ServerVariable is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#serverVariableObject
type ServerVariable struct {
	Extensions map[string]interface{} `json:"-" yaml:"-"`

	Enum        []string `json:"enum,omitempty" yaml:"enum,omitempty"`
	Default     string   `json:"default,omitempty" yaml:"default,omitempty"`
	Description string   `json:"description,omitempty" yaml:"description,omitempty"`
}

func (value *ServerVariable) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{}, 3+len(value.Extensions))
	for k, v := range value.Extensions {
		m[k] = v
	}

	if len(value.Enum) != 0 {
		m["enum"] = value.Enum
	}
	if len(value.Default) != 0 {
		m["default"] = value.Default
	}
	if len(value.Description) != 0 {
		m["description"] = value.Description
	}

	return json.Marshal(m)
}

func (value *ServerVariable) UnmarshalJSON(data []byte) error {
	type ServerVariableBis ServerVariable
	var x ServerVariableBis
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	_ = json.Unmarshal(data, &x.Extensions)

	delete(x.Extensions, "enum")
	delete(x.Extensions, "default")
	delete(x.Extensions, "description")

	*value = ServerVariable(x)

	return nil
}

func (value *ServerVariable) Validate(context.Context) error {
	if value.Default == "" {
		data, err := value.MarshalJSON()
		if err != nil {
			return err
		}

		return fmt.Errorf("field default is required in %s", data)
	}

	return nil
}
