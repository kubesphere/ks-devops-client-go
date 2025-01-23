// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1alpha3OrphanedResourceKey v1alpha3 orphaned resource key
//
// swagger:model v1alpha3.OrphanedResourceKey
type V1alpha3OrphanedResourceKey struct {

	// group
	Group string `json:"group,omitempty"`

	// kind
	Kind string `json:"kind,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this v1alpha3 orphaned resource key
func (m *V1alpha3OrphanedResourceKey) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1alpha3 orphaned resource key based on context it is used
func (m *V1alpha3OrphanedResourceKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha3OrphanedResourceKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha3OrphanedResourceKey) UnmarshalBinary(b []byte) error {
	var res V1alpha3OrphanedResourceKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
