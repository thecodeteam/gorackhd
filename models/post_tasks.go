// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PostTasks post tasks
// swagger:model post_tasks

type PostTasks struct {

	// identifier
	Identifier string `json:"identifier,omitempty"`

	// tasks
	Tasks PostTasksTasks `json:"tasks"`
}

/* polymorph post_tasks identifier false */

/* polymorph post_tasks tasks false */

// Validate validates this post tasks
func (m *PostTasks) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PostTasks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostTasks) UnmarshalBinary(b []byte) error {
	var res PostTasks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}