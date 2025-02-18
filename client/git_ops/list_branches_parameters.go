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

// NewListBranchesParams creates a new ListBranchesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListBranchesParams() *ListBranchesParams {
	return &ListBranchesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListBranchesParamsWithTimeout creates a new ListBranchesParams object
// with the ability to set a timeout on a request.
func NewListBranchesParamsWithTimeout(timeout time.Duration) *ListBranchesParams {
	return &ListBranchesParams{
		timeout: timeout,
	}
}

// NewListBranchesParamsWithContext creates a new ListBranchesParams object
// with the ability to set a context for a request.
func NewListBranchesParamsWithContext(ctx context.Context) *ListBranchesParams {
	return &ListBranchesParams{
		Context: ctx,
	}
}

// NewListBranchesParamsWithHTTPClient creates a new ListBranchesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListBranchesParamsWithHTTPClient(client *http.Client) *ListBranchesParams {
	return &ListBranchesParams{
		HTTPClient: client,
	}
}

/*
ListBranchesParams contains all the parameters to send to the API endpoint

	for the list branches operation.

	Typically these are written to a http.Request.
*/
type ListBranchesParams struct {

	/* Gitrepository.

	   The GitRepository customs resource
	*/
	Gitrepository string

	/* Limit.

	   limit
	*/
	Limit *string

	/* Namespace.

	   The namespace name
	*/
	Namespace string

	/* Page.

	   page

	   Format: page=%d
	   Default: "page=1"
	*/
	Page *string

	/* WithHead.

	   whether get the head reference
	*/
	WithHead *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list branches params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListBranchesParams) WithDefaults() *ListBranchesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list branches params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListBranchesParams) SetDefaults() {
	var (
		pageDefault = string("page=1")
	)

	val := ListBranchesParams{
		Page: &pageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list branches params
func (o *ListBranchesParams) WithTimeout(timeout time.Duration) *ListBranchesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list branches params
func (o *ListBranchesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list branches params
func (o *ListBranchesParams) WithContext(ctx context.Context) *ListBranchesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list branches params
func (o *ListBranchesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list branches params
func (o *ListBranchesParams) WithHTTPClient(client *http.Client) *ListBranchesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list branches params
func (o *ListBranchesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGitrepository adds the gitrepository to the list branches params
func (o *ListBranchesParams) WithGitrepository(gitrepository string) *ListBranchesParams {
	o.SetGitrepository(gitrepository)
	return o
}

// SetGitrepository adds the gitrepository to the list branches params
func (o *ListBranchesParams) SetGitrepository(gitrepository string) {
	o.Gitrepository = gitrepository
}

// WithLimit adds the limit to the list branches params
func (o *ListBranchesParams) WithLimit(limit *string) *ListBranchesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list branches params
func (o *ListBranchesParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithNamespace adds the namespace to the list branches params
func (o *ListBranchesParams) WithNamespace(namespace string) *ListBranchesParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the list branches params
func (o *ListBranchesParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPage adds the page to the list branches params
func (o *ListBranchesParams) WithPage(page *string) *ListBranchesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list branches params
func (o *ListBranchesParams) SetPage(page *string) {
	o.Page = page
}

// WithWithHead adds the withHead to the list branches params
func (o *ListBranchesParams) WithWithHead(withHead *bool) *ListBranchesParams {
	o.SetWithHead(withHead)
	return o
}

// SetWithHead adds the withHead to the list branches params
func (o *ListBranchesParams) SetWithHead(withHead *bool) {
	o.WithHead = withHead
}

// WriteToRequest writes these params to a swagger request
func (o *ListBranchesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param gitrepository
	if err := r.SetPathParam("gitrepository", o.Gitrepository); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit string

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := qrLimit
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if o.Page != nil {

		// query param page
		var qrPage string

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := qrPage
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.WithHead != nil {

		// query param withHead
		var qrWithHead bool

		if o.WithHead != nil {
			qrWithHead = *o.WithHead
		}
		qWithHead := swag.FormatBool(qrWithHead)
		if qWithHead != "" {

			if err := r.SetQueryParam("withHead", qWithHead); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
