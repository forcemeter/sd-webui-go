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

// RemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostReader is a Reader for the RemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPost structure.
type RemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewRemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostOK creates a RemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostOK with default headers values
func NewRemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostOK() *RemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostOK {
	return &RemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostOK{}
}

/*
RemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostOK describes a response with status code 200, with default header values.

Successful Response
*/
type RemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostOK struct {
	Payload interface{}
}

// IsSuccess returns true when this remove custom tag from img infinite image browsing db remove custom tag from img post o k response has a 2xx status code
func (o *RemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this remove custom tag from img infinite image browsing db remove custom tag from img post o k response has a 3xx status code
func (o *RemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove custom tag from img infinite image browsing db remove custom tag from img post o k response has a 4xx status code
func (o *RemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove custom tag from img infinite image browsing db remove custom tag from img post o k response has a 5xx status code
func (o *RemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this remove custom tag from img infinite image browsing db remove custom tag from img post o k response a status code equal to that given
func (o *RemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the remove custom tag from img infinite image browsing db remove custom tag from img post o k response
func (o *RemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostOK) Code() int {
	return 200
}

func (o *RemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostOK) Error() string {
	return fmt.Sprintf("[POST /infinite_image_browsing/db/remove_custom_tag_from_img][%d] removeCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostOK  %+v", 200, o.Payload)
}

func (o *RemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostOK) String() string {
	return fmt.Sprintf("[POST /infinite_image_browsing/db/remove_custom_tag_from_img][%d] removeCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostOK  %+v", 200, o.Payload)
}

func (o *RemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostOK) GetPayload() interface{} {
	return o.Payload
}

func (o *RemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostUnprocessableEntity creates a RemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostUnprocessableEntity with default headers values
func NewRemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostUnprocessableEntity() *RemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostUnprocessableEntity {
	return &RemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostUnprocessableEntity{}
}

/*
RemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type RemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostUnprocessableEntity struct {
	Payload *models.HTTPValidationError
}

// IsSuccess returns true when this remove custom tag from img infinite image browsing db remove custom tag from img post unprocessable entity response has a 2xx status code
func (o *RemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove custom tag from img infinite image browsing db remove custom tag from img post unprocessable entity response has a 3xx status code
func (o *RemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove custom tag from img infinite image browsing db remove custom tag from img post unprocessable entity response has a 4xx status code
func (o *RemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove custom tag from img infinite image browsing db remove custom tag from img post unprocessable entity response has a 5xx status code
func (o *RemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this remove custom tag from img infinite image browsing db remove custom tag from img post unprocessable entity response a status code equal to that given
func (o *RemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the remove custom tag from img infinite image browsing db remove custom tag from img post unprocessable entity response
func (o *RemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostUnprocessableEntity) Code() int {
	return 422
}

func (o *RemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /infinite_image_browsing/db/remove_custom_tag_from_img][%d] removeCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *RemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /infinite_image_browsing/db/remove_custom_tag_from_img][%d] removeCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *RemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostUnprocessableEntity) GetPayload() *models.HTTPValidationError {
	return o.Payload
}

func (o *RemoveCustomTagFromImgInfiniteImageBrowsingDbRemoveCustomTagFromImgPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
