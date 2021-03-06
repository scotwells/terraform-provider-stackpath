// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new workload API client.
func NewCompute(transport runtime.ClientTransport, formats strfmt.Registry) Compute {
	return &compute{transport: transport, formats: formats}
}

/*
compute for workload API
*/

// Compute defines the client interface
type Compute interface {
	CreateWorkload(params *CreateWorkloadParams, authInfo runtime.ClientAuthInfoWriter) (*CreateWorkloadOK, error)
	DeleteWorkload(params *DeleteWorkloadParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteWorkloadNoContent, error)
	GetLocations(params *GetLocationsParams, authInfo runtime.ClientAuthInfoWriter) (*GetLocationsOK, error)
	GetWorkload(params *GetWorkloadParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkloadOK, error)
	GetWorkloads(params *GetWorkloadsParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkloadsOK, error)
	UpdateWorkload(params *UpdateWorkloadParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateWorkloadOK, error)
	GetWorkloadInstances(params *GetWorkloadInstancesParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkloadInstancesOK, error)
	SetTransport(transport runtime.ClientTransport)
}

type compute struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateWorkload creates a new workload
*/
func (a *compute) CreateWorkload(params *CreateWorkloadParams, authInfo runtime.ClientAuthInfoWriter) (*CreateWorkloadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateWorkloadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateWorkload",
		Method:             "POST",
		PathPattern:        "/workload/v1/stacks/{stack_id}/workloads",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateWorkloadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateWorkloadOK), nil

}

/*
DeleteWorkload deletes a workload
*/
func (a *compute) DeleteWorkload(params *DeleteWorkloadParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteWorkloadNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteWorkloadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteWorkload",
		Method:             "DELETE",
		PathPattern:        "/workload/v1/stacks/{stack_id}/workloads/{workload_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteWorkloadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteWorkloadNoContent), nil

}

/*
GetLocations retrieves the locations a workload may be created in
*/
func (a *compute) GetLocations(params *GetLocationsParams, authInfo runtime.ClientAuthInfoWriter) (*GetLocationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLocationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLocations",
		Method:             "GET",
		PathPattern:        "/workload/v1/locations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLocationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLocationsOK), nil

}

/*
GetWorkload retrieves an individual workload
*/
func (a *compute) GetWorkload(params *GetWorkloadParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkloadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkloadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetWorkload",
		Method:             "GET",
		PathPattern:        "/workload/v1/stacks/{stack_id}/workloads/{workload_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWorkloadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWorkloadOK), nil

}

/*
GetWorkloads retrieves a stack s workloads
*/
func (a *compute) GetWorkloads(params *GetWorkloadsParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkloadsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkloadsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetWorkloads",
		Method:             "GET",
		PathPattern:        "/workload/v1/stacks/{stack_id}/workloads",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWorkloadsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWorkloadsOK), nil

}

/*
UpdateWorkload updates a workload
*/
func (a *compute) UpdateWorkload(params *UpdateWorkloadParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateWorkloadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateWorkloadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateWorkload",
		Method:             "PATCH",
		PathPattern:        "/workload/v1/stacks/{stack_id}/workloads/{workload_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateWorkloadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateWorkloadOK), nil

}

/*
GetWorkloadInstances retrieves a workload s instances
*/
func (a *compute) GetWorkloadInstances(params *GetWorkloadInstancesParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkloadInstancesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkloadInstancesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetWorkloadInstances",
		Method:             "GET",
		PathPattern:        "/workload/v1/stacks/{stack_id}/workloads/{workload_id}/instances",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWorkloadInstancesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWorkloadInstancesOK), nil

}

// SetTransport changes the transport on the client
func (a *compute) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
