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

// NewReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetParams creates a new ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetParams() *ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetParams {
	return &ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetParamsWithTimeout creates a new ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetParams object
// with the ability to set a timeout on a request.
func NewReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetParamsWithTimeout(timeout time.Duration) *ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetParams {
	return &ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetParams{
		timeout: timeout,
	}
}

// NewReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetParamsWithContext creates a new ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetParams object
// with the ability to set a context for a request.
func NewReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetParamsWithContext(ctx context.Context) *ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetParams {
	return &ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetParams{
		Context: ctx,
	}
}

// NewReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetParamsWithHTTPClient creates a new ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetParamsWithHTTPClient(client *http.Client) *ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetParams {
	return &ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetParams{
		HTTPClient: client,
	}
}

/*
ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetParams contains all the parameters to send to the API endpoint

	for the read scanned paths infinite image browsing db scanned paths get operation.

	Typically these are written to a http.Request.
*/
type ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the read scanned paths infinite image browsing db scanned paths get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetParams) WithDefaults() *ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the read scanned paths infinite image browsing db scanned paths get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the read scanned paths infinite image browsing db scanned paths get params
func (o *ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetParams) WithTimeout(timeout time.Duration) *ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read scanned paths infinite image browsing db scanned paths get params
func (o *ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read scanned paths infinite image browsing db scanned paths get params
func (o *ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetParams) WithContext(ctx context.Context) *ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read scanned paths infinite image browsing db scanned paths get params
func (o *ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read scanned paths infinite image browsing db scanned paths get params
func (o *ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetParams) WithHTTPClient(client *http.Client) *ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read scanned paths infinite image browsing db scanned paths get params
func (o *ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
