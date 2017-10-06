// Code generated by go-swagger; DO NOT EDIT.

package profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewProfilesGetParams creates a new ProfilesGetParams object
// with the default values initialized.
func NewProfilesGetParams() *ProfilesGetParams {
	var ()
	return &ProfilesGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProfilesGetParamsWithTimeout creates a new ProfilesGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProfilesGetParamsWithTimeout(timeout time.Duration) *ProfilesGetParams {
	var ()
	return &ProfilesGetParams{

		timeout: timeout,
	}
}

// NewProfilesGetParamsWithContext creates a new ProfilesGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewProfilesGetParamsWithContext(ctx context.Context) *ProfilesGetParams {
	var ()
	return &ProfilesGetParams{

		Context: ctx,
	}
}

// NewProfilesGetParamsWithHTTPClient creates a new ProfilesGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProfilesGetParamsWithHTTPClient(client *http.Client) *ProfilesGetParams {
	var ()
	return &ProfilesGetParams{
		HTTPClient: client,
	}
}

/*ProfilesGetParams contains all the parameters to send to the API endpoint
for the profiles get operation typically these are written to a http.Request
*/
type ProfilesGetParams struct {

	/*IP
	  When macs parameters is not passed, passed with MAC adds IP address to lookup

	*/
	IP *string
	/*Ips
	  List of IP addresses to lookup

	*/
	Ips []string
	/*Mac
	  When macs parameter is not passed, passed with IP adds MAC address to lookup

	*/
	Mac *string
	/*Macs
	  List of valid MAC addresses to lookup

	*/
	Macs []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the profiles get params
func (o *ProfilesGetParams) WithTimeout(timeout time.Duration) *ProfilesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the profiles get params
func (o *ProfilesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the profiles get params
func (o *ProfilesGetParams) WithContext(ctx context.Context) *ProfilesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the profiles get params
func (o *ProfilesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the profiles get params
func (o *ProfilesGetParams) WithHTTPClient(client *http.Client) *ProfilesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the profiles get params
func (o *ProfilesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIP adds the ip to the profiles get params
func (o *ProfilesGetParams) WithIP(ip *string) *ProfilesGetParams {
	o.SetIP(ip)
	return o
}

// SetIP adds the ip to the profiles get params
func (o *ProfilesGetParams) SetIP(ip *string) {
	o.IP = ip
}

// WithIps adds the ips to the profiles get params
func (o *ProfilesGetParams) WithIps(ips []string) *ProfilesGetParams {
	o.SetIps(ips)
	return o
}

// SetIps adds the ips to the profiles get params
func (o *ProfilesGetParams) SetIps(ips []string) {
	o.Ips = ips
}

// WithMac adds the mac to the profiles get params
func (o *ProfilesGetParams) WithMac(mac *string) *ProfilesGetParams {
	o.SetMac(mac)
	return o
}

// SetMac adds the mac to the profiles get params
func (o *ProfilesGetParams) SetMac(mac *string) {
	o.Mac = mac
}

// WithMacs adds the macs to the profiles get params
func (o *ProfilesGetParams) WithMacs(macs []string) *ProfilesGetParams {
	o.SetMacs(macs)
	return o
}

// SetMacs adds the macs to the profiles get params
func (o *ProfilesGetParams) SetMacs(macs []string) {
	o.Macs = macs
}

// WriteToRequest writes these params to a swagger request
func (o *ProfilesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IP != nil {

		// query param ip
		var qrIP string
		if o.IP != nil {
			qrIP = *o.IP
		}
		qIP := qrIP
		if qIP != "" {
			if err := r.SetQueryParam("ip", qIP); err != nil {
				return err
			}
		}

	}

	valuesIps := o.Ips

	joinedIps := swag.JoinByFormat(valuesIps, "multi")
	// query array param ips
	if err := r.SetQueryParam("ips", joinedIps...); err != nil {
		return err
	}

	if o.Mac != nil {

		// query param mac
		var qrMac string
		if o.Mac != nil {
			qrMac = *o.Mac
		}
		qMac := qrMac
		if qMac != "" {
			if err := r.SetQueryParam("mac", qMac); err != nil {
				return err
			}
		}

	}

	valuesMacs := o.Macs

	joinedMacs := swag.JoinByFormat(valuesMacs, "multi")
	// query array param macs
	if err := r.SetQueryParam("macs", joinedMacs...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}