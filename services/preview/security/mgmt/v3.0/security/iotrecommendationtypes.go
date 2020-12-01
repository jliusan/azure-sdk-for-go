package security

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// IotRecommendationTypesClient is the API spec for Microsoft.Security (Azure Security Center) resource provider
type IotRecommendationTypesClient struct {
	BaseClient
}

// NewIotRecommendationTypesClient creates an instance of the IotRecommendationTypesClient client.
func NewIotRecommendationTypesClient(subscriptionID string, ascLocation string) IotRecommendationTypesClient {
	return NewIotRecommendationTypesClientWithBaseURI(DefaultBaseURI, subscriptionID, ascLocation)
}

// NewIotRecommendationTypesClientWithBaseURI creates an instance of the IotRecommendationTypesClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewIotRecommendationTypesClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) IotRecommendationTypesClient {
	return IotRecommendationTypesClient{NewWithBaseURI(baseURI, subscriptionID, ascLocation)}
}

// Get get IoT recommendation type
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// solutionName - the name of the IoT Security solution.
// iotRecommendationTypeName - name of the recommendation type
func (client IotRecommendationTypesClient) Get(ctx context.Context, resourceGroupName string, solutionName string, iotRecommendationTypeName string) (result IotRecommendationType, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IotRecommendationTypesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.IotRecommendationTypesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, solutionName, iotRecommendationTypeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.IotRecommendationTypesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.IotRecommendationTypesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.IotRecommendationTypesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client IotRecommendationTypesClient) GetPreparer(ctx context.Context, resourceGroupName string, solutionName string, iotRecommendationTypeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"iotRecommendationTypeName": autorest.Encode("path", iotRecommendationTypeName),
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"solutionName":              autorest.Encode("path", solutionName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}/iotRecommendationTypes/{iotRecommendationTypeName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client IotRecommendationTypesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client IotRecommendationTypesClient) GetResponder(resp *http.Response) (result IotRecommendationType, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get1 get IoT recommendation type
// Parameters:
// iotRecommendationTypeName - name of the recommendation type
func (client IotRecommendationTypesClient) Get1(ctx context.Context, iotRecommendationTypeName string) (result IotRecommendationType, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IotRecommendationTypesClient.Get1")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.IotRecommendationTypesClient", "Get1", err.Error())
	}

	req, err := client.Get1Preparer(ctx, iotRecommendationTypeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.IotRecommendationTypesClient", "Get1", nil, "Failure preparing request")
		return
	}

	resp, err := client.Get1Sender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.IotRecommendationTypesClient", "Get1", resp, "Failure sending request")
		return
	}

	result, err = client.Get1Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.IotRecommendationTypesClient", "Get1", resp, "Failure responding to request")
	}

	return
}

// Get1Preparer prepares the Get1 request.
func (client IotRecommendationTypesClient) Get1Preparer(ctx context.Context, iotRecommendationTypeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"iotRecommendationTypeName": autorest.Encode("path", iotRecommendationTypeName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-06-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/iotRecommendationTypes/{iotRecommendationTypeName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Get1Sender sends the Get1 request. The method will close the
// http.Response Body if it receives an error.
func (client IotRecommendationTypesClient) Get1Sender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// Get1Responder handles the response to the Get1 request. The method always
// closes the http.Response Body.
func (client IotRecommendationTypesClient) Get1Responder(resp *http.Response) (result IotRecommendationType, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list IoT recommendation types
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// solutionName - the name of the IoT Security solution.
func (client IotRecommendationTypesClient) List(ctx context.Context, resourceGroupName string, solutionName string) (result IotRecommendationTypeList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IotRecommendationTypesClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.IotRecommendationTypesClient", "List", err.Error())
	}

	req, err := client.ListPreparer(ctx, resourceGroupName, solutionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.IotRecommendationTypesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.IotRecommendationTypesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.IotRecommendationTypesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client IotRecommendationTypesClient) ListPreparer(ctx context.Context, resourceGroupName string, solutionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"solutionName":      autorest.Encode("path", solutionName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}/iotRecommendationTypes", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client IotRecommendationTypesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client IotRecommendationTypesClient) ListResponder(resp *http.Response) (result IotRecommendationTypeList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List1 list IoT recommendation types
func (client IotRecommendationTypesClient) List1(ctx context.Context) (result IotRecommendationTypeList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IotRecommendationTypesClient.List1")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.IotRecommendationTypesClient", "List1", err.Error())
	}

	req, err := client.List1Preparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.IotRecommendationTypesClient", "List1", nil, "Failure preparing request")
		return
	}

	resp, err := client.List1Sender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.IotRecommendationTypesClient", "List1", resp, "Failure sending request")
		return
	}

	result, err = client.List1Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.IotRecommendationTypesClient", "List1", resp, "Failure responding to request")
	}

	return
}

// List1Preparer prepares the List1 request.
func (client IotRecommendationTypesClient) List1Preparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-06-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/iotRecommendationTypes", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// List1Sender sends the List1 request. The method will close the
// http.Response Body if it receives an error.
func (client IotRecommendationTypesClient) List1Sender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// List1Responder handles the response to the List1 request. The method always
// closes the http.Response Body.
func (client IotRecommendationTypesClient) List1Responder(resp *http.Response) (result IotRecommendationTypeList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
