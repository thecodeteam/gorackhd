package catalogs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetNodesIdentifierCatalogsSourceParams creates a new GetNodesIdentifierCatalogsSourceParams object
// with the default values initialized.
func NewGetNodesIdentifierCatalogsSourceParams() *GetNodesIdentifierCatalogsSourceParams {
	var ()
	return &GetNodesIdentifierCatalogsSourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNodesIdentifierCatalogsSourceParamsWithTimeout creates a new GetNodesIdentifierCatalogsSourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNodesIdentifierCatalogsSourceParamsWithTimeout(timeout time.Duration) *GetNodesIdentifierCatalogsSourceParams {
	var ()
	return &GetNodesIdentifierCatalogsSourceParams{

		timeout: timeout,
	}
}

/*GetNodesIdentifierCatalogsSourceParams contains all the parameters to send to the API endpoint
for the get nodes identifier catalogs source operation typically these are written to a http.Request
*/
type GetNodesIdentifierCatalogsSourceParams struct {

	/*Identifier
	  Mac addresses and unique aliases to identify the node by, |
	expect a string or an array of strings.


	*/
	Identifier string
	/*Source
	  Source catalog name to fetch


	*/
	Source string

	timeout time.Duration
}

// WithIdentifier adds the identifier to the get nodes identifier catalogs source params
func (o *GetNodesIdentifierCatalogsSourceParams) WithIdentifier(identifier string) *GetNodesIdentifierCatalogsSourceParams {
	o.Identifier = identifier
	return o
}

// WithSource adds the source to the get nodes identifier catalogs source params
func (o *GetNodesIdentifierCatalogsSourceParams) WithSource(source string) *GetNodesIdentifierCatalogsSourceParams {
	o.Source = source
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetNodesIdentifierCatalogsSourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param identifier
	if err := r.SetPathParam("identifier", o.Identifier); err != nil {
		return err
	}

	// path param source
	if err := r.SetPathParam("source", o.Source); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
