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
)

// NewGetCommitParams creates a new GetCommitParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCommitParams() *GetCommitParams {
	return &GetCommitParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCommitParamsWithTimeout creates a new GetCommitParams object
// with the ability to set a timeout on a request.
func NewGetCommitParamsWithTimeout(timeout time.Duration) *GetCommitParams {
	return &GetCommitParams{
		timeout: timeout,
	}
}

// NewGetCommitParamsWithContext creates a new GetCommitParams object
// with the ability to set a context for a request.
func NewGetCommitParamsWithContext(ctx context.Context) *GetCommitParams {
	return &GetCommitParams{
		Context: ctx,
	}
}

// NewGetCommitParamsWithHTTPClient creates a new GetCommitParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCommitParamsWithHTTPClient(client *http.Client) *GetCommitParams {
	return &GetCommitParams{
		HTTPClient: client,
	}
}

/*
GetCommitParams contains all the parameters to send to the API endpoint

	for the get commit operation.

	Typically these are written to a http.Request.
*/
type GetCommitParams struct {

	/* Commit.

	   The commit hash
	*/
	Commit string

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

// WithDefaults hydrates default values in the get commit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCommitParams) WithDefaults() *GetCommitParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get commit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCommitParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get commit params
func (o *GetCommitParams) WithTimeout(timeout time.Duration) *GetCommitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get commit params
func (o *GetCommitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get commit params
func (o *GetCommitParams) WithContext(ctx context.Context) *GetCommitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get commit params
func (o *GetCommitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get commit params
func (o *GetCommitParams) WithHTTPClient(client *http.Client) *GetCommitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get commit params
func (o *GetCommitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCommit adds the commit to the get commit params
func (o *GetCommitParams) WithCommit(commit string) *GetCommitParams {
	o.SetCommit(commit)
	return o
}

// SetCommit adds the commit to the get commit params
func (o *GetCommitParams) SetCommit(commit string) {
	o.Commit = commit
}

// WithGitrepository adds the gitrepository to the get commit params
func (o *GetCommitParams) WithGitrepository(gitrepository string) *GetCommitParams {
	o.SetGitrepository(gitrepository)
	return o
}

// SetGitrepository adds the gitrepository to the get commit params
func (o *GetCommitParams) SetGitrepository(gitrepository string) {
	o.Gitrepository = gitrepository
}

// WithNamespace adds the namespace to the get commit params
func (o *GetCommitParams) WithNamespace(namespace string) *GetCommitParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get commit params
func (o *GetCommitParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *GetCommitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param commit
	if err := r.SetPathParam("commit", o.Commit); err != nil {
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
