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

// NewNodesGetTagsByIDParams creates a new NodesGetTagsByIDParams object
// with the default values initialized.
func NewNodesGetTagsByIDParams() *NodesGetTagsByIDParams {
	var ()
	return &NodesGetTagsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNodesGetTagsByIDParamsWithTimeout creates a new NodesGetTagsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNodesGetTagsByIDParamsWithTimeout(timeout time.Duration) *NodesGetTagsByIDParams {
	var ()
	return &NodesGetTagsByIDParams{

		timeout: timeout,
	}
}

// NewNodesGetTagsByIDParamsWithContext creates a new NodesGetTagsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewNodesGetTagsByIDParamsWithContext(ctx context.Context) *NodesGetTagsByIDParams {
	var ()
	return &NodesGetTagsByIDParams{

		Context: ctx,
	}
}

// NewNodesGetTagsByIDParamsWithHTTPClient creates a new NodesGetTagsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNodesGetTagsByIDParamsWithHTTPClient(client *http.Client) *NodesGetTagsByIDParams {
	var ()
	return &NodesGetTagsByIDParams{
		HTTPClient: client,
	}
}

/*NodesGetTagsByIDParams contains all the parameters to send to the API endpoint
for the nodes get tags by Id operation typically these are written to a http.Request
*/
type NodesGetTagsByIDParams struct {

	/*Identifier
	  The node identifier

	*/
	Identifier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the nodes get tags by Id params
func (o *NodesGetTagsByIDParams) WithTimeout(timeout time.Duration) *NodesGetTagsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nodes get tags by Id params
func (o *NodesGetTagsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nodes get tags by Id params
func (o *NodesGetTagsByIDParams) WithContext(ctx context.Context) *NodesGetTagsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nodes get tags by Id params
func (o *NodesGetTagsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nodes get tags by Id params
func (o *NodesGetTagsByIDParams) WithHTTPClient(client *http.Client) *NodesGetTagsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nodes get tags by Id params
func (o *NodesGetTagsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIdentifier adds the identifier to the nodes get tags by Id params
func (o *NodesGetTagsByIDParams) WithIdentifier(identifier string) *NodesGetTagsByIDParams {
	o.SetIdentifier(identifier)
	return o
}

// SetIdentifier adds the identifier to the nodes get tags by Id params
func (o *NodesGetTagsByIDParams) SetIdentifier(identifier string) {
	o.Identifier = identifier
}

// WriteToRequest writes these params to a swagger request
func (o *NodesGetTagsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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