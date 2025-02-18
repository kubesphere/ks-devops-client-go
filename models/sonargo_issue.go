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

// SonargoIssue sonargo issue
//
// swagger:model sonargo.Issue
type SonargoIssue struct {

	// actions
	Actions []string `json:"actions"`

	// assignee
	Assignee string `json:"assignee,omitempty"`

	// author
	Author string `json:"author,omitempty"`

	// close date
	CloseDate string `json:"closeDate,omitempty"`

	// comments
	Comments []*SonargoComment `json:"comments"`

	// component
	Component string `json:"component,omitempty"`

	// creation date
	CreationDate string `json:"creationDate,omitempty"`

	// debt
	Debt string `json:"debt,omitempty"`

	// effort
	Effort string `json:"effort,omitempty"`

	// flows
	Flows []SonargoIssueFlows `json:"flows"`

	// from hotspot
	FromHotspot bool `json:"fromHotspot,omitempty"`

	// hash
	Hash string `json:"hash,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// line
	Line int32 `json:"line,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// organization
	Organization string `json:"organization,omitempty"`

	// project
	Project string `json:"project,omitempty"`

	// resolution
	Resolution string `json:"resolution,omitempty"`

	// rule
	Rule string `json:"rule,omitempty"`

	// severity
	Severity string `json:"severity,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// text range
	TextRange *SonargoTextRange `json:"textRange,omitempty"`

	// transitions
	Transitions []string `json:"transitions"`

	// type
	Type string `json:"type,omitempty"`

	// update date
	UpdateDate string `json:"updateDate,omitempty"`
}

// Validate validates this sonargo issue
func (m *SonargoIssue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTextRange(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SonargoIssue) validateComments(formats strfmt.Registry) error {
	if swag.IsZero(m.Comments) { // not required
		return nil
	}

	for i := 0; i < len(m.Comments); i++ {
		if swag.IsZero(m.Comments[i]) { // not required
			continue
		}

		if m.Comments[i] != nil {
			if err := m.Comments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("comments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("comments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SonargoIssue) validateTextRange(formats strfmt.Registry) error {
	if swag.IsZero(m.TextRange) { // not required
		return nil
	}

	if m.TextRange != nil {
		if err := m.TextRange.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("textRange")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("textRange")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this sonargo issue based on the context it is used
func (m *SonargoIssue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTextRange(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SonargoIssue) contextValidateComments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Comments); i++ {

		if m.Comments[i] != nil {

			if swag.IsZero(m.Comments[i]) { // not required
				return nil
			}

			if err := m.Comments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("comments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("comments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SonargoIssue) contextValidateTextRange(ctx context.Context, formats strfmt.Registry) error {

	if m.TextRange != nil {

		if swag.IsZero(m.TextRange) { // not required
			return nil
		}

		if err := m.TextRange.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("textRange")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("textRange")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SonargoIssue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SonargoIssue) UnmarshalBinary(b []byte) error {
	var res SonargoIssue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
