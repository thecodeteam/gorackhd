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

// NewFilesMd5GetParams creates a new FilesMd5GetParams object
// with the default values initialized.
func NewFilesMd5GetParams() *FilesMd5GetParams {
	var ()
	return &FilesMd5GetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFilesMd5GetParamsWithTimeout creates a new FilesMd5GetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFilesMd5GetParamsWithTimeout(timeout time.Duration) *FilesMd5GetParams {
	var ()
	return &FilesMd5GetParams{

		timeout: timeout,
	}
}

/*FilesMd5GetParams contains all the parameters to send to the API endpoint
for the files md5 get operation typically these are written to a http.Request
*/
type FilesMd5GetParams struct {

	/*Filename
	  File name of a file as provided when you originally stored it

	*/
	Filename string

	timeout time.Duration
}

// WithFilename adds the filename to the files md5 get params
func (o *FilesMd5GetParams) WithFilename(filename string) *FilesMd5GetParams {
	o.Filename = filename
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *FilesMd5GetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param filename
	if err := r.SetPathParam("filename", o.Filename); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
