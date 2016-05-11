package skus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSkusIdentifierNodesParams creates a new GetSkusIdentifierNodesParams object
// with the default values initialized.
func NewGetSkusIdentifierNodesParams() *GetSkusIdentifierNodesParams {
	var ()
	return &GetSkusIdentifierNodesParams{}
}

/*GetSkusIdentifierNodesParams contains all the parameters to send to the API endpoint
for the get skus identifier nodes operation typically these are written to a http.Request
*/
type GetSkusIdentifierNodesParams struct {

	/*Identifier
	  The profile name.


	*/
	Identifier string
}

// WithIdentifier adds the identifier to the get skus identifier nodes params
func (o *GetSkusIdentifierNodesParams) WithIdentifier(Identifier string) *GetSkusIdentifierNodesParams {
	o.Identifier = Identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetSkusIdentifierNodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	// path param identifier
	if err := r.SetPathParam("identifier", o.Identifier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
