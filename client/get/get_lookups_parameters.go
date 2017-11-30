package get

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetLookupsParams creates a new GetLookupsParams object
// with the default values initialized.
func NewGetLookupsParams() *GetLookupsParams {
	var ()
	return &GetLookupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLookupsParamsWithTimeout creates a new GetLookupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLookupsParamsWithTimeout(timeout time.Duration) *GetLookupsParams {
	var ()
	return &GetLookupsParams{

		timeout: timeout,
	}
}

/*GetLookupsParams contains all the parameters to send to the API endpoint
for the get lookups operation typically these are written to a http.Request
*/
type GetLookupsParams struct {

	/*Q
	  query object to pass through to waterline.

	*/
	Q *string

	timeout time.Duration
}

// WithQ adds the q to the get lookups params
func (o *GetLookupsParams) WithQ(q *string) *GetLookupsParams {
	o.Q = q
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetLookupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}