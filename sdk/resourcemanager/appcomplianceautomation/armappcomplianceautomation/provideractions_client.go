//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappcomplianceautomation

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ProviderActionsClient contains the methods for the ProviderActions group.
// Don't use this type directly, use NewProviderActionsClient() instead.
type ProviderActionsClient struct {
	internal *arm.Client
}

// NewProviderActionsClient creates a new instance of ProviderActionsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewProviderActionsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ProviderActionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ProviderActionsClient{
		internal: cl,
	}
	return client, nil
}

// CheckNameAvailability - Check if the given name is available for a report.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-27
//   - body - The content of the action request
//   - options - ProviderActionsClientCheckNameAvailabilityOptions contains the optional parameters for the ProviderActionsClient.CheckNameAvailability
//     method.
func (client *ProviderActionsClient) CheckNameAvailability(ctx context.Context, body CheckNameAvailabilityRequest, options *ProviderActionsClientCheckNameAvailabilityOptions) (ProviderActionsClientCheckNameAvailabilityResponse, error) {
	var err error
	const operationName = "ProviderActionsClient.CheckNameAvailability"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.checkNameAvailabilityCreateRequest(ctx, body, options)
	if err != nil {
		return ProviderActionsClientCheckNameAvailabilityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProviderActionsClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ProviderActionsClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.checkNameAvailabilityHandleResponse(httpResp)
	return resp, err
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *ProviderActionsClient) checkNameAvailabilityCreateRequest(ctx context.Context, body CheckNameAvailabilityRequest, options *ProviderActionsClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AppComplianceAutomation/checkNameAvailability"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *ProviderActionsClient) checkNameAvailabilityHandleResponse(resp *http.Response) (ProviderActionsClientCheckNameAvailabilityResponse, error) {
	result := ProviderActionsClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityResponse); err != nil {
		return ProviderActionsClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// GetCollectionCount - Get the count of reports.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-27
//   - body - The content of the action request
//   - options - ProviderActionsClientGetCollectionCountOptions contains the optional parameters for the ProviderActionsClient.GetCollectionCount
//     method.
func (client *ProviderActionsClient) GetCollectionCount(ctx context.Context, body GetCollectionCountRequest, options *ProviderActionsClientGetCollectionCountOptions) (ProviderActionsClientGetCollectionCountResponse, error) {
	var err error
	const operationName = "ProviderActionsClient.GetCollectionCount"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCollectionCountCreateRequest(ctx, body, options)
	if err != nil {
		return ProviderActionsClientGetCollectionCountResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProviderActionsClientGetCollectionCountResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ProviderActionsClientGetCollectionCountResponse{}, err
	}
	resp, err := client.getCollectionCountHandleResponse(httpResp)
	return resp, err
}

// getCollectionCountCreateRequest creates the GetCollectionCount request.
func (client *ProviderActionsClient) getCollectionCountCreateRequest(ctx context.Context, body GetCollectionCountRequest, options *ProviderActionsClientGetCollectionCountOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AppComplianceAutomation/getCollectionCount"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// getCollectionCountHandleResponse handles the GetCollectionCount response.
func (client *ProviderActionsClient) getCollectionCountHandleResponse(resp *http.Response) (ProviderActionsClientGetCollectionCountResponse, error) {
	result := ProviderActionsClientGetCollectionCountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GetCollectionCountResponse); err != nil {
		return ProviderActionsClientGetCollectionCountResponse{}, err
	}
	return result, nil
}

// GetOverviewStatus - Get the resource overview status.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-27
//   - body - The content of the action request
//   - options - ProviderActionsClientGetOverviewStatusOptions contains the optional parameters for the ProviderActionsClient.GetOverviewStatus
//     method.
func (client *ProviderActionsClient) GetOverviewStatus(ctx context.Context, body GetOverviewStatusRequest, options *ProviderActionsClientGetOverviewStatusOptions) (ProviderActionsClientGetOverviewStatusResponse, error) {
	var err error
	const operationName = "ProviderActionsClient.GetOverviewStatus"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getOverviewStatusCreateRequest(ctx, body, options)
	if err != nil {
		return ProviderActionsClientGetOverviewStatusResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProviderActionsClientGetOverviewStatusResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ProviderActionsClientGetOverviewStatusResponse{}, err
	}
	resp, err := client.getOverviewStatusHandleResponse(httpResp)
	return resp, err
}

// getOverviewStatusCreateRequest creates the GetOverviewStatus request.
func (client *ProviderActionsClient) getOverviewStatusCreateRequest(ctx context.Context, body GetOverviewStatusRequest, options *ProviderActionsClientGetOverviewStatusOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AppComplianceAutomation/getOverviewStatus"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// getOverviewStatusHandleResponse handles the GetOverviewStatus response.
func (client *ProviderActionsClient) getOverviewStatusHandleResponse(resp *http.Response) (ProviderActionsClientGetOverviewStatusResponse, error) {
	result := ProviderActionsClientGetOverviewStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GetOverviewStatusResponse); err != nil {
		return ProviderActionsClientGetOverviewStatusResponse{}, err
	}
	return result, nil
}

// ListInUseStorageAccounts - List the storage accounts which are in use by related reports
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-27
//   - body - The content of the action request
//   - options - ProviderActionsClientListInUseStorageAccountsOptions contains the optional parameters for the ProviderActionsClient.ListInUseStorageAccounts
//     method.
func (client *ProviderActionsClient) ListInUseStorageAccounts(ctx context.Context, body ListInUseStorageAccountsRequest, options *ProviderActionsClientListInUseStorageAccountsOptions) (ProviderActionsClientListInUseStorageAccountsResponse, error) {
	var err error
	const operationName = "ProviderActionsClient.ListInUseStorageAccounts"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listInUseStorageAccountsCreateRequest(ctx, body, options)
	if err != nil {
		return ProviderActionsClientListInUseStorageAccountsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProviderActionsClientListInUseStorageAccountsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ProviderActionsClientListInUseStorageAccountsResponse{}, err
	}
	resp, err := client.listInUseStorageAccountsHandleResponse(httpResp)
	return resp, err
}

// listInUseStorageAccountsCreateRequest creates the ListInUseStorageAccounts request.
func (client *ProviderActionsClient) listInUseStorageAccountsCreateRequest(ctx context.Context, body ListInUseStorageAccountsRequest, options *ProviderActionsClientListInUseStorageAccountsOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AppComplianceAutomation/listInUseStorageAccounts"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// listInUseStorageAccountsHandleResponse handles the ListInUseStorageAccounts response.
func (client *ProviderActionsClient) listInUseStorageAccountsHandleResponse(resp *http.Response) (ProviderActionsClientListInUseStorageAccountsResponse, error) {
	result := ProviderActionsClientListInUseStorageAccountsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListInUseStorageAccountsResponse); err != nil {
		return ProviderActionsClientListInUseStorageAccountsResponse{}, err
	}
	return result, nil
}

// BeginOnboard - Onboard given subscriptions to Microsoft.AppComplianceAutomation provider.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-27
//   - body - The content of the action request
//   - options - ProviderActionsClientBeginOnboardOptions contains the optional parameters for the ProviderActionsClient.BeginOnboard
//     method.
func (client *ProviderActionsClient) BeginOnboard(ctx context.Context, body OnboardRequest, options *ProviderActionsClientBeginOnboardOptions) (*runtime.Poller[ProviderActionsClientOnboardResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.onboard(ctx, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ProviderActionsClientOnboardResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ProviderActionsClientOnboardResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Onboard - Onboard given subscriptions to Microsoft.AppComplianceAutomation provider.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-27
func (client *ProviderActionsClient) onboard(ctx context.Context, body OnboardRequest, options *ProviderActionsClientBeginOnboardOptions) (*http.Response, error) {
	var err error
	const operationName = "ProviderActionsClient.BeginOnboard"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.onboardCreateRequest(ctx, body, options)
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

// onboardCreateRequest creates the Onboard request.
func (client *ProviderActionsClient) onboardCreateRequest(ctx context.Context, body OnboardRequest, options *ProviderActionsClientBeginOnboardOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AppComplianceAutomation/onboard"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginTriggerEvaluation - Trigger quick evaluation for the given subscriptions.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-27
//   - body - The content of the action request
//   - options - ProviderActionsClientBeginTriggerEvaluationOptions contains the optional parameters for the ProviderActionsClient.BeginTriggerEvaluation
//     method.
func (client *ProviderActionsClient) BeginTriggerEvaluation(ctx context.Context, body TriggerEvaluationRequest, options *ProviderActionsClientBeginTriggerEvaluationOptions) (*runtime.Poller[ProviderActionsClientTriggerEvaluationResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.triggerEvaluation(ctx, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ProviderActionsClientTriggerEvaluationResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ProviderActionsClientTriggerEvaluationResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// TriggerEvaluation - Trigger quick evaluation for the given subscriptions.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-27
func (client *ProviderActionsClient) triggerEvaluation(ctx context.Context, body TriggerEvaluationRequest, options *ProviderActionsClientBeginTriggerEvaluationOptions) (*http.Response, error) {
	var err error
	const operationName = "ProviderActionsClient.BeginTriggerEvaluation"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.triggerEvaluationCreateRequest(ctx, body, options)
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

// triggerEvaluationCreateRequest creates the TriggerEvaluation request.
func (client *ProviderActionsClient) triggerEvaluationCreateRequest(ctx context.Context, body TriggerEvaluationRequest, options *ProviderActionsClientBeginTriggerEvaluationOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AppComplianceAutomation/triggerEvaluation"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}