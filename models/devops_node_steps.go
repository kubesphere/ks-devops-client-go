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
	"github.com/go-openapi/validate"
)

// DevopsNodeSteps devops node steps
//
// swagger:model devops.NodeSteps
type DevopsNodeSteps struct {

	// It’s a fully qualified name and is an identifier of the producer of this resource's capability.
	Class string `json:"_class,omitempty"`

	// references the reachable path to this resource
	Links *Links `json:"_links,omitempty"`

	// actions
	Actions []*DevopsNodeStepsActions `json:"actions"`

	// indicate if this step can be approved by current user
	// Required: true
	Aprovable *bool `json:"aprovable"`

	// display description
	DisplayDescription string `json:"displayDescription,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// duration time in mullis
	DurationInMillis int32 `json:"durationInMillis,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// the action should user input
	Input *DevopsInput `json:"input,omitempty"`

	// the result of pipeline run. e.g. SUCCESS
	Result string `json:"result,omitempty"`

	// the time of starts
	StartTime string `json:"startTime,omitempty"`

	// run state. e.g. SKIPPED
	State string `json:"state,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this devops node steps
func (m *DevopsNodeSteps) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAprovable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DevopsNodeSteps) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *DevopsNodeSteps) validateActions(formats strfmt.Registry) error {
	if swag.IsZero(m.Actions) { // not required
		return nil
	}

	for i := 0; i < len(m.Actions); i++ {
		if swag.IsZero(m.Actions[i]) { // not required
			continue
		}

		if m.Actions[i] != nil {
			if err := m.Actions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("actions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DevopsNodeSteps) validateAprovable(formats strfmt.Registry) error {

	if err := validate.Required("aprovable", "body", m.Aprovable); err != nil {
		return err
	}

	return nil
}

func (m *DevopsNodeSteps) validateInput(formats strfmt.Registry) error {
	if swag.IsZero(m.Input) { // not required
		return nil
	}

	if m.Input != nil {
		if err := m.Input.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("input")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this devops node steps based on the context it is used
func (m *DevopsNodeSteps) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateActions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInput(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DevopsNodeSteps) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *DevopsNodeSteps) contextValidateActions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Actions); i++ {

		if m.Actions[i] != nil {

			if swag.IsZero(m.Actions[i]) { // not required
				return nil
			}

			if err := m.Actions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("actions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DevopsNodeSteps) contextValidateInput(ctx context.Context, formats strfmt.Registry) error {

	if m.Input != nil {

		if swag.IsZero(m.Input) { // not required
			return nil
		}

		if err := m.Input.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("input")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DevopsNodeSteps) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DevopsNodeSteps) UnmarshalBinary(b []byte) error {
	var res DevopsNodeSteps
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
