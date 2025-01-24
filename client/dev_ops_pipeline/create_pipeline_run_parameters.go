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

// NewCreatePipelineRunParams creates a new CreatePipelineRunParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePipelineRunParams() *CreatePipelineRunParams {
	return &CreatePipelineRunParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePipelineRunParamsWithTimeout creates a new CreatePipelineRunParams object
// with the ability to set a timeout on a request.
func NewCreatePipelineRunParamsWithTimeout(timeout time.Duration) *CreatePipelineRunParams {
	return &CreatePipelineRunParams{
		timeout: timeout,
	}
}

// NewCreatePipelineRunParamsWithContext creates a new CreatePipelineRunParams object
// with the ability to set a context for a request.
func NewCreatePipelineRunParamsWithContext(ctx context.Context) *CreatePipelineRunParams {
	return &CreatePipelineRunParams{
		Context: ctx,
	}
}

// NewCreatePipelineRunParamsWithHTTPClient creates a new CreatePipelineRunParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePipelineRunParamsWithHTTPClient(client *http.Client) *CreatePipelineRunParams {
	return &CreatePipelineRunParams{
		HTTPClient: client,
	}
}

/*
CreatePipelineRunParams contains all the parameters to send to the API endpoint

	for the create pipeline run operation.

	Typically these are written to a http.Request.
*/
type CreatePipelineRunParams struct {

	// Body.
	Body *models.DevopsRunPayload

	/* Branch.

	   The name of SCM reference, only for multi-branch pipeline
	*/
	Branch *string

	/* Namespace.

	   Namespace of the pipeline
	*/
	Namespace string

	/* Pipeline.

	   Name of the pipeline
	*/
	Pipeline string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create pipeline run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePipelineRunParams) WithDefaults() *CreatePipelineRunParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create pipeline run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePipelineRunParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create pipeline run params
func (o *CreatePipelineRunParams) WithTimeout(timeout time.Duration) *CreatePipelineRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create pipeline run params
func (o *CreatePipelineRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create pipeline run params
func (o *CreatePipelineRunParams) WithContext(ctx context.Context) *CreatePipelineRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create pipeline run params
func (o *CreatePipelineRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create pipeline run params
func (o *CreatePipelineRunParams) WithHTTPClient(client *http.Client) *CreatePipelineRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create pipeline run params
func (o *CreatePipelineRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create pipeline run params
func (o *CreatePipelineRunParams) WithBody(body *models.DevopsRunPayload) *CreatePipelineRunParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create pipeline run params
func (o *CreatePipelineRunParams) SetBody(body *models.DevopsRunPayload) {
	o.Body = body
}

// WithBranch adds the branch to the create pipeline run params
func (o *CreatePipelineRunParams) WithBranch(branch *string) *CreatePipelineRunParams {
	o.SetBranch(branch)
	return o
}

// SetBranch adds the branch to the create pipeline run params
func (o *CreatePipelineRunParams) SetBranch(branch *string) {
	o.Branch = branch
}

// WithNamespace adds the namespace to the create pipeline run params
func (o *CreatePipelineRunParams) WithNamespace(namespace string) *CreatePipelineRunParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the create pipeline run params
func (o *CreatePipelineRunParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPipeline adds the pipeline to the create pipeline run params
func (o *CreatePipelineRunParams) WithPipeline(pipeline string) *CreatePipelineRunParams {
	o.SetPipeline(pipeline)
	return o
}

// SetPipeline adds the pipeline to the create pipeline run params
func (o *CreatePipelineRunParams) SetPipeline(pipeline string) {
	o.Pipeline = pipeline
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePipelineRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Branch != nil {

		// query param branch
		var qrBranch string

		if o.Branch != nil {
			qrBranch = *o.Branch
		}
		qBranch := qrBranch
		if qBranch != "" {

			if err := r.SetQueryParam("branch", qBranch); err != nil {
				return err
			}
		}
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
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
