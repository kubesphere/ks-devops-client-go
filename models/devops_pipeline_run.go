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

// DevopsPipelineRun devops pipeline run
//
// swagger:model devops.PipelineRun
type DevopsPipelineRun struct {

	// It’s a fully qualified name and is an identifier of the producer of this resource's capability.
	Class string `json:"_class,omitempty"`

	// references the reachable path to this resource
	Links *Links `json:"_links,omitempty"`

	// the list of all actions
	Actions []DevopsPipelineRunActions `json:"actions"`

	// the artifacts zip file
	ArtifactsZipFile DevopsPipelineRunArtifactsZipFile `json:"artifactsZipFile,omitempty"`

	// branch
	Branch DevopsPipelineRunBranch `json:"branch,omitempty"`

	// the cause of blockage
	CauseOfBlockage DevopsPipelineRunCauseOfBlockage `json:"causeOfBlockage,omitempty"`

	// causes
	Causes []*DevopsPipelineRunCauses `json:"causes"`

	// changeset information
	ChangeSet []DevopsPipelineRunChangeSet `json:"changeSet"`

	// commit id
	CommitID DevopsPipelineRunCommitID `json:"commitId,omitempty"`

	// commit url
	CommitURL DevopsPipelineRunCommitURL `json:"commitUrl,omitempty"`

	// description
	Description DevopsPipelineRunDescription `json:"description,omitempty"`

	// duration time in millis
	DurationInMillis int32 `json:"durationInMillis,omitempty"`

	// the time of enter the queue
	EnQueueTime string `json:"enQueueTime,omitempty"`

	// the time of end
	EndTime string `json:"endTime,omitempty"`

	// estimated duration time in millis
	EstimatedDurationInMillis int32 `json:"estimatedDurationInMillis,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name DevopsPipelineRunName `json:"name,omitempty"`

	// the name of organization
	Organization string `json:"organization,omitempty"`

	// the name of pipeline
	Pipeline string `json:"pipeline,omitempty"`

	// pull request
	PullRequest DevopsPipelineRunPullRequest `json:"pullRequest,omitempty"`

	// replayable or not
	Replayable bool `json:"replayable,omitempty"`

	// the result of pipeline run. e.g. SUCCESS
	Result string `json:"result,omitempty"`

	// pipeline run summary
	RunSummary string `json:"runSummary,omitempty"`

	// the time of start
	StartTime string `json:"startTime,omitempty"`

	// run state. e.g. RUNNING
	State string `json:"state,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this devops pipeline run
func (m *DevopsPipelineRun) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCauses(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DevopsPipelineRun) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *DevopsPipelineRun) validateCauses(formats strfmt.Registry) error {
	if swag.IsZero(m.Causes) { // not required
		return nil
	}

	for i := 0; i < len(m.Causes); i++ {
		if swag.IsZero(m.Causes[i]) { // not required
			continue
		}

		if m.Causes[i] != nil {
			if err := m.Causes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("causes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("causes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this devops pipeline run based on the context it is used
func (m *DevopsPipelineRun) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCauses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DevopsPipelineRun) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *DevopsPipelineRun) contextValidateCauses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Causes); i++ {

		if m.Causes[i] != nil {

			if swag.IsZero(m.Causes[i]) { // not required
				return nil
			}

			if err := m.Causes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("causes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("causes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DevopsPipelineRun) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DevopsPipelineRun) UnmarshalBinary(b []byte) error {
	var res DevopsPipelineRun
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
