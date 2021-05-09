package spec

import (
	"context"
	"errors"
	"fmt"
	"math"
	"net/url"
	"regexp"
	"strings"

	"github.com/getkin/kin-openapi/jsoninfo"
	"github.com/getkin/kin-openapi/openapi3"
)

var (
	serverKeyRegexp = regexp.MustCompile("^[A-Za-z0-9_\\-]+$")

	ErrServerKeyInvalid = fmt.Errorf("server should match pattern %q", serverKeyRegexp.String())
)

// Servers is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#serversObject
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
	openapi3.ExtensionProps
	URL             string                     `json:"url" yaml:"url"`
	Protocol        string                     `json:"protocol" yaml:"protocol"`
	ProtocolVersion string                     `json:"protocolVersion,omitempty" yaml:"protocolVersion,omitempty"`
	Description     string                     `json:"description,omitempty" yaml:"description,omitempty"`
	Security        []SecurityRequirements     `json:"security,omitempty" yaml:"security,omitempty"`
	Bindings        *ServerBindings            `json:"bindings,omitempty" yaml:"bindings,omitempty"`
	Variables       map[string]*ServerVariable `json:"variables,omitempty" yaml:"variables,omitempty"`
}

func (value *Server) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStrictStruct(value)
}

func (value *Server) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStrictStruct(data, value)
}

func (value Server) ParameterNames() ([]string, error) {
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

func (value Server) MatchRawURL(input string) ([]string, string, bool) {
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
	openapi3.ExtensionProps
	Enum        []string `json:"enum,omitempty" yaml:"enum,omitempty"`
	Default     string   `json:"default,omitempty" yaml:"default,omitempty"`
	Description string   `json:"description,omitempty" yaml:"description,omitempty"`
}

func (value *ServerVariable) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStrictStruct(value)
}

func (value *ServerVariable) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStrictStruct(data, value)
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
