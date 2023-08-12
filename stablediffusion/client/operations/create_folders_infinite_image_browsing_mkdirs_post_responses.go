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

// CreateFoldersInfiniteImageBrowsingMkdirsPostReader is a Reader for the CreateFoldersInfiniteImageBrowsingMkdirsPost structure.
type CreateFoldersInfiniteImageBrowsingMkdirsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateFoldersInfiniteImageBrowsingMkdirsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateFoldersInfiniteImageBrowsingMkdirsPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewCreateFoldersInfiniteImageBrowsingMkdirsPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateFoldersInfiniteImageBrowsingMkdirsPostOK creates a CreateFoldersInfiniteImageBrowsingMkdirsPostOK with default headers values
func NewCreateFoldersInfiniteImageBrowsingMkdirsPostOK() *CreateFoldersInfiniteImageBrowsingMkdirsPostOK {
	return &CreateFoldersInfiniteImageBrowsingMkdirsPostOK{}
}

/*
CreateFoldersInfiniteImageBrowsingMkdirsPostOK describes a response with status code 200, with default header values.

Successful Response
*/
type CreateFoldersInfiniteImageBrowsingMkdirsPostOK struct {
	Payload interface{}
}

// IsSuccess returns true when this create folders infinite image browsing mkdirs post o k response has a 2xx status code
func (o *CreateFoldersInfiniteImageBrowsingMkdirsPostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create folders infinite image browsing mkdirs post o k response has a 3xx status code
func (o *CreateFoldersInfiniteImageBrowsingMkdirsPostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create folders infinite image browsing mkdirs post o k response has a 4xx status code
func (o *CreateFoldersInfiniteImageBrowsingMkdirsPostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create folders infinite image browsing mkdirs post o k response has a 5xx status code
func (o *CreateFoldersInfiniteImageBrowsingMkdirsPostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create folders infinite image browsing mkdirs post o k response a status code equal to that given
func (o *CreateFoldersInfiniteImageBrowsingMkdirsPostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create folders infinite image browsing mkdirs post o k response
func (o *CreateFoldersInfiniteImageBrowsingMkdirsPostOK) Code() int {
	return 200
}

func (o *CreateFoldersInfiniteImageBrowsingMkdirsPostOK) Error() string {
	return fmt.Sprintf("[POST /infinite_image_browsing/mkdirs][%d] createFoldersInfiniteImageBrowsingMkdirsPostOK  %+v", 200, o.Payload)
}

func (o *CreateFoldersInfiniteImageBrowsingMkdirsPostOK) String() string {
	return fmt.Sprintf("[POST /infinite_image_browsing/mkdirs][%d] createFoldersInfiniteImageBrowsingMkdirsPostOK  %+v", 200, o.Payload)
}

func (o *CreateFoldersInfiniteImageBrowsingMkdirsPostOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateFoldersInfiniteImageBrowsingMkdirsPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFoldersInfiniteImageBrowsingMkdirsPostUnprocessableEntity creates a CreateFoldersInfiniteImageBrowsingMkdirsPostUnprocessableEntity with default headers values
func NewCreateFoldersInfiniteImageBrowsingMkdirsPostUnprocessableEntity() *CreateFoldersInfiniteImageBrowsingMkdirsPostUnprocessableEntity {
	return &CreateFoldersInfiniteImageBrowsingMkdirsPostUnprocessableEntity{}
}

/*
CreateFoldersInfiniteImageBrowsingMkdirsPostUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type CreateFoldersInfiniteImageBrowsingMkdirsPostUnprocessableEntity struct {
	Payload *models.HTTPValidationError
}

// IsSuccess returns true when this create folders infinite image browsing mkdirs post unprocessable entity response has a 2xx status code
func (o *CreateFoldersInfiniteImageBrowsingMkdirsPostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create folders infinite image browsing mkdirs post unprocessable entity response has a 3xx status code
func (o *CreateFoldersInfiniteImageBrowsingMkdirsPostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create folders infinite image browsing mkdirs post unprocessable entity response has a 4xx status code
func (o *CreateFoldersInfiniteImageBrowsingMkdirsPostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this create folders infinite image browsing mkdirs post unprocessable entity response has a 5xx status code
func (o *CreateFoldersInfiniteImageBrowsingMkdirsPostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this create folders infinite image browsing mkdirs post unprocessable entity response a status code equal to that given
func (o *CreateFoldersInfiniteImageBrowsingMkdirsPostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the create folders infinite image browsing mkdirs post unprocessable entity response
func (o *CreateFoldersInfiniteImageBrowsingMkdirsPostUnprocessableEntity) Code() int {
	return 422
}

func (o *CreateFoldersInfiniteImageBrowsingMkdirsPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /infinite_image_browsing/mkdirs][%d] createFoldersInfiniteImageBrowsingMkdirsPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateFoldersInfiniteImageBrowsingMkdirsPostUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /infinite_image_browsing/mkdirs][%d] createFoldersInfiniteImageBrowsingMkdirsPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateFoldersInfiniteImageBrowsingMkdirsPostUnprocessableEntity) GetPayload() *models.HTTPValidationError {
	return o.Payload
}

func (o *CreateFoldersInfiniteImageBrowsingMkdirsPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
