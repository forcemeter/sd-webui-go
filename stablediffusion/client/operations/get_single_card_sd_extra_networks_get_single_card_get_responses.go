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

// GetSingleCardSdExtraNetworksGetSingleCardGetReader is a Reader for the GetSingleCardSdExtraNetworksGetSingleCardGet structure.
type GetSingleCardSdExtraNetworksGetSingleCardGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSingleCardSdExtraNetworksGetSingleCardGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSingleCardSdExtraNetworksGetSingleCardGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewGetSingleCardSdExtraNetworksGetSingleCardGetUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSingleCardSdExtraNetworksGetSingleCardGetOK creates a GetSingleCardSdExtraNetworksGetSingleCardGetOK with default headers values
func NewGetSingleCardSdExtraNetworksGetSingleCardGetOK() *GetSingleCardSdExtraNetworksGetSingleCardGetOK {
	return &GetSingleCardSdExtraNetworksGetSingleCardGetOK{}
}

/*
GetSingleCardSdExtraNetworksGetSingleCardGetOK describes a response with status code 200, with default header values.

Successful Response
*/
type GetSingleCardSdExtraNetworksGetSingleCardGetOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get single card sd extra networks get single card get o k response has a 2xx status code
func (o *GetSingleCardSdExtraNetworksGetSingleCardGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get single card sd extra networks get single card get o k response has a 3xx status code
func (o *GetSingleCardSdExtraNetworksGetSingleCardGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get single card sd extra networks get single card get o k response has a 4xx status code
func (o *GetSingleCardSdExtraNetworksGetSingleCardGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get single card sd extra networks get single card get o k response has a 5xx status code
func (o *GetSingleCardSdExtraNetworksGetSingleCardGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get single card sd extra networks get single card get o k response a status code equal to that given
func (o *GetSingleCardSdExtraNetworksGetSingleCardGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get single card sd extra networks get single card get o k response
func (o *GetSingleCardSdExtraNetworksGetSingleCardGetOK) Code() int {
	return 200
}

func (o *GetSingleCardSdExtraNetworksGetSingleCardGetOK) Error() string {
	return fmt.Sprintf("[GET /sd_extra_networks/get-single-card][%d] getSingleCardSdExtraNetworksGetSingleCardGetOK  %+v", 200, o.Payload)
}

func (o *GetSingleCardSdExtraNetworksGetSingleCardGetOK) String() string {
	return fmt.Sprintf("[GET /sd_extra_networks/get-single-card][%d] getSingleCardSdExtraNetworksGetSingleCardGetOK  %+v", 200, o.Payload)
}

func (o *GetSingleCardSdExtraNetworksGetSingleCardGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetSingleCardSdExtraNetworksGetSingleCardGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSingleCardSdExtraNetworksGetSingleCardGetUnprocessableEntity creates a GetSingleCardSdExtraNetworksGetSingleCardGetUnprocessableEntity with default headers values
func NewGetSingleCardSdExtraNetworksGetSingleCardGetUnprocessableEntity() *GetSingleCardSdExtraNetworksGetSingleCardGetUnprocessableEntity {
	return &GetSingleCardSdExtraNetworksGetSingleCardGetUnprocessableEntity{}
}

/*
GetSingleCardSdExtraNetworksGetSingleCardGetUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type GetSingleCardSdExtraNetworksGetSingleCardGetUnprocessableEntity struct {
	Payload *models.HTTPValidationError
}

// IsSuccess returns true when this get single card sd extra networks get single card get unprocessable entity response has a 2xx status code
func (o *GetSingleCardSdExtraNetworksGetSingleCardGetUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get single card sd extra networks get single card get unprocessable entity response has a 3xx status code
func (o *GetSingleCardSdExtraNetworksGetSingleCardGetUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get single card sd extra networks get single card get unprocessable entity response has a 4xx status code
func (o *GetSingleCardSdExtraNetworksGetSingleCardGetUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get single card sd extra networks get single card get unprocessable entity response has a 5xx status code
func (o *GetSingleCardSdExtraNetworksGetSingleCardGetUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get single card sd extra networks get single card get unprocessable entity response a status code equal to that given
func (o *GetSingleCardSdExtraNetworksGetSingleCardGetUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get single card sd extra networks get single card get unprocessable entity response
func (o *GetSingleCardSdExtraNetworksGetSingleCardGetUnprocessableEntity) Code() int {
	return 422
}

func (o *GetSingleCardSdExtraNetworksGetSingleCardGetUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /sd_extra_networks/get-single-card][%d] getSingleCardSdExtraNetworksGetSingleCardGetUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetSingleCardSdExtraNetworksGetSingleCardGetUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /sd_extra_networks/get-single-card][%d] getSingleCardSdExtraNetworksGetSingleCardGetUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetSingleCardSdExtraNetworksGetSingleCardGetUnprocessableEntity) GetPayload() *models.HTTPValidationError {
	return o.Payload
}

func (o *GetSingleCardSdExtraNetworksGetSingleCardGetUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
