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

// NewCreatePipelineParams creates a new CreatePipelineParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePipelineParams() *CreatePipelineParams {
	return &CreatePipelineParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePipelineParamsWithTimeout creates a new CreatePipelineParams object
// with the ability to set a timeout on a request.
func NewCreatePipelineParamsWithTimeout(timeout time.Duration) *CreatePipelineParams {
	return &CreatePipelineParams{
		timeout: timeout,
	}
}

// NewCreatePipelineParamsWithContext creates a new CreatePipelineParams object
// with the ability to set a context for a request.
func NewCreatePipelineParamsWithContext(ctx context.Context) *CreatePipelineParams {
	return &CreatePipelineParams{
		Context: ctx,
	}
}

// NewCreatePipelineParamsWithHTTPClient creates a new CreatePipelineParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePipelineParamsWithHTTPClient(client *http.Client) *CreatePipelineParams {
	return &CreatePipelineParams{
		HTTPClient: client,
	}
}

/*
CreatePipelineParams contains all the parameters to send to the API endpoint

	for the create pipeline operation.

	Typically these are written to a http.Request.
*/
type CreatePipelineParams struct {

	// Body.
	Body *models.V1alpha3Pipeline

	/* Devops.

	   devops name
	*/
	Devops string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create pipeline params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePipelineParams) WithDefaults() *CreatePipelineParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create pipeline params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePipelineParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create pipeline params
func (o *CreatePipelineParams) WithTimeout(timeout time.Duration) *CreatePipelineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create pipeline params
func (o *CreatePipelineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create pipeline params
func (o *CreatePipelineParams) WithContext(ctx context.Context) *CreatePipelineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create pipeline params
func (o *CreatePipelineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create pipeline params
func (o *CreatePipelineParams) WithHTTPClient(client *http.Client) *CreatePipelineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create pipeline params
func (o *CreatePipelineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create pipeline params
func (o *CreatePipelineParams) WithBody(body *models.V1alpha3Pipeline) *CreatePipelineParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create pipeline params
func (o *CreatePipelineParams) SetBody(body *models.V1alpha3Pipeline) {
	o.Body = body
}

// WithDevops adds the devops to the create pipeline params
func (o *CreatePipelineParams) WithDevops(devops string) *CreatePipelineParams {
	o.SetDevops(devops)
	return o
}

// SetDevops adds the devops to the create pipeline params
func (o *CreatePipelineParams) SetDevops(devops string) {
	o.Devops = devops
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePipelineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
