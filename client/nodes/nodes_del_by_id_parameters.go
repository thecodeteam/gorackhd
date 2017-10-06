// Code generated by go-swagger; DO NOT EDIT.

package nodes

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

// NewNodesDelByIDParams creates a new NodesDelByIDParams object
// with the default values initialized.
func NewNodesDelByIDParams() *NodesDelByIDParams {
	var ()
	return &NodesDelByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNodesDelByIDParamsWithTimeout creates a new NodesDelByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNodesDelByIDParamsWithTimeout(timeout time.Duration) *NodesDelByIDParams {
	var ()
	return &NodesDelByIDParams{

		timeout: timeout,
	}
}

// NewNodesDelByIDParamsWithContext creates a new NodesDelByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewNodesDelByIDParamsWithContext(ctx context.Context) *NodesDelByIDParams {
	var ()
	return &NodesDelByIDParams{

		Context: ctx,
	}
}

// NewNodesDelByIDParamsWithHTTPClient creates a new NodesDelByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNodesDelByIDParamsWithHTTPClient(client *http.Client) *NodesDelByIDParams {
	var ()
	return &NodesDelByIDParams{
		HTTPClient: client,
	}
}

/*NodesDelByIDParams contains all the parameters to send to the API endpoint
for the nodes del by Id operation typically these are written to a http.Request
*/
type NodesDelByIDParams struct {

	/*Identifier
	  The node identifier

	*/
	Identifier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the nodes del by Id params
func (o *NodesDelByIDParams) WithTimeout(timeout time.Duration) *NodesDelByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nodes del by Id params
func (o *NodesDelByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nodes del by Id params
func (o *NodesDelByIDParams) WithContext(ctx context.Context) *NodesDelByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nodes del by Id params
func (o *NodesDelByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nodes del by Id params
func (o *NodesDelByIDParams) WithHTTPClient(client *http.Client) *NodesDelByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nodes del by Id params
func (o *NodesDelByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIdentifier adds the identifier to the nodes del by Id params
func (o *NodesDelByIDParams) WithIdentifier(identifier string) *NodesDelByIDParams {
	o.SetIdentifier(identifier)
	return o
}

// SetIdentifier adds the identifier to the nodes del by Id params
func (o *NodesDelByIDParams) SetIdentifier(identifier string) {
	o.Identifier = identifier
}

// WriteToRequest writes these params to a swagger request
func (o *NodesDelByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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