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

// NewGetStepLogParams creates a new GetStepLogParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetStepLogParams() *GetStepLogParams {
	return &GetStepLogParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetStepLogParamsWithTimeout creates a new GetStepLogParams object
// with the ability to set a timeout on a request.
func NewGetStepLogParamsWithTimeout(timeout time.Duration) *GetStepLogParams {
	return &GetStepLogParams{
		timeout: timeout,
	}
}

// NewGetStepLogParamsWithContext creates a new GetStepLogParams object
// with the ability to set a context for a request.
func NewGetStepLogParamsWithContext(ctx context.Context) *GetStepLogParams {
	return &GetStepLogParams{
		Context: ctx,
	}
}

// NewGetStepLogParamsWithHTTPClient creates a new GetStepLogParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetStepLogParamsWithHTTPClient(client *http.Client) *GetStepLogParams {
	return &GetStepLogParams{
		HTTPClient: client,
	}
}

/*
GetStepLogParams contains all the parameters to send to the API endpoint

	for the get step log operation.

	Typically these are written to a http.Request.
*/
type GetStepLogParams struct {

	/* Devops.

	   DevOps project's ID, e.g. project-RRRRAzLBlLEm
	*/
	Devops string

	/* Node.

	   pipeline node ID, the stage in pipeline.
	*/
	Node string

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

	/* Step.

	   pipeline step ID, the step in pipeline.
	*/
	Step string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get step log params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStepLogParams) WithDefaults() *GetStepLogParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get step log params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStepLogParams) SetDefaults() {
	var (
		startDefault = string("start=0")
	)

	val := GetStepLogParams{
		Start: &startDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get step log params
func (o *GetStepLogParams) WithTimeout(timeout time.Duration) *GetStepLogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get step log params
func (o *GetStepLogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get step log params
func (o *GetStepLogParams) WithContext(ctx context.Context) *GetStepLogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get step log params
func (o *GetStepLogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get step log params
func (o *GetStepLogParams) WithHTTPClient(client *http.Client) *GetStepLogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get step log params
func (o *GetStepLogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevops adds the devops to the get step log params
func (o *GetStepLogParams) WithDevops(devops string) *GetStepLogParams {
	o.SetDevops(devops)
	return o
}

// SetDevops adds the devops to the get step log params
func (o *GetStepLogParams) SetDevops(devops string) {
	o.Devops = devops
}

// WithNode adds the node to the get step log params
func (o *GetStepLogParams) WithNode(node string) *GetStepLogParams {
	o.SetNode(node)
	return o
}

// SetNode adds the node to the get step log params
func (o *GetStepLogParams) SetNode(node string) {
	o.Node = node
}

// WithPipeline adds the pipeline to the get step log params
func (o *GetStepLogParams) WithPipeline(pipeline string) *GetStepLogParams {
	o.SetPipeline(pipeline)
	return o
}

// SetPipeline adds the pipeline to the get step log params
func (o *GetStepLogParams) SetPipeline(pipeline string) {
	o.Pipeline = pipeline
}

// WithRun adds the run to the get step log params
func (o *GetStepLogParams) WithRun(run string) *GetStepLogParams {
	o.SetRun(run)
	return o
}

// SetRun adds the run to the get step log params
func (o *GetStepLogParams) SetRun(run string) {
	o.Run = run
}

// WithStart adds the start to the get step log params
func (o *GetStepLogParams) WithStart(start *string) *GetStepLogParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get step log params
func (o *GetStepLogParams) SetStart(start *string) {
	o.Start = start
}

// WithStep adds the step to the get step log params
func (o *GetStepLogParams) WithStep(step string) *GetStepLogParams {
	o.SetStep(step)
	return o
}

// SetStep adds the step to the get step log params
func (o *GetStepLogParams) SetStep(step string) {
	o.Step = step
}

// WriteToRequest writes these params to a swagger request
func (o *GetStepLogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param devops
	if err := r.SetPathParam("devops", o.Devops); err != nil {
		return err
	}

	// path param node
	if err := r.SetPathParam("node", o.Node); err != nil {
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

	// path param step
	if err := r.SetPathParam("step", o.Step); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
