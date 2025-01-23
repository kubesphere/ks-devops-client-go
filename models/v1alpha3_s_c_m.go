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

// V1alpha3SCM v1alpha3 s c m
//
// swagger:model v1alpha3.SCM
type V1alpha3SCM struct {

	// ref name
	// Required: true
	RefName *string `json:"refName"`

	// ref type
	// Required: true
	RefType *string `json:"refType"`
}

// Validate validates this v1alpha3 s c m
func (m *V1alpha3SCM) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRefName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRefType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha3SCM) validateRefName(formats strfmt.Registry) error {

	if err := validate.Required("refName", "body", m.RefName); err != nil {
		return err
	}

	return nil
}

func (m *V1alpha3SCM) validateRefType(formats strfmt.Registry) error {

	if err := validate.Required("refType", "body", m.RefType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1alpha3 s c m based on context it is used
func (m *V1alpha3SCM) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha3SCM) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha3SCM) UnmarshalBinary(b []byte) error {
	var res V1alpha3SCM
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
