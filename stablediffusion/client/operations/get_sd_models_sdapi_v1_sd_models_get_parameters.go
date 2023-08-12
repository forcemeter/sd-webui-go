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

// NewGetSdModelsSdapiV1SdModelsGetParams creates a new GetSdModelsSdapiV1SdModelsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSdModelsSdapiV1SdModelsGetParams() *GetSdModelsSdapiV1SdModelsGetParams {
	return &GetSdModelsSdapiV1SdModelsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSdModelsSdapiV1SdModelsGetParamsWithTimeout creates a new GetSdModelsSdapiV1SdModelsGetParams object
// with the ability to set a timeout on a request.
func NewGetSdModelsSdapiV1SdModelsGetParamsWithTimeout(timeout time.Duration) *GetSdModelsSdapiV1SdModelsGetParams {
	return &GetSdModelsSdapiV1SdModelsGetParams{
		timeout: timeout,
	}
}

// NewGetSdModelsSdapiV1SdModelsGetParamsWithContext creates a new GetSdModelsSdapiV1SdModelsGetParams object
// with the ability to set a context for a request.
func NewGetSdModelsSdapiV1SdModelsGetParamsWithContext(ctx context.Context) *GetSdModelsSdapiV1SdModelsGetParams {
	return &GetSdModelsSdapiV1SdModelsGetParams{
		Context: ctx,
	}
}

// NewGetSdModelsSdapiV1SdModelsGetParamsWithHTTPClient creates a new GetSdModelsSdapiV1SdModelsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSdModelsSdapiV1SdModelsGetParamsWithHTTPClient(client *http.Client) *GetSdModelsSdapiV1SdModelsGetParams {
	return &GetSdModelsSdapiV1SdModelsGetParams{
		HTTPClient: client,
	}
}

/*
GetSdModelsSdapiV1SdModelsGetParams contains all the parameters to send to the API endpoint

	for the get sd models sdapi v1 sd models get operation.

	Typically these are written to a http.Request.
*/
type GetSdModelsSdapiV1SdModelsGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get sd models sdapi v1 sd models get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSdModelsSdapiV1SdModelsGetParams) WithDefaults() *GetSdModelsSdapiV1SdModelsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get sd models sdapi v1 sd models get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSdModelsSdapiV1SdModelsGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get sd models sdapi v1 sd models get params
func (o *GetSdModelsSdapiV1SdModelsGetParams) WithTimeout(timeout time.Duration) *GetSdModelsSdapiV1SdModelsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sd models sdapi v1 sd models get params
func (o *GetSdModelsSdapiV1SdModelsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sd models sdapi v1 sd models get params
func (o *GetSdModelsSdapiV1SdModelsGetParams) WithContext(ctx context.Context) *GetSdModelsSdapiV1SdModelsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sd models sdapi v1 sd models get params
func (o *GetSdModelsSdapiV1SdModelsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sd models sdapi v1 sd models get params
func (o *GetSdModelsSdapiV1SdModelsGetParams) WithHTTPClient(client *http.Client) *GetSdModelsSdapiV1SdModelsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sd models sdapi v1 sd models get params
func (o *GetSdModelsSdapiV1SdModelsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetSdModelsSdapiV1SdModelsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
