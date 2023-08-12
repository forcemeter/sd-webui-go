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

// StaticResourceStaticPathGetReader is a Reader for the StaticResourceStaticPathGet structure.
type StaticResourceStaticPathGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StaticResourceStaticPathGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStaticResourceStaticPathGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewStaticResourceStaticPathGetUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStaticResourceStaticPathGetOK creates a StaticResourceStaticPathGetOK with default headers values
func NewStaticResourceStaticPathGetOK() *StaticResourceStaticPathGetOK {
	return &StaticResourceStaticPathGetOK{}
}

/*
StaticResourceStaticPathGetOK describes a response with status code 200, with default header values.

Successful Response
*/
type StaticResourceStaticPathGetOK struct {
	Payload interface{}
}

// IsSuccess returns true when this static resource static path get o k response has a 2xx status code
func (o *StaticResourceStaticPathGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this static resource static path get o k response has a 3xx status code
func (o *StaticResourceStaticPathGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this static resource static path get o k response has a 4xx status code
func (o *StaticResourceStaticPathGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this static resource static path get o k response has a 5xx status code
func (o *StaticResourceStaticPathGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this static resource static path get o k response a status code equal to that given
func (o *StaticResourceStaticPathGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the static resource static path get o k response
func (o *StaticResourceStaticPathGetOK) Code() int {
	return 200
}

func (o *StaticResourceStaticPathGetOK) Error() string {
	return fmt.Sprintf("[GET /static/{path}][%d] staticResourceStaticPathGetOK  %+v", 200, o.Payload)
}

func (o *StaticResourceStaticPathGetOK) String() string {
	return fmt.Sprintf("[GET /static/{path}][%d] staticResourceStaticPathGetOK  %+v", 200, o.Payload)
}

func (o *StaticResourceStaticPathGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *StaticResourceStaticPathGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStaticResourceStaticPathGetUnprocessableEntity creates a StaticResourceStaticPathGetUnprocessableEntity with default headers values
func NewStaticResourceStaticPathGetUnprocessableEntity() *StaticResourceStaticPathGetUnprocessableEntity {
	return &StaticResourceStaticPathGetUnprocessableEntity{}
}

/*
StaticResourceStaticPathGetUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type StaticResourceStaticPathGetUnprocessableEntity struct {
	Payload *models.HTTPValidationError
}

// IsSuccess returns true when this static resource static path get unprocessable entity response has a 2xx status code
func (o *StaticResourceStaticPathGetUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this static resource static path get unprocessable entity response has a 3xx status code
func (o *StaticResourceStaticPathGetUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this static resource static path get unprocessable entity response has a 4xx status code
func (o *StaticResourceStaticPathGetUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this static resource static path get unprocessable entity response has a 5xx status code
func (o *StaticResourceStaticPathGetUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this static resource static path get unprocessable entity response a status code equal to that given
func (o *StaticResourceStaticPathGetUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the static resource static path get unprocessable entity response
func (o *StaticResourceStaticPathGetUnprocessableEntity) Code() int {
	return 422
}

func (o *StaticResourceStaticPathGetUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /static/{path}][%d] staticResourceStaticPathGetUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *StaticResourceStaticPathGetUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /static/{path}][%d] staticResourceStaticPathGetUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *StaticResourceStaticPathGetUnprocessableEntity) GetPayload() *models.HTTPValidationError {
	return o.Payload
}

func (o *StaticResourceStaticPathGetUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
