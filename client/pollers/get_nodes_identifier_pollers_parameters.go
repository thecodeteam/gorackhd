package pollers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetNodesIdentifierPollersParams creates a new GetNodesIdentifierPollersParams object
// with the default values initialized.
func NewGetNodesIdentifierPollersParams() *GetNodesIdentifierPollersParams {
	var ()
	return &GetNodesIdentifierPollersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNodesIdentifierPollersParamsWithTimeout creates a new GetNodesIdentifierPollersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNodesIdentifierPollersParamsWithTimeout(timeout time.Duration) *GetNodesIdentifierPollersParams {
	var ()
	return &GetNodesIdentifierPollersParams{

		timeout: timeout,
	}
}

/*GetNodesIdentifierPollersParams contains all the parameters to send to the API endpoint
for the get nodes identifier pollers operation typically these are written to a http.Request
*/
type GetNodesIdentifierPollersParams struct {

	/*Identifier
	  Mac addresses and unique aliases to identify the node by, |
	expect a string or an array of strings.


	*/
	Identifier string

	timeout time.Duration
}

// WithIdentifier adds the identifier to the get nodes identifier pollers params
func (o *GetNodesIdentifierPollersParams) WithIdentifier(identifier string) *GetNodesIdentifierPollersParams {
	o.Identifier = identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetNodesIdentifierPollersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
