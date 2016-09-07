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

// NewPostNodesIdentifierWorkflowsParams creates a new PostNodesIdentifierWorkflowsParams object
// with the default values initialized.
func NewPostNodesIdentifierWorkflowsParams() *PostNodesIdentifierWorkflowsParams {
	var ()
	return &PostNodesIdentifierWorkflowsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostNodesIdentifierWorkflowsParamsWithTimeout creates a new PostNodesIdentifierWorkflowsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostNodesIdentifierWorkflowsParamsWithTimeout(timeout time.Duration) *PostNodesIdentifierWorkflowsParams {
	var ()
	return &PostNodesIdentifierWorkflowsParams{

		timeout: timeout,
	}
}

/*PostNodesIdentifierWorkflowsParams contains all the parameters to send to the API endpoint
for the post nodes identifier workflows operation typically these are written to a http.Request
*/
type PostNodesIdentifierWorkflowsParams struct {

	/*Body*/
	Body interface{}
	/*Identifier
	  The node unique identifier


	*/
	Identifier string
	/*Name
	  The injectable Graph name


	*/
	Name string

	timeout time.Duration
}

// WithBody adds the body to the post nodes identifier workflows params
func (o *PostNodesIdentifierWorkflowsParams) WithBody(body interface{}) *PostNodesIdentifierWorkflowsParams {
	o.Body = body
	return o
}

// WithIdentifier adds the identifier to the post nodes identifier workflows params
func (o *PostNodesIdentifierWorkflowsParams) WithIdentifier(identifier string) *PostNodesIdentifierWorkflowsParams {
	o.Identifier = identifier
	return o
}

// WithName adds the name to the post nodes identifier workflows params
func (o *PostNodesIdentifierWorkflowsParams) WithName(name string) *PostNodesIdentifierWorkflowsParams {
	o.Name = name
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostNodesIdentifierWorkflowsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param identifier
	if err := r.SetPathParam("identifier", o.Identifier); err != nil {
		return err
	}

	// query param name
	qrName := o.Name
	qName := qrName
	if qName != "" {
		if err := r.SetQueryParam("name", qName); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
