// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1alpha3TimerTrigger v1alpha3 timer trigger
//
// swagger:model v1alpha3.TimerTrigger
type V1alpha3TimerTrigger struct {

	// jenkins cron script
	Cron string `json:"cron,omitempty"`

	// interval ms
	Interval string `json:"interval,omitempty"`
}

// Validate validates this v1alpha3 timer trigger
func (m *V1alpha3TimerTrigger) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1alpha3 timer trigger based on context it is used
func (m *V1alpha3TimerTrigger) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha3TimerTrigger) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha3TimerTrigger) UnmarshalBinary(b []byte) error {
	var res V1alpha3TimerTrigger
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
