package get

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetObmsLibraryIdentifierParams creates a new GetObmsLibraryIdentifierParams object
// with the default values initialized.
func NewGetObmsLibraryIdentifierParams() *GetObmsLibraryIdentifierParams {
	var ()
	return &GetObmsLibraryIdentifierParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetObmsLibraryIdentifierParamsWithTimeout creates a new GetObmsLibraryIdentifierParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetObmsLibraryIdentifierParamsWithTimeout(timeout time.Duration) *GetObmsLibraryIdentifierParams {
	var ()
	return &GetObmsLibraryIdentifierParams{

		timeout: timeout,
	}
}

/*GetObmsLibraryIdentifierParams contains all the parameters to send to the API endpoint
for the get obms library identifier operation typically these are written to a http.Request
*/
type GetObmsLibraryIdentifierParams struct {

	/*Identifier
	  The obm service name.


	*/
	Identifier string

	timeout time.Duration
}

// WithIdentifier adds the identifier to the get obms library identifier params
func (o *GetObmsLibraryIdentifierParams) WithIdentifier(identifier string) *GetObmsLibraryIdentifierParams {
	o.Identifier = identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetObmsLibraryIdentifierParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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