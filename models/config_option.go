// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConfigOption config option
//
// swagger:model config.Option
type ConfigOption struct {

	// key
	// Required: true
	Key *string `json:"Key"`

	// value
	// Required: true
	Value *string `json:"Value"`
}

// Validate validates this config option
func (m *ConfigOption) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigOption) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("Key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

func (m *ConfigOption) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("Value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this config option based on context it is used
func (m *ConfigOption) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConfigOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigOption) UnmarshalBinary(b []byte) error {
	var res ConfigOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
