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

// MainGetReader is a Reader for the MainGet structure.
type MainGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MainGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMainGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewMainGetOK creates a MainGetOK with default headers values
func NewMainGetOK() *MainGetOK {
	return &MainGetOK{}
}

/*
MainGetOK describes a response with status code 200, with default header values.

Successful Response
*/
type MainGetOK struct {
	Payload string
}

// IsSuccess returns true when this main get o k response has a 2xx status code
func (o *MainGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this main get o k response has a 3xx status code
func (o *MainGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this main get o k response has a 4xx status code
func (o *MainGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this main get o k response has a 5xx status code
func (o *MainGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this main get o k response a status code equal to that given
func (o *MainGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the main get o k response
func (o *MainGetOK) Code() int {
	return 200
}

func (o *MainGetOK) Error() string {
	return fmt.Sprintf("[GET /][%d] mainGetOK  %+v", 200, o.Payload)
}

func (o *MainGetOK) String() string {
	return fmt.Sprintf("[GET /][%d] mainGetOK  %+v", 200, o.Payload)
}

func (o *MainGetOK) GetPayload() string {
	return o.Payload
}

func (o *MainGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
