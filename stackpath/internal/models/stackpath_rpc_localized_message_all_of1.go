// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// StackpathRPCLocalizedMessageAllOf1 stackpath Rpc localized message all of1
// swagger:model stackpathRpcLocalizedMessageAllOf1
type StackpathRPCLocalizedMessageAllOf1 struct {

	// locale
	Locale string `json:"locale,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this stackpath Rpc localized message all of1
func (m *StackpathRPCLocalizedMessageAllOf1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StackpathRPCLocalizedMessageAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackpathRPCLocalizedMessageAllOf1) UnmarshalBinary(b []byte) error {
	var res StackpathRPCLocalizedMessageAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
