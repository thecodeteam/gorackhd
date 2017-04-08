package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// NewNodesDelRelationsParams creates a new NodesDelRelationsParams object
// with the default values initialized.
func NewNodesDelRelationsParams() *NodesDelRelationsParams {
	var ()
	return &NodesDelRelationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNodesDelRelationsParamsWithTimeout creates a new NodesDelRelationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNodesDelRelationsParamsWithTimeout(timeout time.Duration) *NodesDelRelationsParams {
	var ()
	return &NodesDelRelationsParams{

		timeout: timeout,
	}
}

/*NodesDelRelationsParams contains all the parameters to send to the API endpoint
for the nodes del relations operation typically these are written to a http.Request
*/
type NodesDelRelationsParams struct {

	/*Content
	  A json object with relation types as keys and arrays of node ids as values. The nodes given relation types will be updated to remove the nodes given by id


	*/
	Content models.GenericObj
	/*Identifier
	  The node identifier

	*/
	Identifier string

	timeout time.Duration
}

// WithContent adds the content to the nodes del relations params
func (o *NodesDelRelationsParams) WithContent(content models.GenericObj) *NodesDelRelationsParams {
	o.Content = content
	return o
}

// WithIdentifier adds the identifier to the nodes del relations params
func (o *NodesDelRelationsParams) WithIdentifier(identifier string) *NodesDelRelationsParams {
	o.Identifier = identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *NodesDelRelationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Content); err != nil {
		return err
	}

	// path param identifier
	if err := r.SetPathParam("identifier", o.Identifier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
