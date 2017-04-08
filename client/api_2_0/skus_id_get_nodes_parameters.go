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

// NewSkusIDGetNodesParams creates a new SkusIDGetNodesParams object
// with the default values initialized.
func NewSkusIDGetNodesParams() *SkusIDGetNodesParams {
	var ()
	return &SkusIDGetNodesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSkusIDGetNodesParamsWithTimeout creates a new SkusIDGetNodesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSkusIDGetNodesParamsWithTimeout(timeout time.Duration) *SkusIDGetNodesParams {
	var ()
	return &SkusIDGetNodesParams{

		timeout: timeout,
	}
}

/*SkusIDGetNodesParams contains all the parameters to send to the API endpoint
for the skus Id get nodes operation typically these are written to a http.Request
*/
type SkusIDGetNodesParams struct {

	/*Identifier
	  The SKU identifier

	*/
	Identifier string

	timeout time.Duration
}

// WithIdentifier adds the identifier to the skus Id get nodes params
func (o *SkusIDGetNodesParams) WithIdentifier(identifier string) *SkusIDGetNodesParams {
	o.Identifier = identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *SkusIDGetNodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
