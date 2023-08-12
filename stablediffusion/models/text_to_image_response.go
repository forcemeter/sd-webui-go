// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TextToImageResponse TextToImageResponse
//
// swagger:model TextToImageResponse
type TextToImageResponse struct {

	// Image
	//
	// The generated image in base64 format.
	Images []string `json:"images"`

	// Info
	// Required: true
	Info *string `json:"info"`

	// Parameters
	// Required: true
	Parameters interface{} `json:"parameters"`
}

// Validate validates this text to image response
func (m *TextToImageResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TextToImageResponse) validateInfo(formats strfmt.Registry) error {

	if err := validate.Required("info", "body", m.Info); err != nil {
		return err
	}

	return nil
}

func (m *TextToImageResponse) validateParameters(formats strfmt.Registry) error {

	if m.Parameters == nil {
		return errors.Required("parameters", "body", nil)
	}

	return nil
}

// ContextValidate validates this text to image response based on context it is used
func (m *TextToImageResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TextToImageResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TextToImageResponse) UnmarshalBinary(b []byte) error {
	var res TextToImageResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
