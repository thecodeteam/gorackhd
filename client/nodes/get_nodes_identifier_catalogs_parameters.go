package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetNodesIdentifierCatalogsParams creates a new GetNodesIdentifierCatalogsParams object
// with the default values initialized.
func NewGetNodesIdentifierCatalogsParams() *GetNodesIdentifierCatalogsParams {
	var ()
	return &GetNodesIdentifierCatalogsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNodesIdentifierCatalogsParamsWithTimeout creates a new GetNodesIdentifierCatalogsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNodesIdentifierCatalogsParamsWithTimeout(timeout time.Duration) *GetNodesIdentifierCatalogsParams {
	var ()
	return &GetNodesIdentifierCatalogsParams{

		timeout: timeout,
	}
}

/*GetNodesIdentifierCatalogsParams contains all the parameters to send to the API endpoint
for the get nodes identifier catalogs operation typically these are written to a http.Request
*/
type GetNodesIdentifierCatalogsParams struct {

	/*Identifier
	  Mac addresses and unique aliases to identify the node by, |
	expect a string or an array of strings.


	*/
	Identifier string

	timeout time.Duration
}

// WithIdentifier adds the identifier to the get nodes identifier catalogs params
func (o *GetNodesIdentifierCatalogsParams) WithIdentifier(identifier string) *GetNodesIdentifierCatalogsParams {
	o.Identifier = identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetNodesIdentifierCatalogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
