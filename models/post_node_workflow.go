package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

/*PostNodeWorkflow post node workflow

swagger:model post_node_workflow
*/
type PostNodeWorkflow struct {

	/* name
	 */
	Name string `json:"name,omitempty"`

	/* options
	 */
	Options *PostNodeWorkflowOptions `json:"options,omitempty"`
}

// Validate validates this post node workflow
func (m *PostNodeWorkflow) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOptions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostNodeWorkflow) validateOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.Options) { // not required
		return nil
	}

	if m.Options != nil {

		if err := m.Options.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

/*PostNodeWorkflowOptions post node workflow options

swagger:model PostNodeWorkflowOptions
*/
type PostNodeWorkflowOptions struct {

	/* defaults
	 */
	Defaults interface{} `json:"defaults,omitempty"`
}

// Validate validates this post node workflow options
func (m *PostNodeWorkflowOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
