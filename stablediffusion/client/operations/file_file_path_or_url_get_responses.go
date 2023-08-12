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

// FileFilePathOrURLGetReader is a Reader for the FileFilePathOrURLGet structure.
type FileFilePathOrURLGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FileFilePathOrURLGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFileFilePathOrURLGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewFileFilePathOrURLGetUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFileFilePathOrURLGetOK creates a FileFilePathOrURLGetOK with default headers values
func NewFileFilePathOrURLGetOK() *FileFilePathOrURLGetOK {
	return &FileFilePathOrURLGetOK{}
}

/*
FileFilePathOrURLGetOK describes a response with status code 200, with default header values.

Successful Response
*/
type FileFilePathOrURLGetOK struct {
	Payload interface{}
}

// IsSuccess returns true when this file file path or Url get o k response has a 2xx status code
func (o *FileFilePathOrURLGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this file file path or Url get o k response has a 3xx status code
func (o *FileFilePathOrURLGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this file file path or Url get o k response has a 4xx status code
func (o *FileFilePathOrURLGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this file file path or Url get o k response has a 5xx status code
func (o *FileFilePathOrURLGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this file file path or Url get o k response a status code equal to that given
func (o *FileFilePathOrURLGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the file file path or Url get o k response
func (o *FileFilePathOrURLGetOK) Code() int {
	return 200
}

func (o *FileFilePathOrURLGetOK) Error() string {
	return fmt.Sprintf("[GET /file={path_or_url}][%d] fileFilePathOrUrlGetOK  %+v", 200, o.Payload)
}

func (o *FileFilePathOrURLGetOK) String() string {
	return fmt.Sprintf("[GET /file={path_or_url}][%d] fileFilePathOrUrlGetOK  %+v", 200, o.Payload)
}

func (o *FileFilePathOrURLGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *FileFilePathOrURLGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFileFilePathOrURLGetUnprocessableEntity creates a FileFilePathOrURLGetUnprocessableEntity with default headers values
func NewFileFilePathOrURLGetUnprocessableEntity() *FileFilePathOrURLGetUnprocessableEntity {
	return &FileFilePathOrURLGetUnprocessableEntity{}
}

/*
FileFilePathOrURLGetUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type FileFilePathOrURLGetUnprocessableEntity struct {
	Payload *models.HTTPValidationError
}

// IsSuccess returns true when this file file path or Url get unprocessable entity response has a 2xx status code
func (o *FileFilePathOrURLGetUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this file file path or Url get unprocessable entity response has a 3xx status code
func (o *FileFilePathOrURLGetUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this file file path or Url get unprocessable entity response has a 4xx status code
func (o *FileFilePathOrURLGetUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this file file path or Url get unprocessable entity response has a 5xx status code
func (o *FileFilePathOrURLGetUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this file file path or Url get unprocessable entity response a status code equal to that given
func (o *FileFilePathOrURLGetUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the file file path or Url get unprocessable entity response
func (o *FileFilePathOrURLGetUnprocessableEntity) Code() int {
	return 422
}

func (o *FileFilePathOrURLGetUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /file={path_or_url}][%d] fileFilePathOrUrlGetUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *FileFilePathOrURLGetUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /file={path_or_url}][%d] fileFilePathOrUrlGetUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *FileFilePathOrURLGetUnprocessableEntity) GetPayload() *models.HTTPValidationError {
	return o.Payload
}

func (o *FileFilePathOrURLGetUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
