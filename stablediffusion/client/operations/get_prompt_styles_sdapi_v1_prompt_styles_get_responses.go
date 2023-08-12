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

// GetPromptStylesSdapiV1PromptStylesGetReader is a Reader for the GetPromptStylesSdapiV1PromptStylesGet structure.
type GetPromptStylesSdapiV1PromptStylesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPromptStylesSdapiV1PromptStylesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPromptStylesSdapiV1PromptStylesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPromptStylesSdapiV1PromptStylesGetOK creates a GetPromptStylesSdapiV1PromptStylesGetOK with default headers values
func NewGetPromptStylesSdapiV1PromptStylesGetOK() *GetPromptStylesSdapiV1PromptStylesGetOK {
	return &GetPromptStylesSdapiV1PromptStylesGetOK{}
}

/*
GetPromptStylesSdapiV1PromptStylesGetOK describes a response with status code 200, with default header values.

Successful Response
*/
type GetPromptStylesSdapiV1PromptStylesGetOK struct {
	Payload []*models.PromptStyleItem
}

// IsSuccess returns true when this get prompt styles sdapi v1 prompt styles get o k response has a 2xx status code
func (o *GetPromptStylesSdapiV1PromptStylesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get prompt styles sdapi v1 prompt styles get o k response has a 3xx status code
func (o *GetPromptStylesSdapiV1PromptStylesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get prompt styles sdapi v1 prompt styles get o k response has a 4xx status code
func (o *GetPromptStylesSdapiV1PromptStylesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get prompt styles sdapi v1 prompt styles get o k response has a 5xx status code
func (o *GetPromptStylesSdapiV1PromptStylesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get prompt styles sdapi v1 prompt styles get o k response a status code equal to that given
func (o *GetPromptStylesSdapiV1PromptStylesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get prompt styles sdapi v1 prompt styles get o k response
func (o *GetPromptStylesSdapiV1PromptStylesGetOK) Code() int {
	return 200
}

func (o *GetPromptStylesSdapiV1PromptStylesGetOK) Error() string {
	return fmt.Sprintf("[GET /sdapi/v1/prompt-styles][%d] getPromptStylesSdapiV1PromptStylesGetOK  %+v", 200, o.Payload)
}

func (o *GetPromptStylesSdapiV1PromptStylesGetOK) String() string {
	return fmt.Sprintf("[GET /sdapi/v1/prompt-styles][%d] getPromptStylesSdapiV1PromptStylesGetOK  %+v", 200, o.Payload)
}

func (o *GetPromptStylesSdapiV1PromptStylesGetOK) GetPayload() []*models.PromptStyleItem {
	return o.Payload
}

func (o *GetPromptStylesSdapiV1PromptStylesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
