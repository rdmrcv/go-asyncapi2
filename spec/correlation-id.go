package spec

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/rdmrcv/go-asyncapi2/spec/validate"
)

type CorrelationIDs map[string]*CorrelationID

// CorrelationID is defined in AsyncAPI spec: https://github.com/asyncapi/spec/blob/2.0.0/versions/2.0.0/asyncapi.md#correlationIdObject
type CorrelationID struct {
	Extensions map[string]interface{} `json:"-" yaml:"-"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Location    string `json:"location" yaml:"location"`
}

func (value *CorrelationID) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{}, 2+len(value.Extensions))
	for k, v := range value.Extensions {
		m[k] = v
	}

	if len(value.Description) != 0 {
		m["description"] = value.Description
	}

	m["location"] = value.Location

	return json.Marshal(m)
}

func (value *CorrelationID) UnmarshalJSON(data []byte) error {
	type CorrelationIDBis CorrelationID
	var x CorrelationIDBis
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	_ = json.Unmarshal(data, &x.Extensions)

	delete(x.Extensions, "description")
	delete(x.Extensions, "location")

	*value = CorrelationID(x)

	return nil
}

func (value *CorrelationID) Validate(context.Context) error {
	if value.Location == "" {
		return fmt.Errorf("location field is required: %w", validate.ErrWrongField)
	}

	return nil
}
