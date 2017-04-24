package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// IdentityContext Context describing a pair of source and destination identity
// swagger:model IdentityContext
type IdentityContext struct {

	// from
	From Labels `json:"from"`

	// List of l4 egress rules
	L4Egress []*Port `json:"l4-egress"`

	// List of l4 ingress rules
	L4Ingress []*Port `json:"l4-ingress"`

	// to
	To Labels `json:"to"`
}

// Validate validates this identity context
func (m *IdentityContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateL4Egress(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateL4Ingress(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdentityContext) validateL4Egress(formats strfmt.Registry) error {

	if swag.IsZero(m.L4Egress) { // not required
		return nil
	}

	for i := 0; i < len(m.L4Egress); i++ {

		if swag.IsZero(m.L4Egress[i]) { // not required
			continue
		}

		if m.L4Egress[i] != nil {

			if err := m.L4Egress[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("l4-egress" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IdentityContext) validateL4Ingress(formats strfmt.Registry) error {

	if swag.IsZero(m.L4Ingress) { // not required
		return nil
	}

	for i := 0; i < len(m.L4Ingress); i++ {

		if swag.IsZero(m.L4Ingress[i]) { // not required
			continue
		}

		if m.L4Ingress[i] != nil {

			if err := m.L4Ingress[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("l4-ingress" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}
