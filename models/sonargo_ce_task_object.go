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

// SonargoCeTaskObject sonargo ce task object
//
// swagger:model sonargo.CeTaskObject
type SonargoCeTaskObject struct {

	// task
	Task *SonargoTask `json:"task,omitempty"`
}

// Validate validates this sonargo ce task object
func (m *SonargoCeTaskObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTask(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SonargoCeTaskObject) validateTask(formats strfmt.Registry) error {
	if swag.IsZero(m.Task) { // not required
		return nil
	}

	if m.Task != nil {
		if err := m.Task.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("task")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("task")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this sonargo ce task object based on the context it is used
func (m *SonargoCeTaskObject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTask(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SonargoCeTaskObject) contextValidateTask(ctx context.Context, formats strfmt.Registry) error {

	if m.Task != nil {

		if swag.IsZero(m.Task) { // not required
			return nil
		}

		if err := m.Task.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("task")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("task")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SonargoCeTaskObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SonargoCeTaskObject) UnmarshalBinary(b []byte) error {
	var res SonargoCeTaskObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
