package profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetProfilesLibraryIdentifierParams creates a new GetProfilesLibraryIdentifierParams object
// with the default values initialized.
func NewGetProfilesLibraryIdentifierParams() *GetProfilesLibraryIdentifierParams {
	var ()
	return &GetProfilesLibraryIdentifierParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProfilesLibraryIdentifierParamsWithTimeout creates a new GetProfilesLibraryIdentifierParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProfilesLibraryIdentifierParamsWithTimeout(timeout time.Duration) *GetProfilesLibraryIdentifierParams {
	var ()
	return &GetProfilesLibraryIdentifierParams{

		timeout: timeout,
	}
}

/*GetProfilesLibraryIdentifierParams contains all the parameters to send to the API endpoint
for the get profiles library identifier operation typically these are written to a http.Request
*/
type GetProfilesLibraryIdentifierParams struct {

	/*Identifier
	  The profile name.


	*/
	Identifier string

	timeout time.Duration
}

// WithIdentifier adds the identifier to the get profiles library identifier params
func (o *GetProfilesLibraryIdentifierParams) WithIdentifier(identifier string) *GetProfilesLibraryIdentifierParams {
	o.Identifier = identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetProfilesLibraryIdentifierParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
