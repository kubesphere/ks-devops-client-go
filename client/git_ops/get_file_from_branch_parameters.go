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
	"github.com/go-openapi/swag"
)

// NewGetFileFromBranchParams creates a new GetFileFromBranchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFileFromBranchParams() *GetFileFromBranchParams {
	return &GetFileFromBranchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFileFromBranchParamsWithTimeout creates a new GetFileFromBranchParams object
// with the ability to set a timeout on a request.
func NewGetFileFromBranchParamsWithTimeout(timeout time.Duration) *GetFileFromBranchParams {
	return &GetFileFromBranchParams{
		timeout: timeout,
	}
}

// NewGetFileFromBranchParamsWithContext creates a new GetFileFromBranchParams object
// with the ability to set a context for a request.
func NewGetFileFromBranchParamsWithContext(ctx context.Context) *GetFileFromBranchParams {
	return &GetFileFromBranchParams{
		Context: ctx,
	}
}

// NewGetFileFromBranchParamsWithHTTPClient creates a new GetFileFromBranchParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFileFromBranchParamsWithHTTPClient(client *http.Client) *GetFileFromBranchParams {
	return &GetFileFromBranchParams{
		HTTPClient: client,
	}
}

/*
GetFileFromBranchParams contains all the parameters to send to the API endpoint

	for the get file from branch operation.

	Typically these are written to a http.Request.
*/
type GetFileFromBranchParams struct {

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

	/* WithContent.

	   whether get the base64 encoded content of the file
	*/
	WithContent *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get file from branch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFileFromBranchParams) WithDefaults() *GetFileFromBranchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get file from branch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFileFromBranchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get file from branch params
func (o *GetFileFromBranchParams) WithTimeout(timeout time.Duration) *GetFileFromBranchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get file from branch params
func (o *GetFileFromBranchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get file from branch params
func (o *GetFileFromBranchParams) WithContext(ctx context.Context) *GetFileFromBranchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get file from branch params
func (o *GetFileFromBranchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get file from branch params
func (o *GetFileFromBranchParams) WithHTTPClient(client *http.Client) *GetFileFromBranchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get file from branch params
func (o *GetFileFromBranchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBranch adds the branch to the get file from branch params
func (o *GetFileFromBranchParams) WithBranch(branch string) *GetFileFromBranchParams {
	o.SetBranch(branch)
	return o
}

// SetBranch adds the branch to the get file from branch params
func (o *GetFileFromBranchParams) SetBranch(branch string) {
	o.Branch = branch
}

// WithFile adds the file to the get file from branch params
func (o *GetFileFromBranchParams) WithFile(file string) *GetFileFromBranchParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the get file from branch params
func (o *GetFileFromBranchParams) SetFile(file string) {
	o.File = file
}

// WithGitrepository adds the gitrepository to the get file from branch params
func (o *GetFileFromBranchParams) WithGitrepository(gitrepository string) *GetFileFromBranchParams {
	o.SetGitrepository(gitrepository)
	return o
}

// SetGitrepository adds the gitrepository to the get file from branch params
func (o *GetFileFromBranchParams) SetGitrepository(gitrepository string) {
	o.Gitrepository = gitrepository
}

// WithNamespace adds the namespace to the get file from branch params
func (o *GetFileFromBranchParams) WithNamespace(namespace string) *GetFileFromBranchParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get file from branch params
func (o *GetFileFromBranchParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithWithContent adds the withContent to the get file from branch params
func (o *GetFileFromBranchParams) WithWithContent(withContent *bool) *GetFileFromBranchParams {
	o.SetWithContent(withContent)
	return o
}

// SetWithContent adds the withContent to the get file from branch params
func (o *GetFileFromBranchParams) SetWithContent(withContent *bool) {
	o.WithContent = withContent
}

// WriteToRequest writes these params to a swagger request
func (o *GetFileFromBranchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.WithContent != nil {

		// query param withContent
		var qrWithContent bool

		if o.WithContent != nil {
			qrWithContent = *o.WithContent
		}
		qWithContent := swag.FormatBool(qrWithContent)
		if qWithContent != "" {

			if err := r.SetQueryParam("withContent", qWithContent); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
