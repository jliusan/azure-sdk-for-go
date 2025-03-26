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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// ManagedHsmsServer is a fake server for instances of the armkeyvault.ManagedHsmsClient type.
type ManagedHsmsServer struct {
	// CheckMhsmNameAvailability is the fake for method ManagedHsmsClient.CheckMhsmNameAvailability
	// HTTP status codes to indicate success: http.StatusOK
	CheckMhsmNameAvailability func(ctx context.Context, mhsmName armkeyvault.CheckMhsmNameAvailabilityParameters, options *armkeyvault.ManagedHsmsClientCheckMhsmNameAvailabilityOptions) (resp azfake.Responder[armkeyvault.ManagedHsmsClientCheckMhsmNameAvailabilityResponse], errResp azfake.ErrorResponder)

	// BeginCreateOrUpdate is the fake for method ManagedHsmsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, name string, parameters armkeyvault.ManagedHsm, options *armkeyvault.ManagedHsmsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armkeyvault.ManagedHsmsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ManagedHsmsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, name string, options *armkeyvault.ManagedHsmsClientBeginDeleteOptions) (resp azfake.PollerResponder[armkeyvault.ManagedHsmsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ManagedHsmsClient.Get
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Get func(ctx context.Context, resourceGroupName string, name string, options *armkeyvault.ManagedHsmsClientGetOptions) (resp azfake.Responder[armkeyvault.ManagedHsmsClientGetResponse], errResp azfake.ErrorResponder)

	// GetDeleted is the fake for method ManagedHsmsClient.GetDeleted
	// HTTP status codes to indicate success: http.StatusOK
	GetDeleted func(ctx context.Context, name string, location string, options *armkeyvault.ManagedHsmsClientGetDeletedOptions) (resp azfake.Responder[armkeyvault.ManagedHsmsClientGetDeletedResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method ManagedHsmsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armkeyvault.ManagedHsmsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armkeyvault.ManagedHsmsClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method ManagedHsmsClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armkeyvault.ManagedHsmsClientListBySubscriptionOptions) (resp azfake.PagerResponder[armkeyvault.ManagedHsmsClientListBySubscriptionResponse])

	// NewListDeletedPager is the fake for method ManagedHsmsClient.NewListDeletedPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListDeletedPager func(options *armkeyvault.ManagedHsmsClientListDeletedOptions) (resp azfake.PagerResponder[armkeyvault.ManagedHsmsClientListDeletedResponse])

	// BeginPurgeDeleted is the fake for method ManagedHsmsClient.BeginPurgeDeleted
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginPurgeDeleted func(ctx context.Context, name string, location string, options *armkeyvault.ManagedHsmsClientBeginPurgeDeletedOptions) (resp azfake.PollerResponder[armkeyvault.ManagedHsmsClientPurgeDeletedResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method ManagedHsmsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, name string, parameters armkeyvault.ManagedHsm, options *armkeyvault.ManagedHsmsClientBeginUpdateOptions) (resp azfake.PollerResponder[armkeyvault.ManagedHsmsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewManagedHsmsServerTransport creates a new instance of ManagedHsmsServerTransport with the provided implementation.
// The returned ManagedHsmsServerTransport instance is connected to an instance of armkeyvault.ManagedHsmsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewManagedHsmsServerTransport(srv *ManagedHsmsServer) *ManagedHsmsServerTransport {
	return &ManagedHsmsServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armkeyvault.ManagedHsmsClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armkeyvault.ManagedHsmsClientDeleteResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armkeyvault.ManagedHsmsClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armkeyvault.ManagedHsmsClientListBySubscriptionResponse]](),
		newListDeletedPager:         newTracker[azfake.PagerResponder[armkeyvault.ManagedHsmsClientListDeletedResponse]](),
		beginPurgeDeleted:           newTracker[azfake.PollerResponder[armkeyvault.ManagedHsmsClientPurgeDeletedResponse]](),
		beginUpdate:                 newTracker[azfake.PollerResponder[armkeyvault.ManagedHsmsClientUpdateResponse]](),
	}
}

// ManagedHsmsServerTransport connects instances of armkeyvault.ManagedHsmsClient to instances of ManagedHsmsServer.
// Don't use this type directly, use NewManagedHsmsServerTransport instead.
type ManagedHsmsServerTransport struct {
	srv                         *ManagedHsmsServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armkeyvault.ManagedHsmsClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armkeyvault.ManagedHsmsClientDeleteResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armkeyvault.ManagedHsmsClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armkeyvault.ManagedHsmsClientListBySubscriptionResponse]]
	newListDeletedPager         *tracker[azfake.PagerResponder[armkeyvault.ManagedHsmsClientListDeletedResponse]]
	beginPurgeDeleted           *tracker[azfake.PollerResponder[armkeyvault.ManagedHsmsClientPurgeDeletedResponse]]
	beginUpdate                 *tracker[azfake.PollerResponder[armkeyvault.ManagedHsmsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for ManagedHsmsServerTransport.
func (m *ManagedHsmsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return m.dispatchToMethodFake(req, method)
}

func (m *ManagedHsmsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if managedHsmsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = managedHsmsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ManagedHsmsClient.CheckMhsmNameAvailability":
				res.resp, res.err = m.dispatchCheckMhsmNameAvailability(req)
			case "ManagedHsmsClient.BeginCreateOrUpdate":
				res.resp, res.err = m.dispatchBeginCreateOrUpdate(req)
			case "ManagedHsmsClient.BeginDelete":
				res.resp, res.err = m.dispatchBeginDelete(req)
			case "ManagedHsmsClient.Get":
				res.resp, res.err = m.dispatchGet(req)
			case "ManagedHsmsClient.GetDeleted":
				res.resp, res.err = m.dispatchGetDeleted(req)
			case "ManagedHsmsClient.NewListByResourceGroupPager":
				res.resp, res.err = m.dispatchNewListByResourceGroupPager(req)
			case "ManagedHsmsClient.NewListBySubscriptionPager":
				res.resp, res.err = m.dispatchNewListBySubscriptionPager(req)
			case "ManagedHsmsClient.NewListDeletedPager":
				res.resp, res.err = m.dispatchNewListDeletedPager(req)
			case "ManagedHsmsClient.BeginPurgeDeleted":
				res.resp, res.err = m.dispatchBeginPurgeDeleted(req)
			case "ManagedHsmsClient.BeginUpdate":
				res.resp, res.err = m.dispatchBeginUpdate(req)
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

func (m *ManagedHsmsServerTransport) dispatchCheckMhsmNameAvailability(req *http.Request) (*http.Response, error) {
	if m.srv.CheckMhsmNameAvailability == nil {
		return nil, &nonRetriableError{errors.New("fake for method CheckMhsmNameAvailability not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.KeyVault/checkMhsmNameAvailability`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armkeyvault.CheckMhsmNameAvailabilityParameters](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.CheckMhsmNameAvailability(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CheckMhsmNameAvailabilityResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagedHsmsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if m.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := m.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.KeyVault/managedHSMs/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armkeyvault.ManagedHsm](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, nameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		m.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		m.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		m.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (m *ManagedHsmsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if m.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := m.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.KeyVault/managedHSMs/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginDelete(req.Context(), resourceGroupNameParam, nameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		m.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		m.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		m.beginDelete.remove(req)
	}

	return resp, nil
}

func (m *ManagedHsmsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if m.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.KeyVault/managedHSMs/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Get(req.Context(), resourceGroupNameParam, nameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagedHsm, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagedHsmsServerTransport) dispatchGetDeleted(req *http.Request) (*http.Response, error) {
	if m.srv.GetDeleted == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetDeleted not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.KeyVault/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deletedManagedHSMs/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.GetDeleted(req.Context(), nameParam, locationParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DeletedManagedHsm, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagedHsmsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := m.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.KeyVault/managedHSMs`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
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
		var options *armkeyvault.ManagedHsmsClientListByResourceGroupOptions
		if topParam != nil {
			options = &armkeyvault.ManagedHsmsClientListByResourceGroupOptions{
				Top: topParam,
			}
		}
		resp := m.srv.NewListByResourceGroupPager(resourceGroupNameParam, options)
		newListByResourceGroupPager = &resp
		m.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armkeyvault.ManagedHsmsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		m.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (m *ManagedHsmsServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := m.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.KeyVault/managedHSMs`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
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
		var options *armkeyvault.ManagedHsmsClientListBySubscriptionOptions
		if topParam != nil {
			options = &armkeyvault.ManagedHsmsClientListBySubscriptionOptions{
				Top: topParam,
			}
		}
		resp := m.srv.NewListBySubscriptionPager(options)
		newListBySubscriptionPager = &resp
		m.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armkeyvault.ManagedHsmsClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		m.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (m *ManagedHsmsServerTransport) dispatchNewListDeletedPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListDeletedPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListDeletedPager not implemented")}
	}
	newListDeletedPager := m.newListDeletedPager.get(req)
	if newListDeletedPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.KeyVault/deletedManagedHSMs`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := m.srv.NewListDeletedPager(nil)
		newListDeletedPager = &resp
		m.newListDeletedPager.add(req, newListDeletedPager)
		server.PagerResponderInjectNextLinks(newListDeletedPager, req, func(page *armkeyvault.ManagedHsmsClientListDeletedResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListDeletedPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListDeletedPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListDeletedPager) {
		m.newListDeletedPager.remove(req)
	}
	return resp, nil
}

func (m *ManagedHsmsServerTransport) dispatchBeginPurgeDeleted(req *http.Request) (*http.Response, error) {
	if m.srv.BeginPurgeDeleted == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPurgeDeleted not implemented")}
	}
	beginPurgeDeleted := m.beginPurgeDeleted.get(req)
	if beginPurgeDeleted == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.KeyVault/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deletedManagedHSMs/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/purge`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
		if err != nil {
			return nil, err
		}
		locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginPurgeDeleted(req.Context(), nameParam, locationParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPurgeDeleted = &respr
		m.beginPurgeDeleted.add(req, beginPurgeDeleted)
	}

	resp, err := server.PollerResponderNext(beginPurgeDeleted, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		m.beginPurgeDeleted.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPurgeDeleted) {
		m.beginPurgeDeleted.remove(req)
	}

	return resp, nil
}

func (m *ManagedHsmsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if m.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := m.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.KeyVault/managedHSMs/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armkeyvault.ManagedHsm](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginUpdate(req.Context(), resourceGroupNameParam, nameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		m.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		m.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		m.beginUpdate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to ManagedHsmsServerTransport
var managedHsmsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
