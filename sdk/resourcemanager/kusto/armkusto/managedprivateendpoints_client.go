//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkusto

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

// ManagedPrivateEndpointsClient contains the methods for the ManagedPrivateEndpoints group.
// Don't use this type directly, use NewManagedPrivateEndpointsClient() instead.
type ManagedPrivateEndpointsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewManagedPrivateEndpointsClient creates a new instance of ManagedPrivateEndpointsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagedPrivateEndpointsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedPrivateEndpointsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagedPrivateEndpointsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CheckNameAvailability - Checks that the managed private endpoints resource name is valid and is not already in use.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-13
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the Kusto cluster.
//   - resourceName - The name of the resource.
//   - options - ManagedPrivateEndpointsClientCheckNameAvailabilityOptions contains the optional parameters for the ManagedPrivateEndpointsClient.CheckNameAvailability
//     method.
func (client *ManagedPrivateEndpointsClient) CheckNameAvailability(ctx context.Context, resourceGroupName string, clusterName string, resourceName ManagedPrivateEndpointsCheckNameRequest, options *ManagedPrivateEndpointsClientCheckNameAvailabilityOptions) (ManagedPrivateEndpointsClientCheckNameAvailabilityResponse, error) {
	var err error
	const operationName = "ManagedPrivateEndpointsClient.CheckNameAvailability"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.checkNameAvailabilityCreateRequest(ctx, resourceGroupName, clusterName, resourceName, options)
	if err != nil {
		return ManagedPrivateEndpointsClientCheckNameAvailabilityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedPrivateEndpointsClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagedPrivateEndpointsClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.checkNameAvailabilityHandleResponse(httpResp)
	return resp, err
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *ManagedPrivateEndpointsClient) checkNameAvailabilityCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, resourceName ManagedPrivateEndpointsCheckNameRequest, options *ManagedPrivateEndpointsClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/managedPrivateEndpointsCheckNameAvailability"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-13")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resourceName); err != nil {
		return nil, err
	}
	return req, nil
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *ManagedPrivateEndpointsClient) checkNameAvailabilityHandleResponse(resp *http.Response) (ManagedPrivateEndpointsClientCheckNameAvailabilityResponse, error) {
	result := ManagedPrivateEndpointsClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameResult); err != nil {
		return ManagedPrivateEndpointsClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// BeginCreateOrUpdate - Creates a managed private endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-13
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the Kusto cluster.
//   - managedPrivateEndpointName - The name of the managed private endpoint.
//   - parameters - The managed private endpoint parameters.
//   - options - ManagedPrivateEndpointsClientBeginCreateOrUpdateOptions contains the optional parameters for the ManagedPrivateEndpointsClient.BeginCreateOrUpdate
//     method.
func (client *ManagedPrivateEndpointsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, managedPrivateEndpointName string, parameters ManagedPrivateEndpoint, options *ManagedPrivateEndpointsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ManagedPrivateEndpointsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, clusterName, managedPrivateEndpointName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ManagedPrivateEndpointsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ManagedPrivateEndpointsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates a managed private endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-13
func (client *ManagedPrivateEndpointsClient) createOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, managedPrivateEndpointName string, parameters ManagedPrivateEndpoint, options *ManagedPrivateEndpointsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ManagedPrivateEndpointsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, clusterName, managedPrivateEndpointName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ManagedPrivateEndpointsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, managedPrivateEndpointName string, parameters ManagedPrivateEndpoint, options *ManagedPrivateEndpointsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/managedPrivateEndpoints/{managedPrivateEndpointName}"
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
	if managedPrivateEndpointName == "" {
		return nil, errors.New("parameter managedPrivateEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedPrivateEndpointName}", url.PathEscape(managedPrivateEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-13")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a managed private endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-13
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the Kusto cluster.
//   - managedPrivateEndpointName - The name of the managed private endpoint.
//   - options - ManagedPrivateEndpointsClientBeginDeleteOptions contains the optional parameters for the ManagedPrivateEndpointsClient.BeginDelete
//     method.
func (client *ManagedPrivateEndpointsClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, managedPrivateEndpointName string, options *ManagedPrivateEndpointsClientBeginDeleteOptions) (*runtime.Poller[ManagedPrivateEndpointsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, managedPrivateEndpointName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ManagedPrivateEndpointsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ManagedPrivateEndpointsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes a managed private endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-13
func (client *ManagedPrivateEndpointsClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, managedPrivateEndpointName string, options *ManagedPrivateEndpointsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ManagedPrivateEndpointsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, managedPrivateEndpointName, options)
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
func (client *ManagedPrivateEndpointsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, managedPrivateEndpointName string, options *ManagedPrivateEndpointsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/managedPrivateEndpoints/{managedPrivateEndpointName}"
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
	if managedPrivateEndpointName == "" {
		return nil, errors.New("parameter managedPrivateEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedPrivateEndpointName}", url.PathEscape(managedPrivateEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-13")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a managed private endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-13
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the Kusto cluster.
//   - managedPrivateEndpointName - The name of the managed private endpoint.
//   - options - ManagedPrivateEndpointsClientGetOptions contains the optional parameters for the ManagedPrivateEndpointsClient.Get
//     method.
func (client *ManagedPrivateEndpointsClient) Get(ctx context.Context, resourceGroupName string, clusterName string, managedPrivateEndpointName string, options *ManagedPrivateEndpointsClientGetOptions) (ManagedPrivateEndpointsClientGetResponse, error) {
	var err error
	const operationName = "ManagedPrivateEndpointsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, managedPrivateEndpointName, options)
	if err != nil {
		return ManagedPrivateEndpointsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedPrivateEndpointsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagedPrivateEndpointsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ManagedPrivateEndpointsClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, managedPrivateEndpointName string, options *ManagedPrivateEndpointsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/managedPrivateEndpoints/{managedPrivateEndpointName}"
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
	if managedPrivateEndpointName == "" {
		return nil, errors.New("parameter managedPrivateEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedPrivateEndpointName}", url.PathEscape(managedPrivateEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-13")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagedPrivateEndpointsClient) getHandleResponse(resp *http.Response) (ManagedPrivateEndpointsClientGetResponse, error) {
	result := ManagedPrivateEndpointsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedPrivateEndpoint); err != nil {
		return ManagedPrivateEndpointsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Returns the list of managed private endpoints.
//
// Generated from API version 2024-04-13
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the Kusto cluster.
//   - options - ManagedPrivateEndpointsClientListOptions contains the optional parameters for the ManagedPrivateEndpointsClient.NewListPager
//     method.
func (client *ManagedPrivateEndpointsClient) NewListPager(resourceGroupName string, clusterName string, options *ManagedPrivateEndpointsClientListOptions) *runtime.Pager[ManagedPrivateEndpointsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedPrivateEndpointsClientListResponse]{
		More: func(page ManagedPrivateEndpointsClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ManagedPrivateEndpointsClientListResponse) (ManagedPrivateEndpointsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ManagedPrivateEndpointsClient.NewListPager")
			req, err := client.listCreateRequest(ctx, resourceGroupName, clusterName, options)
			if err != nil {
				return ManagedPrivateEndpointsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ManagedPrivateEndpointsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ManagedPrivateEndpointsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ManagedPrivateEndpointsClient) listCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *ManagedPrivateEndpointsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/managedPrivateEndpoints"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-13")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ManagedPrivateEndpointsClient) listHandleResponse(resp *http.Response) (ManagedPrivateEndpointsClientListResponse, error) {
	result := ManagedPrivateEndpointsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedPrivateEndpointListResult); err != nil {
		return ManagedPrivateEndpointsClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates a managed private endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-13
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the Kusto cluster.
//   - managedPrivateEndpointName - The name of the managed private endpoint.
//   - parameters - The managed private endpoint parameters.
//   - options - ManagedPrivateEndpointsClientBeginUpdateOptions contains the optional parameters for the ManagedPrivateEndpointsClient.BeginUpdate
//     method.
func (client *ManagedPrivateEndpointsClient) BeginUpdate(ctx context.Context, resourceGroupName string, clusterName string, managedPrivateEndpointName string, parameters ManagedPrivateEndpoint, options *ManagedPrivateEndpointsClientBeginUpdateOptions) (*runtime.Poller[ManagedPrivateEndpointsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, clusterName, managedPrivateEndpointName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ManagedPrivateEndpointsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ManagedPrivateEndpointsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Updates a managed private endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-13
func (client *ManagedPrivateEndpointsClient) update(ctx context.Context, resourceGroupName string, clusterName string, managedPrivateEndpointName string, parameters ManagedPrivateEndpoint, options *ManagedPrivateEndpointsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ManagedPrivateEndpointsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, clusterName, managedPrivateEndpointName, parameters, options)
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

// updateCreateRequest creates the Update request.
func (client *ManagedPrivateEndpointsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, managedPrivateEndpointName string, parameters ManagedPrivateEndpoint, options *ManagedPrivateEndpointsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/managedPrivateEndpoints/{managedPrivateEndpointName}"
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
	if managedPrivateEndpointName == "" {
		return nil, errors.New("parameter managedPrivateEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedPrivateEndpointName}", url.PathEscape(managedPrivateEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-13")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}
