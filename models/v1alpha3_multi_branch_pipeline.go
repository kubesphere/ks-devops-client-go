// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1alpha3MultiBranchPipeline v1alpha3 multi branch pipeline
//
// swagger:model v1alpha3.MultiBranchPipeline
type V1alpha3MultiBranchPipeline struct {

	// bitbucket server scm defile
	BitbucketServerSource *V1alpha3BitbucketServerSource `json:"bitbucket_server_source,omitempty"`

	// description of pipeline
	Description string `json:"description,omitempty"`

	// Discarder of pipeline, managing when to drop a pipeline
	Discarder *V1alpha3DiscarderProperty `json:"discarder,omitempty"`

	// git scm define
	GitSource *V1alpha3GitSource `json:"git_source,omitempty"`

	// github scm define
	GithubSource *V1alpha3GithubSource `json:"github_source,omitempty"`

	// gitlab scm define
	GitlabSource *V1alpha3GitlabSource `json:"gitlab_source,omitempty"`

	// Pipeline tasks that need to be triggered when branch creation/deletion
	MultibranchJobTrigger *V1alpha3MultiBranchJobTrigger `json:"multibranch_job_trigger,omitempty"`

	// name of pipeline
	// Required: true
	Name *string `json:"name"`

	// script path in scm
	// Required: true
	ScriptPath *string `json:"script_path"`

	// single branch svn scm define
	SingleSvnSource *V1alpha3SingleSvnSource `json:"single_svn_source,omitempty"`

	// type of scm, such as github/git/svn
	// Required: true
	SourceType *string `json:"source_type"`

	// multi branch svn scm define
	SvnSource *V1alpha3SvnSource `json:"svn_source,omitempty"`

	// Timer to trigger pipeline run
	TimerTrigger *V1alpha3TimerTrigger `json:"timer_trigger,omitempty"`
}

