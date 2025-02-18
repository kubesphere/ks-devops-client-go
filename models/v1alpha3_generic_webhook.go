// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1alpha3GenericWebhook v1alpha3 generic webhook
//
// swagger:model v1alpha3.GenericWebhook
type V1alpha3GenericWebhook struct {

	// Indicate the reason why a webhook triggered
	Cause string `json:"cause,omitempty"`

	// Indicate if the generic webhook is enabled
	Enable bool `json:"enable,omitempty"`

	// Filter expression which against the filter name
	FilterExpression string `json:"filter_expression,omitempty"`

	// Filter name for the generic webhook, it could be a variable name
	FilterText string `json:"filter_text,omitempty"`

	// Define variables which come from the HTTP request header
	HeaderVariables []*V1alpha3GenericVariable `json:"header_variables"`

	// Indicate if print the post content
	PrintPostContent bool `json:"print_post_content,omitempty"`

	// Indicate if print the variables
	PrintVariables bool `json:"print_variables,omitempty"`

	// Define variables which come from the HTTP request
	RequestVariables []*V1alpha3GenericVariable `json:"request_variables"`

	// The token of generic webhook
	Token string `json:"token,omitempty"`
}

// Validate validates this v1alpha3 generic webhook
func (m *V1alpha3GenericWebhook) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHeaderVariables(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestVariables(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha3GenericWebhook) validateHeaderVariables(formats strfmt.Registry) error {
	if swag.IsZero(m.HeaderVariables) { // not required
		return nil
	}

	for i := 0; i < len(m.HeaderVariables); i++ {
		if swag.IsZero(m.HeaderVariables[i]) { // not required
			continue
		}

		if m.HeaderVariables[i] != nil {
			if err := m.HeaderVariables[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("header_variables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("header_variables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1alpha3GenericWebhook) validateRequestVariables(formats strfmt.Registry) error {
	if swag.IsZero(m.RequestVariables) { // not required
		return nil
	}

	for i := 0; i < len(m.RequestVariables); i++ {
		if swag.IsZero(m.RequestVariables[i]) { // not required
			continue
		}

		if m.RequestVariables[i] != nil {
			if err := m.RequestVariables[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("request_variables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("request_variables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1alpha3 generic webhook based on the context it is used
func (m *V1alpha3GenericWebhook) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHeaderVariables(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequestVariables(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha3GenericWebhook) contextValidateHeaderVariables(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.HeaderVariables); i++ {

		if m.HeaderVariables[i] != nil {

			if swag.IsZero(m.HeaderVariables[i]) { // not required
				return nil
			}

			if err := m.HeaderVariables[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("header_variables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("header_variables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1alpha3GenericWebhook) contextValidateRequestVariables(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RequestVariables); i++ {

		if m.RequestVariables[i] != nil {

			if swag.IsZero(m.RequestVariables[i]) { // not required
				return nil
			}

			if err := m.RequestVariables[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("request_variables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("request_variables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha3GenericWebhook) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha3GenericWebhook) UnmarshalBinary(b []byte) error {
	var res V1alpha3GenericWebhook
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
