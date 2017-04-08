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

// NewGetTasksByIDParams creates a new GetTasksByIDParams object
// with the default values initialized.
func NewGetTasksByIDParams() *GetTasksByIDParams {
	var ()
	return &GetTasksByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTasksByIDParamsWithTimeout creates a new GetTasksByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTasksByIDParamsWithTimeout(timeout time.Duration) *GetTasksByIDParams {
	var ()
	return &GetTasksByIDParams{

		timeout: timeout,
	}
}

/*GetTasksByIDParams contains all the parameters to send to the API endpoint
for the get tasks by Id operation typically these are written to a http.Request
*/
type GetTasksByIDParams struct {

	/*Identifier
	  The task identifier

	*/
	Identifier string

	timeout time.Duration
}

// WithIdentifier adds the identifier to the get tasks by Id params
func (o *GetTasksByIDParams) WithIdentifier(identifier string) *GetTasksByIDParams {
	o.Identifier = identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetTasksByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
