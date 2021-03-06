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

// StackpathRPCRequestInfo stackpath rpc request info
// swagger:model stackpath.rpc.RequestInfo
type StackpathRPCRequestInfo struct {
	StackpathRPCRequestInfoAllOf1
}

// AtType gets the at type of this subtype
func (m *StackpathRPCRequestInfo) AtType() string {
	return "type.stackpathapis.com/stackpath.rpc.RequestInfo"
}

// SetAtType sets the at type of this subtype
func (m *StackpathRPCRequestInfo) SetAtType(val string) {

}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *StackpathRPCRequestInfo) UnmarshalJSON(raw []byte) error {
	var data struct {
		StackpathRPCRequestInfoAllOf1
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

	var result StackpathRPCRequestInfo

	if base.AtType != result.AtType() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid @type value: %q", base.AtType)
	}

	result.StackpathRPCRequestInfoAllOf1 = data.StackpathRPCRequestInfoAllOf1

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m StackpathRPCRequestInfo) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		StackpathRPCRequestInfoAllOf1
	}{

		StackpathRPCRequestInfoAllOf1: m.StackpathRPCRequestInfoAllOf1,
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

// Validate validates this stackpath rpc request info
func (m *StackpathRPCRequestInfo) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with StackpathRPCRequestInfoAllOf1
	if err := m.StackpathRPCRequestInfoAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *StackpathRPCRequestInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackpathRPCRequestInfo) UnmarshalBinary(b []byte) error {
	var res StackpathRPCRequestInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
