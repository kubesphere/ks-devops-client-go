// Code generated by go-swagger; DO NOT EDIT.

package dev_ops_template

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

// NewHandleGetTemplateParams creates a new HandleGetTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHandleGetTemplateParams() *HandleGetTemplateParams {
	return &HandleGetTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHandleGetTemplateParamsWithTimeout creates a new HandleGetTemplateParams object
// with the ability to set a timeout on a request.
func NewHandleGetTemplateParamsWithTimeout(timeout time.Duration) *HandleGetTemplateParams {
	return &HandleGetTemplateParams{
		timeout: timeout,
	}
}

// NewHandleGetTemplateParamsWithContext creates a new HandleGetTemplateParams object
// with the ability to set a context for a request.
func NewHandleGetTemplateParamsWithContext(ctx context.Context) *HandleGetTemplateParams {
	return &HandleGetTemplateParams{
		Context: ctx,
	}
}

// NewHandleGetTemplateParamsWithHTTPClient creates a new HandleGetTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewHandleGetTemplateParamsWithHTTPClient(client *http.Client) *HandleGetTemplateParams {
	return &HandleGetTemplateParams{
		HTTPClient: client,
	}
}

/*
HandleGetTemplateParams contains all the parameters to send to the API endpoint

	for the handle get template operation.

	Typically these are written to a http.Request.
*/
type HandleGetTemplateParams struct {

	/* Devops.

	   DevOps project name
	*/
	Devops string

	/* Template.

	   Template name
	*/
	Template string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the handle get template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HandleGetTemplateParams) WithDefaults() *HandleGetTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the handle get template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HandleGetTemplateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the handle get template params
func (o *HandleGetTemplateParams) WithTimeout(timeout time.Duration) *HandleGetTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the handle get template params
func (o *HandleGetTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the handle get template params
func (o *HandleGetTemplateParams) WithContext(ctx context.Context) *HandleGetTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the handle get template params
func (o *HandleGetTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the handle get template params
func (o *HandleGetTemplateParams) WithHTTPClient(client *http.Client) *HandleGetTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the handle get template params
func (o *HandleGetTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevops adds the devops to the handle get template params
func (o *HandleGetTemplateParams) WithDevops(devops string) *HandleGetTemplateParams {
	o.SetDevops(devops)
	return o
}

// SetDevops adds the devops to the handle get template params
func (o *HandleGetTemplateParams) SetDevops(devops string) {
	o.Devops = devops
}

// WithTemplate adds the template to the handle get template params
func (o *HandleGetTemplateParams) WithTemplate(template string) *HandleGetTemplateParams {
	o.SetTemplate(template)
	return o
}

// SetTemplate adds the template to the handle get template params
func (o *HandleGetTemplateParams) SetTemplate(template string) {
	o.Template = template
}

// WriteToRequest writes these params to a swagger request
func (o *HandleGetTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param devops
	if err := r.SetPathParam("devops", o.Devops); err != nil {
		return err
	}

	// path param template
	if err := r.SetPathParam("template", o.Template); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
