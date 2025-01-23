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
)

// NewDeleteDevOpsProjectParams creates a new DeleteDevOpsProjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDevOpsProjectParams() *DeleteDevOpsProjectParams {
	return &DeleteDevOpsProjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDevOpsProjectParamsWithTimeout creates a new DeleteDevOpsProjectParams object
// with the ability to set a timeout on a request.
func NewDeleteDevOpsProjectParamsWithTimeout(timeout time.Duration) *DeleteDevOpsProjectParams {
	return &DeleteDevOpsProjectParams{
		timeout: timeout,
	}
}

// NewDeleteDevOpsProjectParamsWithContext creates a new DeleteDevOpsProjectParams object
// with the ability to set a context for a request.
func NewDeleteDevOpsProjectParamsWithContext(ctx context.Context) *DeleteDevOpsProjectParams {
	return &DeleteDevOpsProjectParams{
		Context: ctx,
	}
}

// NewDeleteDevOpsProjectParamsWithHTTPClient creates a new DeleteDevOpsProjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDevOpsProjectParamsWithHTTPClient(client *http.Client) *DeleteDevOpsProjectParams {
	return &DeleteDevOpsProjectParams{
		HTTPClient: client,
	}
}

/*
DeleteDevOpsProjectParams contains all the parameters to send to the API endpoint

	for the delete dev ops project operation.

	Typically these are written to a http.Request.
*/
type DeleteDevOpsProjectParams struct {

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

// WithDefaults hydrates default values in the delete dev ops project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDevOpsProjectParams) WithDefaults() *DeleteDevOpsProjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete dev ops project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDevOpsProjectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete dev ops project params
func (o *DeleteDevOpsProjectParams) WithTimeout(timeout time.Duration) *DeleteDevOpsProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete dev ops project params
func (o *DeleteDevOpsProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete dev ops project params
func (o *DeleteDevOpsProjectParams) WithContext(ctx context.Context) *DeleteDevOpsProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete dev ops project params
func (o *DeleteDevOpsProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete dev ops project params
func (o *DeleteDevOpsProjectParams) WithHTTPClient(client *http.Client) *DeleteDevOpsProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete dev ops project params
func (o *DeleteDevOpsProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevops adds the devops to the delete dev ops project params
func (o *DeleteDevOpsProjectParams) WithDevops(devops string) *DeleteDevOpsProjectParams {
	o.SetDevops(devops)
	return o
}

// SetDevops adds the devops to the delete dev ops project params
func (o *DeleteDevOpsProjectParams) SetDevops(devops string) {
	o.Devops = devops
}

// WithWorkspace adds the workspace to the delete dev ops project params
func (o *DeleteDevOpsProjectParams) WithWorkspace(workspace string) *DeleteDevOpsProjectParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the delete dev ops project params
func (o *DeleteDevOpsProjectParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDevOpsProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
