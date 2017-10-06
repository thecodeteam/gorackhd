// Code generated by go-swagger; DO NOT EDIT.

package tasks

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

	"github.com/spiegela/gorackhd/models"
)

// NewPostTaskByIDParams creates a new PostTaskByIDParams object
// with the default values initialized.
func NewPostTaskByIDParams() *PostTaskByIDParams {
	var ()
	return &PostTaskByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostTaskByIDParamsWithTimeout creates a new PostTaskByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostTaskByIDParamsWithTimeout(timeout time.Duration) *PostTaskByIDParams {
	var ()
	return &PostTaskByIDParams{

		timeout: timeout,
	}
}

// NewPostTaskByIDParamsWithContext creates a new PostTaskByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostTaskByIDParamsWithContext(ctx context.Context) *PostTaskByIDParams {
	var ()
	return &PostTaskByIDParams{

		Context: ctx,
	}
}

// NewPostTaskByIDParamsWithHTTPClient creates a new PostTaskByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostTaskByIDParamsWithHTTPClient(client *http.Client) *PostTaskByIDParams {
	var ()
	return &PostTaskByIDParams{
		HTTPClient: client,
	}
}

/*PostTaskByIDParams contains all the parameters to send to the API endpoint
for the post task by Id operation typically these are written to a http.Request
*/
type PostTaskByIDParams struct {

	/*Body
	  The obm settings to apply

	*/
	Body *models.PostTasks
	/*Identifier
	  The task identifier

	*/
	Identifier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post task by Id params
func (o *PostTaskByIDParams) WithTimeout(timeout time.Duration) *PostTaskByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post task by Id params
func (o *PostTaskByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post task by Id params
func (o *PostTaskByIDParams) WithContext(ctx context.Context) *PostTaskByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post task by Id params
func (o *PostTaskByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post task by Id params
func (o *PostTaskByIDParams) WithHTTPClient(client *http.Client) *PostTaskByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post task by Id params
func (o *PostTaskByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post task by Id params
func (o *PostTaskByIDParams) WithBody(body *models.PostTasks) *PostTaskByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post task by Id params
func (o *PostTaskByIDParams) SetBody(body *models.PostTasks) {
	o.Body = body
}

// WithIdentifier adds the identifier to the post task by Id params
func (o *PostTaskByIDParams) WithIdentifier(identifier string) *PostTaskByIDParams {
	o.SetIdentifier(identifier)
	return o
}

// SetIdentifier adds the identifier to the post task by Id params
func (o *PostTaskByIDParams) SetIdentifier(identifier string) {
	o.Identifier = identifier
}

// WriteToRequest writes these params to a swagger request
func (o *PostTaskByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models.PostTasks)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param identifier
	if err := r.SetPathParam("identifier", o.Identifier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}