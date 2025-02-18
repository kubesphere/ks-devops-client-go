// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DevopsValidates devops validates
//
// swagger:model devops.Validates
type DevopsValidates struct {

	// the id of credential
	CredentialID string `json:"credentialId,omitempty"`
}

// Validate validates this devops validates
func (m *DevopsValidates) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this devops validates based on context it is used
func (m *DevopsValidates) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DevopsValidates) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DevopsValidates) UnmarshalBinary(b []byte) error {
	var res DevopsValidates
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
