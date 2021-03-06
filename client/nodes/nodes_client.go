package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new nodes API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for nodes API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteNodesIdentifier deletes specified node

Delete specified node.

*/
func (a *Client) DeleteNodesIdentifier(params *DeleteNodesIdentifierParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNodesIdentifierOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNodesIdentifierParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteNodesIdentifier",
		Method:             "DELETE",
		PathPattern:        "/nodes/{identifier}",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteNodesIdentifierReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteNodesIdentifierOK), nil
}

/*
DeleteNodesIdentifierTagsTagname removes tag from specified node

Remove tag from specified node.

*/
func (a *Client) DeleteNodesIdentifierTagsTagname(params *DeleteNodesIdentifierTagsTagnameParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNodesIdentifierTagsTagnameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNodesIdentifierTagsTagnameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteNodesIdentifierTagsTagname",
		Method:             "DELETE",
		PathPattern:        "/nodes/{identifier}/tags/{tagname}",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteNodesIdentifierTagsTagnameReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteNodesIdentifierTagsTagnameOK), nil
}

/*
DeleteNodesIdentifierWorkflowsActive cancels currently running workflows for specified node

Cancel currently running workflows for specified node

*/
func (a *Client) DeleteNodesIdentifierWorkflowsActive(params *DeleteNodesIdentifierWorkflowsActiveParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNodesIdentifierWorkflowsActiveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNodesIdentifierWorkflowsActiveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteNodesIdentifierWorkflowsActive",
		Method:             "DELETE",
		PathPattern:        "/nodes/{identifier}/workflows/active",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteNodesIdentifierWorkflowsActiveReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteNodesIdentifierWorkflowsActiveOK), nil
}

/*
DeleteNodesMacaddressDhcpWhitelist removes a whitelist of specified mac address

Remove a whitelist of specified mac address

*/
func (a *Client) DeleteNodesMacaddressDhcpWhitelist(params *DeleteNodesMacaddressDhcpWhitelistParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNodesMacaddressDhcpWhitelistNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNodesMacaddressDhcpWhitelistParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteNodesMacaddressDhcpWhitelist",
		Method:             "DELETE",
		PathPattern:        "/nodes/{macaddress}/dhcp/whitelist",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteNodesMacaddressDhcpWhitelistReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteNodesMacaddressDhcpWhitelistNoContent), nil
}

/*
GetNodes lists of all nodes or if there are none an empty object

List of all nodes or if there are none an empty object

*/
func (a *Client) GetNodes(params *GetNodesParams, authInfo runtime.ClientAuthInfoWriter) (*GetNodesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNodes",
		Method:             "GET",
		PathPattern:        "/nodes",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNodesReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNodesOK), nil
}

/*
GetNodesIdentifier thes node by identifier

The node by identifier

*/
func (a *Client) GetNodesIdentifier(params *GetNodesIdentifierParams, authInfo runtime.ClientAuthInfoWriter) (*GetNodesIdentifierOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodesIdentifierParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNodesIdentifier",
		Method:             "GET",
		PathPattern:        "/nodes/{identifier}",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNodesIdentifierReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNodesIdentifierOK), nil
}

/*
GetNodesIdentifierCatalogs fetches catalog of specified node

Fetch catalog of specified node

*/
func (a *Client) GetNodesIdentifierCatalogs(params *GetNodesIdentifierCatalogsParams, authInfo runtime.ClientAuthInfoWriter) (*GetNodesIdentifierCatalogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodesIdentifierCatalogsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNodesIdentifierCatalogs",
		Method:             "GET",
		PathPattern:        "/nodes/{identifier}/catalogs",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNodesIdentifierCatalogsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNodesIdentifierCatalogsOK), nil
}

