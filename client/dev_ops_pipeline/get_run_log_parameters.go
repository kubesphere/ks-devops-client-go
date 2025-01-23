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

// NewGetRunLogParams creates a new GetRunLogParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRunLogParams() *GetRunLogParams {
	return &GetRunLogParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRunLogParamsWithTimeout creates a new GetRunLogParams object
// with the ability to set a timeout on a request.
func NewGetRunLogParamsWithTimeout(timeout time.Duration) *GetRunLogParams {
	return &GetRunLogParams{
		timeout: timeout,
	}
}

// NewGetRunLogParamsWithContext creates a new GetRunLogParams object
// with the ability to set a context for a request.
func NewGetRunLogParamsWithContext(ctx context.Context) *GetRunLogParams {
	return &GetRunLogParams{
		Context: ctx,
	}
}

// NewGetRunLogParamsWithHTTPClient creates a new GetRunLogParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRunLogParamsWithHTTPClient(client *http.Client) *GetRunLogParams {
	return &GetRunLogParams{
		HTTPClient: client,
	}
}

/*
GetRunLogParams contains all the parameters to send to the API endpoint

	for the get run log operation.

	Typically these are written to a http.Request.
*/
type GetRunLogParams struct {

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

	/* Start.

	   the item number that the search starts from.

	   Format: start=%d
	   Default: "start=0"
	*/
	Start *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get run log params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRunLogParams) WithDefaults() *GetRunLogParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get run log params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRunLogParams) SetDefaults() {
	var (
		startDefault = string("start=0")
	)

	val := GetRunLogParams{
		Start: &startDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get run log params
func (o *GetRunLogParams) WithTimeout(timeout time.Duration) *GetRunLogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get run log params
func (o *GetRunLogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get run log params
func (o *GetRunLogParams) WithContext(ctx context.Context) *GetRunLogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get run log params
func (o *GetRunLogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get run log params
func (o *GetRunLogParams) WithHTTPClient(client *http.Client) *GetRunLogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get run log params
func (o *GetRunLogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevops adds the devops to the get run log params
func (o *GetRunLogParams) WithDevops(devops string) *GetRunLogParams {
	o.SetDevops(devops)
	return o
}

// SetDevops adds the devops to the get run log params
func (o *GetRunLogParams) SetDevops(devops string) {
	o.Devops = devops
}

// WithPipeline adds the pipeline to the get run log params
func (o *GetRunLogParams) WithPipeline(pipeline string) *GetRunLogParams {
	o.SetPipeline(pipeline)
	return o
}

// SetPipeline adds the pipeline to the get run log params
func (o *GetRunLogParams) SetPipeline(pipeline string) {
	o.Pipeline = pipeline
}

// WithRun adds the run to the get run log params
func (o *GetRunLogParams) WithRun(run string) *GetRunLogParams {
	o.SetRun(run)
	return o
}

// SetRun adds the run to the get run log params
func (o *GetRunLogParams) SetRun(run string) {
	o.Run = run
}

// WithStart adds the start to the get run log params
func (o *GetRunLogParams) WithStart(start *string) *GetRunLogParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get run log params
func (o *GetRunLogParams) SetStart(start *string) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *GetRunLogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Start != nil {

		// query param start
		var qrStart string

		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := qrStart
		if qStart != "" {

			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
