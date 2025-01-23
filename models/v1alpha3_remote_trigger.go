// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1alpha3RemoteTrigger v1alpha3 remote trigger
//
// swagger:model v1alpha3.RemoteTrigger
type V1alpha3RemoteTrigger struct {

	// remote trigger token
	Token string `json:"token,omitempty"`
}

// Validate validates this v1alpha3 remote trigger
func (m *V1alpha3RemoteTrigger) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1alpha3 remote trigger based on context it is used
func (m *V1alpha3RemoteTrigger) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha3RemoteTrigger) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha3RemoteTrigger) UnmarshalBinary(b []byte) error {
	var res V1alpha3RemoteTrigger
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
