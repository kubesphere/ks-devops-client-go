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

// NewDownloadFileFromCommitParams creates a new DownloadFileFromCommitParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDownloadFileFromCommitParams() *DownloadFileFromCommitParams {
	return &DownloadFileFromCommitParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDownloadFileFromCommitParamsWithTimeout creates a new DownloadFileFromCommitParams object
// with the ability to set a timeout on a request.
func NewDownloadFileFromCommitParamsWithTimeout(timeout time.Duration) *DownloadFileFromCommitParams {
	return &DownloadFileFromCommitParams{
		timeout: timeout,
	}
}

// NewDownloadFileFromCommitParamsWithContext creates a new DownloadFileFromCommitParams object
// with the ability to set a context for a request.
func NewDownloadFileFromCommitParamsWithContext(ctx context.Context) *DownloadFileFromCommitParams {
	return &DownloadFileFromCommitParams{
		Context: ctx,
	}
}

// NewDownloadFileFromCommitParamsWithHTTPClient creates a new DownloadFileFromCommitParams object
// with the ability to set a custom HTTPClient for a request.
func NewDownloadFileFromCommitParamsWithHTTPClient(client *http.Client) *DownloadFileFromCommitParams {
	return &DownloadFileFromCommitParams{
		HTTPClient: client,
	}
}

/*
DownloadFileFromCommitParams contains all the parameters to send to the API endpoint

	for the download file from commit operation.

	Typically these are written to a http.Request.
*/
type DownloadFileFromCommitParams struct {

	/* Commit.

	   The commit hash
	*/
	Commit string

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

// WithDefaults hydrates default values in the download file from commit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownloadFileFromCommitParams) WithDefaults() *DownloadFileFromCommitParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the download file from commit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownloadFileFromCommitParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the download file from commit params
func (o *DownloadFileFromCommitParams) WithTimeout(timeout time.Duration) *DownloadFileFromCommitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the download file from commit params
func (o *DownloadFileFromCommitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the download file from commit params
func (o *DownloadFileFromCommitParams) WithContext(ctx context.Context) *DownloadFileFromCommitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the download file from commit params
func (o *DownloadFileFromCommitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the download file from commit params
func (o *DownloadFileFromCommitParams) WithHTTPClient(client *http.Client) *DownloadFileFromCommitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the download file from commit params
func (o *DownloadFileFromCommitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCommit adds the commit to the download file from commit params
func (o *DownloadFileFromCommitParams) WithCommit(commit string) *DownloadFileFromCommitParams {
	o.SetCommit(commit)
	return o
}

// SetCommit adds the commit to the download file from commit params
func (o *DownloadFileFromCommitParams) SetCommit(commit string) {
	o.Commit = commit
}

// WithFile adds the file to the download file from commit params
func (o *DownloadFileFromCommitParams) WithFile(file string) *DownloadFileFromCommitParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the download file from commit params
func (o *DownloadFileFromCommitParams) SetFile(file string) {
	o.File = file
}

// WithGitrepository adds the gitrepository to the download file from commit params
func (o *DownloadFileFromCommitParams) WithGitrepository(gitrepository string) *DownloadFileFromCommitParams {
	o.SetGitrepository(gitrepository)
	return o
}

// SetGitrepository adds the gitrepository to the download file from commit params
func (o *DownloadFileFromCommitParams) SetGitrepository(gitrepository string) {
	o.Gitrepository = gitrepository
}

// WithNamespace adds the namespace to the download file from commit params
func (o *DownloadFileFromCommitParams) WithNamespace(namespace string) *DownloadFileFromCommitParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the download file from commit params
func (o *DownloadFileFromCommitParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *DownloadFileFromCommitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param commit
	if err := r.SetPathParam("commit", o.Commit); err != nil {
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
