//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicefabricmanagedclusters

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// NodeTypesClient contains the methods for the NodeTypes group.
// Don't use this type directly, use NewNodeTypesClient() instead.
type NodeTypesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewNodeTypesClient creates a new instance of NodeTypesClient with the specified values.
//   - subscriptionID - The customer subscription identifier.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNodeTypesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NodeTypesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NodeTypesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update a Service Fabric node type of a given managed cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group.
//   - clusterName - The name of the cluster resource.
//   - nodeTypeName - The name of the node type.
//   - parameters - The node type resource.
//   - options - NodeTypesClientBeginCreateOrUpdateOptions contains the optional parameters for the NodeTypesClient.BeginCreateOrUpdate
//     method.
func (client *NodeTypesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, nodeTypeName string, parameters NodeType, options *NodeTypesClientBeginCreateOrUpdateOptions) (*runtime.Poller[NodeTypesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, clusterName, nodeTypeName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NodeTypesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NodeTypesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create or update a Service Fabric node type of a given managed cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
func (client *NodeTypesClient) createOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, nodeTypeName string, parameters NodeType, options *NodeTypesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "NodeTypesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, clusterName, nodeTypeName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *NodeTypesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, nodeTypeName string, parameters NodeType, options *NodeTypesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/managedClusters/{clusterName}/nodeTypes/{nodeTypeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if nodeTypeName == "" {
		return nil, errors.New("parameter nodeTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nodeTypeName}", url.PathEscape(nodeTypeName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a Service Fabric node type of a given managed cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group.
//   - clusterName - The name of the cluster resource.
//   - nodeTypeName - The name of the node type.
//   - options - NodeTypesClientBeginDeleteOptions contains the optional parameters for the NodeTypesClient.BeginDelete method.
func (client *NodeTypesClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, nodeTypeName string, options *NodeTypesClientBeginDeleteOptions) (*runtime.Poller[NodeTypesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, nodeTypeName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NodeTypesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NodeTypesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a Service Fabric node type of a given managed cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
func (client *NodeTypesClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, nodeTypeName string, options *NodeTypesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "NodeTypesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, nodeTypeName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *NodeTypesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, nodeTypeName string, options *NodeTypesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/managedClusters/{clusterName}/nodeTypes/{nodeTypeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if nodeTypeName == "" {
		return nil, errors.New("parameter nodeTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nodeTypeName}", url.PathEscape(nodeTypeName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginDeleteNode - Deletes one or more nodes on the node type. It will disable the fabric nodes, trigger a delete on the
// VMs and removes the state from the cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group.
//   - clusterName - The name of the cluster resource.
//   - nodeTypeName - The name of the node type.
//   - parameters - parameters for delete action.
//   - options - NodeTypesClientBeginDeleteNodeOptions contains the optional parameters for the NodeTypesClient.BeginDeleteNode
//     method.
func (client *NodeTypesClient) BeginDeleteNode(ctx context.Context, resourceGroupName string, clusterName string, nodeTypeName string, parameters NodeTypeActionParameters, options *NodeTypesClientBeginDeleteNodeOptions) (*runtime.Poller[NodeTypesClientDeleteNodeResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteNode(ctx, resourceGroupName, clusterName, nodeTypeName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NodeTypesClientDeleteNodeResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NodeTypesClientDeleteNodeResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// DeleteNode - Deletes one or more nodes on the node type. It will disable the fabric nodes, trigger a delete on the VMs
// and removes the state from the cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
func (client *NodeTypesClient) deleteNode(ctx context.Context, resourceGroupName string, clusterName string, nodeTypeName string, parameters NodeTypeActionParameters, options *NodeTypesClientBeginDeleteNodeOptions) (*http.Response, error) {
	var err error
	const operationName = "NodeTypesClient.BeginDeleteNode"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteNodeCreateRequest(ctx, resourceGroupName, clusterName, nodeTypeName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteNodeCreateRequest creates the DeleteNode request.
func (client *NodeTypesClient) deleteNodeCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, nodeTypeName string, parameters NodeTypeActionParameters, options *NodeTypesClientBeginDeleteNodeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/managedClusters/{clusterName}/nodeTypes/{nodeTypeName}/deleteNode"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if nodeTypeName == "" {
		return nil, errors.New("parameter nodeTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nodeTypeName}", url.PathEscape(nodeTypeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Get a Service Fabric node type of a given managed cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group.
//   - clusterName - The name of the cluster resource.
//   - nodeTypeName - The name of the node type.
//   - options - NodeTypesClientGetOptions contains the optional parameters for the NodeTypesClient.Get method.
func (client *NodeTypesClient) Get(ctx context.Context, resourceGroupName string, clusterName string, nodeTypeName string, options *NodeTypesClientGetOptions) (NodeTypesClientGetResponse, error) {
	var err error
	const operationName = "NodeTypesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, nodeTypeName, options)
	if err != nil {
		return NodeTypesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NodeTypesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NodeTypesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *NodeTypesClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, nodeTypeName string, options *NodeTypesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/managedClusters/{clusterName}/nodeTypes/{nodeTypeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if nodeTypeName == "" {
		return nil, errors.New("parameter nodeTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nodeTypeName}", url.PathEscape(nodeTypeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *NodeTypesClient) getHandleResponse(resp *http.Response) (NodeTypesClientGetResponse, error) {
	result := NodeTypesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NodeType); err != nil {
		return NodeTypesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByManagedClustersPager - Gets all Node types of the specified managed cluster.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group.
//   - clusterName - The name of the cluster resource.
//   - options - NodeTypesClientListByManagedClustersOptions contains the optional parameters for the NodeTypesClient.NewListByManagedClustersPager
//     method.
func (client *NodeTypesClient) NewListByManagedClustersPager(resourceGroupName string, clusterName string, options *NodeTypesClientListByManagedClustersOptions) *runtime.Pager[NodeTypesClientListByManagedClustersResponse] {
	return runtime.NewPager(runtime.PagingHandler[NodeTypesClientListByManagedClustersResponse]{
		More: func(page NodeTypesClientListByManagedClustersResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NodeTypesClientListByManagedClustersResponse) (NodeTypesClientListByManagedClustersResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "NodeTypesClient.NewListByManagedClustersPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByManagedClustersCreateRequest(ctx, resourceGroupName, clusterName, options)
			}, nil)
			if err != nil {
				return NodeTypesClientListByManagedClustersResponse{}, err
			}
			return client.listByManagedClustersHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByManagedClustersCreateRequest creates the ListByManagedClusters request.
func (client *NodeTypesClient) listByManagedClustersCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *NodeTypesClientListByManagedClustersOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/managedClusters/{clusterName}/nodeTypes"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByManagedClustersHandleResponse handles the ListByManagedClusters response.
func (client *NodeTypesClient) listByManagedClustersHandleResponse(resp *http.Response) (NodeTypesClientListByManagedClustersResponse, error) {
	result := NodeTypesClientListByManagedClustersResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NodeTypeListResult); err != nil {
		return NodeTypesClientListByManagedClustersResponse{}, err
	}
	return result, nil
}

// BeginReimage - Reimages one or more nodes on the node type. It will disable the fabric nodes, trigger a reimage on the
// VMs and activate the nodes back again.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group.
//   - clusterName - The name of the cluster resource.
//   - nodeTypeName - The name of the node type.
//   - parameters - parameters for reimage action.
//   - options - NodeTypesClientBeginReimageOptions contains the optional parameters for the NodeTypesClient.BeginReimage method.
func (client *NodeTypesClient) BeginReimage(ctx context.Context, resourceGroupName string, clusterName string, nodeTypeName string, parameters NodeTypeActionParameters, options *NodeTypesClientBeginReimageOptions) (*runtime.Poller[NodeTypesClientReimageResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.reimage(ctx, resourceGroupName, clusterName, nodeTypeName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NodeTypesClientReimageResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NodeTypesClientReimageResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Reimage - Reimages one or more nodes on the node type. It will disable the fabric nodes, trigger a reimage on the VMs and
// activate the nodes back again.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
func (client *NodeTypesClient) reimage(ctx context.Context, resourceGroupName string, clusterName string, nodeTypeName string, parameters NodeTypeActionParameters, options *NodeTypesClientBeginReimageOptions) (*http.Response, error) {
	var err error
	const operationName = "NodeTypesClient.BeginReimage"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.reimageCreateRequest(ctx, resourceGroupName, clusterName, nodeTypeName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// reimageCreateRequest creates the Reimage request.
func (client *NodeTypesClient) reimageCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, nodeTypeName string, parameters NodeTypeActionParameters, options *NodeTypesClientBeginReimageOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/managedClusters/{clusterName}/nodeTypes/{nodeTypeName}/reimage"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if nodeTypeName == "" {
		return nil, errors.New("parameter nodeTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nodeTypeName}", url.PathEscape(nodeTypeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginRestart - Restarts one or more nodes on the node type. It will disable the fabric nodes, trigger a restart on the
// VMs and activate the nodes back again.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group.
//   - clusterName - The name of the cluster resource.
//   - nodeTypeName - The name of the node type.
//   - parameters - parameters for restart action.
//   - options - NodeTypesClientBeginRestartOptions contains the optional parameters for the NodeTypesClient.BeginRestart method.
func (client *NodeTypesClient) BeginRestart(ctx context.Context, resourceGroupName string, clusterName string, nodeTypeName string, parameters NodeTypeActionParameters, options *NodeTypesClientBeginRestartOptions) (*runtime.Poller[NodeTypesClientRestartResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.restart(ctx, resourceGroupName, clusterName, nodeTypeName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NodeTypesClientRestartResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NodeTypesClientRestartResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Restart - Restarts one or more nodes on the node type. It will disable the fabric nodes, trigger a restart on the VMs and
// activate the nodes back again.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
func (client *NodeTypesClient) restart(ctx context.Context, resourceGroupName string, clusterName string, nodeTypeName string, parameters NodeTypeActionParameters, options *NodeTypesClientBeginRestartOptions) (*http.Response, error) {
	var err error
	const operationName = "NodeTypesClient.BeginRestart"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.restartCreateRequest(ctx, resourceGroupName, clusterName, nodeTypeName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// restartCreateRequest creates the Restart request.
func (client *NodeTypesClient) restartCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, nodeTypeName string, parameters NodeTypeActionParameters, options *NodeTypesClientBeginRestartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/managedClusters/{clusterName}/nodeTypes/{nodeTypeName}/restart"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if nodeTypeName == "" {
		return nil, errors.New("parameter nodeTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nodeTypeName}", url.PathEscape(nodeTypeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// Update - Update the configuration of a node type of a given managed cluster, only updating tags.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group.
//   - clusterName - The name of the cluster resource.
//   - nodeTypeName - The name of the node type.
//   - parameters - The parameters to update the node type configuration.
//   - options - NodeTypesClientUpdateOptions contains the optional parameters for the NodeTypesClient.Update method.
func (client *NodeTypesClient) Update(ctx context.Context, resourceGroupName string, clusterName string, nodeTypeName string, parameters NodeTypeUpdateParameters, options *NodeTypesClientUpdateOptions) (NodeTypesClientUpdateResponse, error) {
	var err error
	const operationName = "NodeTypesClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, clusterName, nodeTypeName, parameters, options)
	if err != nil {
		return NodeTypesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NodeTypesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NodeTypesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *NodeTypesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, nodeTypeName string, parameters NodeTypeUpdateParameters, options *NodeTypesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/managedClusters/{clusterName}/nodeTypes/{nodeTypeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if nodeTypeName == "" {
		return nil, errors.New("parameter nodeTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nodeTypeName}", url.PathEscape(nodeTypeName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *NodeTypesClient) updateHandleResponse(resp *http.Response) (NodeTypesClientUpdateResponse, error) {
	result := NodeTypesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NodeType); err != nil {
		return NodeTypesClientUpdateResponse{}, err
	}
	return result, nil
}