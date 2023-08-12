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

// ReloadapiSdapiV1ReloadCheckpointPostReader is a Reader for the ReloadapiSdapiV1ReloadCheckpointPost structure.
type ReloadapiSdapiV1ReloadCheckpointPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReloadapiSdapiV1ReloadCheckpointPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReloadapiSdapiV1ReloadCheckpointPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReloadapiSdapiV1ReloadCheckpointPostOK creates a ReloadapiSdapiV1ReloadCheckpointPostOK with default headers values
func NewReloadapiSdapiV1ReloadCheckpointPostOK() *ReloadapiSdapiV1ReloadCheckpointPostOK {
	return &ReloadapiSdapiV1ReloadCheckpointPostOK{}
}

/*
ReloadapiSdapiV1ReloadCheckpointPostOK describes a response with status code 200, with default header values.

Successful Response
*/
type ReloadapiSdapiV1ReloadCheckpointPostOK struct {
	Payload interface{}
}

// IsSuccess returns true when this reloadapi sdapi v1 reload checkpoint post o k response has a 2xx status code
func (o *ReloadapiSdapiV1ReloadCheckpointPostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this reloadapi sdapi v1 reload checkpoint post o k response has a 3xx status code
func (o *ReloadapiSdapiV1ReloadCheckpointPostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this reloadapi sdapi v1 reload checkpoint post o k response has a 4xx status code
func (o *ReloadapiSdapiV1ReloadCheckpointPostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this reloadapi sdapi v1 reload checkpoint post o k response has a 5xx status code
func (o *ReloadapiSdapiV1ReloadCheckpointPostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this reloadapi sdapi v1 reload checkpoint post o k response a status code equal to that given
func (o *ReloadapiSdapiV1ReloadCheckpointPostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the reloadapi sdapi v1 reload checkpoint post o k response
func (o *ReloadapiSdapiV1ReloadCheckpointPostOK) Code() int {
	return 200
}

func (o *ReloadapiSdapiV1ReloadCheckpointPostOK) Error() string {
	return fmt.Sprintf("[POST /sdapi/v1/reload-checkpoint][%d] reloadapiSdapiV1ReloadCheckpointPostOK  %+v", 200, o.Payload)
}

func (o *ReloadapiSdapiV1ReloadCheckpointPostOK) String() string {
	return fmt.Sprintf("[POST /sdapi/v1/reload-checkpoint][%d] reloadapiSdapiV1ReloadCheckpointPostOK  %+v", 200, o.Payload)
}

func (o *ReloadapiSdapiV1ReloadCheckpointPostOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ReloadapiSdapiV1ReloadCheckpointPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
