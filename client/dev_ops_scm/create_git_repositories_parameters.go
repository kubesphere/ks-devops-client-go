// Code generated by go-swagger; DO NOT EDIT.

package dev_ops_scm

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

// NewCreateGitRepositoriesParams creates a new CreateGitRepositoriesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateGitRepositoriesParams() *CreateGitRepositoriesParams {
	return &CreateGitRepositoriesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateGitRepositoriesParamsWithTimeout creates a new CreateGitRepositoriesParams object
// with the ability to set a timeout on a request.
func NewCreateGitRepositoriesParamsWithTimeout(timeout time.Duration) *CreateGitRepositoriesParams {
	return &CreateGitRepositoriesParams{
		timeout: timeout,
	}
}

// NewCreateGitRepositoriesParamsWithContext creates a new CreateGitRepositoriesParams object
// with the ability to set a context for a request.
func NewCreateGitRepositoriesParamsWithContext(ctx context.Context) *CreateGitRepositoriesParams {
	return &CreateGitRepositoriesParams{
		Context: ctx,
	}
}

// NewCreateGitRepositoriesParamsWithHTTPClient creates a new CreateGitRepositoriesParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateGitRepositoriesParamsWithHTTPClient(client *http.Client) *CreateGitRepositoriesParams {
	return &CreateGitRepositoriesParams{
		HTTPClient: client,
	}
}

/*
CreateGitRepositoriesParams contains all the parameters to send to the API endpoint

	for the create git repositories operation.

	Typically these are written to a http.Request.
*/
type CreateGitRepositoriesParams struct {

	// Body.
	Body *models.V1alpha3GitRepository

	/* Namespace.

	   The namespace name
	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create git repositories params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateGitRepositoriesParams) WithDefaults() *CreateGitRepositoriesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create git repositories params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateGitRepositoriesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create git repositories params
func (o *CreateGitRepositoriesParams) WithTimeout(timeout time.Duration) *CreateGitRepositoriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create git repositories params
func (o *CreateGitRepositoriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create git repositories params
func (o *CreateGitRepositoriesParams) WithContext(ctx context.Context) *CreateGitRepositoriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create git repositories params
func (o *CreateGitRepositoriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create git repositories params
func (o *CreateGitRepositoriesParams) WithHTTPClient(client *http.Client) *CreateGitRepositoriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create git repositories params
func (o *CreateGitRepositoriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create git repositories params
func (o *CreateGitRepositoriesParams) WithBody(body *models.V1alpha3GitRepository) *CreateGitRepositoriesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create git repositories params
func (o *CreateGitRepositoriesParams) SetBody(body *models.V1alpha3GitRepository) {
	o.Body = body
}

// WithNamespace adds the namespace to the create git repositories params
func (o *CreateGitRepositoriesParams) WithNamespace(namespace string) *CreateGitRepositoriesParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the create git repositories params
func (o *CreateGitRepositoriesParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *CreateGitRepositoriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
