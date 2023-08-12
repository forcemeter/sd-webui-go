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

// SDModelItem SDModelItem
//
// swagger:model SDModelItem
type SDModelItem struct {

	// Config file
	Config string `json:"config,omitempty"`

	// Filename
	// Required: true
	Filename *string `json:"filename"`

	// Short hash
	Hash string `json:"hash,omitempty"`

	// Model Name
	// Required: true
	ModelName *string `json:"model_name"`

	// sha256 hash
	Sha256 string `json:"sha256,omitempty"`

	// Title
	// Required: true
	Title *string `json:"title"`
}

// Validate validates this s d model item
func (m *SDModelItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilename(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModelName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SDModelItem) validateFilename(formats strfmt.Registry) error {

	if err := validate.Required("filename", "body", m.Filename); err != nil {
		return err
	}

	return nil
}

func (m *SDModelItem) validateModelName(formats strfmt.Registry) error {

	if err := validate.Required("model_name", "body", m.ModelName); err != nil {
		return err
	}

	return nil
}

func (m *SDModelItem) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this s d model item based on context it is used
func (m *SDModelItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SDModelItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SDModelItem) UnmarshalBinary(b []byte) error {
	var res SDModelItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
