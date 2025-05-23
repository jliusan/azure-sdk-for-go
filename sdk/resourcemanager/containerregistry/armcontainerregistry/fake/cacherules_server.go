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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
	"net/http"
	"net/url"
	"regexp"
)

// CacheRulesServer is a fake server for instances of the armcontainerregistry.CacheRulesClient type.
type CacheRulesServer struct {
	// BeginCreate is the fake for method CacheRulesClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, registryName string, cacheRuleName string, cacheRuleCreateParameters armcontainerregistry.CacheRule, options *armcontainerregistry.CacheRulesClientBeginCreateOptions) (resp azfake.PollerResponder[armcontainerregistry.CacheRulesClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method CacheRulesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, registryName string, cacheRuleName string, options *armcontainerregistry.CacheRulesClientBeginDeleteOptions) (resp azfake.PollerResponder[armcontainerregistry.CacheRulesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method CacheRulesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, registryName string, cacheRuleName string, options *armcontainerregistry.CacheRulesClientGetOptions) (resp azfake.Responder[armcontainerregistry.CacheRulesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method CacheRulesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, registryName string, options *armcontainerregistry.CacheRulesClientListOptions) (resp azfake.PagerResponder[armcontainerregistry.CacheRulesClientListResponse])

	// BeginUpdate is the fake for method CacheRulesClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginUpdate func(ctx context.Context, resourceGroupName string, registryName string, cacheRuleName string, cacheRuleUpdateParameters armcontainerregistry.CacheRuleUpdateParameters, options *armcontainerregistry.CacheRulesClientBeginUpdateOptions) (resp azfake.PollerResponder[armcontainerregistry.CacheRulesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewCacheRulesServerTransport creates a new instance of CacheRulesServerTransport with the provided implementation.
// The returned CacheRulesServerTransport instance is connected to an instance of armcontainerregistry.CacheRulesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCacheRulesServerTransport(srv *CacheRulesServer) *CacheRulesServerTransport {
	return &CacheRulesServerTransport{
		srv:          srv,
		beginCreate:  newTracker[azfake.PollerResponder[armcontainerregistry.CacheRulesClientCreateResponse]](),
		beginDelete:  newTracker[azfake.PollerResponder[armcontainerregistry.CacheRulesClientDeleteResponse]](),
		newListPager: newTracker[azfake.PagerResponder[armcontainerregistry.CacheRulesClientListResponse]](),
		beginUpdate:  newTracker[azfake.PollerResponder[armcontainerregistry.CacheRulesClientUpdateResponse]](),
	}
}

// CacheRulesServerTransport connects instances of armcontainerregistry.CacheRulesClient to instances of CacheRulesServer.
// Don't use this type directly, use NewCacheRulesServerTransport instead.
type CacheRulesServerTransport struct {
	srv          *CacheRulesServer
	beginCreate  *tracker[azfake.PollerResponder[armcontainerregistry.CacheRulesClientCreateResponse]]
	beginDelete  *tracker[azfake.PollerResponder[armcontainerregistry.CacheRulesClientDeleteResponse]]
	newListPager *tracker[azfake.PagerResponder[armcontainerregistry.CacheRulesClientListResponse]]
	beginUpdate  *tracker[azfake.PollerResponder[armcontainerregistry.CacheRulesClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for CacheRulesServerTransport.
func (c *CacheRulesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return c.dispatchToMethodFake(req, method)
}

func (c *CacheRulesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if cacheRulesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = cacheRulesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "CacheRulesClient.BeginCreate":
				res.resp, res.err = c.dispatchBeginCreate(req)
			case "CacheRulesClient.BeginDelete":
				res.resp, res.err = c.dispatchBeginDelete(req)
			case "CacheRulesClient.Get":
				res.resp, res.err = c.dispatchGet(req)
			case "CacheRulesClient.NewListPager":
				res.resp, res.err = c.dispatchNewListPager(req)
			case "CacheRulesClient.BeginUpdate":
				res.resp, res.err = c.dispatchBeginUpdate(req)
			default:
				res.err = fmt.Errorf("unhandled API %s", method)
			}

		}
		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (c *CacheRulesServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := c.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerRegistry/registries/(?P<registryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cacheRules/(?P<cacheRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcontainerregistry.CacheRule](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		registryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("registryName")])
		if err != nil {
			return nil, err
		}
		cacheRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cacheRuleName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginCreate(req.Context(), resourceGroupNameParam, registryNameParam, cacheRuleNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		c.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		c.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		c.beginCreate.remove(req)
	}

	return resp, nil
}

func (c *CacheRulesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := c.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerRegistry/registries/(?P<registryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cacheRules/(?P<cacheRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		registryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("registryName")])
		if err != nil {
			return nil, err
		}
		cacheRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cacheRuleName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), resourceGroupNameParam, registryNameParam, cacheRuleNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		c.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		c.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		c.beginDelete.remove(req)
	}

	return resp, nil
}

func (c *CacheRulesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerRegistry/registries/(?P<registryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cacheRules/(?P<cacheRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	registryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("registryName")])
	if err != nil {
		return nil, err
	}
	cacheRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cacheRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, registryNameParam, cacheRuleNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CacheRule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CacheRulesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := c.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerRegistry/registries/(?P<registryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cacheRules`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		registryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("registryName")])
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListPager(resourceGroupNameParam, registryNameParam, nil)
		newListPager = &resp
		c.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armcontainerregistry.CacheRulesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		c.newListPager.remove(req)
	}
	return resp, nil
}

func (c *CacheRulesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := c.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerRegistry/registries/(?P<registryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cacheRules/(?P<cacheRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcontainerregistry.CacheRuleUpdateParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		registryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("registryName")])
		if err != nil {
			return nil, err
		}
		cacheRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cacheRuleName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginUpdate(req.Context(), resourceGroupNameParam, registryNameParam, cacheRuleNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		c.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		c.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		c.beginUpdate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to CacheRulesServerTransport
var cacheRulesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
