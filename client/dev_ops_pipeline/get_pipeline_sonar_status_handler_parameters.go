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

// NewGetPipelineSonarStatusHandlerParams creates a new GetPipelineSonarStatusHandlerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPipelineSonarStatusHandlerParams() *GetPipelineSonarStatusHandlerParams {
	return &GetPipelineSonarStatusHandlerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPipelineSonarStatusHandlerParamsWithTimeout creates a new GetPipelineSonarStatusHandlerParams object
// with the ability to set a timeout on a request.
func NewGetPipelineSonarStatusHandlerParamsWithTimeout(timeout time.Duration) *GetPipelineSonarStatusHandlerParams {
	return &GetPipelineSonarStatusHandlerParams{
		timeout: timeout,
	}
}

// NewGetPipelineSonarStatusHandlerParamsWithContext creates a new GetPipelineSonarStatusHandlerParams object
// with the ability to set a context for a request.
func NewGetPipelineSonarStatusHandlerParamsWithContext(ctx context.Context) *GetPipelineSonarStatusHandlerParams {
	return &GetPipelineSonarStatusHandlerParams{
		Context: ctx,
	}
}

// NewGetPipelineSonarStatusHandlerParamsWithHTTPClient creates a new GetPipelineSonarStatusHandlerParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPipelineSonarStatusHandlerParamsWithHTTPClient(client *http.Client) *GetPipelineSonarStatusHandlerParams {
	return &GetPipelineSonarStatusHandlerParams{
		HTTPClient: client,
	}
}

/*
GetPipelineSonarStatusHandlerParams contains all the parameters to send to the API endpoint

	for the get pipeline sonar status handler operation.

	Typically these are written to a http.Request.
*/
type GetPipelineSonarStatusHandlerParams struct {

	/* Devops.

	   DevOps project's ID, e.g. project-RRRRAzLBlLEm
	*/
	Devops string

	/* Pipeline.

	   the name of pipeline, e.g. sample-pipeline
	*/
	Pipeline string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get pipeline sonar status handler params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPipelineSonarStatusHandlerParams) WithDefaults() *GetPipelineSonarStatusHandlerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get pipeline sonar status handler params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPipelineSonarStatusHandlerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get pipeline sonar status handler params
func (o *GetPipelineSonarStatusHandlerParams) WithTimeout(timeout time.Duration) *GetPipelineSonarStatusHandlerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get pipeline sonar status handler params
func (o *GetPipelineSonarStatusHandlerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get pipeline sonar status handler params
func (o *GetPipelineSonarStatusHandlerParams) WithContext(ctx context.Context) *GetPipelineSonarStatusHandlerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get pipeline sonar status handler params
func (o *GetPipelineSonarStatusHandlerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get pipeline sonar status handler params
func (o *GetPipelineSonarStatusHandlerParams) WithHTTPClient(client *http.Client) *GetPipelineSonarStatusHandlerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get pipeline sonar status handler params
func (o *GetPipelineSonarStatusHandlerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevops adds the devops to the get pipeline sonar status handler params
func (o *GetPipelineSonarStatusHandlerParams) WithDevops(devops string) *GetPipelineSonarStatusHandlerParams {
	o.SetDevops(devops)
	return o
}

// SetDevops adds the devops to the get pipeline sonar status handler params
func (o *GetPipelineSonarStatusHandlerParams) SetDevops(devops string) {
	o.Devops = devops
}

// WithPipeline adds the pipeline to the get pipeline sonar status handler params
func (o *GetPipelineSonarStatusHandlerParams) WithPipeline(pipeline string) *GetPipelineSonarStatusHandlerParams {
	o.SetPipeline(pipeline)
	return o
}

// SetPipeline adds the pipeline to the get pipeline sonar status handler params
func (o *GetPipelineSonarStatusHandlerParams) SetPipeline(pipeline string) {
	o.Pipeline = pipeline
}

// WriteToRequest writes these params to a swagger request
func (o *GetPipelineSonarStatusHandlerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
