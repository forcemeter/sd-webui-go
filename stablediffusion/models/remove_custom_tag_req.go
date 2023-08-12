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

// RemoveCustomTagReq RemoveCustomTagReq
//
// swagger:model RemoveCustomTagReq
type RemoveCustomTagReq struct {

	// Tag Id
	// Required: true
	TagID *string `json:"tag_id"`
}

// Validate validates this remove custom tag req
func (m *RemoveCustomTagReq) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTagID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemoveCustomTagReq) validateTagID(formats strfmt.Registry) error {

	if err := validate.Required("tag_id", "body", m.TagID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this remove custom tag req based on context it is used
func (m *RemoveCustomTagReq) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RemoveCustomTagReq) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemoveCustomTagReq) UnmarshalBinary(b []byte) error {
	var res RemoveCustomTagReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
