// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SonargoHistory sonargo history
//
// swagger:model sonargo.History
type SonargoHistory struct {

	// date
	Date string `json:"date,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this sonargo history
func (m *SonargoHistory) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sonargo history based on context it is used
func (m *SonargoHistory) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SonargoHistory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SonargoHistory) UnmarshalBinary(b []byte) error {
	var res SonargoHistory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
