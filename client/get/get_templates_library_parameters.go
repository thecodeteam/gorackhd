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

// NewGetTemplatesLibraryParams creates a new GetTemplatesLibraryParams object
// with the default values initialized.
func NewGetTemplatesLibraryParams() *GetTemplatesLibraryParams {

	return &GetTemplatesLibraryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTemplatesLibraryParamsWithTimeout creates a new GetTemplatesLibraryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTemplatesLibraryParamsWithTimeout(timeout time.Duration) *GetTemplatesLibraryParams {

	return &GetTemplatesLibraryParams{

		timeout: timeout,
	}
}

/*GetTemplatesLibraryParams contains all the parameters to send to the API endpoint
for the get templates library operation typically these are written to a http.Request
*/
type GetTemplatesLibraryParams struct {
	timeout time.Duration
}

// WriteToRequest writes these params to a swagger request
func (o *GetTemplatesLibraryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
