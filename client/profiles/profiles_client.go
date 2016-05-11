package profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new profiles API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for profiles API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteSkusIdentifier deletes specific sku

Delete specific sku.

*/
func (a *Client) DeleteSkusIdentifier(params *DeleteSkusIdentifierParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteSkusIdentifierNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSkusIdentifierParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteSkusIdentifier",
		Method:             "DELETE",
		PathPattern:        "/skus/{identifier}",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteSkusIdentifierReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSkusIdentifierNoContent), nil
}

/*
GetProfilesLibrary gets list of possible profiles

get list of possible profiles

*/
func (a *Client) GetProfilesLibrary(params *GetProfilesLibraryParams, authInfo runtime.ClientAuthInfoWriter) (*GetProfilesLibraryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProfilesLibraryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetProfilesLibrary",
		Method:             "GET",
		PathPattern:        "/profiles/library",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetProfilesLibraryReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetProfilesLibraryOK), nil
}

/*
GetProfilesLibraryIdentifier gets a single profile

get a single profile

*/
func (a *Client) GetProfilesLibraryIdentifier(params *GetProfilesLibraryIdentifierParams, authInfo runtime.ClientAuthInfoWriter) (*GetProfilesLibraryIdentifierOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProfilesLibraryIdentifierParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetProfilesLibraryIdentifier",
		Method:             "GET",
		PathPattern:        "/profiles/library/{identifier}",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetProfilesLibraryIdentifierReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetProfilesLibraryIdentifierOK), nil
}

/*
PutProfilesLibraryIdentifier puts a single profile

put a single profile

*/
func (a *Client) PutProfilesLibraryIdentifier(params *PutProfilesLibraryIdentifierParams, authInfo runtime.ClientAuthInfoWriter) (*PutProfilesLibraryIdentifierOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutProfilesLibraryIdentifierParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutProfilesLibraryIdentifier",
		Method:             "PUT",
		PathPattern:        "/profiles/library/{identifier}",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutProfilesLibraryIdentifierReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutProfilesLibraryIdentifierOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
