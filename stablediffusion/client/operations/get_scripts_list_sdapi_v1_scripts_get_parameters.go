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

// NewGetScriptsListSdapiV1ScriptsGetParams creates a new GetScriptsListSdapiV1ScriptsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetScriptsListSdapiV1ScriptsGetParams() *GetScriptsListSdapiV1ScriptsGetParams {
	return &GetScriptsListSdapiV1ScriptsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetScriptsListSdapiV1ScriptsGetParamsWithTimeout creates a new GetScriptsListSdapiV1ScriptsGetParams object
// with the ability to set a timeout on a request.
func NewGetScriptsListSdapiV1ScriptsGetParamsWithTimeout(timeout time.Duration) *GetScriptsListSdapiV1ScriptsGetParams {
	return &GetScriptsListSdapiV1ScriptsGetParams{
		timeout: timeout,
	}
}

// NewGetScriptsListSdapiV1ScriptsGetParamsWithContext creates a new GetScriptsListSdapiV1ScriptsGetParams object
// with the ability to set a context for a request.
func NewGetScriptsListSdapiV1ScriptsGetParamsWithContext(ctx context.Context) *GetScriptsListSdapiV1ScriptsGetParams {
	return &GetScriptsListSdapiV1ScriptsGetParams{
		Context: ctx,
	}
}

// NewGetScriptsListSdapiV1ScriptsGetParamsWithHTTPClient creates a new GetScriptsListSdapiV1ScriptsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetScriptsListSdapiV1ScriptsGetParamsWithHTTPClient(client *http.Client) *GetScriptsListSdapiV1ScriptsGetParams {
	return &GetScriptsListSdapiV1ScriptsGetParams{
		HTTPClient: client,
	}
}

/*
GetScriptsListSdapiV1ScriptsGetParams contains all the parameters to send to the API endpoint

	for the get scripts list sdapi v1 scripts get operation.

	Typically these are written to a http.Request.
*/
type GetScriptsListSdapiV1ScriptsGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get scripts list sdapi v1 scripts get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScriptsListSdapiV1ScriptsGetParams) WithDefaults() *GetScriptsListSdapiV1ScriptsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get scripts list sdapi v1 scripts get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScriptsListSdapiV1ScriptsGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get scripts list sdapi v1 scripts get params
func (o *GetScriptsListSdapiV1ScriptsGetParams) WithTimeout(timeout time.Duration) *GetScriptsListSdapiV1ScriptsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get scripts list sdapi v1 scripts get params
func (o *GetScriptsListSdapiV1ScriptsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get scripts list sdapi v1 scripts get params
func (o *GetScriptsListSdapiV1ScriptsGetParams) WithContext(ctx context.Context) *GetScriptsListSdapiV1ScriptsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get scripts list sdapi v1 scripts get params
func (o *GetScriptsListSdapiV1ScriptsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get scripts list sdapi v1 scripts get params
func (o *GetScriptsListSdapiV1ScriptsGetParams) WithHTTPClient(client *http.Client) *GetScriptsListSdapiV1ScriptsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get scripts list sdapi v1 scripts get params
func (o *GetScriptsListSdapiV1ScriptsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetScriptsListSdapiV1ScriptsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
