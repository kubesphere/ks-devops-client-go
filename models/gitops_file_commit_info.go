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

// GitopsFileCommitInfo gitops file commit info
//
// swagger:model gitops.FileCommitInfo
type GitopsFileCommitInfo struct {

	// commit
	// Required: true
	Commit *GitopsCommit `json:"commit"`

	// data
	// Required: true
	Data *string `json:"data"`

	// is binary
	// Required: true
	IsBinary *bool `json:"isBinary"`

	// is dir
	// Required: true
	IsDir *bool `json:"isDir"`

	// is symlink
	// Required: true
	IsSymlink *bool `json:"isSymlink"`

	// name
	// Required: true
	Name *string `json:"name"`

	// size
	// Required: true
	Size *int64 `json:"size"`
}

// Validate validates this gitops file commit info
func (m *GitopsFileCommitInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsBinary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsDir(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsSymlink(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitopsFileCommitInfo) validateCommit(formats strfmt.Registry) error {

	if err := validate.Required("commit", "body", m.Commit); err != nil {
		return err
	}

	if m.Commit != nil {
		if err := m.Commit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("commit")
			}
			return err
		}
	}

	return nil
}

func (m *GitopsFileCommitInfo) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	return nil
}

func (m *GitopsFileCommitInfo) validateIsBinary(formats strfmt.Registry) error {

	if err := validate.Required("isBinary", "body", m.IsBinary); err != nil {
		return err
	}

	return nil
}

func (m *GitopsFileCommitInfo) validateIsDir(formats strfmt.Registry) error {

	if err := validate.Required("isDir", "body", m.IsDir); err != nil {
		return err
	}

	return nil
}

func (m *GitopsFileCommitInfo) validateIsSymlink(formats strfmt.Registry) error {

	if err := validate.Required("isSymlink", "body", m.IsSymlink); err != nil {
		return err
	}

	return nil
}

func (m *GitopsFileCommitInfo) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *GitopsFileCommitInfo) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this gitops file commit info based on the context it is used
func (m *GitopsFileCommitInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCommit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitopsFileCommitInfo) contextValidateCommit(ctx context.Context, formats strfmt.Registry) error {

	if m.Commit != nil {

		if err := m.Commit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("commit")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GitopsFileCommitInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitopsFileCommitInfo) UnmarshalBinary(b []byte) error {
	var res GitopsFileCommitInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
