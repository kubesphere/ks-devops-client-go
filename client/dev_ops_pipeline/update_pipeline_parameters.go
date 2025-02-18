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

	"github.com/kubesphere/ks-devops-client-go/models"
)

// NewUpdatePipelineParams creates a new UpdatePipelineParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdatePipelineParams() *UpdatePipelineParams {
	return &UpdatePipelineParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePipelineParamsWithTimeout creates a new UpdatePipelineParams object
// with the ability to set a timeout on a request.
func NewUpdatePipelineParamsWithTimeout(timeout time.Duration) *UpdatePipelineParams {
	return &UpdatePipelineParams{
		timeout: timeout,
	}
}

// NewUpdatePipelineParamsWithContext creates a new UpdatePipelineParams object
// with the ability to set a context for a request.
func NewUpdatePipelineParamsWithContext(ctx context.Context) *UpdatePipelineParams {
	return &UpdatePipelineParams{
		Context: ctx,
	}
}

// NewUpdatePipelineParamsWithHTTPClient creates a new UpdatePipelineParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdatePipelineParamsWithHTTPClient(client *http.Client) *UpdatePipelineParams {
	return &UpdatePipelineParams{
		HTTPClient: client,
	}
}

/*
UpdatePipelineParams contains all the parameters to send to the API endpoint

	for the update pipeline operation.

	Typically these are written to a http.Request.
*/
type UpdatePipelineParams struct {

	// Body.
	Body *models.V1alpha3Pipeline

	/* Devops.

	   project name
	*/
	Devops string

	/* Pipeline.

	   pipeline name
	*/
	Pipeline string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update pipeline params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePipelineParams) WithDefaults() *UpdatePipelineParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update pipeline params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePipelineParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update pipeline params
func (o *UpdatePipelineParams) WithTimeout(timeout time.Duration) *UpdatePipelineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update pipeline params
func (o *UpdatePipelineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update pipeline params
func (o *UpdatePipelineParams) WithContext(ctx context.Context) *UpdatePipelineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update pipeline params
func (o *UpdatePipelineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update pipeline params
func (o *UpdatePipelineParams) WithHTTPClient(client *http.Client) *UpdatePipelineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update pipeline params
func (o *UpdatePipelineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update pipeline params
func (o *UpdatePipelineParams) WithBody(body *models.V1alpha3Pipeline) *UpdatePipelineParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update pipeline params
func (o *UpdatePipelineParams) SetBody(body *models.V1alpha3Pipeline) {
	o.Body = body
}

// WithDevops adds the devops to the update pipeline params
func (o *UpdatePipelineParams) WithDevops(devops string) *UpdatePipelineParams {
	o.SetDevops(devops)
	return o
}

// SetDevops adds the devops to the update pipeline params
func (o *UpdatePipelineParams) SetDevops(devops string) {
	o.Devops = devops
}

// WithPipeline adds the pipeline to the update pipeline params
func (o *UpdatePipelineParams) WithPipeline(pipeline string) *UpdatePipelineParams {
	o.SetPipeline(pipeline)
	return o
}

// SetPipeline adds the pipeline to the update pipeline params
func (o *UpdatePipelineParams) SetPipeline(pipeline string) {
	o.Pipeline = pipeline
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePipelineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param pipeline
	if err := r.SetPathParam("pipeline", o.Pipeline); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
