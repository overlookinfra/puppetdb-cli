// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServerError server error
//
// swagger:model ServerError
type ServerError struct {

	// cause0
	Cause0 string `json:"cause0,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this server error
func (m *ServerError) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServerError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerError) UnmarshalBinary(b []byte) error {
	var res ServerError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
