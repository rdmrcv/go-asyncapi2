package spec

import (
	"os"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/ghodss/yaml"

	"github.com/rdmrcv/go-asyncapi2/spec/bindings"
)

func TestT_MarshalJSON(t *testing.T) {
	var doc = &T{
		AsyncAPI: "2.0.0",

		Info: &openapi3.Info{
			Title:   "Streetlights API",
			Version: "1.0.0",
			Description: `The Smartylighting Streetlights API allows you to remotely manage the city lights.

### Check out its awesome features:

* Turn a specific streetlight on/off 🌃
* Dim a specific streetlight 😎
* Receive real-time information about environmental lighting conditions 📈`,
			License: &openapi3.License{
				Name: "Apache 2.0",
				URL:  "https://www.apache.org/licenses/LICENSE-2.0",
			},
			Contact: &openapi3.Contact{
				Name:  "Roman Domrachev",
				Email: "ligser@gmail.com",
				URL:   "https://github.com/ligser",
			},
		},

		Servers: Servers{
			"production": &Server{
				URL:         "api.streetlights.smartylighting.com:{port}",
				Protocol:    "mqtt",
				Description: "Test broker",
				Variables: map[string]*ServerVariable{
					"port": {
						Description: "Secure connection (TLS) is available through port 8883.",
						Default:     "1883",
						Enum:        []string{"1883", "8883"},
					},
				},
				Security: []SecurityRequirements{
					{"apiKey": []string{}},
					{"supportedOauthFlows": []string{
						"streetlights:on",
						"streetlights:off",
						"streetlights:dim",
					}},
					{"openIdConnectWellKnown": []string{}},
				},
			},
		},

		DefaultContentType: "application/json",

		Channels: Channels{
			"smartylighting/streetlights/1/0/event/{streetlightId}/lighting/measured": &Channel{
				Description: "The topic on which measured values may be produced and consumed.",
				Parameters: ParametersRefs{
					"streetlightId": &ParameterRef{Ref: "#/components/parameters/streetlightId"},
				},
				Subscribe: &OperationRef{
					Value: &Operation{
						OperationTrait: OperationTrait{
							Description: "Receive information about environmental lighting conditions of a particular streetlight.",
							OperationID: "receiveLightMeasurement",
						},
						Traits: []*OperationTraitRef{
							{Ref: "#/components/operationTraits/kafka"},
						},
						Message: &MessageOneOf{MessageRef: MessageRef{Ref: "#/components/messages/lightMeasured"}},
					},
				},
			},

			"smartylighting/streetlights/1/0/action/{streetlightId}/turn/on": {
				Description: "The topic on which turn on values may be produced and consumed.",
				Parameters: map[string]*ParameterRef{
					"streetlightId": {Ref: "#/components/parameters/streetlightId"},
				},
				Publish: &OperationRef{
					Value: &Operation{
						OperationTrait: OperationTrait{
							Description: "Turn environmental lighting condition to on.",
							OperationID: "turnOn",
						},
						Traits: []*OperationTraitRef{
							{Ref: "#/components/operationTraits/kafka"},
						},
						Message: &MessageOneOf{MessageRef: MessageRef{Ref: "#/components/messages/turnOnOff"}},
					},
				},
			},

			"smartylighting/streetlights/1/0/action/{streetlightId}/turn/off": {
				Description: "The topic on which turn off values may be produced and consumed.",
				Parameters: map[string]*ParameterRef{
					"streetlightId": {Ref: "#/components/parameters/streetlightId"},
				},
				Publish: &OperationRef{
					Value: &Operation{
						OperationTrait: OperationTrait{
							Description: "Turn environmental lighting condition to off.",
							OperationID: "turnOff",
						},
						Traits: []*OperationTraitRef{
							{Ref: "#/components/operationTraits/kafka"},
						},
						Message: &MessageOneOf{MessageRef: MessageRef{Ref: "#/components/messages/turnOnOff"}},
					},
				},
			},

			"smartylighting/streetlights/1/0/action/{streetlightId}/dim": {
				Description: "The topic on which dim values may be produced and consumed.",
				Parameters: map[string]*ParameterRef{
					"streetlightId": {Ref: "#/components/parameters/streetlightId"},
				},
				Publish: &OperationRef{
					Value: &Operation{
						OperationTrait: OperationTrait{
							Description: "Dim environmental lighting condition to passed level.",
							OperationID: "dimLight",
						},
						Traits: []*OperationTraitRef{
							{Ref: "#/components/operationTraits/kafka"},
						},
						Message: &MessageOneOf{MessageRef: MessageRef{Ref: "#/components/messages/dimLight"}},
					},
				},
			},
		},

		Components: &Components{
			Messages: Messages{
				"lightMeasured": &Message{
					MessageTrait: MessageTrait{
						Name:        "lightMeasured",
						Title:       "Light measured",
						Summary:     "Inform about environmental lighting conditions for a particular streetlight.",
						ContentType: "application/json",
					},
					Traits: []*MessageTraitRef{
						{Ref: "#/components/messageTraits/commonHeaders"},
					},
					Payload: &openapi3.SchemaRef{
						Ref: "#/components/schemas/lightMeasuredPayload",
					},
				},
				"turnOnOff": &Message{
					MessageTrait: MessageTrait{
						Name:    "turnOnOff",
						Title:   "Turn on/off",
						Summary: "Command a particular streetlight to turn the lights on or off.",
					},
					Traits: []*MessageTraitRef{
						{Ref: "#/components/messageTraits/commonHeaders"},
					},
					Payload: &openapi3.SchemaRef{
						Ref: "#/components/schemas/turnOnOffPayload",
					},
				},
				"dimLight": &Message{
					MessageTrait: MessageTrait{
						Name:    "dimLight",
						Title:   "Dim light",
						Summary: "Command a particular streetlight to dim the lights.",
					},
					Traits: []*MessageTraitRef{
						{Ref: "#/components/messageTraits/commonHeaders"},
					},
					Payload: &openapi3.SchemaRef{
						Ref: "#/components/schemas/dimLightPayload",
					},
				},
			},

			Schemas: map[string]*openapi3.SchemaRef{
				"lightMeasuredPayload": {
					Value: &openapi3.Schema{
						Type: "object",
						Properties: openapi3.Schemas{
							"command": &openapi3.SchemaRef{Value: &openapi3.Schema{
								Type:        "integer",
								Min:         openapi3.Float64Ptr(0),
								Description: "Light intensity measured in lumens.",
							}},
							"sentAt": &openapi3.SchemaRef{Ref: "#/components/schemas/sentAt"},
						},
					},
				},
				"turnOnOffPayload": {
					Value: &openapi3.Schema{
						Type: "object",
						Properties: openapi3.Schemas{
							"command": &openapi3.SchemaRef{Value: &openapi3.Schema{
								Type:        "string",
								Enum:        []interface{}{"on", "off"},
								Description: "Whether to turn on or off the light.",
							}},
							"sentAt": &openapi3.SchemaRef{Ref: "#/components/schemas/sentAt"},
						},
					},
				},
				"dimLightPayload": {
					Value: &openapi3.Schema{
						Type: "object",
						Properties: openapi3.Schemas{
							"percentage": &openapi3.SchemaRef{Value: &openapi3.Schema{
								Type:        "integer",
								Description: "Percentage to which the light should be dimmed to.",
								Min:         openapi3.Float64Ptr(0),
								Max:         openapi3.Float64Ptr(100),
							}},
							"sentAt": &openapi3.SchemaRef{Ref: "#/components/schemas/sentAt"},
						},
					},
				},
				"sentAt": {
					Value: &openapi3.Schema{
						Type:        "string",
						Format:      "date-time",
						Description: "Date and time when the message was sent.",
					},
				},
			},
			SecuritySchemes: map[string]*SecurityScheme{
				"apiKey": {
					Type:        "apiKey",
					In:          "user",
					Description: "Provide your API key as the user and leave the password empty.",
				},
				"supportedOauthFlows": {
					Type:        "oauth2",
					Description: "Flows to support OAuth 2.0",
					Flows: &OAuthFlows{
						Implicit: &OAuthFlowObject{
							AuthorizationUrl: "https://authserver.example/auth",
							Scopes: map[string]string{
								"streetlights:on":  "Ability to switch lights on",
								"streetlights:off": "Ability to switch lights off",
								"streetlights:dim": "Ability to dim the lights",
							},
						},
						Password: &OAuthFlowObject{
							TokenUrl: "https://authserver.example/token",
							Scopes: map[string]string{
								"streetlights:on":  "Ability to switch lights on",
								"streetlights:off": "Ability to switch lights off",
								"streetlights:dim": "Ability to dim the lights",
							},
						},
						ClientCredentials: &OAuthFlowObject{
							TokenUrl: "https://authserver.example/token",
							Scopes: map[string]string{
								"streetlights:on":  "Ability to switch lights on",
								"streetlights:off": "Ability to switch lights off",
								"streetlights:dim": "Ability to dim the lights",
							},
						},
						AuthorizationCode: &OAuthFlowObject{
							AuthorizationUrl: "https://authserver.example/auth",
							TokenUrl:         "https://authserver.example/token",
							RefreshUrl:       "https://authserver.example/refresh",
							Scopes: map[string]string{
								"streetlights:on":  "Ability to switch lights on",
								"streetlights:off": "Ability to switch lights off",
								"streetlights:dim": "Ability to dim the lights",
							},
						},
					},
				},
				"openIdConnectWellKnown": {
					Type:             "openIdConnect",
					OpenIDConnectUrl: "https://authserver.example/.well-known",
				},
			},

			Parameters: map[string]*Parameter{
				"streetlightId": {
					Description: "The ID of the streetlight.",
					Schema: &openapi3.Schema{
						Type: "string",
					},
				},
			},

			MessageTraits: map[string]*MessageTrait{
				"commonHeaders": {
					Headers: &openapi3.SchemaRef{Value: &openapi3.Schema{
						Type: "object",
						Properties: map[string]*openapi3.SchemaRef{
							"my-app-header": {Value: &openapi3.Schema{
								Type: "integer",
								Min:  openapi3.Float64Ptr(0),
								Max:  openapi3.Float64Ptr(100),
							}},
						},
					}},
				},
			},

			OperationTraits: map[string]*OperationTrait{
				"kafka": {
					Bindings: &OperationBindings{
						Kafka: &bindings.KafkaOperation{
							ClientID: &openapi3.Schema{Type: "string"},
						},
					},
				},
			},
		},

		Tags: openapi3.Tags{
			{Name: "demo", Description: "Tag placed for demo purpose."},
		},
	}

	bts, err := doc.MarshalJSON()
	if err != nil {
		panic(err)
	}

	bts, err = yaml.JSONToYAML(bts)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("./test/streetlights.yml")
	if err != nil {
		panic(err)
	}

	if _, err := file.Write(bts); err != nil {
		panic(err)
	}
}
