// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// V1VirtualMachineSpecMapEntry A string to virtual machine configuration key/value pair
// swagger:model v1VirtualMachineSpecMapEntry
type V1VirtualMachineSpecMapEntry map[string]V1VirtualMachineSpec

// Validate validates this v1 virtual machine spec map entry
func (m V1VirtualMachineSpecMapEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Required("", "body", V1VirtualMachineSpecMapEntry(m)); err != nil {
		return err
	}

	for k := range m {

		if val, ok := m[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
