package lookups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteLookupsIDParams creates a new DeleteLookupsIDParams object
// with the default values initialized.
func NewDeleteLookupsIDParams() *DeleteLookupsIDParams {
	var ()
	return &DeleteLookupsIDParams{}
}

/*DeleteLookupsIDParams contains all the parameters to send to the API endpoint
for the delete lookups ID operation typically these are written to a http.Request
*/
type DeleteLookupsIDParams struct {

	/*ID
	  id to delete

	*/
	ID string
}

// WithID adds the id to the delete lookups ID params
func (o *DeleteLookupsIDParams) WithID(ID string) *DeleteLookupsIDParams {
	o.ID = ID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLookupsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
