// Code generated by go-swagger; DO NOT EDIT.

package git_ops

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

// NewAddFilesParams creates a new AddFilesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddFilesParams() *AddFilesParams {
	return &AddFilesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddFilesParamsWithTimeout creates a new AddFilesParams object
// with the ability to set a timeout on a request.
func NewAddFilesParamsWithTimeout(timeout time.Duration) *AddFilesParams {
	return &AddFilesParams{
		timeout: timeout,
	}
}

// NewAddFilesParamsWithContext creates a new AddFilesParams object
// with the ability to set a context for a request.
func NewAddFilesParamsWithContext(ctx context.Context) *AddFilesParams {
	return &AddFilesParams{
		Context: ctx,
	}
}

// NewAddFilesParamsWithHTTPClient creates a new AddFilesParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddFilesParamsWithHTTPClient(client *http.Client) *AddFilesParams {
	return &AddFilesParams{
		HTTPClient: client,
	}
}

/*
AddFilesParams contains all the parameters to send to the API endpoint

	for the add files operation.

	Typically these are written to a http.Request.
*/
type AddFilesParams struct {

	// Body.
	Body *models.GitopsAddFilesInput

	/* Branch.

	   The branch of git repository
	*/
	Branch string

	/* Gitrepository.

	   The GitRepository customs resource
	*/
	Gitrepository string

	/* Namespace.

	   The namespace name
	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add files params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddFilesParams) WithDefaults() *AddFilesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add files params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddFilesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add files params
func (o *AddFilesParams) WithTimeout(timeout time.Duration) *AddFilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add files params
func (o *AddFilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add files params
func (o *AddFilesParams) WithContext(ctx context.Context) *AddFilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add files params
func (o *AddFilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add files params
func (o *AddFilesParams) WithHTTPClient(client *http.Client) *AddFilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add files params
func (o *AddFilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add files params
func (o *AddFilesParams) WithBody(body *models.GitopsAddFilesInput) *AddFilesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add files params
func (o *AddFilesParams) SetBody(body *models.GitopsAddFilesInput) {
	o.Body = body
}

// WithBranch adds the branch to the add files params
func (o *AddFilesParams) WithBranch(branch string) *AddFilesParams {
	o.SetBranch(branch)
	return o
}

// SetBranch adds the branch to the add files params
func (o *AddFilesParams) SetBranch(branch string) {
	o.Branch = branch
}

// WithGitrepository adds the gitrepository to the add files params
func (o *AddFilesParams) WithGitrepository(gitrepository string) *AddFilesParams {
	o.SetGitrepository(gitrepository)
	return o
}

// SetGitrepository adds the gitrepository to the add files params
func (o *AddFilesParams) SetGitrepository(gitrepository string) {
	o.Gitrepository = gitrepository
}

// WithNamespace adds the namespace to the add files params
func (o *AddFilesParams) WithNamespace(namespace string) *AddFilesParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the add files params
func (o *AddFilesParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *AddFilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param branch
	if err := r.SetPathParam("branch", o.Branch); err != nil {
		return err
	}

	// path param gitrepository
	if err := r.SetPathParam("gitrepository", o.Gitrepository); err != nil {
		return err
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
