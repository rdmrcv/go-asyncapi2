package spec

import (
	"context"
	"encoding/json"
)

type oneOfField struct {
	OneOf []json.RawMessage `json:"oneOf,omitempty" yaml:"oneOf,omitempty"`
}

type MessageOneOf struct {
	MessageRef

	OneOf []*MessageRef
}

func (value *MessageOneOf) MarshalJSON() ([]byte, error) {
	if len(value.OneOf) == 0 {
		if len(value.Ref) != 0 {
			return json.Marshal(Ref{Ref: value.Ref})
		}

		return json.Marshal(value.Value)
	}

	ents := make([]json.RawMessage, 0, len(value.OneOf))
	for _, ent := range value.OneOf {
		if ent == nil {
			continue
		}

		var (
			entJson []byte
			err     error
		)

		if len(ent.Ref) != 0 {
			entJson, err = json.Marshal(Ref{Ref: ent.Ref})
		} else {
			entJson, err = json.Marshal(ent.Value)
		}

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
		var ref Ref
		if err := json.Unmarshal(data, &ref); err != nil {
			return err
		}

		if len(ref.Ref) != 0 {
			value.Ref = ref.Ref

			return nil
		}

		return json.Unmarshal(data, value.Value)
	}

	ents := make([]*MessageRef, 0, len(oneof.OneOf))
	for _, ent := range oneof.OneOf {
		entModel := MessageRef{
			Value: &Message{},
		}

		var ref Ref
		if err := json.Unmarshal(ent, &ref); err != nil {
			return err
		}

		if len(ref.Ref) != 0 {
			entModel.Ref = ref.Ref

			continue
		}

		if err := json.Unmarshal(ent, value.Value); err != nil {
			return err
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
