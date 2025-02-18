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
)

// NewGetCredentialParams creates a new GetCredentialParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCredentialParams() *GetCredentialParams {
	return &GetCredentialParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCredentialParamsWithTimeout creates a new GetCredentialParams object
// with the ability to set a timeout on a request.
func NewGetCredentialParamsWithTimeout(timeout time.Duration) *GetCredentialParams {
	return &GetCredentialParams{
		timeout: timeout,
	}
}

// NewGetCredentialParamsWithContext creates a new GetCredentialParams object
// with the ability to set a context for a request.
func NewGetCredentialParamsWithContext(ctx context.Context) *GetCredentialParams {
	return &GetCredentialParams{
		Context: ctx,
	}
}

// NewGetCredentialParamsWithHTTPClient creates a new GetCredentialParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCredentialParamsWithHTTPClient(client *http.Client) *GetCredentialParams {
	return &GetCredentialParams{
		HTTPClient: client,
	}
}

/*
GetCredentialParams contains all the parameters to send to the API endpoint

	for the get credential operation.

	Typically these are written to a http.Request.
*/
type GetCredentialParams struct {

	/* Credential.

	   pipeline name
	*/
	Credential string

	/* Devops.

	   project name
	*/
	Devops string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get credential params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCredentialParams) WithDefaults() *GetCredentialParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get credential params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCredentialParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get credential params
func (o *GetCredentialParams) WithTimeout(timeout time.Duration) *GetCredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get credential params
func (o *GetCredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get credential params
func (o *GetCredentialParams) WithContext(ctx context.Context) *GetCredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get credential params
func (o *GetCredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get credential params
func (o *GetCredentialParams) WithHTTPClient(client *http.Client) *GetCredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get credential params
func (o *GetCredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredential adds the credential to the get credential params
func (o *GetCredentialParams) WithCredential(credential string) *GetCredentialParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the get credential params
func (o *GetCredentialParams) SetCredential(credential string) {
	o.Credential = credential
}

// WithDevops adds the devops to the get credential params
func (o *GetCredentialParams) WithDevops(devops string) *GetCredentialParams {
	o.SetDevops(devops)
	return o
}

// SetDevops adds the devops to the get credential params
func (o *GetCredentialParams) SetDevops(devops string) {
	o.Devops = devops
}

// WriteToRequest writes these params to a swagger request
func (o *GetCredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param credential
	if err := r.SetPathParam("credential", o.Credential); err != nil {
		return err
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
