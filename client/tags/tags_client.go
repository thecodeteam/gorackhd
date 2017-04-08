package tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new tags API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tags API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateTag posts a tag

Create a new tag.
*/
func (a *Client) CreateTag(params *CreateTagParams) (*CreateTagCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateTagParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createTag",
		Method:             "POST",
		PathPattern:        "/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateTagReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateTagCreated), nil
}

/*
DeleteTag deletes the specified tag

Delete the specified tag.
*/
func (a *Client) DeleteTag(params *DeleteTagParams) (*DeleteTagNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTagParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteTag",
		Method:             "DELETE",
		PathPattern:        "/tags/{tagName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteTagReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteTagNoContent), nil
}

/*
GetAllTags gets all tags

Get a list of all tags currently stored.
*/
func (a *Client) GetAllTags(params *GetAllTagsParams) (*GetAllTagsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllTagsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAllTags",
		Method:             "GET",
		PathPattern:        "/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAllTagsReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAllTagsOK), nil
}

/*
GetNodesByTag gets nodes with the specified tag

Get a list of nodes with the specified tag.
*/
func (a *Client) GetNodesByTag(params *GetNodesByTagParams) (*GetNodesByTagOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodesByTagParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNodesByTag",
		Method:             "GET",
		PathPattern:        "/tags/{tagName}/nodes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNodesByTagReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNodesByTagOK), nil
}

/*
GetTag gets a tag

Get the data associated with the specified tag.
*/
func (a *Client) GetTag(params *GetTagParams) (*GetTagOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTagParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTag",
		Method:             "GET",
		PathPattern:        "/tags/{tagName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetTagReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTagOK), nil
}

/*
PostWorkflowByID posts workflow to node by tag

Start a workflow against all nodes with the specified tag.
*/
func (a *Client) PostWorkflowByID(params *PostWorkflowByIDParams) (*PostWorkflowByIDAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostWorkflowByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postWorkflowById",
		Method:             "POST",
		PathPattern:        "/tags/{tagName}/nodes/workflows",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostWorkflowByIDReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostWorkflowByIDAccepted), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