/*
GetNodesIdentifierCatalogsSource fetches catalog of specified node for given source

Fetch catalog of specified node for given source

*/
func (a *Client) GetNodesIdentifierCatalogsSource(params *GetNodesIdentifierCatalogsSourceParams, authInfo runtime.ClientAuthInfoWriter) (*GetNodesIdentifierCatalogsSourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodesIdentifierCatalogsSourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNodesIdentifierCatalogsSource",
		Method:             "GET",
		PathPattern:        "/nodes/{identifier}/catalogs/{source}",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNodesIdentifierCatalogsSourceReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNodesIdentifierCatalogsSourceOK), nil
}

/*
GetNodesIdentifierObm gets the obm settings associated with a node

get the obm settings associated with a node.

*/
func (a *Client) GetNodesIdentifierObm(params *GetNodesIdentifierObmParams, authInfo runtime.ClientAuthInfoWriter) (*GetNodesIdentifierObmOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodesIdentifierObmParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNodesIdentifierObm",
		Method:             "GET",
		PathPattern:        "/nodes/{identifier}/obm",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNodesIdentifierObmReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNodesIdentifierObmOK), nil
}

/*
GetNodesIdentifierObmIdentify fetches status of identify light on node through o b m if supported

Fetch status of identify light on node through OBM (if supported)

*/
func (a *Client) GetNodesIdentifierObmIdentify(params *GetNodesIdentifierObmIdentifyParams, authInfo runtime.ClientAuthInfoWriter) (*GetNodesIdentifierObmIdentifyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodesIdentifierObmIdentifyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNodesIdentifierObmIdentify",
		Method:             "GET",
		PathPattern:        "/nodes/{identifier}/obm/identify",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNodesIdentifierObmIdentifyReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNodesIdentifierObmIdentifyOK), nil
}

/*
GetNodesIdentifierPollers fetches all pollers for specified node

Fetch all pollers for specified node

*/
func (a *Client) GetNodesIdentifierPollers(params *GetNodesIdentifierPollersParams, authInfo runtime.ClientAuthInfoWriter) (*GetNodesIdentifierPollersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodesIdentifierPollersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNodesIdentifierPollers",
		Method:             "GET",
		PathPattern:        "/nodes/{identifier}/pollers",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNodesIdentifierPollersReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNodesIdentifierPollersOK), nil
}

/*
GetNodesIdentifierTags gets the tags associated with a node

get the tags associated with a node.

*/
func (a *Client) GetNodesIdentifierTags(params *GetNodesIdentifierTagsParams, authInfo runtime.ClientAuthInfoWriter) (*GetNodesIdentifierTagsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodesIdentifierTagsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNodesIdentifierTags",
		Method:             "GET",
		PathPattern:        "/nodes/{identifier}/tags",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNodesIdentifierTagsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNodesIdentifierTagsOK), nil
}

/*
GetNodesIdentifierWorkflows fetches all workflows for specified node

Fetch all workflows for specified node

*/
func (a *Client) GetNodesIdentifierWorkflows(params *GetNodesIdentifierWorkflowsParams, authInfo runtime.ClientAuthInfoWriter) (*GetNodesIdentifierWorkflowsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodesIdentifierWorkflowsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNodesIdentifierWorkflows",
		Method:             "GET",
		PathPattern:        "/nodes/{identifier}/workflows",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNodesIdentifierWorkflowsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNodesIdentifierWorkflowsOK), nil
}

/*
GetNodesIdentifierWorkflowsActive fetches currently running workflows for specified node

Fetch currently running workflows for specified node

*/
func (a *Client) GetNodesIdentifierWorkflowsActive(params *GetNodesIdentifierWorkflowsActiveParams, authInfo runtime.ClientAuthInfoWriter) (*GetNodesIdentifierWorkflowsActiveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodesIdentifierWorkflowsActiveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNodesIdentifierWorkflowsActive",
		Method:             "GET",
		PathPattern:        "/nodes/{identifier}/workflows/active",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNodesIdentifierWorkflowsActiveReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNodesIdentifierWorkflowsActiveOK), nil
}

