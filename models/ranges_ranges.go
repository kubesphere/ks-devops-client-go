// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RangesRanges ranges ranges
//
// swagger:model .ranges.ranges
type RangesRanges struct {

	// End build number
	End int32 `json:"end,omitempty"`

	// Start build number
	Start int32 `json:"start,omitempty"`
}

// Validate validates this ranges ranges
func (m *RangesRanges) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ranges ranges based on context it is used
func (m *RangesRanges) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RangesRanges) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RangesRanges) UnmarshalBinary(b []byte) error {
	var res RangesRanges
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
