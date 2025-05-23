// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatafactory

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

// TriggerRunsClient contains the methods for the TriggerRuns group.
// Don't use this type directly, use NewTriggerRunsClient() instead.
type TriggerRunsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTriggerRunsClient creates a new instance of TriggerRunsClient with the specified values.
//   - subscriptionID - The subscription identifier.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTriggerRunsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TriggerRunsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TriggerRunsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Cancel - Cancel a single trigger instance by runId.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - triggerName - The trigger name.
//   - runID - The pipeline run identifier.
//   - options - TriggerRunsClientCancelOptions contains the optional parameters for the TriggerRunsClient.Cancel method.
func (client *TriggerRunsClient) Cancel(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, runID string, options *TriggerRunsClientCancelOptions) (TriggerRunsClientCancelResponse, error) {
	var err error
	const operationName = "TriggerRunsClient.Cancel"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.cancelCreateRequest(ctx, resourceGroupName, factoryName, triggerName, runID, options)
	if err != nil {
		return TriggerRunsClientCancelResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TriggerRunsClientCancelResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TriggerRunsClientCancelResponse{}, err
	}
	return TriggerRunsClientCancelResponse{}, nil
}

// cancelCreateRequest creates the Cancel request.
func (client *TriggerRunsClient) cancelCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, runID string, _ *TriggerRunsClientCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/triggers/{triggerName}/triggerRuns/{runId}/cancel"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	if runID == "" {
		return nil, errors.New("parameter runID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// QueryByFactory - Query trigger runs.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - filterParameters - Parameters to filter the pipeline run.
//   - options - TriggerRunsClientQueryByFactoryOptions contains the optional parameters for the TriggerRunsClient.QueryByFactory
//     method.
func (client *TriggerRunsClient) QueryByFactory(ctx context.Context, resourceGroupName string, factoryName string, filterParameters RunFilterParameters, options *TriggerRunsClientQueryByFactoryOptions) (TriggerRunsClientQueryByFactoryResponse, error) {
	var err error
	const operationName = "TriggerRunsClient.QueryByFactory"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.queryByFactoryCreateRequest(ctx, resourceGroupName, factoryName, filterParameters, options)
	if err != nil {
		return TriggerRunsClientQueryByFactoryResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TriggerRunsClientQueryByFactoryResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TriggerRunsClientQueryByFactoryResponse{}, err
	}
	resp, err := client.queryByFactoryHandleResponse(httpResp)
	return resp, err
}

// queryByFactoryCreateRequest creates the QueryByFactory request.
func (client *TriggerRunsClient) queryByFactoryCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, filterParameters RunFilterParameters, _ *TriggerRunsClientQueryByFactoryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/queryTriggerRuns"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, filterParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// queryByFactoryHandleResponse handles the QueryByFactory response.
func (client *TriggerRunsClient) queryByFactoryHandleResponse(resp *http.Response) (TriggerRunsClientQueryByFactoryResponse, error) {
	result := TriggerRunsClientQueryByFactoryResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TriggerRunsQueryResponse); err != nil {
		return TriggerRunsClientQueryByFactoryResponse{}, err
	}
	return result, nil
}

// Rerun - Rerun single trigger instance by runId.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - triggerName - The trigger name.
//   - runID - The pipeline run identifier.
//   - options - TriggerRunsClientRerunOptions contains the optional parameters for the TriggerRunsClient.Rerun method.
func (client *TriggerRunsClient) Rerun(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, runID string, options *TriggerRunsClientRerunOptions) (TriggerRunsClientRerunResponse, error) {
	var err error
	const operationName = "TriggerRunsClient.Rerun"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.rerunCreateRequest(ctx, resourceGroupName, factoryName, triggerName, runID, options)
	if err != nil {
		return TriggerRunsClientRerunResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TriggerRunsClientRerunResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TriggerRunsClientRerunResponse{}, err
	}
	return TriggerRunsClientRerunResponse{}, nil
}

// rerunCreateRequest creates the Rerun request.
func (client *TriggerRunsClient) rerunCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, runID string, _ *TriggerRunsClientRerunOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/triggers/{triggerName}/triggerRuns/{runId}/rerun"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	if runID == "" {
		return nil, errors.New("parameter runID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
