// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewLambdaInternalProfileStartupGetParams creates a new LambdaInternalProfileStartupGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLambdaInternalProfileStartupGetParams() *LambdaInternalProfileStartupGetParams {
	return &LambdaInternalProfileStartupGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLambdaInternalProfileStartupGetParamsWithTimeout creates a new LambdaInternalProfileStartupGetParams object
// with the ability to set a timeout on a request.
func NewLambdaInternalProfileStartupGetParamsWithTimeout(timeout time.Duration) *LambdaInternalProfileStartupGetParams {
	return &LambdaInternalProfileStartupGetParams{
		timeout: timeout,
	}
}

// NewLambdaInternalProfileStartupGetParamsWithContext creates a new LambdaInternalProfileStartupGetParams object
// with the ability to set a context for a request.
func NewLambdaInternalProfileStartupGetParamsWithContext(ctx context.Context) *LambdaInternalProfileStartupGetParams {
	return &LambdaInternalProfileStartupGetParams{
		Context: ctx,
	}
}

// NewLambdaInternalProfileStartupGetParamsWithHTTPClient creates a new LambdaInternalProfileStartupGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewLambdaInternalProfileStartupGetParamsWithHTTPClient(client *http.Client) *LambdaInternalProfileStartupGetParams {
	return &LambdaInternalProfileStartupGetParams{
		HTTPClient: client,
	}
}

/*
LambdaInternalProfileStartupGetParams contains all the parameters to send to the API endpoint

	for the lambda internal profile startup get operation.

	Typically these are written to a http.Request.
*/
type LambdaInternalProfileStartupGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the lambda internal profile startup get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LambdaInternalProfileStartupGetParams) WithDefaults() *LambdaInternalProfileStartupGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the lambda internal profile startup get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LambdaInternalProfileStartupGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the lambda internal profile startup get params
func (o *LambdaInternalProfileStartupGetParams) WithTimeout(timeout time.Duration) *LambdaInternalProfileStartupGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lambda internal profile startup get params
func (o *LambdaInternalProfileStartupGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lambda internal profile startup get params
func (o *LambdaInternalProfileStartupGetParams) WithContext(ctx context.Context) *LambdaInternalProfileStartupGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lambda internal profile startup get params
func (o *LambdaInternalProfileStartupGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lambda internal profile startup get params
func (o *LambdaInternalProfileStartupGetParams) WithHTTPClient(client *http.Client) *LambdaInternalProfileStartupGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lambda internal profile startup get params
func (o *LambdaInternalProfileStartupGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *LambdaInternalProfileStartupGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
