package patch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new patch API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for patch API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
PatchConfig patches update server configuration

Patch/update server configurationm and then return the patched configuration.

*/
func (a *Client) PatchConfig(params *PatchConfigParams, authInfo runtime.ClientAuthInfoWriter) (*PatchConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchConfig",
		Method:             "PATCH",
		PathPattern:        "/config",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PatchConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchConfigOK), nil
}

/*
PatchNodesIdentifier patches specified node

Patch specified node

*/
func (a *Client) PatchNodesIdentifier(params *PatchNodesIdentifierParams, authInfo runtime.ClientAuthInfoWriter) (*PatchNodesIdentifierOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchNodesIdentifierParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchNodesIdentifier",
		Method:             "PATCH",
		PathPattern:        "/nodes/{identifier}",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PatchNodesIdentifierReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchNodesIdentifierOK), nil
}

/*
PatchPollersIdentifier patches specifics of the specified poller

patch specifics of the specified poller

*/
func (a *Client) PatchPollersIdentifier(params *PatchPollersIdentifierParams, authInfo runtime.ClientAuthInfoWriter) (*PatchPollersIdentifierOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchPollersIdentifierParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchPollersIdentifier",
		Method:             "PATCH",
		PathPattern:        "/pollers/{identifier}",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PatchPollersIdentifierReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchPollersIdentifierOK), nil
}

/*
PatchSkusIdentifier patches a single sku

patch a single sku

*/
func (a *Client) PatchSkusIdentifier(params *PatchSkusIdentifierParams, authInfo runtime.ClientAuthInfoWriter) (*PatchSkusIdentifierOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchSkusIdentifierParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchSkusIdentifier",
		Method:             "PATCH",
		PathPattern:        "/skus/{identifier}",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PatchSkusIdentifierReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchSkusIdentifierOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
