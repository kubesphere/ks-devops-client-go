// Code generated by go-swagger; DO NOT EDIT.

package dev_ops_credential

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

	"kubesphere.io/devops-client/models"
)

// NewCreateCredentialParams creates a new CreateCredentialParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateCredentialParams() *CreateCredentialParams {
	return &CreateCredentialParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCredentialParamsWithTimeout creates a new CreateCredentialParams object
// with the ability to set a timeout on a request.
func NewCreateCredentialParamsWithTimeout(timeout time.Duration) *CreateCredentialParams {
	return &CreateCredentialParams{
		timeout: timeout,
	}
}

// NewCreateCredentialParamsWithContext creates a new CreateCredentialParams object
// with the ability to set a context for a request.
func NewCreateCredentialParamsWithContext(ctx context.Context) *CreateCredentialParams {
	return &CreateCredentialParams{
		Context: ctx,
	}
}

// NewCreateCredentialParamsWithHTTPClient creates a new CreateCredentialParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateCredentialParamsWithHTTPClient(client *http.Client) *CreateCredentialParams {
	return &CreateCredentialParams{
		HTTPClient: client,
	}
}

/*
CreateCredentialParams contains all the parameters to send to the API endpoint

	for the create credential operation.

	Typically these are written to a http.Request.
*/
type CreateCredentialParams struct {

	// Body.
	Body *models.V1Secret

	/* Devops.

	   devops name
	*/
	Devops string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create credential params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCredentialParams) WithDefaults() *CreateCredentialParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create credential params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCredentialParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create credential params
func (o *CreateCredentialParams) WithTimeout(timeout time.Duration) *CreateCredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create credential params
func (o *CreateCredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create credential params
func (o *CreateCredentialParams) WithContext(ctx context.Context) *CreateCredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create credential params
func (o *CreateCredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create credential params
func (o *CreateCredentialParams) WithHTTPClient(client *http.Client) *CreateCredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create credential params
func (o *CreateCredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create credential params
func (o *CreateCredentialParams) WithBody(body *models.V1Secret) *CreateCredentialParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create credential params
func (o *CreateCredentialParams) SetBody(body *models.V1Secret) {
	o.Body = body
}

// WithDevops adds the devops to the create credential params
func (o *CreateCredentialParams) WithDevops(devops string) *CreateCredentialParams {
	o.SetDevops(devops)
	return o
}

// SetDevops adds the devops to the create credential params
func (o *CreateCredentialParams) SetDevops(devops string) {
	o.Devops = devops
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
