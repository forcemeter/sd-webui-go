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

// NewFileFilePathOrURLGetParams creates a new FileFilePathOrURLGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFileFilePathOrURLGetParams() *FileFilePathOrURLGetParams {
	return &FileFilePathOrURLGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFileFilePathOrURLGetParamsWithTimeout creates a new FileFilePathOrURLGetParams object
// with the ability to set a timeout on a request.
func NewFileFilePathOrURLGetParamsWithTimeout(timeout time.Duration) *FileFilePathOrURLGetParams {
	return &FileFilePathOrURLGetParams{
		timeout: timeout,
	}
}

// NewFileFilePathOrURLGetParamsWithContext creates a new FileFilePathOrURLGetParams object
// with the ability to set a context for a request.
func NewFileFilePathOrURLGetParamsWithContext(ctx context.Context) *FileFilePathOrURLGetParams {
	return &FileFilePathOrURLGetParams{
		Context: ctx,
	}
}

// NewFileFilePathOrURLGetParamsWithHTTPClient creates a new FileFilePathOrURLGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewFileFilePathOrURLGetParamsWithHTTPClient(client *http.Client) *FileFilePathOrURLGetParams {
	return &FileFilePathOrURLGetParams{
		HTTPClient: client,
	}
}

/*
FileFilePathOrURLGetParams contains all the parameters to send to the API endpoint

	for the file file path or url get operation.

	Typically these are written to a http.Request.
*/
type FileFilePathOrURLGetParams struct {

	// PathOrURL.
	PathOrURL string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the file file path or url get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FileFilePathOrURLGetParams) WithDefaults() *FileFilePathOrURLGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the file file path or url get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FileFilePathOrURLGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the file file path or url get params
func (o *FileFilePathOrURLGetParams) WithTimeout(timeout time.Duration) *FileFilePathOrURLGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the file file path or url get params
func (o *FileFilePathOrURLGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the file file path or url get params
func (o *FileFilePathOrURLGetParams) WithContext(ctx context.Context) *FileFilePathOrURLGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the file file path or url get params
func (o *FileFilePathOrURLGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the file file path or url get params
func (o *FileFilePathOrURLGetParams) WithHTTPClient(client *http.Client) *FileFilePathOrURLGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the file file path or url get params
func (o *FileFilePathOrURLGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPathOrURL adds the pathOrURL to the file file path or url get params
func (o *FileFilePathOrURLGetParams) WithPathOrURL(pathOrURL string) *FileFilePathOrURLGetParams {
	o.SetPathOrURL(pathOrURL)
	return o
}

// SetPathOrURL adds the pathOrUrl to the file file path or url get params
func (o *FileFilePathOrURLGetParams) SetPathOrURL(pathOrURL string) {
	o.PathOrURL = pathOrURL
}

// WriteToRequest writes these params to a swagger request
func (o *FileFilePathOrURLGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param path_or_url
	if err := r.SetPathParam("path_or_url", o.PathOrURL); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
