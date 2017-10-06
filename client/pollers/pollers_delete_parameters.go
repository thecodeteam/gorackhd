// Code generated by go-swagger; DO NOT EDIT.

package pollers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPollersDeleteParams creates a new PollersDeleteParams object
// with the default values initialized.
func NewPollersDeleteParams() *PollersDeleteParams {
	var ()
	return &PollersDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPollersDeleteParamsWithTimeout creates a new PollersDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPollersDeleteParamsWithTimeout(timeout time.Duration) *PollersDeleteParams {
	var ()
	return &PollersDeleteParams{

		timeout: timeout,
	}
}

// NewPollersDeleteParamsWithContext creates a new PollersDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewPollersDeleteParamsWithContext(ctx context.Context) *PollersDeleteParams {
	var ()
	return &PollersDeleteParams{

		Context: ctx,
	}
}

// NewPollersDeleteParamsWithHTTPClient creates a new PollersDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPollersDeleteParamsWithHTTPClient(client *http.Client) *PollersDeleteParams {
	var ()
	return &PollersDeleteParams{
		HTTPClient: client,
	}
}

/*PollersDeleteParams contains all the parameters to send to the API endpoint
for the pollers delete operation typically these are written to a http.Request
*/
type PollersDeleteParams struct {

	/*Identifier
	  The poller identifier

	*/
	Identifier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pollers delete params
func (o *PollersDeleteParams) WithTimeout(timeout time.Duration) *PollersDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pollers delete params
func (o *PollersDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pollers delete params
func (o *PollersDeleteParams) WithContext(ctx context.Context) *PollersDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pollers delete params
func (o *PollersDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pollers delete params
func (o *PollersDeleteParams) WithHTTPClient(client *http.Client) *PollersDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pollers delete params
func (o *PollersDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIdentifier adds the identifier to the pollers delete params
func (o *PollersDeleteParams) WithIdentifier(identifier string) *PollersDeleteParams {
	o.SetIdentifier(identifier)
	return o
}

// SetIdentifier adds the identifier to the pollers delete params
func (o *PollersDeleteParams) SetIdentifier(identifier string) {
	o.Identifier = identifier
}

// WriteToRequest writes these params to a swagger request
func (o *PollersDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
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