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

// NewThumbnailInfiniteImageBrowsingImageThumbnailGetParams creates a new ThumbnailInfiniteImageBrowsingImageThumbnailGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThumbnailInfiniteImageBrowsingImageThumbnailGetParams() *ThumbnailInfiniteImageBrowsingImageThumbnailGetParams {
	return &ThumbnailInfiniteImageBrowsingImageThumbnailGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewThumbnailInfiniteImageBrowsingImageThumbnailGetParamsWithTimeout creates a new ThumbnailInfiniteImageBrowsingImageThumbnailGetParams object
// with the ability to set a timeout on a request.
func NewThumbnailInfiniteImageBrowsingImageThumbnailGetParamsWithTimeout(timeout time.Duration) *ThumbnailInfiniteImageBrowsingImageThumbnailGetParams {
	return &ThumbnailInfiniteImageBrowsingImageThumbnailGetParams{
		timeout: timeout,
	}
}

// NewThumbnailInfiniteImageBrowsingImageThumbnailGetParamsWithContext creates a new ThumbnailInfiniteImageBrowsingImageThumbnailGetParams object
// with the ability to set a context for a request.
func NewThumbnailInfiniteImageBrowsingImageThumbnailGetParamsWithContext(ctx context.Context) *ThumbnailInfiniteImageBrowsingImageThumbnailGetParams {
	return &ThumbnailInfiniteImageBrowsingImageThumbnailGetParams{
		Context: ctx,
	}
}

// NewThumbnailInfiniteImageBrowsingImageThumbnailGetParamsWithHTTPClient creates a new ThumbnailInfiniteImageBrowsingImageThumbnailGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewThumbnailInfiniteImageBrowsingImageThumbnailGetParamsWithHTTPClient(client *http.Client) *ThumbnailInfiniteImageBrowsingImageThumbnailGetParams {
	return &ThumbnailInfiniteImageBrowsingImageThumbnailGetParams{
		HTTPClient: client,
	}
}

/*
ThumbnailInfiniteImageBrowsingImageThumbnailGetParams contains all the parameters to send to the API endpoint

	for the thumbnail infinite image browsing image thumbnail get operation.

	Typically these are written to a http.Request.
*/
type ThumbnailInfiniteImageBrowsingImageThumbnailGetParams struct {

	// Path.
	Path string

	// Size.
	//
	// Default: "256x256"
	Size *string

	// T.
	T string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thumbnail infinite image browsing image thumbnail get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetParams) WithDefaults() *ThumbnailInfiniteImageBrowsingImageThumbnailGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thumbnail infinite image browsing image thumbnail get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetParams) SetDefaults() {
	var (
		sizeDefault = string("256x256")
	)

	val := ThumbnailInfiniteImageBrowsingImageThumbnailGetParams{
		Size: &sizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the thumbnail infinite image browsing image thumbnail get params
func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetParams) WithTimeout(timeout time.Duration) *ThumbnailInfiniteImageBrowsingImageThumbnailGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thumbnail infinite image browsing image thumbnail get params
func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thumbnail infinite image browsing image thumbnail get params
func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetParams) WithContext(ctx context.Context) *ThumbnailInfiniteImageBrowsingImageThumbnailGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thumbnail infinite image browsing image thumbnail get params
func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thumbnail infinite image browsing image thumbnail get params
func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetParams) WithHTTPClient(client *http.Client) *ThumbnailInfiniteImageBrowsingImageThumbnailGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thumbnail infinite image browsing image thumbnail get params
func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPath adds the path to the thumbnail infinite image browsing image thumbnail get params
func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetParams) WithPath(path string) *ThumbnailInfiniteImageBrowsingImageThumbnailGetParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the thumbnail infinite image browsing image thumbnail get params
func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetParams) SetPath(path string) {
	o.Path = path
}

// WithSize adds the size to the thumbnail infinite image browsing image thumbnail get params
func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetParams) WithSize(size *string) *ThumbnailInfiniteImageBrowsingImageThumbnailGetParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the thumbnail infinite image browsing image thumbnail get params
func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetParams) SetSize(size *string) {
	o.Size = size
}

// WithT adds the t to the thumbnail infinite image browsing image thumbnail get params
func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetParams) WithT(t string) *ThumbnailInfiniteImageBrowsingImageThumbnailGetParams {
	o.SetT(t)
	return o
}

// SetT adds the t to the thumbnail infinite image browsing image thumbnail get params
func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetParams) SetT(t string) {
	o.T = t
}

// WriteToRequest writes these params to a swagger request
func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param path
	qrPath := o.Path
	qPath := qrPath
	if qPath != "" {

		if err := r.SetQueryParam("path", qPath); err != nil {
			return err
		}
	}

	if o.Size != nil {

		// query param size
		var qrSize string

		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := qrSize
		if qSize != "" {

			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}
	}

	// query param t
	qrT := o.T
	qT := qrT
	if qT != "" {

		if err := r.SetQueryParam("t", qT); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
