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

// Links links
//
// swagger:model ._links
type Links struct {

	// actions
	Actions *Actions `json:"actions,omitempty"`

	// branches
	Branches *Branches `json:"branches,omitempty"`

	// queue
	Queue *Queue `json:"queue,omitempty"`

	// runs
	Runs *Runs `json:"runs,omitempty"`

	// scm
	Scm *Scm `json:"scm,omitempty"`

	// self
	Self *Self `json:"self,omitempty"`

	// trends
	Trends *Trends `json:"trends,omitempty"`
}

// Validate validates this links
func (m *Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBranches(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrends(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Links) validateActions(formats strfmt.Registry) error {
	if swag.IsZero(m.Actions) { // not required
		return nil
	}

	if m.Actions != nil {
		if err := m.Actions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("actions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("actions")
			}
			return err
		}
	}

	return nil
}

func (m *Links) validateBranches(formats strfmt.Registry) error {
	if swag.IsZero(m.Branches) { // not required
		return nil
	}

	if m.Branches != nil {
		if err := m.Branches.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("branches")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("branches")
			}
			return err
		}
	}

	return nil
}

func (m *Links) validateQueue(formats strfmt.Registry) error {
	if swag.IsZero(m.Queue) { // not required
		return nil
	}

	if m.Queue != nil {
		if err := m.Queue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("queue")
			}
			return err
		}
	}

	return nil
}

func (m *Links) validateRuns(formats strfmt.Registry) error {
	if swag.IsZero(m.Runs) { // not required
		return nil
	}

	if m.Runs != nil {
		if err := m.Runs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("runs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("runs")
			}
			return err
		}
	}

	return nil
}

func (m *Links) validateScm(formats strfmt.Registry) error {
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

func (m *Links) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("self")
			}
			return err
		}
	}

	return nil
}

func (m *Links) validateTrends(formats strfmt.Registry) error {
	if swag.IsZero(m.Trends) { // not required
		return nil
	}

	if m.Trends != nil {
		if err := m.Trends.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trends")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trends")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this links based on the context it is used
func (m *Links) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBranches(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQueue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRuns(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTrends(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Links) contextValidateActions(ctx context.Context, formats strfmt.Registry) error {

	if m.Actions != nil {

		if swag.IsZero(m.Actions) { // not required
			return nil
		}

		if err := m.Actions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("actions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("actions")
			}
			return err
		}
	}

	return nil
}

func (m *Links) contextValidateBranches(ctx context.Context, formats strfmt.Registry) error {

	if m.Branches != nil {

		if swag.IsZero(m.Branches) { // not required
			return nil
		}

		if err := m.Branches.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("branches")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("branches")
			}
			return err
		}
	}

	return nil
}

func (m *Links) contextValidateQueue(ctx context.Context, formats strfmt.Registry) error {

	if m.Queue != nil {

		if swag.IsZero(m.Queue) { // not required
			return nil
		}

		if err := m.Queue.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("queue")
			}
			return err
		}
	}

	return nil
}

func (m *Links) contextValidateRuns(ctx context.Context, formats strfmt.Registry) error {

	if m.Runs != nil {

		if swag.IsZero(m.Runs) { // not required
			return nil
		}

		if err := m.Runs.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("runs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("runs")
			}
			return err
		}
	}

	return nil
}

func (m *Links) contextValidateScm(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Links) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("self")
			}
			return err
		}
	}

	return nil
}

func (m *Links) contextValidateTrends(ctx context.Context, formats strfmt.Registry) error {

	if m.Trends != nil {

		if swag.IsZero(m.Trends) { // not required
			return nil
		}

		if err := m.Trends.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trends")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trends")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Links) UnmarshalBinary(b []byte) error {
	var res Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
