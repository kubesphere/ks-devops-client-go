// Code generated by go-swagger; DO NOT EDIT.

package dev_ops_jenkins

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

// NewUpdateJenkinsfileParams creates a new UpdateJenkinsfileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateJenkinsfileParams() *UpdateJenkinsfileParams {
	return &UpdateJenkinsfileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateJenkinsfileParamsWithTimeout creates a new UpdateJenkinsfileParams object
// with the ability to set a timeout on a request.
func NewUpdateJenkinsfileParamsWithTimeout(timeout time.Duration) *UpdateJenkinsfileParams {
	return &UpdateJenkinsfileParams{
		timeout: timeout,
	}
}

// NewUpdateJenkinsfileParamsWithContext creates a new UpdateJenkinsfileParams object
// with the ability to set a context for a request.
func NewUpdateJenkinsfileParamsWithContext(ctx context.Context) *UpdateJenkinsfileParams {
	return &UpdateJenkinsfileParams{
		Context: ctx,
	}
}

// NewUpdateJenkinsfileParamsWithHTTPClient creates a new UpdateJenkinsfileParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateJenkinsfileParamsWithHTTPClient(client *http.Client) *UpdateJenkinsfileParams {
	return &UpdateJenkinsfileParams{
		HTTPClient: client,
	}
}

/*
UpdateJenkinsfileParams contains all the parameters to send to the API endpoint

	for the update jenkinsfile operation.

	Typically these are written to a http.Request.
*/
type UpdateJenkinsfileParams struct {

	/* Body.

	   The Jenkinsfile content should be in the 'data' field
	*/
	Body *models.V1alpha3GenericPayload

	/* Devops.

	   project name
	*/
	Devops string

	/* Mode.

	   the mode(json or raw) that you expect to update the Jenkinsfile
	*/
	Mode *string

	/* Pipeline.

	   pipeline name
	*/
	Pipeline string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update jenkinsfile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateJenkinsfileParams) WithDefaults() *UpdateJenkinsfileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update jenkinsfile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateJenkinsfileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update jenkinsfile params
func (o *UpdateJenkinsfileParams) WithTimeout(timeout time.Duration) *UpdateJenkinsfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update jenkinsfile params
func (o *UpdateJenkinsfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update jenkinsfile params
func (o *UpdateJenkinsfileParams) WithContext(ctx context.Context) *UpdateJenkinsfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update jenkinsfile params
func (o *UpdateJenkinsfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update jenkinsfile params
func (o *UpdateJenkinsfileParams) WithHTTPClient(client *http.Client) *UpdateJenkinsfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update jenkinsfile params
func (o *UpdateJenkinsfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update jenkinsfile params
func (o *UpdateJenkinsfileParams) WithBody(body *models.V1alpha3GenericPayload) *UpdateJenkinsfileParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update jenkinsfile params
func (o *UpdateJenkinsfileParams) SetBody(body *models.V1alpha3GenericPayload) {
	o.Body = body
}

// WithDevops adds the devops to the update jenkinsfile params
func (o *UpdateJenkinsfileParams) WithDevops(devops string) *UpdateJenkinsfileParams {
	o.SetDevops(devops)
	return o
}

// SetDevops adds the devops to the update jenkinsfile params
func (o *UpdateJenkinsfileParams) SetDevops(devops string) {
	o.Devops = devops
}

// WithMode adds the mode to the update jenkinsfile params
func (o *UpdateJenkinsfileParams) WithMode(mode *string) *UpdateJenkinsfileParams {
	o.SetMode(mode)
	return o
}

// SetMode adds the mode to the update jenkinsfile params
func (o *UpdateJenkinsfileParams) SetMode(mode *string) {
	o.Mode = mode
}

// WithPipeline adds the pipeline to the update jenkinsfile params
func (o *UpdateJenkinsfileParams) WithPipeline(pipeline string) *UpdateJenkinsfileParams {
	o.SetPipeline(pipeline)
	return o
}

// SetPipeline adds the pipeline to the update jenkinsfile params
func (o *UpdateJenkinsfileParams) SetPipeline(pipeline string) {
	o.Pipeline = pipeline
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateJenkinsfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Mode != nil {

		// query param mode
		var qrMode string

		if o.Mode != nil {
			qrMode = *o.Mode
		}
		qMode := qrMode
		if qMode != "" {

			if err := r.SetQueryParam("mode", qMode); err != nil {
				return err
			}
		}
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
