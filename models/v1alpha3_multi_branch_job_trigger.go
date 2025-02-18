// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1alpha3MultiBranchJobTrigger v1alpha3 multi branch job trigger
//
// swagger:model v1alpha3.MultiBranchJobTrigger
type V1alpha3MultiBranchJobTrigger struct {

	// pipeline name to trigger
	CreateActionJobToTrigger string `json:"create_action_job_to_trigger,omitempty"`

	// pipeline name to trigger
	DeleteActionJobToTrigger string `json:"delete_action_job_to_trigger,omitempty"`
}

// Validate validates this v1alpha3 multi branch job trigger
func (m *V1alpha3MultiBranchJobTrigger) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1alpha3 multi branch job trigger based on context it is used
func (m *V1alpha3MultiBranchJobTrigger) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha3MultiBranchJobTrigger) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha3MultiBranchJobTrigger) UnmarshalBinary(b []byte) error {
	var res V1alpha3MultiBranchJobTrigger
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
