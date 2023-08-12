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

// NewStartupEventsStartupEventsGetParams creates a new StartupEventsStartupEventsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStartupEventsStartupEventsGetParams() *StartupEventsStartupEventsGetParams {
	return &StartupEventsStartupEventsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStartupEventsStartupEventsGetParamsWithTimeout creates a new StartupEventsStartupEventsGetParams object
// with the ability to set a timeout on a request.
func NewStartupEventsStartupEventsGetParamsWithTimeout(timeout time.Duration) *StartupEventsStartupEventsGetParams {
	return &StartupEventsStartupEventsGetParams{
		timeout: timeout,
	}
}

// NewStartupEventsStartupEventsGetParamsWithContext creates a new StartupEventsStartupEventsGetParams object
// with the ability to set a context for a request.
func NewStartupEventsStartupEventsGetParamsWithContext(ctx context.Context) *StartupEventsStartupEventsGetParams {
	return &StartupEventsStartupEventsGetParams{
		Context: ctx,
	}
}

// NewStartupEventsStartupEventsGetParamsWithHTTPClient creates a new StartupEventsStartupEventsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewStartupEventsStartupEventsGetParamsWithHTTPClient(client *http.Client) *StartupEventsStartupEventsGetParams {
	return &StartupEventsStartupEventsGetParams{
		HTTPClient: client,
	}
}

/*
StartupEventsStartupEventsGetParams contains all the parameters to send to the API endpoint

	for the startup events startup events get operation.

	Typically these are written to a http.Request.
*/
type StartupEventsStartupEventsGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the startup events startup events get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartupEventsStartupEventsGetParams) WithDefaults() *StartupEventsStartupEventsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the startup events startup events get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartupEventsStartupEventsGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the startup events startup events get params
func (o *StartupEventsStartupEventsGetParams) WithTimeout(timeout time.Duration) *StartupEventsStartupEventsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the startup events startup events get params
func (o *StartupEventsStartupEventsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the startup events startup events get params
func (o *StartupEventsStartupEventsGetParams) WithContext(ctx context.Context) *StartupEventsStartupEventsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the startup events startup events get params
func (o *StartupEventsStartupEventsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the startup events startup events get params
func (o *StartupEventsStartupEventsGetParams) WithHTTPClient(client *http.Client) *StartupEventsStartupEventsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the startup events startup events get params
func (o *StartupEventsStartupEventsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StartupEventsStartupEventsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
