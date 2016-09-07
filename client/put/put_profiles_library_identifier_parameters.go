package put

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPutProfilesLibraryIdentifierParams creates a new PutProfilesLibraryIdentifierParams object
// with the default values initialized.
func NewPutProfilesLibraryIdentifierParams() *PutProfilesLibraryIdentifierParams {
	var ()
	return &PutProfilesLibraryIdentifierParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutProfilesLibraryIdentifierParamsWithTimeout creates a new PutProfilesLibraryIdentifierParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutProfilesLibraryIdentifierParamsWithTimeout(timeout time.Duration) *PutProfilesLibraryIdentifierParams {
	var ()
	return &PutProfilesLibraryIdentifierParams{

		timeout: timeout,
	}
}

/*PutProfilesLibraryIdentifierParams contains all the parameters to send to the API endpoint
for the put profiles library identifier operation typically these are written to a http.Request
*/
type PutProfilesLibraryIdentifierParams struct {

	/*Identifier
	  The profile name.


	*/
	Identifier string

	timeout time.Duration
}

// WithIdentifier adds the identifier to the put profiles library identifier params
func (o *PutProfilesLibraryIdentifierParams) WithIdentifier(identifier string) *PutProfilesLibraryIdentifierParams {
	o.Identifier = identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PutProfilesLibraryIdentifierParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
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
