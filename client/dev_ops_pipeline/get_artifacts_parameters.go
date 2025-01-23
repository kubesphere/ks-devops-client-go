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

// NewGetArtifactsParams creates a new GetArtifactsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetArtifactsParams() *GetArtifactsParams {
	return &GetArtifactsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetArtifactsParamsWithTimeout creates a new GetArtifactsParams object
// with the ability to set a timeout on a request.
func NewGetArtifactsParamsWithTimeout(timeout time.Duration) *GetArtifactsParams {
	return &GetArtifactsParams{
		timeout: timeout,
	}
}

// NewGetArtifactsParamsWithContext creates a new GetArtifactsParams object
// with the ability to set a context for a request.
func NewGetArtifactsParamsWithContext(ctx context.Context) *GetArtifactsParams {
	return &GetArtifactsParams{
		Context: ctx,
	}
}

// NewGetArtifactsParamsWithHTTPClient creates a new GetArtifactsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetArtifactsParamsWithHTTPClient(client *http.Client) *GetArtifactsParams {
	return &GetArtifactsParams{
		HTTPClient: client,
	}
}

/*
GetArtifactsParams contains all the parameters to send to the API endpoint

	for the get artifacts operation.

	Typically these are written to a http.Request.
*/
type GetArtifactsParams struct {

	/* Devops.

	   DevOps project's ID, e.g. project-RRRRAzLBlLEm
	*/
	Devops string

	/* Limit.

	   the limit item count of the search.

	   Format: limit=%d
	*/
	Limit *string

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
	*/
	Start *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get artifacts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetArtifactsParams) WithDefaults() *GetArtifactsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get artifacts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetArtifactsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get artifacts params
func (o *GetArtifactsParams) WithTimeout(timeout time.Duration) *GetArtifactsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get artifacts params
func (o *GetArtifactsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get artifacts params
func (o *GetArtifactsParams) WithContext(ctx context.Context) *GetArtifactsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get artifacts params
func (o *GetArtifactsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get artifacts params
func (o *GetArtifactsParams) WithHTTPClient(client *http.Client) *GetArtifactsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get artifacts params
func (o *GetArtifactsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevops adds the devops to the get artifacts params
func (o *GetArtifactsParams) WithDevops(devops string) *GetArtifactsParams {
	o.SetDevops(devops)
	return o
}

// SetDevops adds the devops to the get artifacts params
func (o *GetArtifactsParams) SetDevops(devops string) {
	o.Devops = devops
}

// WithLimit adds the limit to the get artifacts params
func (o *GetArtifactsParams) WithLimit(limit *string) *GetArtifactsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get artifacts params
func (o *GetArtifactsParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithPipeline adds the pipeline to the get artifacts params
func (o *GetArtifactsParams) WithPipeline(pipeline string) *GetArtifactsParams {
	o.SetPipeline(pipeline)
	return o
}

// SetPipeline adds the pipeline to the get artifacts params
func (o *GetArtifactsParams) SetPipeline(pipeline string) {
	o.Pipeline = pipeline
}

// WithRun adds the run to the get artifacts params
func (o *GetArtifactsParams) WithRun(run string) *GetArtifactsParams {
	o.SetRun(run)
	return o
}

// SetRun adds the run to the get artifacts params
func (o *GetArtifactsParams) SetRun(run string) {
	o.Run = run
}

// WithStart adds the start to the get artifacts params
func (o *GetArtifactsParams) WithStart(start *string) *GetArtifactsParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get artifacts params
func (o *GetArtifactsParams) SetStart(start *string) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *GetArtifactsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param devops
	if err := r.SetPathParam("devops", o.Devops); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit string

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := qrLimit
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
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
