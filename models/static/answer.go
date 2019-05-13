package static

import (
	"encoding/json"

	"github.com/go-openapi/strfmt"
)

// JSONNumber is an alias of json.Number. Required for addition of Validate() method
// which is required by go-swagger generated code
type JSONNumber json.Number

// Validate fulfils requirement by generated code
func (m *JSONNumber) Validate(formats strfmt.Registry) error {
	return nil
}

// UnmarshalJSON attempts to unmarshal as json.Number and converts to JSONNumber
func (m *JSONNumber) UnmarshalJSON(b []byte) error {
	var jn json.Number

	if err := json.Unmarshal(b, &jn); err != nil {
		return err
	}
	*m = JSONNumber(jn)
	return nil
}
