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

// NewStopBranchPipelineParams creates a new StopBranchPipelineParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStopBranchPipelineParams() *StopBranchPipelineParams {
	return &StopBranchPipelineParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStopBranchPipelineParamsWithTimeout creates a new StopBranchPipelineParams object
// with the ability to set a timeout on a request.
func NewStopBranchPipelineParamsWithTimeout(timeout time.Duration) *StopBranchPipelineParams {
	return &StopBranchPipelineParams{
		timeout: timeout,
	}
}

// NewStopBranchPipelineParamsWithContext creates a new StopBranchPipelineParams object
// with the ability to set a context for a request.
func NewStopBranchPipelineParamsWithContext(ctx context.Context) *StopBranchPipelineParams {
	return &StopBranchPipelineParams{
		Context: ctx,
	}
}

// NewStopBranchPipelineParamsWithHTTPClient creates a new StopBranchPipelineParams object
// with the ability to set a custom HTTPClient for a request.
func NewStopBranchPipelineParamsWithHTTPClient(client *http.Client) *StopBranchPipelineParams {
	return &StopBranchPipelineParams{
		HTTPClient: client,
	}
}

/*
StopBranchPipelineParams contains all the parameters to send to the API endpoint

	for the stop branch pipeline operation.

	Typically these are written to a http.Request.
*/
type StopBranchPipelineParams struct {

	/* Blocking.

	   stop and between each retries will sleep.

	   Format: blocking=%t
	   Default: "blocking=false"
	*/
	Blocking *string

	// Body.
	Body []int64

	/* Branch.

	   the name of branch, same as repository branch.
	*/
	Branch string

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

	/* TimeOutInSecs.

	   the time of stop and between each retries sleep.

	   Format: timeOutInSecs=%d
	   Default: "timeOutInSecs=10"
	*/
	TimeOutInSecs *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stop branch pipeline params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopBranchPipelineParams) WithDefaults() *StopBranchPipelineParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stop branch pipeline params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopBranchPipelineParams) SetDefaults() {
	var (
		blockingDefault = string("blocking=false")

		timeOutInSecsDefault = string("timeOutInSecs=10")
	)

	val := StopBranchPipelineParams{
		Blocking:      &blockingDefault,
		TimeOutInSecs: &timeOutInSecsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the stop branch pipeline params
func (o *StopBranchPipelineParams) WithTimeout(timeout time.Duration) *StopBranchPipelineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop branch pipeline params
func (o *StopBranchPipelineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop branch pipeline params
func (o *StopBranchPipelineParams) WithContext(ctx context.Context) *StopBranchPipelineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop branch pipeline params
func (o *StopBranchPipelineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop branch pipeline params
func (o *StopBranchPipelineParams) WithHTTPClient(client *http.Client) *StopBranchPipelineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop branch pipeline params
func (o *StopBranchPipelineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBlocking adds the blocking to the stop branch pipeline params
func (o *StopBranchPipelineParams) WithBlocking(blocking *string) *StopBranchPipelineParams {
	o.SetBlocking(blocking)
	return o
}

// SetBlocking adds the blocking to the stop branch pipeline params
func (o *StopBranchPipelineParams) SetBlocking(blocking *string) {
	o.Blocking = blocking
}

// WithBody adds the body to the stop branch pipeline params
func (o *StopBranchPipelineParams) WithBody(body []int64) *StopBranchPipelineParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the stop branch pipeline params
func (o *StopBranchPipelineParams) SetBody(body []int64) {
	o.Body = body
}

// WithBranch adds the branch to the stop branch pipeline params
func (o *StopBranchPipelineParams) WithBranch(branch string) *StopBranchPipelineParams {
	o.SetBranch(branch)
	return o
}

// SetBranch adds the branch to the stop branch pipeline params
func (o *StopBranchPipelineParams) SetBranch(branch string) {
	o.Branch = branch
}

// WithDevops adds the devops to the stop branch pipeline params
func (o *StopBranchPipelineParams) WithDevops(devops string) *StopBranchPipelineParams {
	o.SetDevops(devops)
	return o
}

// SetDevops adds the devops to the stop branch pipeline params
func (o *StopBranchPipelineParams) SetDevops(devops string) {
	o.Devops = devops
}

// WithPipeline adds the pipeline to the stop branch pipeline params
func (o *StopBranchPipelineParams) WithPipeline(pipeline string) *StopBranchPipelineParams {
	o.SetPipeline(pipeline)
	return o
}

// SetPipeline adds the pipeline to the stop branch pipeline params
func (o *StopBranchPipelineParams) SetPipeline(pipeline string) {
	o.Pipeline = pipeline
}

// WithRun adds the run to the stop branch pipeline params
func (o *StopBranchPipelineParams) WithRun(run string) *StopBranchPipelineParams {
	o.SetRun(run)
	return o
}

// SetRun adds the run to the stop branch pipeline params
func (o *StopBranchPipelineParams) SetRun(run string) {
	o.Run = run
}

// WithTimeOutInSecs adds the timeOutInSecs to the stop branch pipeline params
func (o *StopBranchPipelineParams) WithTimeOutInSecs(timeOutInSecs *string) *StopBranchPipelineParams {
	o.SetTimeOutInSecs(timeOutInSecs)
	return o
}

// SetTimeOutInSecs adds the timeOutInSecs to the stop branch pipeline params
func (o *StopBranchPipelineParams) SetTimeOutInSecs(timeOutInSecs *string) {
	o.TimeOutInSecs = timeOutInSecs
}

// WriteToRequest writes these params to a swagger request
func (o *StopBranchPipelineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Blocking != nil {

		// query param blocking
		var qrBlocking string

		if o.Blocking != nil {
			qrBlocking = *o.Blocking
		}
		qBlocking := qrBlocking
		if qBlocking != "" {

			if err := r.SetQueryParam("blocking", qBlocking); err != nil {
				return err
			}
		}
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param branch
	if err := r.SetPathParam("branch", o.Branch); err != nil {
		return err
	}

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

	if o.TimeOutInSecs != nil {

		// query param timeOutInSecs
		var qrTimeOutInSecs string

		if o.TimeOutInSecs != nil {
			qrTimeOutInSecs = *o.TimeOutInSecs
		}
		qTimeOutInSecs := qrTimeOutInSecs
		if qTimeOutInSecs != "" {

			if err := r.SetQueryParam("timeOutInSecs", qTimeOutInSecs); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
