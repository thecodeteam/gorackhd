package patch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPatchPollersIdentifierParams creates a new PatchPollersIdentifierParams object
// with the default values initialized.
func NewPatchPollersIdentifierParams() *PatchPollersIdentifierParams {
	var ()
	return &PatchPollersIdentifierParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchPollersIdentifierParamsWithTimeout creates a new PatchPollersIdentifierParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchPollersIdentifierParamsWithTimeout(timeout time.Duration) *PatchPollersIdentifierParams {
	var ()
	return &PatchPollersIdentifierParams{

		timeout: timeout,
	}
}

/*PatchPollersIdentifierParams contains all the parameters to send to the API endpoint
for the patch pollers identifier operation typically these are written to a http.Request
*/
type PatchPollersIdentifierParams struct {

	/*Identifier
	  poller identifier


	*/
	Identifier string

	timeout time.Duration
}

// WithIdentifier adds the identifier to the patch pollers identifier params
func (o *PatchPollersIdentifierParams) WithIdentifier(identifier string) *PatchPollersIdentifierParams {
	o.Identifier = identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PatchPollersIdentifierParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
