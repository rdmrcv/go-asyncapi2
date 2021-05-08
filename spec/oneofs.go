package spec

import (
	"context"
	"encoding/json"

	"github.com/getkin/kin-openapi/jsoninfo"
	"github.com/go-openapi/jsonpointer"
)

type oneOfField struct {
	OneOf []json.RawMessage `json:"oneOf,omitempty" yaml:"oneOf,omitempty"`
}

type MessageOneOf struct {
	MessageRef
	OneOf []*MessageRef
}

var _ jsonpointer.JSONPointable = (*MessageOneOf)(nil)

func (value *MessageOneOf) MarshalJSON() ([]byte, error) {
	if len(value.OneOf) == 0 {
		return jsoninfo.MarshalRef(value.Ref, value.Value)
	}

	ents := make([]json.RawMessage, 0, len(value.OneOf))
	for _, ent := range value.OneOf {
		if ent == nil {
			continue
		}

		entJson, err := jsoninfo.MarshalRef(ent.Ref, ent.Value)
		if err != nil {
			return nil, err
		}

		ents = append(ents, entJson)
	}

	return json.Marshal(&oneOfField{OneOf: ents})
}

func (value *MessageOneOf) UnmarshalJSON(data []byte) error {
	oneof := oneOfField{}
	if err := json.Unmarshal(data, &oneof); err != nil {
		return err
	}

	if len(oneof.OneOf) == 0 {
		return jsoninfo.UnmarshalRef(data, &value.Ref, &value.Value)
	}

	ents := make([]*MessageRef, 0, len(oneof.OneOf))
	for _, ent := range oneof.OneOf {
		entModel := MessageRef{
			Value: &Message{},
		}

		if err := jsoninfo.UnmarshalRef(ent, &entModel.Ref, entModel.Value); err != nil {
			return err
		}

		if len(entModel.Ref) != 0 {
			entModel.Value = nil
		}

		ents = append(ents, &entModel)
	}

	return nil
}

func (value *MessageOneOf) Validate(ctx context.Context) error {
	if v := value.Value; v != nil {
		return v.Validate(ctx)
	}

	if v := value.OneOf; len(v) > 0 {
		for _, ent := range v {
			if err := ent.Validate(ctx); err != nil {
				return err
			}
		}
	}

	return foundUnresolvedRef(value.Ref)
}

func (value MessageOneOf) JSONLookup(token string) (interface{}, error) {
	if token == "$ref" {
		return value.Ref, nil
	}

	if token == "oneOf" {
		return value.OneOf, nil
	}

	ptr, _, err := jsonpointer.GetForToken(value.Value, token)
	return ptr, err
}
