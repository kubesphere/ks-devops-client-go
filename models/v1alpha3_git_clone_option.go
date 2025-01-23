// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1alpha3GitCloneOption v1alpha3 git clone option
//
// swagger:model v1alpha3.GitCloneOption
type V1alpha3GitCloneOption struct {

	// git clone depth
	Depth int32 `json:"depth,omitempty"`

	// Whether to use git shallow clone
	Shallow bool `json:"shallow,omitempty"`

	// git clone timeout mins
	Timeout int32 `json:"timeout,omitempty"`
}

// Validate validates this v1alpha3 git clone option
func (m *V1alpha3GitCloneOption) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1alpha3 git clone option based on context it is used
func (m *V1alpha3GitCloneOption) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha3GitCloneOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha3GitCloneOption) UnmarshalBinary(b []byte) error {
	var res V1alpha3GitCloneOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
