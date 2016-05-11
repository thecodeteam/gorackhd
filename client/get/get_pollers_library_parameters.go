package get

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPollersLibraryParams creates a new GetPollersLibraryParams object
// with the default values initialized.
func NewGetPollersLibraryParams() *GetPollersLibraryParams {

	return &GetPollersLibraryParams{}
}

/*GetPollersLibraryParams contains all the parameters to send to the API endpoint
for the get pollers library operation typically these are written to a http.Request
*/
type GetPollersLibraryParams struct {
}

// WriteToRequest writes these params to a swagger request
func (o *GetPollersLibraryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
