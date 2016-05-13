package put

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPutTemplatesLibraryIdentifierParams creates a new PutTemplatesLibraryIdentifierParams object
// with the default values initialized.
func NewPutTemplatesLibraryIdentifierParams() *PutTemplatesLibraryIdentifierParams {
	var ()
	return &PutTemplatesLibraryIdentifierParams{}
}

/*PutTemplatesLibraryIdentifierParams contains all the parameters to send to the API endpoint
for the put templates library identifier operation typically these are written to a http.Request
*/
type PutTemplatesLibraryIdentifierParams struct {

	/*Identifier
	  objectid of template


	*/
	Identifier string
}

// WithIdentifier adds the identifier to the put templates library identifier params
func (o *PutTemplatesLibraryIdentifierParams) WithIdentifier(Identifier string) *PutTemplatesLibraryIdentifierParams {
	o.Identifier = Identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PutTemplatesLibraryIdentifierParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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