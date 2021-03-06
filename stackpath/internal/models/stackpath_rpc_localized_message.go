// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// StackpathRPCLocalizedMessage stackpath rpc localized message
// swagger:model stackpath.rpc.LocalizedMessage
type StackpathRPCLocalizedMessage struct {
	StackpathRPCLocalizedMessageAllOf1
}

// AtType gets the at type of this subtype
func (m *StackpathRPCLocalizedMessage) AtType() string {
	return "type.stackpathapis.com/stackpath.rpc.LocalizedMessage"
}

// SetAtType sets the at type of this subtype
func (m *StackpathRPCLocalizedMessage) SetAtType(val string) {

}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *StackpathRPCLocalizedMessage) UnmarshalJSON(raw []byte) error {
	var data struct {
		StackpathRPCLocalizedMessageAllOf1
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		AtType string `json:"@type"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result StackpathRPCLocalizedMessage

	if base.AtType != result.AtType() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid @type value: %q", base.AtType)
	}

	result.StackpathRPCLocalizedMessageAllOf1 = data.StackpathRPCLocalizedMessageAllOf1

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m StackpathRPCLocalizedMessage) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		StackpathRPCLocalizedMessageAllOf1
	}{

		StackpathRPCLocalizedMessageAllOf1: m.StackpathRPCLocalizedMessageAllOf1,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		AtType string `json:"@type"`
	}{

		AtType: m.AtType(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this stackpath rpc localized message
func (m *StackpathRPCLocalizedMessage) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with StackpathRPCLocalizedMessageAllOf1
	if err := m.StackpathRPCLocalizedMessageAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *StackpathRPCLocalizedMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackpathRPCLocalizedMessage) UnmarshalBinary(b []byte) error {
	var res StackpathRPCLocalizedMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
