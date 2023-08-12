// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InterrogateRequest InterrogateRequest
//
// swagger:model InterrogateRequest
type InterrogateRequest struct {

	// Image
	//
	// Image to work on, must be a Base64 string containing the image's data.
	Image string `json:"image,omitempty"`

	// Model
	//
	// The interrogate model used.
	Model *string `json:"model,omitempty"`
}

// Validate validates this interrogate request
func (m *InterrogateRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this interrogate request based on context it is used
func (m *InterrogateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InterrogateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InterrogateRequest) UnmarshalBinary(b []byte) error {
	var res InterrogateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
