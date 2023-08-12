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

	"github.com/SpenserCai/sd-webui-go/stablediffusion/models"
)

// NewRemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParams creates a new RemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParams() *RemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParams {
	return &RemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParamsWithTimeout creates a new RemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParams object
// with the ability to set a timeout on a request.
func NewRemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParamsWithTimeout(timeout time.Duration) *RemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParams {
	return &RemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParams{
		timeout: timeout,
	}
}

// NewRemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParamsWithContext creates a new RemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParams object
// with the ability to set a context for a request.
func NewRemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParamsWithContext(ctx context.Context) *RemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParams {
	return &RemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParams{
		Context: ctx,
	}
}

// NewRemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParamsWithHTTPClient creates a new RemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParamsWithHTTPClient(client *http.Client) *RemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParams {
	return &RemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParams{
		HTTPClient: client,
	}
}

/*
RemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParams contains all the parameters to send to the API endpoint

	for the remove custom tag infinite image browsing db remove custom tag post operation.

	Typically these are written to a http.Request.
*/
type RemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParams struct {

	// Body.
	Body *models.RemoveCustomTagReq

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove custom tag infinite image browsing db remove custom tag post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParams) WithDefaults() *RemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove custom tag infinite image browsing db remove custom tag post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the remove custom tag infinite image browsing db remove custom tag post params
func (o *RemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParams) WithTimeout(timeout time.Duration) *RemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove custom tag infinite image browsing db remove custom tag post params
func (o *RemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove custom tag infinite image browsing db remove custom tag post params
func (o *RemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParams) WithContext(ctx context.Context) *RemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove custom tag infinite image browsing db remove custom tag post params
func (o *RemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove custom tag infinite image browsing db remove custom tag post params
func (o *RemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParams) WithHTTPClient(client *http.Client) *RemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove custom tag infinite image browsing db remove custom tag post params
func (o *RemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the remove custom tag infinite image browsing db remove custom tag post params
func (o *RemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParams) WithBody(body *models.RemoveCustomTagReq) *RemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the remove custom tag infinite image browsing db remove custom tag post params
func (o *RemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParams) SetBody(body *models.RemoveCustomTagReq) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveCustomTagInfiniteImageBrowsingDbRemoveCustomTagPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
