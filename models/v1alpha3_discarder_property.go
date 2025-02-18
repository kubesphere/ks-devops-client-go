// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1alpha3DiscarderProperty v1alpha3 discarder property
//
// swagger:model v1alpha3.DiscarderProperty
type V1alpha3DiscarderProperty struct {

	// days to keep pipeline
	DaysToKeep string `json:"days_to_keep,omitempty"`

	// nums to keep pipeline
	NumToKeep string `json:"num_to_keep,omitempty"`
}

// Validate validates this v1alpha3 discarder property
func (m *V1alpha3DiscarderProperty) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1alpha3 discarder property based on context it is used
func (m *V1alpha3DiscarderProperty) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha3DiscarderProperty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha3DiscarderProperty) UnmarshalBinary(b []byte) error {
	var res V1alpha3DiscarderProperty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
