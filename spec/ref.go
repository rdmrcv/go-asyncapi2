package spec

import (
	"context"
	"encoding/json"
)

type Ref struct {
	Ref string `json:"$ref" yaml:"$ref"`
}

type RefG[V interface {
	comparable

	json.Marshaler
	json.Unmarshaler

	Validate(ctx context.Context) error
}] struct {
	Ref   string
	Value V
}

func (value *RefG[V]) MarshalJSON() ([]byte, error) {
	if ref := value.Ref; ref != "" {
		return json.Marshal(Ref{Ref: ref})
	}

	return value.Value.MarshalJSON()
}

func (value *RefG[V]) UnmarshalJSON(data []byte) error {
	refProps := &Ref{}

	if err := json.Unmarshal(data, refProps); err == nil {
		if len(refProps.Ref) > 0 {
			value.Ref = refProps.Ref

			return nil
		}
	}

	return json.Unmarshal(data, value.Value)
}

func (value *RefG[V]) Validate(ctx context.Context) error {
	var zero V

	if v := value.Value; v != zero {
		return v.Validate(ctx)
	}

	return foundUnresolvedRef(value.Ref)
}
