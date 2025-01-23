// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SonargoRuleParam sonargo rule param
//
// swagger:model sonargo.RuleParam
type SonargoRuleParam struct {

	// default value
	DefaultValue string `json:"defaultValue,omitempty"`

	// html desc
	HTMLDesc string `json:"htmlDesc,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this sonargo rule param
func (m *SonargoRuleParam) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sonargo rule param based on context it is used
func (m *SonargoRuleParam) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SonargoRuleParam) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SonargoRuleParam) UnmarshalBinary(b []byte) error {
	var res SonargoRuleParam
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
