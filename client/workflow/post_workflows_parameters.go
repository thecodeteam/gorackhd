package workflow

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostWorkflowsParams creates a new PostWorkflowsParams object
// with the default values initialized.
func NewPostWorkflowsParams() *PostWorkflowsParams {
	var ()
	return &PostWorkflowsParams{}
}

/*PostWorkflowsParams contains all the parameters to send to the API endpoint
for the post workflows operation typically these are written to a http.Request
*/
type PostWorkflowsParams struct {

	/*Body*/
	Body interface{}
	/*Name
	  The injectable Graph name


	*/
	Name string
}

// WithBody adds the body to the post workflows params
func (o *PostWorkflowsParams) WithBody(Body interface{}) *PostWorkflowsParams {
	o.Body = Body
	return o
}

// WithName adds the name to the post workflows params
func (o *PostWorkflowsParams) WithName(Name string) *PostWorkflowsParams {
	o.Name = Name
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostWorkflowsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
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
