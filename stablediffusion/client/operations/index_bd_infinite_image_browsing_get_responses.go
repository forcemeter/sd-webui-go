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

// IndexBdInfiniteImageBrowsingGetReader is a Reader for the IndexBdInfiniteImageBrowsingGet structure.
type IndexBdInfiniteImageBrowsingGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IndexBdInfiniteImageBrowsingGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIndexBdInfiniteImageBrowsingGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIndexBdInfiniteImageBrowsingGetOK creates a IndexBdInfiniteImageBrowsingGetOK with default headers values
func NewIndexBdInfiniteImageBrowsingGetOK() *IndexBdInfiniteImageBrowsingGetOK {
	return &IndexBdInfiniteImageBrowsingGetOK{}
}

/*
IndexBdInfiniteImageBrowsingGetOK describes a response with status code 200, with default header values.

Successful Response
*/
type IndexBdInfiniteImageBrowsingGetOK struct {
	Payload interface{}
}

// IsSuccess returns true when this index bd infinite image browsing get o k response has a 2xx status code
func (o *IndexBdInfiniteImageBrowsingGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this index bd infinite image browsing get o k response has a 3xx status code
func (o *IndexBdInfiniteImageBrowsingGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this index bd infinite image browsing get o k response has a 4xx status code
func (o *IndexBdInfiniteImageBrowsingGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this index bd infinite image browsing get o k response has a 5xx status code
func (o *IndexBdInfiniteImageBrowsingGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this index bd infinite image browsing get o k response a status code equal to that given
func (o *IndexBdInfiniteImageBrowsingGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the index bd infinite image browsing get o k response
func (o *IndexBdInfiniteImageBrowsingGetOK) Code() int {
	return 200
}

func (o *IndexBdInfiniteImageBrowsingGetOK) Error() string {
	return fmt.Sprintf("[GET /infinite_image_browsing][%d] indexBdInfiniteImageBrowsingGetOK  %+v", 200, o.Payload)
}

func (o *IndexBdInfiniteImageBrowsingGetOK) String() string {
	return fmt.Sprintf("[GET /infinite_image_browsing][%d] indexBdInfiniteImageBrowsingGetOK  %+v", 200, o.Payload)
}

func (o *IndexBdInfiniteImageBrowsingGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *IndexBdInfiniteImageBrowsingGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