/*
GetSkusIdentifierNodes gets nodes for specific sku

get nodes for specific sku

*/
func (a *Client) GetSkusIdentifierNodes(params *GetSkusIdentifierNodesParams, authInfo runtime.ClientAuthInfoWriter) (*GetSkusIdentifierNodesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSkusIdentifierNodesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSkusIdentifierNodes",
		Method:             "GET",
		PathPattern:        "/skus/{identifier}/nodes",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSkusIdentifierNodesReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSkusIdentifierNodesOK), nil
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
PatchNodesIdentifierTags patches specified node s tags

Patch specified node's tags

*/
func (a *Client) PatchNodesIdentifierTags(params *PatchNodesIdentifierTagsParams, authInfo runtime.ClientAuthInfoWriter) (*PatchNodesIdentifierTagsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchNodesIdentifierTagsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchNodesIdentifierTags",
		Method:             "PATCH",
		PathPattern:        "/nodes/{identifier}/tags",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PatchNodesIdentifierTagsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchNodesIdentifierTagsOK), nil
}

/*
PostNodes posts

post

*/
func (a *Client) PostNodes(params *PostNodesParams, authInfo runtime.ClientAuthInfoWriter) (*PostNodesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostNodesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostNodes",
		Method:             "POST",
		PathPattern:        "/nodes",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostNodesReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostNodesCreated), nil
}

/*
PostNodesIdentifierObm sets the obm settings associated with a node

set he obm settings associated with a node.

*/
func (a *Client) PostNodesIdentifierObm(params *PostNodesIdentifierObmParams, authInfo runtime.ClientAuthInfoWriter) (*PostNodesIdentifierObmCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostNodesIdentifierObmParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostNodesIdentifierObm",
		Method:             "POST",
		PathPattern:        "/nodes/{identifier}/obm",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostNodesIdentifierObmReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostNodesIdentifierObmCreated), nil
}

/*
PostNodesIdentifierObmIdentify enables or disable identify light on node through o b m if supported

Enable or disable identify light on node through OBM (if supported)

*/
func (a *Client) PostNodesIdentifierObmIdentify(params *PostNodesIdentifierObmIdentifyParams, authInfo runtime.ClientAuthInfoWriter) (*PostNodesIdentifierObmIdentifyCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostNodesIdentifierObmIdentifyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostNodesIdentifierObmIdentify",
		Method:             "POST",
		PathPattern:        "/nodes/{identifier}/obm/identify",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostNodesIdentifierObmIdentifyReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostNodesIdentifierObmIdentifyCreated), nil
}

/*
PostNodesIdentifierWorkflows creates workflow for specified node

create workflow for specified node

*/
func (a *Client) PostNodesIdentifierWorkflows(params *PostNodesIdentifierWorkflowsParams, authInfo runtime.ClientAuthInfoWriter) (*PostNodesIdentifierWorkflowsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostNodesIdentifierWorkflowsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostNodesIdentifierWorkflows",
		Method:             "POST",
		PathPattern:        "/nodes/{identifier}/workflows",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostNodesIdentifierWorkflowsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostNodesIdentifierWorkflowsCreated), nil
}

/*
PostNodesMacaddressDhcpWhitelist adds a whitelist of specified mac address

Add a whitelist of specified mac address

*/
func (a *Client) PostNodesMacaddressDhcpWhitelist(params *PostNodesMacaddressDhcpWhitelistParams, authInfo runtime.ClientAuthInfoWriter) (*PostNodesMacaddressDhcpWhitelistCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostNodesMacaddressDhcpWhitelistParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostNodesMacaddressDhcpWhitelist",
		Method:             "POST",
		PathPattern:        "/nodes/{macaddress}/dhcp/whitelist",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostNodesMacaddressDhcpWhitelistReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostNodesMacaddressDhcpWhitelistCreated), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
