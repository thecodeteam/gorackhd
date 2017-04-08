package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// NewNodesPatchTagByIDParams creates a new NodesPatchTagByIDParams object
// with the default values initialized.
func NewNodesPatchTagByIDParams() *NodesPatchTagByIDParams {
	var ()
	return &NodesPatchTagByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNodesPatchTagByIDParamsWithTimeout creates a new NodesPatchTagByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNodesPatchTagByIDParamsWithTimeout(timeout time.Duration) *NodesPatchTagByIDParams {
	var ()
	return &NodesPatchTagByIDParams{

		timeout: timeout,
	}
}

/*NodesPatchTagByIDParams contains all the parameters to send to the API endpoint
for the nodes patch tag by Id operation typically these are written to a http.Request
*/
type NodesPatchTagByIDParams struct {

	/*Body
	  The tag properties to patch

	*/
	Body *models.PostTags
	/*Identifier
	  The node identifier

	*/
	Identifier string

	timeout time.Duration
}

// WithBody adds the body to the nodes patch tag by Id params
func (o *NodesPatchTagByIDParams) WithBody(body *models.PostTags) *NodesPatchTagByIDParams {
	o.Body = body
	return o
}

// WithIdentifier adds the identifier to the nodes patch tag by Id params
func (o *NodesPatchTagByIDParams) WithIdentifier(identifier string) *NodesPatchTagByIDParams {
	o.Identifier = identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *NodesPatchTagByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.PostTags)
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
