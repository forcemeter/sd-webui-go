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

// NewText2imgapiSdapiV1Txt2imgPostParams creates a new Text2imgapiSdapiV1Txt2imgPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewText2imgapiSdapiV1Txt2imgPostParams() *Text2imgapiSdapiV1Txt2imgPostParams {
	return &Text2imgapiSdapiV1Txt2imgPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewText2imgapiSdapiV1Txt2imgPostParamsWithTimeout creates a new Text2imgapiSdapiV1Txt2imgPostParams object
// with the ability to set a timeout on a request.
func NewText2imgapiSdapiV1Txt2imgPostParamsWithTimeout(timeout time.Duration) *Text2imgapiSdapiV1Txt2imgPostParams {
	return &Text2imgapiSdapiV1Txt2imgPostParams{
		timeout: timeout,
	}
}

// NewText2imgapiSdapiV1Txt2imgPostParamsWithContext creates a new Text2imgapiSdapiV1Txt2imgPostParams object
// with the ability to set a context for a request.
func NewText2imgapiSdapiV1Txt2imgPostParamsWithContext(ctx context.Context) *Text2imgapiSdapiV1Txt2imgPostParams {
	return &Text2imgapiSdapiV1Txt2imgPostParams{
		Context: ctx,
	}
}

// NewText2imgapiSdapiV1Txt2imgPostParamsWithHTTPClient creates a new Text2imgapiSdapiV1Txt2imgPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewText2imgapiSdapiV1Txt2imgPostParamsWithHTTPClient(client *http.Client) *Text2imgapiSdapiV1Txt2imgPostParams {
	return &Text2imgapiSdapiV1Txt2imgPostParams{
		HTTPClient: client,
	}
}

/*
Text2imgapiSdapiV1Txt2imgPostParams contains all the parameters to send to the API endpoint

	for the text2imgapi sdapi v1 txt2img post operation.

	Typically these are written to a http.Request.
*/
type Text2imgapiSdapiV1Txt2imgPostParams struct {

	// Body.
	Body *models.StableDiffusionProcessingTxt2Img

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the text2imgapi sdapi v1 txt2img post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Text2imgapiSdapiV1Txt2imgPostParams) WithDefaults() *Text2imgapiSdapiV1Txt2imgPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the text2imgapi sdapi v1 txt2img post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Text2imgapiSdapiV1Txt2imgPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the text2imgapi sdapi v1 txt2img post params
func (o *Text2imgapiSdapiV1Txt2imgPostParams) WithTimeout(timeout time.Duration) *Text2imgapiSdapiV1Txt2imgPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the text2imgapi sdapi v1 txt2img post params
func (o *Text2imgapiSdapiV1Txt2imgPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the text2imgapi sdapi v1 txt2img post params
func (o *Text2imgapiSdapiV1Txt2imgPostParams) WithContext(ctx context.Context) *Text2imgapiSdapiV1Txt2imgPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the text2imgapi sdapi v1 txt2img post params
func (o *Text2imgapiSdapiV1Txt2imgPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the text2imgapi sdapi v1 txt2img post params
func (o *Text2imgapiSdapiV1Txt2imgPostParams) WithHTTPClient(client *http.Client) *Text2imgapiSdapiV1Txt2imgPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the text2imgapi sdapi v1 txt2img post params
func (o *Text2imgapiSdapiV1Txt2imgPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the text2imgapi sdapi v1 txt2img post params
func (o *Text2imgapiSdapiV1Txt2imgPostParams) WithBody(body *models.StableDiffusionProcessingTxt2Img) *Text2imgapiSdapiV1Txt2imgPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the text2imgapi sdapi v1 txt2img post params
func (o *Text2imgapiSdapiV1Txt2imgPostParams) SetBody(body *models.StableDiffusionProcessingTxt2Img) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *Text2imgapiSdapiV1Txt2imgPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
