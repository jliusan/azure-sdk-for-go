//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappcontainers

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

// DaprSubscriptionsClient contains the methods for the DaprSubscriptions group.
// Don't use this type directly, use NewDaprSubscriptionsClient() instead.
type DaprSubscriptionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDaprSubscriptionsClient creates a new instance of DaprSubscriptionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDaprSubscriptionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DaprSubscriptionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DaprSubscriptionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a Dapr subscription in a Managed Environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - environmentName - Name of the Managed Environment.
//   - name - Name of the Dapr subscription.
//   - daprSubscriptionEnvelope - Configuration details of the Dapr subscription.
//   - options - DaprSubscriptionsClientCreateOrUpdateOptions contains the optional parameters for the DaprSubscriptionsClient.CreateOrUpdate
//     method.
func (client *DaprSubscriptionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, environmentName string, name string, daprSubscriptionEnvelope DaprSubscription, options *DaprSubscriptionsClientCreateOrUpdateOptions) (DaprSubscriptionsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "DaprSubscriptionsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, environmentName, name, daprSubscriptionEnvelope, options)
	if err != nil {
		return DaprSubscriptionsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DaprSubscriptionsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return DaprSubscriptionsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DaprSubscriptionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, name string, daprSubscriptionEnvelope DaprSubscription, options *DaprSubscriptionsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{environmentName}/daprSubscriptions/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, daprSubscriptionEnvelope); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DaprSubscriptionsClient) createOrUpdateHandleResponse(resp *http.Response) (DaprSubscriptionsClientCreateOrUpdateResponse, error) {
	result := DaprSubscriptionsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DaprSubscription); err != nil {
		return DaprSubscriptionsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a Dapr subscription from a Managed Environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - environmentName - Name of the Managed Environment.
//   - name - Name of the Dapr subscription.
//   - options - DaprSubscriptionsClientDeleteOptions contains the optional parameters for the DaprSubscriptionsClient.Delete
//     method.
func (client *DaprSubscriptionsClient) Delete(ctx context.Context, resourceGroupName string, environmentName string, name string, options *DaprSubscriptionsClientDeleteOptions) (DaprSubscriptionsClientDeleteResponse, error) {
	var err error
	const operationName = "DaprSubscriptionsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, environmentName, name, options)
	if err != nil {
		return DaprSubscriptionsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DaprSubscriptionsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return DaprSubscriptionsClientDeleteResponse{}, err
	}
	return DaprSubscriptionsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DaprSubscriptionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, name string, options *DaprSubscriptionsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{environmentName}/daprSubscriptions/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a dapr subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - environmentName - Name of the Managed Environment.
//   - name - Name of the Dapr subscription.
//   - options - DaprSubscriptionsClientGetOptions contains the optional parameters for the DaprSubscriptionsClient.Get method.
func (client *DaprSubscriptionsClient) Get(ctx context.Context, resourceGroupName string, environmentName string, name string, options *DaprSubscriptionsClientGetOptions) (DaprSubscriptionsClientGetResponse, error) {
	var err error
	const operationName = "DaprSubscriptionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, environmentName, name, options)
	if err != nil {
		return DaprSubscriptionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DaprSubscriptionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DaprSubscriptionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DaprSubscriptionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, name string, options *DaprSubscriptionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{environmentName}/daprSubscriptions/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DaprSubscriptionsClient) getHandleResponse(resp *http.Response) (DaprSubscriptionsClientGetResponse, error) {
	result := DaprSubscriptionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DaprSubscription); err != nil {
		return DaprSubscriptionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get the Dapr subscriptions for a managed environment.
//
// Generated from API version 2023-11-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - environmentName - Name of the Managed Environment.
//   - options - DaprSubscriptionsClientListOptions contains the optional parameters for the DaprSubscriptionsClient.NewListPager
//     method.
func (client *DaprSubscriptionsClient) NewListPager(resourceGroupName string, environmentName string, options *DaprSubscriptionsClientListOptions) *runtime.Pager[DaprSubscriptionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DaprSubscriptionsClientListResponse]{
		More: func(page DaprSubscriptionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DaprSubscriptionsClientListResponse) (DaprSubscriptionsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DaprSubscriptionsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, environmentName, options)
			}, nil)
			if err != nil {
				return DaprSubscriptionsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *DaprSubscriptionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, options *DaprSubscriptionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{environmentName}/daprSubscriptions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DaprSubscriptionsClient) listHandleResponse(resp *http.Response) (DaprSubscriptionsClientListResponse, error) {
	result := DaprSubscriptionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DaprSubscriptionsCollection); err != nil {
		return DaprSubscriptionsClientListResponse{}, err
	}
	return result, nil
}