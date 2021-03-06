package workflow

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new workflow API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for workflow API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
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
GetWorkflows fetches workflows

Fetch workflows

*/
func (a *Client) GetWorkflows(params *GetWorkflowsParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkflowsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkflowsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetWorkflows",
		Method:             "GET",
		PathPattern:        "/workflows",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetWorkflowsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWorkflowsOK), nil
}

/*
GetWorkflowsInstanceID fetches workflow by instance ID

Fetch workflow by instance ID

*/
func (a *Client) GetWorkflowsInstanceID(params *GetWorkflowsInstanceIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkflowsInstanceIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkflowsInstanceIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetWorkflowsInstanceID",
		Method:             "GET",
		PathPattern:        "/workflows/{instanceId}",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetWorkflowsInstanceIDReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWorkflowsInstanceIDOK), nil
}

/*
GetWorkflowsLibrary lists all workflows available to run

List all workflows available to run

*/
func (a *Client) GetWorkflowsLibrary(params *GetWorkflowsLibraryParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkflowsLibraryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkflowsLibraryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetWorkflowsLibrary",
		Method:             "GET",
		PathPattern:        "/workflows/library",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetWorkflowsLibraryReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWorkflowsLibraryOK), nil
}

/*
GetWorkflowsLibraryInjectableName fetches workflow by injectable name

Fetch workflow by injectable name

*/
func (a *Client) GetWorkflowsLibraryInjectableName(params *GetWorkflowsLibraryInjectableNameParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkflowsLibraryInjectableNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkflowsLibraryInjectableNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetWorkflowsLibraryInjectableName",
		Method:             "GET",
		PathPattern:        "/workflows/library/{injectableName}",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetWorkflowsLibraryInjectableNameReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWorkflowsLibraryInjectableNameOK), nil
}

/*
GetWorkflowsTasks fetches tasks from task library

Fetch tasks from task library

*/
func (a *Client) GetWorkflowsTasks(params *GetWorkflowsTasksParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkflowsTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkflowsTasksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetWorkflowsTasks",
		Method:             "GET",
		PathPattern:        "/workflows/tasks",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetWorkflowsTasksReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWorkflowsTasksOK), nil
}

/*
GetWorkflowsTasksLibrary lists workflow tasks library

List workflow tasks library

*/
func (a *Client) GetWorkflowsTasksLibrary(params *GetWorkflowsTasksLibraryParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkflowsTasksLibraryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkflowsTasksLibraryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetWorkflowsTasksLibrary",
		Method:             "GET",
		PathPattern:        "/workflows/tasks/library",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetWorkflowsTasksLibraryReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWorkflowsTasksLibraryOK), nil
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
PostWorkflows creates an unbound workflow

create an unbound workflow

*/
func (a *Client) PostWorkflows(params *PostWorkflowsParams, authInfo runtime.ClientAuthInfoWriter) (*PostWorkflowsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostWorkflowsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostWorkflows",
		Method:             "POST",
		PathPattern:        "/workflows",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostWorkflowsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostWorkflowsOK), nil
}

/*
PutWorkflows defines new workflow

define new workflow

*/
func (a *Client) PutWorkflows(params *PutWorkflowsParams, authInfo runtime.ClientAuthInfoWriter) (*PutWorkflowsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutWorkflowsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutWorkflows",
		Method:             "PUT",
		PathPattern:        "/workflows",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutWorkflowsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutWorkflowsOK), nil
}

/*
PutWorkflowsTasks adds tasks to task library

Add tasks to task library

*/
func (a *Client) PutWorkflowsTasks(params *PutWorkflowsTasksParams, authInfo runtime.ClientAuthInfoWriter) (*PutWorkflowsTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutWorkflowsTasksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutWorkflowsTasks",
		Method:             "PUT",
		PathPattern:        "/workflows/tasks",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutWorkflowsTasksReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutWorkflowsTasksOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
