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

// ThumbnailInfiniteImageBrowsingImageThumbnailGetReader is a Reader for the ThumbnailInfiniteImageBrowsingImageThumbnailGet structure.
type ThumbnailInfiniteImageBrowsingImageThumbnailGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewThumbnailInfiniteImageBrowsingImageThumbnailGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewThumbnailInfiniteImageBrowsingImageThumbnailGetUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewThumbnailInfiniteImageBrowsingImageThumbnailGetOK creates a ThumbnailInfiniteImageBrowsingImageThumbnailGetOK with default headers values
func NewThumbnailInfiniteImageBrowsingImageThumbnailGetOK() *ThumbnailInfiniteImageBrowsingImageThumbnailGetOK {
	return &ThumbnailInfiniteImageBrowsingImageThumbnailGetOK{}
}

/*
ThumbnailInfiniteImageBrowsingImageThumbnailGetOK describes a response with status code 200, with default header values.

Successful Response
*/
type ThumbnailInfiniteImageBrowsingImageThumbnailGetOK struct {
	Payload interface{}
}

// IsSuccess returns true when this thumbnail infinite image browsing image thumbnail get o k response has a 2xx status code
func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this thumbnail infinite image browsing image thumbnail get o k response has a 3xx status code
func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this thumbnail infinite image browsing image thumbnail get o k response has a 4xx status code
func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this thumbnail infinite image browsing image thumbnail get o k response has a 5xx status code
func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this thumbnail infinite image browsing image thumbnail get o k response a status code equal to that given
func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the thumbnail infinite image browsing image thumbnail get o k response
func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetOK) Code() int {
	return 200
}

func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetOK) Error() string {
	return fmt.Sprintf("[GET /infinite_image_browsing/image-thumbnail][%d] thumbnailInfiniteImageBrowsingImageThumbnailGetOK  %+v", 200, o.Payload)
}

func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetOK) String() string {
	return fmt.Sprintf("[GET /infinite_image_browsing/image-thumbnail][%d] thumbnailInfiniteImageBrowsingImageThumbnailGetOK  %+v", 200, o.Payload)
}

func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThumbnailInfiniteImageBrowsingImageThumbnailGetUnprocessableEntity creates a ThumbnailInfiniteImageBrowsingImageThumbnailGetUnprocessableEntity with default headers values
func NewThumbnailInfiniteImageBrowsingImageThumbnailGetUnprocessableEntity() *ThumbnailInfiniteImageBrowsingImageThumbnailGetUnprocessableEntity {
	return &ThumbnailInfiniteImageBrowsingImageThumbnailGetUnprocessableEntity{}
}

/*
ThumbnailInfiniteImageBrowsingImageThumbnailGetUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type ThumbnailInfiniteImageBrowsingImageThumbnailGetUnprocessableEntity struct {
	Payload *models.HTTPValidationError
}

// IsSuccess returns true when this thumbnail infinite image browsing image thumbnail get unprocessable entity response has a 2xx status code
func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this thumbnail infinite image browsing image thumbnail get unprocessable entity response has a 3xx status code
func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this thumbnail infinite image browsing image thumbnail get unprocessable entity response has a 4xx status code
func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this thumbnail infinite image browsing image thumbnail get unprocessable entity response has a 5xx status code
func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this thumbnail infinite image browsing image thumbnail get unprocessable entity response a status code equal to that given
func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the thumbnail infinite image browsing image thumbnail get unprocessable entity response
func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetUnprocessableEntity) Code() int {
	return 422
}

func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /infinite_image_browsing/image-thumbnail][%d] thumbnailInfiniteImageBrowsingImageThumbnailGetUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /infinite_image_browsing/image-thumbnail][%d] thumbnailInfiniteImageBrowsingImageThumbnailGetUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetUnprocessableEntity) GetPayload() *models.HTTPValidationError {
	return o.Payload
}

func (o *ThumbnailInfiniteImageBrowsingImageThumbnailGetUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
