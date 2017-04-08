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

// NewHooksPostParams creates a new HooksPostParams object
// with the default values initialized.
func NewHooksPostParams() *HooksPostParams {
	var ()
	return &HooksPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHooksPostParamsWithTimeout creates a new HooksPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHooksPostParamsWithTimeout(timeout time.Duration) *HooksPostParams {
	var ()
	return &HooksPostParams{

		timeout: timeout,
	}
}

/*HooksPostParams contains all the parameters to send to the API endpoint
for the hooks post operation typically these are written to a http.Request
*/
type HooksPostParams struct {

	/*Body
	  configuration hook to be created

	*/
	Body *models.Hooks20HookPost

	timeout time.Duration
}

// WithBody adds the body to the hooks post params
func (o *HooksPostParams) WithBody(body *models.Hooks20HookPost) *HooksPostParams {
	o.Body = body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *HooksPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.Hooks20HookPost)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
