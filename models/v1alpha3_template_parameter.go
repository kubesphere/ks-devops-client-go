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

// V1alpha3TemplateParameter v1alpha3 template parameter
//
// swagger:model v1alpha3.TemplateParameter
type V1alpha3TemplateParameter struct {

	// default
	Default V1JSON `json:"default,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// required
	Required bool `json:"required,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// validation
	Validation *V1alpha3ParameterValidation `json:"validation,omitempty"`
}

// Validate validates this v1alpha3 template parameter
func (m *V1alpha3TemplateParameter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha3TemplateParameter) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1alpha3TemplateParameter) validateValidation(formats strfmt.Registry) error {
	if swag.IsZero(m.Validation) { // not required
		return nil
	}

	if m.Validation != nil {
		if err := m.Validation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("validation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("validation")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1alpha3 template parameter based on the context it is used
func (m *V1alpha3TemplateParameter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateValidation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha3TemplateParameter) contextValidateValidation(ctx context.Context, formats strfmt.Registry) error {

	if m.Validation != nil {

		if swag.IsZero(m.Validation) { // not required
			return nil
		}

		if err := m.Validation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("validation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("validation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha3TemplateParameter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha3TemplateParameter) UnmarshalBinary(b []byte) error {
	var res V1alpha3TemplateParameter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
