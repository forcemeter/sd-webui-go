// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// LambdaInternalProfileStartupGetReader is a Reader for the LambdaInternalProfileStartupGet structure.
type LambdaInternalProfileStartupGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LambdaInternalProfileStartupGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLambdaInternalProfileStartupGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLambdaInternalProfileStartupGetOK creates a LambdaInternalProfileStartupGetOK with default headers values
func NewLambdaInternalProfileStartupGetOK() *LambdaInternalProfileStartupGetOK {
	return &LambdaInternalProfileStartupGetOK{}
}

/*
LambdaInternalProfileStartupGetOK describes a response with status code 200, with default header values.

Successful Response
*/
type LambdaInternalProfileStartupGetOK struct {
	Payload interface{}
}

// IsSuccess returns true when this lambda internal profile startup get o k response has a 2xx status code
func (o *LambdaInternalProfileStartupGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this lambda internal profile startup get o k response has a 3xx status code
func (o *LambdaInternalProfileStartupGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this lambda internal profile startup get o k response has a 4xx status code
func (o *LambdaInternalProfileStartupGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this lambda internal profile startup get o k response has a 5xx status code
func (o *LambdaInternalProfileStartupGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this lambda internal profile startup get o k response a status code equal to that given
func (o *LambdaInternalProfileStartupGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the lambda internal profile startup get o k response
func (o *LambdaInternalProfileStartupGetOK) Code() int {
	return 200
}

func (o *LambdaInternalProfileStartupGetOK) Error() string {
	return fmt.Sprintf("[GET /internal/profile-startup][%d] lambdaInternalProfileStartupGetOK  %+v", 200, o.Payload)
}

func (o *LambdaInternalProfileStartupGetOK) String() string {
	return fmt.Sprintf("[GET /internal/profile-startup][%d] lambdaInternalProfileStartupGetOK  %+v", 200, o.Payload)
}

func (o *LambdaInternalProfileStartupGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *LambdaInternalProfileStartupGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
