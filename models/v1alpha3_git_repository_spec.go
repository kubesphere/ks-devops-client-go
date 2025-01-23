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

// V1alpha3GitRepositorySpec v1alpha3 git repository spec
//
// swagger:model v1alpha3.GitRepositorySpec
type V1alpha3GitRepositorySpec struct {

	// owner
	Owner string `json:"owner,omitempty"`

	// provider
	Provider string `json:"provider,omitempty"`

	// repo
	Repo string `json:"repo,omitempty"`

	// secret
	Secret *V1SecretReference `json:"secret,omitempty"`

	// server
	Server string `json:"server,omitempty"`

	// url
	URL string `json:"url,omitempty"`

	// webhooks
	Webhooks []*V1LocalObjectReference `json:"webhooks"`
}

// Validate validates this v1alpha3 git repository spec
func (m *V1alpha3GitRepositorySpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebhooks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha3GitRepositorySpec) validateSecret(formats strfmt.Registry) error {
	if swag.IsZero(m.Secret) { // not required
		return nil
	}

	if m.Secret != nil {
		if err := m.Secret.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secret")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("secret")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha3GitRepositorySpec) validateWebhooks(formats strfmt.Registry) error {
	if swag.IsZero(m.Webhooks) { // not required
		return nil
	}

	for i := 0; i < len(m.Webhooks); i++ {
		if swag.IsZero(m.Webhooks[i]) { // not required
			continue
		}

		if m.Webhooks[i] != nil {
			if err := m.Webhooks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("webhooks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("webhooks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1alpha3 git repository spec based on the context it is used
func (m *V1alpha3GitRepositorySpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSecret(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWebhooks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha3GitRepositorySpec) contextValidateSecret(ctx context.Context, formats strfmt.Registry) error {

	if m.Secret != nil {

		if swag.IsZero(m.Secret) { // not required
			return nil
		}

		if err := m.Secret.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secret")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("secret")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha3GitRepositorySpec) contextValidateWebhooks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Webhooks); i++ {

		if m.Webhooks[i] != nil {

			if swag.IsZero(m.Webhooks[i]) { // not required
				return nil
			}

			if err := m.Webhooks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("webhooks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("webhooks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha3GitRepositorySpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha3GitRepositorySpec) UnmarshalBinary(b []byte) error {
	var res V1alpha3GitRepositorySpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
