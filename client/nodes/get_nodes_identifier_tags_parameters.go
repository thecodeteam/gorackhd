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

// NewGetNodesIdentifierTagsParams creates a new GetNodesIdentifierTagsParams object
// with the default values initialized.
func NewGetNodesIdentifierTagsParams() *GetNodesIdentifierTagsParams {
	var ()
	return &GetNodesIdentifierTagsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNodesIdentifierTagsParamsWithTimeout creates a new GetNodesIdentifierTagsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNodesIdentifierTagsParamsWithTimeout(timeout time.Duration) *GetNodesIdentifierTagsParams {
	var ()
	return &GetNodesIdentifierTagsParams{

		timeout: timeout,
	}
}

/*GetNodesIdentifierTagsParams contains all the parameters to send to the API endpoint
for the get nodes identifier tags operation typically these are written to a http.Request
*/
type GetNodesIdentifierTagsParams struct {

	/*Identifier
	  Mac addresses and unique aliases to identify the node by


	*/
	Identifier string

	timeout time.Duration
}

// WithIdentifier adds the identifier to the get nodes identifier tags params
func (o *GetNodesIdentifierTagsParams) WithIdentifier(identifier string) *GetNodesIdentifierTagsParams {
	o.Identifier = identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetNodesIdentifierTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
