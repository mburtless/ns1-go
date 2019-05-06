// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Secondaries secondaries
// swagger:model Secondaries
type Secondaries struct {

	// ip
	IP string `json:"ip,omitempty"`

	// networks
	Networks []int64 `json:"networks"`

	// notify
	Notify bool `json:"notify,omitempty"`

	// port
	Port int64 `json:"port,omitempty"`
}

// Validate validates this secondaries
func (m *Secondaries) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Secondaries) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secondaries) UnmarshalBinary(b []byte) error {
	var res Secondaries
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
