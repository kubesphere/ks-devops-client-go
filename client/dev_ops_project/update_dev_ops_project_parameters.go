// Code generated by go-swagger; DO NOT EDIT.

package dev_ops_project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/kubesphere/ks-devops-client-go/models"
)

// NewUpdateDevOpsProjectParams creates a new UpdateDevOpsProjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDevOpsProjectParams() *UpdateDevOpsProjectParams {
	return &UpdateDevOpsProjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDevOpsProjectParamsWithTimeout creates a new UpdateDevOpsProjectParams object
// with the ability to set a timeout on a request.
func NewUpdateDevOpsProjectParamsWithTimeout(timeout time.Duration) *UpdateDevOpsProjectParams {
	return &UpdateDevOpsProjectParams{
		timeout: timeout,
	}
}

// NewUpdateDevOpsProjectParamsWithContext creates a new UpdateDevOpsProjectParams object
// with the ability to set a context for a request.
func NewUpdateDevOpsProjectParamsWithContext(ctx context.Context) *UpdateDevOpsProjectParams {
	return &UpdateDevOpsProjectParams{
		Context: ctx,
	}
}

// NewUpdateDevOpsProjectParamsWithHTTPClient creates a new UpdateDevOpsProjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDevOpsProjectParamsWithHTTPClient(client *http.Client) *UpdateDevOpsProjectParams {
	return &UpdateDevOpsProjectParams{
		HTTPClient: client,
	}
}

/*
UpdateDevOpsProjectParams contains all the parameters to send to the API endpoint

	for the update dev ops project operation.

	Typically these are written to a http.Request.
*/
type UpdateDevOpsProjectParams struct {

	// Body.
	Body *models.V1alpha3DevOpsProject

	/* Devops.

	   project name
	*/
	Devops string

	/* Workspace.

	   workspace name
	*/
	Workspace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update dev ops project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDevOpsProjectParams) WithDefaults() *UpdateDevOpsProjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update dev ops project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDevOpsProjectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update dev ops project params
func (o *UpdateDevOpsProjectParams) WithTimeout(timeout time.Duration) *UpdateDevOpsProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update dev ops project params
func (o *UpdateDevOpsProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update dev ops project params
func (o *UpdateDevOpsProjectParams) WithContext(ctx context.Context) *UpdateDevOpsProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update dev ops project params
func (o *UpdateDevOpsProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update dev ops project params
func (o *UpdateDevOpsProjectParams) WithHTTPClient(client *http.Client) *UpdateDevOpsProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update dev ops project params
func (o *UpdateDevOpsProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update dev ops project params
func (o *UpdateDevOpsProjectParams) WithBody(body *models.V1alpha3DevOpsProject) *UpdateDevOpsProjectParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update dev ops project params
func (o *UpdateDevOpsProjectParams) SetBody(body *models.V1alpha3DevOpsProject) {
	o.Body = body
}

// WithDevops adds the devops to the update dev ops project params
func (o *UpdateDevOpsProjectParams) WithDevops(devops string) *UpdateDevOpsProjectParams {
	o.SetDevops(devops)
	return o
}

// SetDevops adds the devops to the update dev ops project params
func (o *UpdateDevOpsProjectParams) SetDevops(devops string) {
	o.Devops = devops
}

// WithWorkspace adds the workspace to the update dev ops project params
func (o *UpdateDevOpsProjectParams) WithWorkspace(workspace string) *UpdateDevOpsProjectParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the update dev ops project params
func (o *UpdateDevOpsProjectParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDevOpsProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param devops
	if err := r.SetPathParam("devops", o.Devops); err != nil {
		return err
	}

	// path param workspace
	if err := r.SetPathParam("workspace", o.Workspace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
