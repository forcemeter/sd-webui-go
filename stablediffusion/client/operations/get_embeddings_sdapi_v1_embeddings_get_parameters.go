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

// NewGetEmbeddingsSdapiV1EmbeddingsGetParams creates a new GetEmbeddingsSdapiV1EmbeddingsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEmbeddingsSdapiV1EmbeddingsGetParams() *GetEmbeddingsSdapiV1EmbeddingsGetParams {
	return &GetEmbeddingsSdapiV1EmbeddingsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEmbeddingsSdapiV1EmbeddingsGetParamsWithTimeout creates a new GetEmbeddingsSdapiV1EmbeddingsGetParams object
// with the ability to set a timeout on a request.
func NewGetEmbeddingsSdapiV1EmbeddingsGetParamsWithTimeout(timeout time.Duration) *GetEmbeddingsSdapiV1EmbeddingsGetParams {
	return &GetEmbeddingsSdapiV1EmbeddingsGetParams{
		timeout: timeout,
	}
}

// NewGetEmbeddingsSdapiV1EmbeddingsGetParamsWithContext creates a new GetEmbeddingsSdapiV1EmbeddingsGetParams object
// with the ability to set a context for a request.
func NewGetEmbeddingsSdapiV1EmbeddingsGetParamsWithContext(ctx context.Context) *GetEmbeddingsSdapiV1EmbeddingsGetParams {
	return &GetEmbeddingsSdapiV1EmbeddingsGetParams{
		Context: ctx,
	}
}

// NewGetEmbeddingsSdapiV1EmbeddingsGetParamsWithHTTPClient creates a new GetEmbeddingsSdapiV1EmbeddingsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEmbeddingsSdapiV1EmbeddingsGetParamsWithHTTPClient(client *http.Client) *GetEmbeddingsSdapiV1EmbeddingsGetParams {
	return &GetEmbeddingsSdapiV1EmbeddingsGetParams{
		HTTPClient: client,
	}
}

/*
GetEmbeddingsSdapiV1EmbeddingsGetParams contains all the parameters to send to the API endpoint

	for the get embeddings sdapi v1 embeddings get operation.

	Typically these are written to a http.Request.
*/
type GetEmbeddingsSdapiV1EmbeddingsGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get embeddings sdapi v1 embeddings get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEmbeddingsSdapiV1EmbeddingsGetParams) WithDefaults() *GetEmbeddingsSdapiV1EmbeddingsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get embeddings sdapi v1 embeddings get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEmbeddingsSdapiV1EmbeddingsGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get embeddings sdapi v1 embeddings get params
func (o *GetEmbeddingsSdapiV1EmbeddingsGetParams) WithTimeout(timeout time.Duration) *GetEmbeddingsSdapiV1EmbeddingsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get embeddings sdapi v1 embeddings get params
func (o *GetEmbeddingsSdapiV1EmbeddingsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get embeddings sdapi v1 embeddings get params
func (o *GetEmbeddingsSdapiV1EmbeddingsGetParams) WithContext(ctx context.Context) *GetEmbeddingsSdapiV1EmbeddingsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get embeddings sdapi v1 embeddings get params
func (o *GetEmbeddingsSdapiV1EmbeddingsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get embeddings sdapi v1 embeddings get params
func (o *GetEmbeddingsSdapiV1EmbeddingsGetParams) WithHTTPClient(client *http.Client) *GetEmbeddingsSdapiV1EmbeddingsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get embeddings sdapi v1 embeddings get params
func (o *GetEmbeddingsSdapiV1EmbeddingsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetEmbeddingsSdapiV1EmbeddingsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
