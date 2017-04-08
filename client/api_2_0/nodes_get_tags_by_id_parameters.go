package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewNodesGetTagsByIDParams creates a new NodesGetTagsByIDParams object
// with the default values initialized.
func NewNodesGetTagsByIDParams() *NodesGetTagsByIDParams {
	var ()
	return &NodesGetTagsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNodesGetTagsByIDParamsWithTimeout creates a new NodesGetTagsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNodesGetTagsByIDParamsWithTimeout(timeout time.Duration) *NodesGetTagsByIDParams {
	var ()
	return &NodesGetTagsByIDParams{

		timeout: timeout,
	}
}

/*NodesGetTagsByIDParams contains all the parameters to send to the API endpoint
for the nodes get tags by Id operation typically these are written to a http.Request
*/
type NodesGetTagsByIDParams struct {

	/*Identifier
	  The node identifier

	*/
	Identifier string

	timeout time.Duration
}

// WithIdentifier adds the identifier to the nodes get tags by Id params
func (o *NodesGetTagsByIDParams) WithIdentifier(identifier string) *NodesGetTagsByIDParams {
	o.Identifier = identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *NodesGetTagsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
