// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// JobEdge job edge
//
// swagger:model job.Edge
type JobEdge struct {

	// id
	ID string `json:"id,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this job edge
func (m *JobEdge) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this job edge based on context it is used
func (m *JobEdge) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *JobEdge) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobEdge) UnmarshalBinary(b []byte) error {
	var res JobEdge
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
