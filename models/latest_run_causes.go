// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LatestRunCauses latest run causes
//
// swagger:model .latestRun.causes
type LatestRunCauses struct {

	// It’s a fully qualified name and is an identifier of the producer of this resource's capability.
	Class string `json:"_class,omitempty"`

	// short description
	ShortDescription string `json:"shortDescription,omitempty"`

	// user id
	UserID string `json:"userId,omitempty"`

	// user name
	UserName string `json:"userName,omitempty"`
}

// Validate validates this latest run causes
func (m *LatestRunCauses) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this latest run causes based on context it is used
func (m *LatestRunCauses) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LatestRunCauses) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LatestRunCauses) UnmarshalBinary(b []byte) error {
	var res LatestRunCauses
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
