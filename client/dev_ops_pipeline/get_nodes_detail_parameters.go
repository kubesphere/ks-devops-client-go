// Code generated by go-swagger; DO NOT EDIT.

package dev_ops_pipeline

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

// NewGetNodesDetailParams creates a new GetNodesDetailParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNodesDetailParams() *GetNodesDetailParams {
	return &GetNodesDetailParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNodesDetailParamsWithTimeout creates a new GetNodesDetailParams object
// with the ability to set a timeout on a request.
func NewGetNodesDetailParamsWithTimeout(timeout time.Duration) *GetNodesDetailParams {
	return &GetNodesDetailParams{
		timeout: timeout,
	}
}

// NewGetNodesDetailParamsWithContext creates a new GetNodesDetailParams object
// with the ability to set a context for a request.
func NewGetNodesDetailParamsWithContext(ctx context.Context) *GetNodesDetailParams {
	return &GetNodesDetailParams{
		Context: ctx,
	}
}

// NewGetNodesDetailParamsWithHTTPClient creates a new GetNodesDetailParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNodesDetailParamsWithHTTPClient(client *http.Client) *GetNodesDetailParams {
	return &GetNodesDetailParams{
		HTTPClient: client,
	}
}

/*
GetNodesDetailParams contains all the parameters to send to the API endpoint

	for the get nodes detail operation.

	Typically these are written to a http.Request.
*/
type GetNodesDetailParams struct {

	/* Devops.

	   DevOps project's ID, e.g. project-RRRRAzLBlLEm
	*/
	Devops string

	/* Pipeline.

	   the name of the CI/CD pipeline
	*/
	Pipeline string

	/* Run.

	   pipeline run ID, the unique ID for a pipeline once build.
	*/
	Run string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get nodes detail params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNodesDetailParams) WithDefaults() *GetNodesDetailParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get nodes detail params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNodesDetailParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get nodes detail params
func (o *GetNodesDetailParams) WithTimeout(timeout time.Duration) *GetNodesDetailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nodes detail params
func (o *GetNodesDetailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nodes detail params
func (o *GetNodesDetailParams) WithContext(ctx context.Context) *GetNodesDetailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nodes detail params
func (o *GetNodesDetailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nodes detail params
func (o *GetNodesDetailParams) WithHTTPClient(client *http.Client) *GetNodesDetailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nodes detail params
func (o *GetNodesDetailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevops adds the devops to the get nodes detail params
func (o *GetNodesDetailParams) WithDevops(devops string) *GetNodesDetailParams {
	o.SetDevops(devops)
	return o
}

// SetDevops adds the devops to the get nodes detail params
func (o *GetNodesDetailParams) SetDevops(devops string) {
	o.Devops = devops
}

// WithPipeline adds the pipeline to the get nodes detail params
func (o *GetNodesDetailParams) WithPipeline(pipeline string) *GetNodesDetailParams {
	o.SetPipeline(pipeline)
	return o
}

// SetPipeline adds the pipeline to the get nodes detail params
func (o *GetNodesDetailParams) SetPipeline(pipeline string) {
	o.Pipeline = pipeline
}

// WithRun adds the run to the get nodes detail params
func (o *GetNodesDetailParams) WithRun(run string) *GetNodesDetailParams {
	o.SetRun(run)
	return o
}

// SetRun adds the run to the get nodes detail params
func (o *GetNodesDetailParams) SetRun(run string) {
	o.Run = run
}

// WriteToRequest writes these params to a swagger request
func (o *GetNodesDetailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param devops
	if err := r.SetPathParam("devops", o.Devops); err != nil {
		return err
	}

	// path param pipeline
	if err := r.SetPathParam("pipeline", o.Pipeline); err != nil {
		return err
	}

	// path param run
	if err := r.SetPathParam("run", o.Run); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
