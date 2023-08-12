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

// DeoldifyImageDeoldifyImagePostReader is a Reader for the DeoldifyImageDeoldifyImagePost structure.
type DeoldifyImageDeoldifyImagePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeoldifyImageDeoldifyImagePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeoldifyImageDeoldifyImagePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewDeoldifyImageDeoldifyImagePostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeoldifyImageDeoldifyImagePostOK creates a DeoldifyImageDeoldifyImagePostOK with default headers values
func NewDeoldifyImageDeoldifyImagePostOK() *DeoldifyImageDeoldifyImagePostOK {
	return &DeoldifyImageDeoldifyImagePostOK{}
}

/*
DeoldifyImageDeoldifyImagePostOK describes a response with status code 200, with default header values.

Successful Response
*/
type DeoldifyImageDeoldifyImagePostOK struct {
	Payload interface{}
}

// IsSuccess returns true when this deoldify image deoldify image post o k response has a 2xx status code
func (o *DeoldifyImageDeoldifyImagePostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this deoldify image deoldify image post o k response has a 3xx status code
func (o *DeoldifyImageDeoldifyImagePostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this deoldify image deoldify image post o k response has a 4xx status code
func (o *DeoldifyImageDeoldifyImagePostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this deoldify image deoldify image post o k response has a 5xx status code
func (o *DeoldifyImageDeoldifyImagePostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this deoldify image deoldify image post o k response a status code equal to that given
func (o *DeoldifyImageDeoldifyImagePostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the deoldify image deoldify image post o k response
func (o *DeoldifyImageDeoldifyImagePostOK) Code() int {
	return 200
}

func (o *DeoldifyImageDeoldifyImagePostOK) Error() string {
	return fmt.Sprintf("[POST /deoldify/image][%d] deoldifyImageDeoldifyImagePostOK  %+v", 200, o.Payload)
}

func (o *DeoldifyImageDeoldifyImagePostOK) String() string {
	return fmt.Sprintf("[POST /deoldify/image][%d] deoldifyImageDeoldifyImagePostOK  %+v", 200, o.Payload)
}

func (o *DeoldifyImageDeoldifyImagePostOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeoldifyImageDeoldifyImagePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeoldifyImageDeoldifyImagePostUnprocessableEntity creates a DeoldifyImageDeoldifyImagePostUnprocessableEntity with default headers values
func NewDeoldifyImageDeoldifyImagePostUnprocessableEntity() *DeoldifyImageDeoldifyImagePostUnprocessableEntity {
	return &DeoldifyImageDeoldifyImagePostUnprocessableEntity{}
}

/*
DeoldifyImageDeoldifyImagePostUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type DeoldifyImageDeoldifyImagePostUnprocessableEntity struct {
	Payload *models.HTTPValidationError
}

// IsSuccess returns true when this deoldify image deoldify image post unprocessable entity response has a 2xx status code
func (o *DeoldifyImageDeoldifyImagePostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this deoldify image deoldify image post unprocessable entity response has a 3xx status code
func (o *DeoldifyImageDeoldifyImagePostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this deoldify image deoldify image post unprocessable entity response has a 4xx status code
func (o *DeoldifyImageDeoldifyImagePostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this deoldify image deoldify image post unprocessable entity response has a 5xx status code
func (o *DeoldifyImageDeoldifyImagePostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this deoldify image deoldify image post unprocessable entity response a status code equal to that given
func (o *DeoldifyImageDeoldifyImagePostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the deoldify image deoldify image post unprocessable entity response
func (o *DeoldifyImageDeoldifyImagePostUnprocessableEntity) Code() int {
	return 422
}

func (o *DeoldifyImageDeoldifyImagePostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /deoldify/image][%d] deoldifyImageDeoldifyImagePostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *DeoldifyImageDeoldifyImagePostUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /deoldify/image][%d] deoldifyImageDeoldifyImagePostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *DeoldifyImageDeoldifyImagePostUnprocessableEntity) GetPayload() *models.HTTPValidationError {
	return o.Payload
}

func (o *DeoldifyImageDeoldifyImagePostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
