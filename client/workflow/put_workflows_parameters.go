package workflow

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPutWorkflowsParams creates a new PutWorkflowsParams object
// with the default values initialized.
func NewPutWorkflowsParams() *PutWorkflowsParams {
	var ()
	return &PutWorkflowsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutWorkflowsParamsWithTimeout creates a new PutWorkflowsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutWorkflowsParamsWithTimeout(timeout time.Duration) *PutWorkflowsParams {
	var ()
	return &PutWorkflowsParams{

		timeout: timeout,
	}
}

/*PutWorkflowsParams contains all the parameters to send to the API endpoint
for the put workflows operation typically these are written to a http.Request
*/
type PutWorkflowsParams struct {

	/*Body*/
	Body interface{}

	timeout time.Duration
}

// WithBody adds the body to the put workflows params
func (o *PutWorkflowsParams) WithBody(body interface{}) *PutWorkflowsParams {
	o.Body = body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PutWorkflowsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}