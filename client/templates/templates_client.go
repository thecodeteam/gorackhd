package templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new templates API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for templates API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
TemplatesGetByName gets the specified template

Get the template with the specified name.
*/
func (a *Client) TemplatesGetByName(params *TemplatesGetByNameParams) (*TemplatesGetByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTemplatesGetByNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "templatesGetByName",
		Method:             "GET",
		PathPattern:        "/templates/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TemplatesGetByNameReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*TemplatesGetByNameOK), nil
}

/*
TemplatesHeadByName gets headers of the specified template

Get the template headers with the specified name. Only gets headers and returns no body.
*/
func (a *Client) TemplatesHeadByName(params *TemplatesHeadByNameParams) (*TemplatesHeadByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTemplatesHeadByNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "templatesHeadByName",
		Method:             "HEAD",
		PathPattern:        "/templates/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TemplatesHeadByNameReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*TemplatesHeadByNameOK), nil
}

/*
TemplatesLibDelete deletes the specified template

Delete the template with the specified name.
*/
func (a *Client) TemplatesLibDelete(params *TemplatesLibDeleteParams) (*TemplatesLibDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTemplatesLibDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "templatesLibDelete",
		Method:             "DELETE",
		PathPattern:        "/templates/library/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TemplatesLibDeleteReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*TemplatesLibDeleteNoContent), nil
}

/*
TemplatesLibGet gets the specified template

Get the contents of the specified template file.
*/
func (a *Client) TemplatesLibGet(params *TemplatesLibGetParams) (*TemplatesLibGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTemplatesLibGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "templatesLibGet",
		Method:             "GET",
		PathPattern:        "/templates/library/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TemplatesLibGetReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*TemplatesLibGetOK), nil
}

/*
TemplatesLibPut creates or update the specified template

Create or update the metadata and the content of the specified template.
*/
func (a *Client) TemplatesLibPut(params *TemplatesLibPutParams) (*TemplatesLibPutCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTemplatesLibPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "templatesLibPut",
		Method:             "PUT",
		PathPattern:        "/templates/library/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TemplatesLibPutReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*TemplatesLibPutCreated), nil
}

/*
TemplatesMetaGet gets a list of metadata for all templates

Retrieves metadata for each defined template. The metadata includes the template name, and scope, not the actual contents of the template.

*/
func (a *Client) TemplatesMetaGet(params *TemplatesMetaGetParams) (*TemplatesMetaGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTemplatesMetaGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "templatesMetaGet",
		Method:             "GET",
		PathPattern:        "/templates/metadata",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TemplatesMetaGetReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*TemplatesMetaGetOK), nil
}

/*
TemplatesMetaGetByName gets the metadata for the specified template

Get metadata for the template with the specified name.
*/
func (a *Client) TemplatesMetaGetByName(params *TemplatesMetaGetByNameParams) (*TemplatesMetaGetByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTemplatesMetaGetByNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "templatesMetaGetByName",
		Method:             "GET",
		PathPattern:        "/templates/metadata/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TemplatesMetaGetByNameReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*TemplatesMetaGetByNameOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
