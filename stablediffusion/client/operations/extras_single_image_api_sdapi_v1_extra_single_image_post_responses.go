// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/SpenserCai/sd-webui-go/stablediffusion/models"
)

// ExtrasSingleImageAPISdapiV1ExtraSingleImagePostReader is a Reader for the ExtrasSingleImageAPISdapiV1ExtraSingleImagePost structure.
type ExtrasSingleImageAPISdapiV1ExtraSingleImagePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasSingleImageAPISdapiV1ExtraSingleImagePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasSingleImageAPISdapiV1ExtraSingleImagePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewExtrasSingleImageAPISdapiV1ExtraSingleImagePostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasSingleImageAPISdapiV1ExtraSingleImagePostOK creates a ExtrasSingleImageAPISdapiV1ExtraSingleImagePostOK with default headers values
func NewExtrasSingleImageAPISdapiV1ExtraSingleImagePostOK() *ExtrasSingleImageAPISdapiV1ExtraSingleImagePostOK {
	return &ExtrasSingleImageAPISdapiV1ExtraSingleImagePostOK{}
}

/*
ExtrasSingleImageAPISdapiV1ExtraSingleImagePostOK describes a response with status code 200, with default header values.

Successful Response
*/
type ExtrasSingleImageAPISdapiV1ExtraSingleImagePostOK struct {
	Payload *models.ExtrasSingleImageResponse
}

// IsSuccess returns true when this extras single image Api sdapi v1 extra single image post o k response has a 2xx status code
func (o *ExtrasSingleImageAPISdapiV1ExtraSingleImagePostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras single image Api sdapi v1 extra single image post o k response has a 3xx status code
func (o *ExtrasSingleImageAPISdapiV1ExtraSingleImagePostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras single image Api sdapi v1 extra single image post o k response has a 4xx status code
func (o *ExtrasSingleImageAPISdapiV1ExtraSingleImagePostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras single image Api sdapi v1 extra single image post o k response has a 5xx status code
func (o *ExtrasSingleImageAPISdapiV1ExtraSingleImagePostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras single image Api sdapi v1 extra single image post o k response a status code equal to that given
func (o *ExtrasSingleImageAPISdapiV1ExtraSingleImagePostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras single image Api sdapi v1 extra single image post o k response
func (o *ExtrasSingleImageAPISdapiV1ExtraSingleImagePostOK) Code() int {
	return 200
}

func (o *ExtrasSingleImageAPISdapiV1ExtraSingleImagePostOK) Error() string {
	return fmt.Sprintf("[POST /sdapi/v1/extra-single-image][%d] extrasSingleImageApiSdapiV1ExtraSingleImagePostOK  %+v", 200, o.Payload)
}

func (o *ExtrasSingleImageAPISdapiV1ExtraSingleImagePostOK) String() string {
	return fmt.Sprintf("[POST /sdapi/v1/extra-single-image][%d] extrasSingleImageApiSdapiV1ExtraSingleImagePostOK  %+v", 200, o.Payload)
}

func (o *ExtrasSingleImageAPISdapiV1ExtraSingleImagePostOK) GetPayload() *models.ExtrasSingleImageResponse {
	return o.Payload
}

func (o *ExtrasSingleImageAPISdapiV1ExtraSingleImagePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExtrasSingleImageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasSingleImageAPISdapiV1ExtraSingleImagePostUnprocessableEntity creates a ExtrasSingleImageAPISdapiV1ExtraSingleImagePostUnprocessableEntity with default headers values
func NewExtrasSingleImageAPISdapiV1ExtraSingleImagePostUnprocessableEntity() *ExtrasSingleImageAPISdapiV1ExtraSingleImagePostUnprocessableEntity {
	return &ExtrasSingleImageAPISdapiV1ExtraSingleImagePostUnprocessableEntity{}
}

/*
ExtrasSingleImageAPISdapiV1ExtraSingleImagePostUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type ExtrasSingleImageAPISdapiV1ExtraSingleImagePostUnprocessableEntity struct {
	Payload *models.HTTPValidationError
}

// IsSuccess returns true when this extras single image Api sdapi v1 extra single image post unprocessable entity response has a 2xx status code
func (o *ExtrasSingleImageAPISdapiV1ExtraSingleImagePostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this extras single image Api sdapi v1 extra single image post unprocessable entity response has a 3xx status code
func (o *ExtrasSingleImageAPISdapiV1ExtraSingleImagePostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras single image Api sdapi v1 extra single image post unprocessable entity response has a 4xx status code
func (o *ExtrasSingleImageAPISdapiV1ExtraSingleImagePostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this extras single image Api sdapi v1 extra single image post unprocessable entity response has a 5xx status code
func (o *ExtrasSingleImageAPISdapiV1ExtraSingleImagePostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this extras single image Api sdapi v1 extra single image post unprocessable entity response a status code equal to that given
func (o *ExtrasSingleImageAPISdapiV1ExtraSingleImagePostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the extras single image Api sdapi v1 extra single image post unprocessable entity response
func (o *ExtrasSingleImageAPISdapiV1ExtraSingleImagePostUnprocessableEntity) Code() int {
	return 422
}

func (o *ExtrasSingleImageAPISdapiV1ExtraSingleImagePostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /sdapi/v1/extra-single-image][%d] extrasSingleImageApiSdapiV1ExtraSingleImagePostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ExtrasSingleImageAPISdapiV1ExtraSingleImagePostUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /sdapi/v1/extra-single-image][%d] extrasSingleImageApiSdapiV1ExtraSingleImagePostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ExtrasSingleImageAPISdapiV1ExtraSingleImagePostUnprocessableEntity) GetPayload() *models.HTTPValidationError {
	return o.Payload
}

func (o *ExtrasSingleImageAPISdapiV1ExtraSingleImagePostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
