package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*GetUserObj get user obj

swagger:model get_user_obj
*/
type GetUserObj struct {

	/* role
	 */
	Role string `json:"role,omitempty"`

	/* username
	 */
	Username string `json:"username,omitempty"`
}

// Validate validates this get user obj
func (m *GetUserObj) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
