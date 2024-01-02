//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// ProjectAllowedEnvironmentTypesServer is a fake server for instances of the armdevcenter.ProjectAllowedEnvironmentTypesClient type.
type ProjectAllowedEnvironmentTypesServer struct {
	// Get is the fake for method ProjectAllowedEnvironmentTypesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, projectName string, environmentTypeName string, options *armdevcenter.ProjectAllowedEnvironmentTypesClientGetOptions) (resp azfake.Responder[armdevcenter.ProjectAllowedEnvironmentTypesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ProjectAllowedEnvironmentTypesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, projectName string, options *armdevcenter.ProjectAllowedEnvironmentTypesClientListOptions) (resp azfake.PagerResponder[armdevcenter.ProjectAllowedEnvironmentTypesClientListResponse])
}

// NewProjectAllowedEnvironmentTypesServerTransport creates a new instance of ProjectAllowedEnvironmentTypesServerTransport with the provided implementation.
// The returned ProjectAllowedEnvironmentTypesServerTransport instance is connected to an instance of armdevcenter.ProjectAllowedEnvironmentTypesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewProjectAllowedEnvironmentTypesServerTransport(srv *ProjectAllowedEnvironmentTypesServer) *ProjectAllowedEnvironmentTypesServerTransport {
	return &ProjectAllowedEnvironmentTypesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armdevcenter.ProjectAllowedEnvironmentTypesClientListResponse]](),
	}
}

// ProjectAllowedEnvironmentTypesServerTransport connects instances of armdevcenter.ProjectAllowedEnvironmentTypesClient to instances of ProjectAllowedEnvironmentTypesServer.
// Don't use this type directly, use NewProjectAllowedEnvironmentTypesServerTransport instead.
type ProjectAllowedEnvironmentTypesServerTransport struct {
	srv          *ProjectAllowedEnvironmentTypesServer
	newListPager *tracker[azfake.PagerResponder[armdevcenter.ProjectAllowedEnvironmentTypesClientListResponse]]
}

// Do implements the policy.Transporter interface for ProjectAllowedEnvironmentTypesServerTransport.
func (p *ProjectAllowedEnvironmentTypesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ProjectAllowedEnvironmentTypesClient.Get":
		resp, err = p.dispatchGet(req)
	case "ProjectAllowedEnvironmentTypesClient.NewListPager":
		resp, err = p.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *ProjectAllowedEnvironmentTypesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/projects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/allowedEnvironmentTypes/(?P<environmentTypeName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	projectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("projectName")])
	if err != nil {
		return nil, err
	}
	environmentTypeNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("environmentTypeName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, projectNameParam, environmentTypeNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AllowedEnvironmentType, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *ProjectAllowedEnvironmentTypesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := p.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/projects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/allowedEnvironmentTypes`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		projectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("projectName")])
		if err != nil {
			return nil, err
		}
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *armdevcenter.ProjectAllowedEnvironmentTypesClientListOptions
		if topParam != nil {
			options = &armdevcenter.ProjectAllowedEnvironmentTypesClientListOptions{
				Top: topParam,
			}
		}
		resp := p.srv.NewListPager(resourceGroupNameParam, projectNameParam, options)
		newListPager = &resp
		p.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armdevcenter.ProjectAllowedEnvironmentTypesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		p.newListPager.remove(req)
	}
	return resp, nil
}