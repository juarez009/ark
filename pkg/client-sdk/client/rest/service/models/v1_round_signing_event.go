// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1RoundSigningEvent v1 round signing event
// swagger:model v1RoundSigningEvent
type V1RoundSigningEvent struct {

	// cosigners pubkeys
	CosignersPubkeys []string `json:"cosignersPubkeys"`

	// id
	ID string `json:"id,omitempty"`

	// unsigned tree
	UnsignedTree *V1Tree `json:"unsignedTree,omitempty"`
}

// Validate validates this v1 round signing event
func (m *V1RoundSigningEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUnsignedTree(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1RoundSigningEvent) validateUnsignedTree(formats strfmt.Registry) error {

	if swag.IsZero(m.UnsignedTree) { // not required
		return nil
	}

	if m.UnsignedTree != nil {
		if err := m.UnsignedTree.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unsignedTree")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1RoundSigningEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1RoundSigningEvent) UnmarshalBinary(b []byte) error {
	var res V1RoundSigningEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
