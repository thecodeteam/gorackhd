package post

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostNodesParams creates a new PostNodesParams object
// with the default values initialized.
func NewPostNodesParams() *PostNodesParams {
	var ()
	return &PostNodesParams{}
}

/*PostNodesParams contains all the parameters to send to the API endpoint
for the post nodes operation typically these are written to a http.Request
*/
type PostNodesParams struct {

	/*Identifiers
	  Mac addresses and unique aliases to identify the node by, |
	expect a string or an array of strings.


	*/
	Identifiers interface{}
}

// WithIdentifiers adds the identifiers to the post nodes params
func (o *PostNodesParams) WithIdentifiers(Identifiers interface{}) *PostNodesParams {
	o.Identifiers = Identifiers
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostNodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	if err := r.SetBodyParam(o.Identifiers); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