// Validate validates this v1alpha3 multi branch pipeline
func (m *V1alpha3MultiBranchPipeline) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBitbucketServerSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiscarder(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGitSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGithubSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGitlabSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMultibranchJobTrigger(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScriptPath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSingleSvnSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvnSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimerTrigger(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha3MultiBranchPipeline) validateBitbucketServerSource(formats strfmt.Registry) error {
	if swag.IsZero(m.BitbucketServerSource) { // not required
		return nil
	}

	if m.BitbucketServerSource != nil {
		if err := m.BitbucketServerSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bitbucket_server_source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bitbucket_server_source")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha3MultiBranchPipeline) validateDiscarder(formats strfmt.Registry) error {
	if swag.IsZero(m.Discarder) { // not required
		return nil
	}

	if m.Discarder != nil {
		if err := m.Discarder.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("discarder")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("discarder")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha3MultiBranchPipeline) validateGitSource(formats strfmt.Registry) error {
	if swag.IsZero(m.GitSource) { // not required
		return nil
	}

	if m.GitSource != nil {
		if err := m.GitSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("git_source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("git_source")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha3MultiBranchPipeline) validateGithubSource(formats strfmt.Registry) error {
	if swag.IsZero(m.GithubSource) { // not required
		return nil
	}

	if m.GithubSource != nil {
		if err := m.GithubSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("github_source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("github_source")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha3MultiBranchPipeline) validateGitlabSource(formats strfmt.Registry) error {
	if swag.IsZero(m.GitlabSource) { // not required
		return nil
	}

	if m.GitlabSource != nil {
		if err := m.GitlabSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gitlab_source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gitlab_source")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha3MultiBranchPipeline) validateMultibranchJobTrigger(formats strfmt.Registry) error {
	if swag.IsZero(m.MultibranchJobTrigger) { // not required
		return nil
	}

	if m.MultibranchJobTrigger != nil {
		if err := m.MultibranchJobTrigger.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("multibranch_job_trigger")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("multibranch_job_trigger")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha3MultiBranchPipeline) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1alpha3MultiBranchPipeline) validateScriptPath(formats strfmt.Registry) error {

	if err := validate.Required("script_path", "body", m.ScriptPath); err != nil {
		return err
	}

	return nil
}

func (m *V1alpha3MultiBranchPipeline) validateSingleSvnSource(formats strfmt.Registry) error {
	if swag.IsZero(m.SingleSvnSource) { // not required
		return nil
	}

	if m.SingleSvnSource != nil {
		if err := m.SingleSvnSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("single_svn_source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("single_svn_source")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha3MultiBranchPipeline) validateSourceType(formats strfmt.Registry) error {

	if err := validate.Required("source_type", "body", m.SourceType); err != nil {
		return err
	}

	return nil
}

func (m *V1alpha3MultiBranchPipeline) validateSvnSource(formats strfmt.Registry) error {
	if swag.IsZero(m.SvnSource) { // not required
		return nil
	}

	if m.SvnSource != nil {
		if err := m.SvnSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svn_source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("svn_source")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha3MultiBranchPipeline) validateTimerTrigger(formats strfmt.Registry) error {
	if swag.IsZero(m.TimerTrigger) { // not required
		return nil
	}

	if m.TimerTrigger != nil {
		if err := m.TimerTrigger.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timer_trigger")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("timer_trigger")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1alpha3 multi branch pipeline based on the context it is used
func (m *V1alpha3MultiBranchPipeline) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBitbucketServerSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDiscarder(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGitSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGithubSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGitlabSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMultibranchJobTrigger(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSingleSvnSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvnSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimerTrigger(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha3MultiBranchPipeline) contextValidateBitbucketServerSource(ctx context.Context, formats strfmt.Registry) error {

	if m.BitbucketServerSource != nil {

		if swag.IsZero(m.BitbucketServerSource) { // not required
			return nil
		}

		if err := m.BitbucketServerSource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bitbucket_server_source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bitbucket_server_source")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha3MultiBranchPipeline) contextValidateDiscarder(ctx context.Context, formats strfmt.Registry) error {

	if m.Discarder != nil {

		if swag.IsZero(m.Discarder) { // not required
			return nil
		}

		if err := m.Discarder.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("discarder")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("discarder")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha3MultiBranchPipeline) contextValidateGitSource(ctx context.Context, formats strfmt.Registry) error {

	if m.GitSource != nil {

		if swag.IsZero(m.GitSource) { // not required
			return nil
		}

		if err := m.GitSource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("git_source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("git_source")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha3MultiBranchPipeline) contextValidateGithubSource(ctx context.Context, formats strfmt.Registry) error {

	if m.GithubSource != nil {

		if swag.IsZero(m.GithubSource) { // not required
			return nil
		}

		if err := m.GithubSource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("github_source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("github_source")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha3MultiBranchPipeline) contextValidateGitlabSource(ctx context.Context, formats strfmt.Registry) error {

	if m.GitlabSource != nil {

		if swag.IsZero(m.GitlabSource) { // not required
			return nil
		}

		if err := m.GitlabSource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gitlab_source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gitlab_source")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha3MultiBranchPipeline) contextValidateMultibranchJobTrigger(ctx context.Context, formats strfmt.Registry) error {

	if m.MultibranchJobTrigger != nil {

		if swag.IsZero(m.MultibranchJobTrigger) { // not required
			return nil
		}

		if err := m.MultibranchJobTrigger.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("multibranch_job_trigger")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("multibranch_job_trigger")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha3MultiBranchPipeline) contextValidateSingleSvnSource(ctx context.Context, formats strfmt.Registry) error {

	if m.SingleSvnSource != nil {

		if swag.IsZero(m.SingleSvnSource) { // not required
			return nil
		}

		if err := m.SingleSvnSource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("single_svn_source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("single_svn_source")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha3MultiBranchPipeline) contextValidateSvnSource(ctx context.Context, formats strfmt.Registry) error {

	if m.SvnSource != nil {

		if swag.IsZero(m.SvnSource) { // not required
			return nil
		}

		if err := m.SvnSource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svn_source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("svn_source")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha3MultiBranchPipeline) contextValidateTimerTrigger(ctx context.Context, formats strfmt.Registry) error {

	if m.TimerTrigger != nil {

		if swag.IsZero(m.TimerTrigger) { // not required
			return nil
		}

		if err := m.TimerTrigger.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timer_trigger")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("timer_trigger")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha3MultiBranchPipeline) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha3MultiBranchPipeline) UnmarshalBinary(b []byte) error {
	var res V1alpha3MultiBranchPipeline
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
