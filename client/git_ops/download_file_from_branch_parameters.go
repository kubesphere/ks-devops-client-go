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

// NewDownloadFileFromBranchParams creates a new DownloadFileFromBranchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDownloadFileFromBranchParams() *DownloadFileFromBranchParams {
	return &DownloadFileFromBranchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDownloadFileFromBranchParamsWithTimeout creates a new DownloadFileFromBranchParams object
// with the ability to set a timeout on a request.
func NewDownloadFileFromBranchParamsWithTimeout(timeout time.Duration) *DownloadFileFromBranchParams {
	return &DownloadFileFromBranchParams{
		timeout: timeout,
	}
}

// NewDownloadFileFromBranchParamsWithContext creates a new DownloadFileFromBranchParams object
// with the ability to set a context for a request.
func NewDownloadFileFromBranchParamsWithContext(ctx context.Context) *DownloadFileFromBranchParams {
	return &DownloadFileFromBranchParams{
		Context: ctx,
	}
}

// NewDownloadFileFromBranchParamsWithHTTPClient creates a new DownloadFileFromBranchParams object
// with the ability to set a custom HTTPClient for a request.
func NewDownloadFileFromBranchParamsWithHTTPClient(client *http.Client) *DownloadFileFromBranchParams {
	return &DownloadFileFromBranchParams{
		HTTPClient: client,
	}
}

/*
DownloadFileFromBranchParams contains all the parameters to send to the API endpoint

	for the download file from branch operation.

	Typically these are written to a http.Request.
*/
type DownloadFileFromBranchParams struct {

	/* Branch.

	   The branch of git repository
	*/
	Branch string

	/* File.

	   base64 encoded file path
	*/
	File string

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

// WithDefaults hydrates default values in the download file from branch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownloadFileFromBranchParams) WithDefaults() *DownloadFileFromBranchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the download file from branch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownloadFileFromBranchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the download file from branch params
func (o *DownloadFileFromBranchParams) WithTimeout(timeout time.Duration) *DownloadFileFromBranchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the download file from branch params
func (o *DownloadFileFromBranchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the download file from branch params
func (o *DownloadFileFromBranchParams) WithContext(ctx context.Context) *DownloadFileFromBranchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the download file from branch params
func (o *DownloadFileFromBranchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the download file from branch params
func (o *DownloadFileFromBranchParams) WithHTTPClient(client *http.Client) *DownloadFileFromBranchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the download file from branch params
func (o *DownloadFileFromBranchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBranch adds the branch to the download file from branch params
func (o *DownloadFileFromBranchParams) WithBranch(branch string) *DownloadFileFromBranchParams {
	o.SetBranch(branch)
	return o
}

// SetBranch adds the branch to the download file from branch params
func (o *DownloadFileFromBranchParams) SetBranch(branch string) {
	o.Branch = branch
}

// WithFile adds the file to the download file from branch params
func (o *DownloadFileFromBranchParams) WithFile(file string) *DownloadFileFromBranchParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the download file from branch params
func (o *DownloadFileFromBranchParams) SetFile(file string) {
	o.File = file
}

// WithGitrepository adds the gitrepository to the download file from branch params
func (o *DownloadFileFromBranchParams) WithGitrepository(gitrepository string) *DownloadFileFromBranchParams {
	o.SetGitrepository(gitrepository)
	return o
}

// SetGitrepository adds the gitrepository to the download file from branch params
func (o *DownloadFileFromBranchParams) SetGitrepository(gitrepository string) {
	o.Gitrepository = gitrepository
}

// WithNamespace adds the namespace to the download file from branch params
func (o *DownloadFileFromBranchParams) WithNamespace(namespace string) *DownloadFileFromBranchParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the download file from branch params
func (o *DownloadFileFromBranchParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *DownloadFileFromBranchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param branch
	if err := r.SetPathParam("branch", o.Branch); err != nil {
		return err
	}

	// path param file
	if err := r.SetPathParam("file", o.File); err != nil {
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
