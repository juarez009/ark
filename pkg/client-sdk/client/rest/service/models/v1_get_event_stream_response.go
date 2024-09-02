// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1GetEventStreamResponse v1 get event stream response
// swagger:model v1GetEventStreamResponse
type V1GetEventStreamResponse struct {

	// round failed
	RoundFailed *V1RoundFailed `json:"roundFailed,omitempty"`

	// round finalization
	RoundFinalization *V1RoundFinalizationEvent `json:"roundFinalization,omitempty"`

	// round finalized
	RoundFinalized *V1RoundFinalizedEvent `json:"roundFinalized,omitempty"`

	// round signing
	RoundSigning *V1RoundSigningEvent `json:"roundSigning,omitempty"`

	// round signing nonces generated
	RoundSigningNoncesGenerated *V1RoundSigningNoncesGeneratedEvent `json:"roundSigningNoncesGenerated,omitempty"`
}

// Validate validates this v1 get event stream response
func (m *V1GetEventStreamResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRoundFailed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoundFinalization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoundFinalized(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoundSigning(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoundSigningNoncesGenerated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1GetEventStreamResponse) validateRoundFailed(formats strfmt.Registry) error {

	if swag.IsZero(m.RoundFailed) { // not required
		return nil
	}

	if m.RoundFailed != nil {
		if err := m.RoundFailed.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("roundFailed")
			}
			return err
		}
	}

	return nil
}

func (m *V1GetEventStreamResponse) validateRoundFinalization(formats strfmt.Registry) error {

	if swag.IsZero(m.RoundFinalization) { // not required
		return nil
	}

	if m.RoundFinalization != nil {
		if err := m.RoundFinalization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("roundFinalization")
			}
			return err
		}
	}

	return nil
}

func (m *V1GetEventStreamResponse) validateRoundFinalized(formats strfmt.Registry) error {

	if swag.IsZero(m.RoundFinalized) { // not required
		return nil
	}

	if m.RoundFinalized != nil {
		if err := m.RoundFinalized.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("roundFinalized")
			}
			return err
		}
	}

	return nil
}

func (m *V1GetEventStreamResponse) validateRoundSigning(formats strfmt.Registry) error {

	if swag.IsZero(m.RoundSigning) { // not required
		return nil
	}

	if m.RoundSigning != nil {
		if err := m.RoundSigning.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("roundSigning")
			}
			return err
		}
	}

	return nil
}

func (m *V1GetEventStreamResponse) validateRoundSigningNoncesGenerated(formats strfmt.Registry) error {

	if swag.IsZero(m.RoundSigningNoncesGenerated) { // not required
		return nil
	}

	if m.RoundSigningNoncesGenerated != nil {
		if err := m.RoundSigningNoncesGenerated.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("roundSigningNoncesGenerated")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1GetEventStreamResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1GetEventStreamResponse) UnmarshalBinary(b []byte) error {
	var res V1GetEventStreamResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
