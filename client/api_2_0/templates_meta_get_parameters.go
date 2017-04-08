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

// NewTemplatesMetaGetParams creates a new TemplatesMetaGetParams object
// with the default values initialized.
func NewTemplatesMetaGetParams() *TemplatesMetaGetParams {
	var ()
	return &TemplatesMetaGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTemplatesMetaGetParamsWithTimeout creates a new TemplatesMetaGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTemplatesMetaGetParamsWithTimeout(timeout time.Duration) *TemplatesMetaGetParams {
	var ()
	return &TemplatesMetaGetParams{

		timeout: timeout,
	}
}

/*TemplatesMetaGetParams contains all the parameters to send to the API endpoint
for the templates meta get operation typically these are written to a http.Request
*/
type TemplatesMetaGetParams struct {

	/*Sort
	  Query string specifying properties to sort with

	*/
	Sort *string

	timeout time.Duration
}

// WithSort adds the sort to the templates meta get params
func (o *TemplatesMetaGetParams) WithSort(sort *string) *TemplatesMetaGetParams {
	o.Sort = sort
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *TemplatesMetaGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Sort != nil {

		// query param sort
		var qrSort string
		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {
			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
