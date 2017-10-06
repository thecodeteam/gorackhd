// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserObj user obj
// swagger:model user_obj

type UserObj struct {

	// password
	Password string `json:"password,omitempty"`

	// role
	Role string `json:"role,omitempty"`

	// username
	// Pattern: ^[A-Za-z][0-9A-Za-z._-]+$
	Username string `json:"username,omitempty"`
}

/* polymorph user_obj password false */

/* polymorph user_obj role false */

/* polymorph user_obj username false */

// Validate validates this user obj
func (m *UserObj) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUsername(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserObj) validateUsername(formats strfmt.Registry) error {

	if swag.IsZero(m.Username) { // not required
		return nil
	}

	if err := validate.Pattern("username", "body", string(m.Username), `^[A-Za-z][0-9A-Za-z._-]+$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserObj) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserObj) UnmarshalBinary(b []byte) error {
	var res UserObj
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}