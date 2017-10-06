// Code generated by go-swagger; DO NOT EDIT.

package workflows

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

// NewWorkflowsGetGraphsByNameParams creates a new WorkflowsGetGraphsByNameParams object
// with the default values initialized.
func NewWorkflowsGetGraphsByNameParams() *WorkflowsGetGraphsByNameParams {
	var ()
	return &WorkflowsGetGraphsByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWorkflowsGetGraphsByNameParamsWithTimeout creates a new WorkflowsGetGraphsByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWorkflowsGetGraphsByNameParamsWithTimeout(timeout time.Duration) *WorkflowsGetGraphsByNameParams {
	var ()
	return &WorkflowsGetGraphsByNameParams{

		timeout: timeout,
	}
}

// NewWorkflowsGetGraphsByNameParamsWithContext creates a new WorkflowsGetGraphsByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewWorkflowsGetGraphsByNameParamsWithContext(ctx context.Context) *WorkflowsGetGraphsByNameParams {
	var ()
	return &WorkflowsGetGraphsByNameParams{

		Context: ctx,
	}
}

// NewWorkflowsGetGraphsByNameParamsWithHTTPClient creates a new WorkflowsGetGraphsByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWorkflowsGetGraphsByNameParamsWithHTTPClient(client *http.Client) *WorkflowsGetGraphsByNameParams {
	var ()
	return &WorkflowsGetGraphsByNameParams{
		HTTPClient: client,
	}
}

/*WorkflowsGetGraphsByNameParams contains all the parameters to send to the API endpoint
for the workflows get graphs by name operation typically these are written to a http.Request
*/
type WorkflowsGetGraphsByNameParams struct {

	/*InjectableName
	  The workflow graph injectable name

	*/
	InjectableName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the workflows get graphs by name params
func (o *WorkflowsGetGraphsByNameParams) WithTimeout(timeout time.Duration) *WorkflowsGetGraphsByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the workflows get graphs by name params
func (o *WorkflowsGetGraphsByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the workflows get graphs by name params
func (o *WorkflowsGetGraphsByNameParams) WithContext(ctx context.Context) *WorkflowsGetGraphsByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the workflows get graphs by name params
func (o *WorkflowsGetGraphsByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the workflows get graphs by name params
func (o *WorkflowsGetGraphsByNameParams) WithHTTPClient(client *http.Client) *WorkflowsGetGraphsByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the workflows get graphs by name params
func (o *WorkflowsGetGraphsByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInjectableName adds the injectableName to the workflows get graphs by name params
func (o *WorkflowsGetGraphsByNameParams) WithInjectableName(injectableName string) *WorkflowsGetGraphsByNameParams {
	o.SetInjectableName(injectableName)
	return o
}

// SetInjectableName adds the injectableName to the workflows get graphs by name params
func (o *WorkflowsGetGraphsByNameParams) SetInjectableName(injectableName string) {
	o.InjectableName = injectableName
}

// WriteToRequest writes these params to a swagger request
func (o *WorkflowsGetGraphsByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param injectableName
	if err := r.SetPathParam("injectableName", o.InjectableName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}