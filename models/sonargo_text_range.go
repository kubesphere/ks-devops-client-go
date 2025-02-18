// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SonargoTextRange sonargo text range
//
// swagger:model sonargo.TextRange
type SonargoTextRange struct {

	// end line
	EndLine int32 `json:"endLine,omitempty"`

	// end offset
	EndOffset int32 `json:"endOffset,omitempty"`

	// start line
	StartLine int32 `json:"startLine,omitempty"`

	// start offset
	StartOffset int32 `json:"startOffset,omitempty"`
}

// Validate validates this sonargo text range
func (m *SonargoTextRange) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sonargo text range based on context it is used
func (m *SonargoTextRange) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SonargoTextRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SonargoTextRange) UnmarshalBinary(b []byte) error {
	var res SonargoTextRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
