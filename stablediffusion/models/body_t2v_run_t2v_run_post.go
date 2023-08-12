// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BodyT2vRunT2vRunPost Body_t2v_run_t2v_run_post
//
// swagger:model Body_t2v_run_t2v_run_post
type BodyT2vRunT2vRunPost struct {

	// Inpainting Image
	// Format: binary
	InpaintingImage io.ReadCloser `json:"inpainting_image,omitempty"`

	// Vid2Vid Input
	// Format: binary
	Vid2vidInput io.ReadCloser `json:"vid2vid_input,omitempty"`
}

// Validate validates this body t2v run t2v run post
func (m *BodyT2vRunT2vRunPost) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this body t2v run t2v run post based on context it is used
func (m *BodyT2vRunT2vRunPost) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BodyT2vRunT2vRunPost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BodyT2vRunT2vRunPost) UnmarshalBinary(b []byte) error {
	var res BodyT2vRunT2vRunPost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
