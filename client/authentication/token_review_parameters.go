// Code generated by go-swagger; DO NOT EDIT.

package authentication

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

// NewTokenReviewParams creates a new TokenReviewParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTokenReviewParams() *TokenReviewParams {
	return &TokenReviewParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTokenReviewParamsWithTimeout creates a new TokenReviewParams object
// with the ability to set a timeout on a request.
func NewTokenReviewParamsWithTimeout(timeout time.Duration) *TokenReviewParams {
	return &TokenReviewParams{
		timeout: timeout,
	}
}

// NewTokenReviewParamsWithContext creates a new TokenReviewParams object
// with the ability to set a context for a request.
func NewTokenReviewParamsWithContext(ctx context.Context) *TokenReviewParams {
	return &TokenReviewParams{
		Context: ctx,
	}
}

// NewTokenReviewParamsWithHTTPClient creates a new TokenReviewParams object
// with the ability to set a custom HTTPClient for a request.
func NewTokenReviewParamsWithHTTPClient(client *http.Client) *TokenReviewParams {
	return &TokenReviewParams{
		HTTPClient: client,
	}
}

/*
TokenReviewParams contains all the parameters to send to the API endpoint

	for the token review operation.

	Typically these are written to a http.Request.
*/
type TokenReviewParams struct {

	// Body.
	Body *models.OauthTokenReview

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the token review params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TokenReviewParams) WithDefaults() *TokenReviewParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the token review params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TokenReviewParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the token review params
func (o *TokenReviewParams) WithTimeout(timeout time.Duration) *TokenReviewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the token review params
func (o *TokenReviewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the token review params
func (o *TokenReviewParams) WithContext(ctx context.Context) *TokenReviewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the token review params
func (o *TokenReviewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the token review params
func (o *TokenReviewParams) WithHTTPClient(client *http.Client) *TokenReviewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the token review params
func (o *TokenReviewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the token review params
func (o *TokenReviewParams) WithBody(body *models.OauthTokenReview) *TokenReviewParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the token review params
func (o *TokenReviewParams) SetBody(body *models.OauthTokenReview) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *TokenReviewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
