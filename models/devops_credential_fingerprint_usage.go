// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DevopsCredentialFingerprintUsage devops credential fingerprint usage
//
// swagger:model devops.Credential.fingerprint.usage
type DevopsCredentialFingerprintUsage struct {

	// pipeline full name
	Name string `json:"name,omitempty"`

	// The build number of all pipelines that use this credential
	Ranges *Ranges `json:"ranges,omitempty"`
}

// Validate validates this devops credential fingerprint usage
func (m *DevopsCredentialFingerprintUsage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRanges(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DevopsCredentialFingerprintUsage) validateRanges(formats strfmt.Registry) error {
	if swag.IsZero(m.Ranges) { // not required
		return nil
	}

	if m.Ranges != nil {
		if err := m.Ranges.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ranges")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ranges")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this devops credential fingerprint usage based on the context it is used
func (m *DevopsCredentialFingerprintUsage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRanges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DevopsCredentialFingerprintUsage) contextValidateRanges(ctx context.Context, formats strfmt.Registry) error {

	if m.Ranges != nil {

		if swag.IsZero(m.Ranges) { // not required
			return nil
		}

		if err := m.Ranges.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ranges")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ranges")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DevopsCredentialFingerprintUsage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DevopsCredentialFingerprintUsage) UnmarshalBinary(b []byte) error {
	var res DevopsCredentialFingerprintUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
