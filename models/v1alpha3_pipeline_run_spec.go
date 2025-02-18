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

// V1alpha3PipelineRunSpec v1alpha3 pipeline run spec
//
// swagger:model v1alpha3.PipelineRunSpec
type V1alpha3PipelineRunSpec struct {

	// action
	Action string `json:"action,omitempty"`

	// parameters
	Parameters []*V1alpha3Parameter `json:"parameters"`

	// pipeline ref
	// Required: true
	PipelineRef *V1ObjectReference `json:"pipelineRef"`

	// pipeline spec
	PipelineSpec *V1alpha3PipelineSpec `json:"pipelineSpec,omitempty"`

	// scm
	Scm *V1alpha3SCM `json:"scm,omitempty"`
}

// Validate validates this v1alpha3 pipeline run spec
func (m *V1alpha3PipelineRunSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePipelineRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePipelineSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha3PipelineRunSpec) validateParameters(formats strfmt.Registry) error {
	if swag.IsZero(m.Parameters) { // not required
		return nil
	}

	for i := 0; i < len(m.Parameters); i++ {
		if swag.IsZero(m.Parameters[i]) { // not required
			continue
		}

		if m.Parameters[i] != nil {
			if err := m.Parameters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1alpha3PipelineRunSpec) validatePipelineRef(formats strfmt.Registry) error {

	if err := validate.Required("pipelineRef", "body", m.PipelineRef); err != nil {
		return err
	}

	if m.PipelineRef != nil {
		if err := m.PipelineRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pipelineRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pipelineRef")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha3PipelineRunSpec) validatePipelineSpec(formats strfmt.Registry) error {
	if swag.IsZero(m.PipelineSpec) { // not required
		return nil
	}

	if m.PipelineSpec != nil {
		if err := m.PipelineSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pipelineSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pipelineSpec")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha3PipelineRunSpec) validateScm(formats strfmt.Registry) error {
	if swag.IsZero(m.Scm) { // not required
		return nil
	}

	if m.Scm != nil {
		if err := m.Scm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1alpha3 pipeline run spec based on the context it is used
func (m *V1alpha3PipelineRunSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateParameters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePipelineRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePipelineSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha3PipelineRunSpec) contextValidateParameters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Parameters); i++ {

		if m.Parameters[i] != nil {

			if swag.IsZero(m.Parameters[i]) { // not required
				return nil
			}

			if err := m.Parameters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1alpha3PipelineRunSpec) contextValidatePipelineRef(ctx context.Context, formats strfmt.Registry) error {

	if m.PipelineRef != nil {

		if err := m.PipelineRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pipelineRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pipelineRef")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha3PipelineRunSpec) contextValidatePipelineSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.PipelineSpec != nil {

		if swag.IsZero(m.PipelineSpec) { // not required
			return nil
		}

		if err := m.PipelineSpec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pipelineSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pipelineSpec")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha3PipelineRunSpec) contextValidateScm(ctx context.Context, formats strfmt.Registry) error {

	if m.Scm != nil {

		if swag.IsZero(m.Scm) { // not required
			return nil
		}

		if err := m.Scm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha3PipelineRunSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha3PipelineRunSpec) UnmarshalBinary(b []byte) error {
	var res V1alpha3PipelineRunSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
