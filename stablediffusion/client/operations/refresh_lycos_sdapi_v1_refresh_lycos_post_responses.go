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

// RefreshLycosSdapiV1RefreshLycosPostReader is a Reader for the RefreshLycosSdapiV1RefreshLycosPost structure.
type RefreshLycosSdapiV1RefreshLycosPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RefreshLycosSdapiV1RefreshLycosPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRefreshLycosSdapiV1RefreshLycosPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRefreshLycosSdapiV1RefreshLycosPostOK creates a RefreshLycosSdapiV1RefreshLycosPostOK with default headers values
func NewRefreshLycosSdapiV1RefreshLycosPostOK() *RefreshLycosSdapiV1RefreshLycosPostOK {
	return &RefreshLycosSdapiV1RefreshLycosPostOK{}
}

/*
RefreshLycosSdapiV1RefreshLycosPostOK describes a response with status code 200, with default header values.

Successful Response
*/
type RefreshLycosSdapiV1RefreshLycosPostOK struct {
	Payload interface{}
}

// IsSuccess returns true when this refresh lycos sdapi v1 refresh lycos post o k response has a 2xx status code
func (o *RefreshLycosSdapiV1RefreshLycosPostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this refresh lycos sdapi v1 refresh lycos post o k response has a 3xx status code
func (o *RefreshLycosSdapiV1RefreshLycosPostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this refresh lycos sdapi v1 refresh lycos post o k response has a 4xx status code
func (o *RefreshLycosSdapiV1RefreshLycosPostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this refresh lycos sdapi v1 refresh lycos post o k response has a 5xx status code
func (o *RefreshLycosSdapiV1RefreshLycosPostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this refresh lycos sdapi v1 refresh lycos post o k response a status code equal to that given
func (o *RefreshLycosSdapiV1RefreshLycosPostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the refresh lycos sdapi v1 refresh lycos post o k response
func (o *RefreshLycosSdapiV1RefreshLycosPostOK) Code() int {
	return 200
}

func (o *RefreshLycosSdapiV1RefreshLycosPostOK) Error() string {
	return fmt.Sprintf("[POST /sdapi/v1/refresh-lycos][%d] refreshLycosSdapiV1RefreshLycosPostOK  %+v", 200, o.Payload)
}

func (o *RefreshLycosSdapiV1RefreshLycosPostOK) String() string {
	return fmt.Sprintf("[POST /sdapi/v1/refresh-lycos][%d] refreshLycosSdapiV1RefreshLycosPostOK  %+v", 200, o.Payload)
}

func (o *RefreshLycosSdapiV1RefreshLycosPostOK) GetPayload() interface{} {
	return o.Payload
}

func (o *RefreshLycosSdapiV1RefreshLycosPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
