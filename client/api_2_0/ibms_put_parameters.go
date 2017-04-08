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

// NewIbmsPutParams creates a new IbmsPutParams object
// with the default values initialized.
func NewIbmsPutParams() *IbmsPutParams {
	var ()
	return &IbmsPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIbmsPutParamsWithTimeout creates a new IbmsPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIbmsPutParamsWithTimeout(timeout time.Duration) *IbmsPutParams {
	var ()
	return &IbmsPutParams{

		timeout: timeout,
	}
}

/*IbmsPutParams contains all the parameters to send to the API endpoint
for the ibms put operation typically these are written to a http.Request
*/
type IbmsPutParams struct {

	/*Body
	  The IBM settings information to create

	*/
	Body *models.SSHIbmServiceIbm

	timeout time.Duration
}

// WithBody adds the body to the ibms put params
func (o *IbmsPutParams) WithBody(body *models.SSHIbmServiceIbm) *IbmsPutParams {
	o.Body = body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *IbmsPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.SSHIbmServiceIbm)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
