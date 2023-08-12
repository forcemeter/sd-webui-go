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

// NewGetQueueStatusQueueStatusGetParams creates a new GetQueueStatusQueueStatusGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetQueueStatusQueueStatusGetParams() *GetQueueStatusQueueStatusGetParams {
	return &GetQueueStatusQueueStatusGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetQueueStatusQueueStatusGetParamsWithTimeout creates a new GetQueueStatusQueueStatusGetParams object
// with the ability to set a timeout on a request.
func NewGetQueueStatusQueueStatusGetParamsWithTimeout(timeout time.Duration) *GetQueueStatusQueueStatusGetParams {
	return &GetQueueStatusQueueStatusGetParams{
		timeout: timeout,
	}
}

// NewGetQueueStatusQueueStatusGetParamsWithContext creates a new GetQueueStatusQueueStatusGetParams object
// with the ability to set a context for a request.
func NewGetQueueStatusQueueStatusGetParamsWithContext(ctx context.Context) *GetQueueStatusQueueStatusGetParams {
	return &GetQueueStatusQueueStatusGetParams{
		Context: ctx,
	}
}

// NewGetQueueStatusQueueStatusGetParamsWithHTTPClient creates a new GetQueueStatusQueueStatusGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetQueueStatusQueueStatusGetParamsWithHTTPClient(client *http.Client) *GetQueueStatusQueueStatusGetParams {
	return &GetQueueStatusQueueStatusGetParams{
		HTTPClient: client,
	}
}

/*
GetQueueStatusQueueStatusGetParams contains all the parameters to send to the API endpoint

	for the get queue status queue status get operation.

	Typically these are written to a http.Request.
*/
type GetQueueStatusQueueStatusGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get queue status queue status get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetQueueStatusQueueStatusGetParams) WithDefaults() *GetQueueStatusQueueStatusGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get queue status queue status get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetQueueStatusQueueStatusGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get queue status queue status get params
func (o *GetQueueStatusQueueStatusGetParams) WithTimeout(timeout time.Duration) *GetQueueStatusQueueStatusGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get queue status queue status get params
func (o *GetQueueStatusQueueStatusGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get queue status queue status get params
func (o *GetQueueStatusQueueStatusGetParams) WithContext(ctx context.Context) *GetQueueStatusQueueStatusGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get queue status queue status get params
func (o *GetQueueStatusQueueStatusGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get queue status queue status get params
func (o *GetQueueStatusQueueStatusGetParams) WithHTTPClient(client *http.Client) *GetQueueStatusQueueStatusGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get queue status queue status get params
func (o *GetQueueStatusQueueStatusGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetQueueStatusQueueStatusGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
