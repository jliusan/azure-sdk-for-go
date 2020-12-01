package qnamaker

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// EndpointKeysClient is the an API for QnAMaker Service
type EndpointKeysClient struct {
	BaseClient
}

// NewEndpointKeysClient creates an instance of the EndpointKeysClient client.
func NewEndpointKeysClient(endpoint string) EndpointKeysClient {
	return EndpointKeysClient{New(endpoint)}
}

// GetKeys sends the get keys request.
func (client EndpointKeysClient) GetKeys(ctx context.Context) (result EndpointKeysDTO, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EndpointKeysClient.GetKeys")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetKeysPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "qnamaker.EndpointKeysClient", "GetKeys", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetKeysSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "qnamaker.EndpointKeysClient", "GetKeys", resp, "Failure sending request")
		return
	}

	result, err = client.GetKeysResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "qnamaker.EndpointKeysClient", "GetKeys", resp, "Failure responding to request")
	}

	return
}

// GetKeysPreparer prepares the GetKeys request.
func (client EndpointKeysClient) GetKeysPreparer(ctx context.Context) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}/qnamaker/v5.0-preview.1", urlParameters),
		autorest.WithPath("/endpointkeys"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetKeysSender sends the GetKeys request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointKeysClient) GetKeysSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetKeysResponder handles the response to the GetKeys request. The method always
// closes the http.Response Body.
func (client EndpointKeysClient) GetKeysResponder(resp *http.Response) (result EndpointKeysDTO, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// RefreshKeys sends the refresh keys request.
// Parameters:
// keyType - type of Key
func (client EndpointKeysClient) RefreshKeys(ctx context.Context, keyType string) (result EndpointKeysDTO, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EndpointKeysClient.RefreshKeys")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RefreshKeysPreparer(ctx, keyType)
	if err != nil {
		err = autorest.NewErrorWithError(err, "qnamaker.EndpointKeysClient", "RefreshKeys", nil, "Failure preparing request")
		return
	}

	resp, err := client.RefreshKeysSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "qnamaker.EndpointKeysClient", "RefreshKeys", resp, "Failure sending request")
		return
	}

	result, err = client.RefreshKeysResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "qnamaker.EndpointKeysClient", "RefreshKeys", resp, "Failure responding to request")
	}

	return
}

// RefreshKeysPreparer prepares the RefreshKeys request.
func (client EndpointKeysClient) RefreshKeysPreparer(ctx context.Context, keyType string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"keyType": autorest.Encode("path", keyType),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPatch(),
		autorest.WithCustomBaseURL("{Endpoint}/qnamaker/v5.0-preview.1", urlParameters),
		autorest.WithPathParameters("/endpointkeys/{keyType}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RefreshKeysSender sends the RefreshKeys request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointKeysClient) RefreshKeysSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// RefreshKeysResponder handles the response to the RefreshKeys request. The method always
// closes the http.Response Body.
func (client EndpointKeysClient) RefreshKeysResponder(resp *http.Response) (result EndpointKeysDTO, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
