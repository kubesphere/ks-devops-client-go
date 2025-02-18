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
	"github.com/go-openapi/validate"
)

// DevopsGeneralAction devops general action
//
// swagger:model devops.GeneralAction
type DevopsGeneralAction struct {

	// total count
	TotalCount int64 `json:"TotalCount,omitempty"`

	// Url name
	URLName string `json:"UrlName,omitempty"`

	// class
	Class string `json:"_class,omitempty"`

	// builds by branch name
	BuildsByBranchName map[string]DevopsBuilds `json:"buildsByBranchName,omitempty"`

	// causes
	Causes []DevopsGeneralActionCauses `json:"causes"`

	// ce task Id
	CeTaskID string `json:"ceTaskId,omitempty"`

	// last built revision
	LastBuiltRevision *DevopsBuildRevision `json:"lastBuiltRevision,omitempty"`

	// parameters
	Parameters []*DevopsGeneralParameter `json:"parameters"`

	// remote urls
	RemoteUrls []string `json:"remoteUrls"`

	// scm name
	ScmName string `json:"scmName,omitempty"`

	// server Url
	ServerURL string `json:"serverUrl,omitempty"`

	// sonarqube dashboard Url
	SonarqubeDashboardURL string `json:"sonarqubeDashboardUrl,omitempty"`

	// subdir
	Subdir DevopsGeneralActionSubdir `json:"subdir,omitempty"`
}

// Validate validates this devops general action
func (m *DevopsGeneralAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildsByBranchName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastBuiltRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DevopsGeneralAction) validateBuildsByBranchName(formats strfmt.Registry) error {
	if swag.IsZero(m.BuildsByBranchName) { // not required
		return nil
	}

	for k := range m.BuildsByBranchName {

		if err := validate.Required("buildsByBranchName"+"."+k, "body", m.BuildsByBranchName[k]); err != nil {
			return err
		}
		if val, ok := m.BuildsByBranchName[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("buildsByBranchName" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("buildsByBranchName" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *DevopsGeneralAction) validateLastBuiltRevision(formats strfmt.Registry) error {
	if swag.IsZero(m.LastBuiltRevision) { // not required
		return nil
	}

	if m.LastBuiltRevision != nil {
		if err := m.LastBuiltRevision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastBuiltRevision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastBuiltRevision")
			}
			return err
		}
	}

	return nil
}

func (m *DevopsGeneralAction) validateParameters(formats strfmt.Registry) error {
	if swag.IsZero(m.Parameters) { // not required
		return nil
	}

	for i := 0; i < len(m.Parameters); i++ {
		if swag.IsZero(m.Parameters[i]) { // not required
			continue
		}

		if m.Parameters[i] != nil {
			if err := m.Parameters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this devops general action based on the context it is used
func (m *DevopsGeneralAction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBuildsByBranchName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastBuiltRevision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParameters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DevopsGeneralAction) contextValidateBuildsByBranchName(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.BuildsByBranchName {

		if val, ok := m.BuildsByBranchName[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *DevopsGeneralAction) contextValidateLastBuiltRevision(ctx context.Context, formats strfmt.Registry) error {

	if m.LastBuiltRevision != nil {

		if swag.IsZero(m.LastBuiltRevision) { // not required
			return nil
		}

		if err := m.LastBuiltRevision.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastBuiltRevision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastBuiltRevision")
			}
			return err
		}
	}

	return nil
}

func (m *DevopsGeneralAction) contextValidateParameters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Parameters); i++ {

		if m.Parameters[i] != nil {

			if swag.IsZero(m.Parameters[i]) { // not required
				return nil
			}

			if err := m.Parameters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DevopsGeneralAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DevopsGeneralAction) UnmarshalBinary(b []byte) error {
	var res DevopsGeneralAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
