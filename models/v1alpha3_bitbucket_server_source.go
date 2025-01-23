// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1alpha3BitbucketServerSource v1alpha3 bitbucket server source
//
// swagger:model v1alpha3.BitbucketServerSource
type V1alpha3BitbucketServerSource struct {

	// Allow Jenkins send build status notification to Bitbucket
	AcceptJenkinsNotification bool `json:"accept_jenkins_notification,omitempty"`

	// The api url can specify the location of the github apiserver.For private cloud configuration
	APIURI string `json:"api_uri,omitempty"`

	// credential id to access github source
	CredentialID string `json:"credential_id,omitempty"`

	// Discover branch configuration
	DiscoverBranches int32 `json:"discover_branches,omitempty"`

	// Discover fork PR configuration
	DiscoverPrFromForks *V1alpha3DiscoverPRFromForks `json:"discover_pr_from_forks,omitempty"`

	// Discover origin PR configuration
	DiscoverPrFromOrigin int32 `json:"discover_pr_from_origin,omitempty"`

	// Discover tag configuration
	DiscoverTags bool `json:"discover_tags,omitempty"`

	// advavced git clone options
	GitCloneOption *V1alpha3GitCloneOption `json:"git_clone_option,omitempty"`

	// owner of github repo
	Owner string `json:"owner,omitempty"`

	// Regex used to match the name of the branch that needs to be run
	RegexFilter string `json:"regex_filter,omitempty"`

	// repo name of github repo
	Repo string `json:"repo,omitempty"`

	// uid of scm
	ScmID string `json:"scm_id,omitempty"`
}

// Validate validates this v1alpha3 bitbucket server source
func (m *V1alpha3BitbucketServerSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDiscoverPrFromForks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGitCloneOption(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha3BitbucketServerSource) validateDiscoverPrFromForks(formats strfmt.Registry) error {
	if swag.IsZero(m.DiscoverPrFromForks) { // not required
		return nil
	}

	if m.DiscoverPrFromForks != nil {
		if err := m.DiscoverPrFromForks.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("discover_pr_from_forks")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("discover_pr_from_forks")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha3BitbucketServerSource) validateGitCloneOption(formats strfmt.Registry) error {
	if swag.IsZero(m.GitCloneOption) { // not required
		return nil
	}

	if m.GitCloneOption != nil {
		if err := m.GitCloneOption.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("git_clone_option")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("git_clone_option")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1alpha3 bitbucket server source based on the context it is used
func (m *V1alpha3BitbucketServerSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDiscoverPrFromForks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGitCloneOption(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha3BitbucketServerSource) contextValidateDiscoverPrFromForks(ctx context.Context, formats strfmt.Registry) error {

	if m.DiscoverPrFromForks != nil {

		if swag.IsZero(m.DiscoverPrFromForks) { // not required
			return nil
		}

		if err := m.DiscoverPrFromForks.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("discover_pr_from_forks")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("discover_pr_from_forks")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha3BitbucketServerSource) contextValidateGitCloneOption(ctx context.Context, formats strfmt.Registry) error {

	if m.GitCloneOption != nil {

		if swag.IsZero(m.GitCloneOption) { // not required
			return nil
		}

		if err := m.GitCloneOption.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("git_clone_option")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("git_clone_option")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha3BitbucketServerSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha3BitbucketServerSource) UnmarshalBinary(b []byte) error {
	var res V1alpha3BitbucketServerSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
