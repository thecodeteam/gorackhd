package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*NodesPostObmByID OBM settings

swagger:model nodes_post_obm_by_id
*/
type NodesPostObmByID struct {

	/* config

	Required: true
	*/
	Config interface{} `json:"config"`

	/* service

	Required: true
	*/
	Service *string `json:"service"`
}

// Validate validates this nodes post obm by id
func (m *NodesPostObmByID) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateService(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodesPostObmByID) validateConfig(formats strfmt.Registry) error {

	return nil
}

func (m *NodesPostObmByID) validateService(formats strfmt.Registry) error {

	if err := validate.Required("service", "body", m.Service); err != nil {
		return err
	}

	return nil
}
