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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare"
	"net/http"
	"net/url"
	"regexp"
)

// ShareSubscriptionsServer is a fake server for instances of the armdatashare.ShareSubscriptionsClient type.
type ShareSubscriptionsServer struct {
	// BeginCancelSynchronization is the fake for method ShareSubscriptionsClient.BeginCancelSynchronization
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCancelSynchronization func(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, shareSubscriptionSynchronization armdatashare.ShareSubscriptionSynchronization, options *armdatashare.ShareSubscriptionsClientBeginCancelSynchronizationOptions) (resp azfake.PollerResponder[armdatashare.ShareSubscriptionsClientCancelSynchronizationResponse], errResp azfake.ErrorResponder)

	// Create is the fake for method ShareSubscriptionsClient.Create
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	Create func(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, shareSubscription armdatashare.ShareSubscription, options *armdatashare.ShareSubscriptionsClientCreateOptions) (resp azfake.Responder[armdatashare.ShareSubscriptionsClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ShareSubscriptionsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, options *armdatashare.ShareSubscriptionsClientBeginDeleteOptions) (resp azfake.PollerResponder[armdatashare.ShareSubscriptionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ShareSubscriptionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, options *armdatashare.ShareSubscriptionsClientGetOptions) (resp azfake.Responder[armdatashare.ShareSubscriptionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByAccountPager is the fake for method ShareSubscriptionsClient.NewListByAccountPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByAccountPager func(resourceGroupName string, accountName string, options *armdatashare.ShareSubscriptionsClientListByAccountOptions) (resp azfake.PagerResponder[armdatashare.ShareSubscriptionsClientListByAccountResponse])

	// NewListSourceShareSynchronizationSettingsPager is the fake for method ShareSubscriptionsClient.NewListSourceShareSynchronizationSettingsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListSourceShareSynchronizationSettingsPager func(resourceGroupName string, accountName string, shareSubscriptionName string, options *armdatashare.ShareSubscriptionsClientListSourceShareSynchronizationSettingsOptions) (resp azfake.PagerResponder[armdatashare.ShareSubscriptionsClientListSourceShareSynchronizationSettingsResponse])

	// NewListSynchronizationDetailsPager is the fake for method ShareSubscriptionsClient.NewListSynchronizationDetailsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListSynchronizationDetailsPager func(resourceGroupName string, accountName string, shareSubscriptionName string, shareSubscriptionSynchronization armdatashare.ShareSubscriptionSynchronization, options *armdatashare.ShareSubscriptionsClientListSynchronizationDetailsOptions) (resp azfake.PagerResponder[armdatashare.ShareSubscriptionsClientListSynchronizationDetailsResponse])

	// NewListSynchronizationsPager is the fake for method ShareSubscriptionsClient.NewListSynchronizationsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListSynchronizationsPager func(resourceGroupName string, accountName string, shareSubscriptionName string, options *armdatashare.ShareSubscriptionsClientListSynchronizationsOptions) (resp azfake.PagerResponder[armdatashare.ShareSubscriptionsClientListSynchronizationsResponse])

	// BeginSynchronize is the fake for method ShareSubscriptionsClient.BeginSynchronize
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginSynchronize func(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, synchronize armdatashare.Synchronize, options *armdatashare.ShareSubscriptionsClientBeginSynchronizeOptions) (resp azfake.PollerResponder[armdatashare.ShareSubscriptionsClientSynchronizeResponse], errResp azfake.ErrorResponder)
}

// NewShareSubscriptionsServerTransport creates a new instance of ShareSubscriptionsServerTransport with the provided implementation.
// The returned ShareSubscriptionsServerTransport instance is connected to an instance of armdatashare.ShareSubscriptionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewShareSubscriptionsServerTransport(srv *ShareSubscriptionsServer) *ShareSubscriptionsServerTransport {
	return &ShareSubscriptionsServerTransport{
		srv:                        srv,
		beginCancelSynchronization: newTracker[azfake.PollerResponder[armdatashare.ShareSubscriptionsClientCancelSynchronizationResponse]](),
		beginDelete:                newTracker[azfake.PollerResponder[armdatashare.ShareSubscriptionsClientDeleteResponse]](),
		newListByAccountPager:      newTracker[azfake.PagerResponder[armdatashare.ShareSubscriptionsClientListByAccountResponse]](),
		newListSourceShareSynchronizationSettingsPager: newTracker[azfake.PagerResponder[armdatashare.ShareSubscriptionsClientListSourceShareSynchronizationSettingsResponse]](),
		newListSynchronizationDetailsPager:             newTracker[azfake.PagerResponder[armdatashare.ShareSubscriptionsClientListSynchronizationDetailsResponse]](),
		newListSynchronizationsPager:                   newTracker[azfake.PagerResponder[armdatashare.ShareSubscriptionsClientListSynchronizationsResponse]](),
		beginSynchronize:                               newTracker[azfake.PollerResponder[armdatashare.ShareSubscriptionsClientSynchronizeResponse]](),
	}
}

// ShareSubscriptionsServerTransport connects instances of armdatashare.ShareSubscriptionsClient to instances of ShareSubscriptionsServer.
// Don't use this type directly, use NewShareSubscriptionsServerTransport instead.
type ShareSubscriptionsServerTransport struct {
	srv                                            *ShareSubscriptionsServer
	beginCancelSynchronization                     *tracker[azfake.PollerResponder[armdatashare.ShareSubscriptionsClientCancelSynchronizationResponse]]
	beginDelete                                    *tracker[azfake.PollerResponder[armdatashare.ShareSubscriptionsClientDeleteResponse]]
	newListByAccountPager                          *tracker[azfake.PagerResponder[armdatashare.ShareSubscriptionsClientListByAccountResponse]]
	newListSourceShareSynchronizationSettingsPager *tracker[azfake.PagerResponder[armdatashare.ShareSubscriptionsClientListSourceShareSynchronizationSettingsResponse]]
	newListSynchronizationDetailsPager             *tracker[azfake.PagerResponder[armdatashare.ShareSubscriptionsClientListSynchronizationDetailsResponse]]
	newListSynchronizationsPager                   *tracker[azfake.PagerResponder[armdatashare.ShareSubscriptionsClientListSynchronizationsResponse]]
	beginSynchronize                               *tracker[azfake.PollerResponder[armdatashare.ShareSubscriptionsClientSynchronizeResponse]]
}

// Do implements the policy.Transporter interface for ShareSubscriptionsServerTransport.
func (s *ShareSubscriptionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ShareSubscriptionsClient.BeginCancelSynchronization":
		resp, err = s.dispatchBeginCancelSynchronization(req)
	case "ShareSubscriptionsClient.Create":
		resp, err = s.dispatchCreate(req)
	case "ShareSubscriptionsClient.BeginDelete":
		resp, err = s.dispatchBeginDelete(req)
	case "ShareSubscriptionsClient.Get":
		resp, err = s.dispatchGet(req)
	case "ShareSubscriptionsClient.NewListByAccountPager":
		resp, err = s.dispatchNewListByAccountPager(req)
	case "ShareSubscriptionsClient.NewListSourceShareSynchronizationSettingsPager":
		resp, err = s.dispatchNewListSourceShareSynchronizationSettingsPager(req)
	case "ShareSubscriptionsClient.NewListSynchronizationDetailsPager":
		resp, err = s.dispatchNewListSynchronizationDetailsPager(req)
	case "ShareSubscriptionsClient.NewListSynchronizationsPager":
		resp, err = s.dispatchNewListSynchronizationsPager(req)
	case "ShareSubscriptionsClient.BeginSynchronize":
		resp, err = s.dispatchBeginSynchronize(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ShareSubscriptionsServerTransport) dispatchBeginCancelSynchronization(req *http.Request) (*http.Response, error) {
	if s.srv.BeginCancelSynchronization == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCancelSynchronization not implemented")}
	}
	beginCancelSynchronization := s.beginCancelSynchronization.get(req)
	if beginCancelSynchronization == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataShare/accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/shareSubscriptions/(?P<shareSubscriptionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cancelSynchronization`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdatashare.ShareSubscriptionSynchronization](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		shareSubscriptionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("shareSubscriptionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginCancelSynchronization(req.Context(), resourceGroupNameParam, accountNameParam, shareSubscriptionNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCancelSynchronization = &respr
		s.beginCancelSynchronization.add(req, beginCancelSynchronization)
	}

	resp, err := server.PollerResponderNext(beginCancelSynchronization, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginCancelSynchronization.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCancelSynchronization) {
		s.beginCancelSynchronization.remove(req)
	}

	return resp, nil
}

func (s *ShareSubscriptionsServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if s.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataShare/accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/shareSubscriptions/(?P<shareSubscriptionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdatashare.ShareSubscription](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	shareSubscriptionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("shareSubscriptionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Create(req.Context(), resourceGroupNameParam, accountNameParam, shareSubscriptionNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ShareSubscription, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ShareSubscriptionsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if s.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := s.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataShare/accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/shareSubscriptions/(?P<shareSubscriptionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		shareSubscriptionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("shareSubscriptionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginDelete(req.Context(), resourceGroupNameParam, accountNameParam, shareSubscriptionNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		s.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		s.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		s.beginDelete.remove(req)
	}

	return resp, nil
}

func (s *ShareSubscriptionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataShare/accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/shareSubscriptions/(?P<shareSubscriptionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	shareSubscriptionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("shareSubscriptionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, accountNameParam, shareSubscriptionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ShareSubscription, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ShareSubscriptionsServerTransport) dispatchNewListByAccountPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByAccountPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByAccountPager not implemented")}
	}
	newListByAccountPager := s.newListByAccountPager.get(req)
	if newListByAccountPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataShare/accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/shareSubscriptions`
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
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		orderbyUnescaped, err := url.QueryUnescape(qp.Get("$orderby"))
		if err != nil {
			return nil, err
		}
		orderbyParam := getOptional(orderbyUnescaped)
		var options *armdatashare.ShareSubscriptionsClientListByAccountOptions
		if skipTokenParam != nil || filterParam != nil || orderbyParam != nil {
			options = &armdatashare.ShareSubscriptionsClientListByAccountOptions{
				SkipToken: skipTokenParam,
				Filter:    filterParam,
				Orderby:   orderbyParam,
			}
		}
		resp := s.srv.NewListByAccountPager(resourceGroupNameParam, accountNameParam, options)
		newListByAccountPager = &resp
		s.newListByAccountPager.add(req, newListByAccountPager)
		server.PagerResponderInjectNextLinks(newListByAccountPager, req, func(page *armdatashare.ShareSubscriptionsClientListByAccountResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByAccountPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListByAccountPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByAccountPager) {
		s.newListByAccountPager.remove(req)
	}
	return resp, nil
}

func (s *ShareSubscriptionsServerTransport) dispatchNewListSourceShareSynchronizationSettingsPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListSourceShareSynchronizationSettingsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListSourceShareSynchronizationSettingsPager not implemented")}
	}
	newListSourceShareSynchronizationSettingsPager := s.newListSourceShareSynchronizationSettingsPager.get(req)
	if newListSourceShareSynchronizationSettingsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataShare/accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/shareSubscriptions/(?P<shareSubscriptionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listSourceShareSynchronizationSettings`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		shareSubscriptionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("shareSubscriptionName")])
		if err != nil {
			return nil, err
		}
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
		var options *armdatashare.ShareSubscriptionsClientListSourceShareSynchronizationSettingsOptions
		if skipTokenParam != nil {
			options = &armdatashare.ShareSubscriptionsClientListSourceShareSynchronizationSettingsOptions{
				SkipToken: skipTokenParam,
			}
		}
		resp := s.srv.NewListSourceShareSynchronizationSettingsPager(resourceGroupNameParam, accountNameParam, shareSubscriptionNameParam, options)
		newListSourceShareSynchronizationSettingsPager = &resp
		s.newListSourceShareSynchronizationSettingsPager.add(req, newListSourceShareSynchronizationSettingsPager)
		server.PagerResponderInjectNextLinks(newListSourceShareSynchronizationSettingsPager, req, func(page *armdatashare.ShareSubscriptionsClientListSourceShareSynchronizationSettingsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListSourceShareSynchronizationSettingsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListSourceShareSynchronizationSettingsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListSourceShareSynchronizationSettingsPager) {
		s.newListSourceShareSynchronizationSettingsPager.remove(req)
	}
	return resp, nil
}

func (s *ShareSubscriptionsServerTransport) dispatchNewListSynchronizationDetailsPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListSynchronizationDetailsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListSynchronizationDetailsPager not implemented")}
	}
	newListSynchronizationDetailsPager := s.newListSynchronizationDetailsPager.get(req)
	if newListSynchronizationDetailsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataShare/accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/shareSubscriptions/(?P<shareSubscriptionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listSynchronizationDetails`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		body, err := server.UnmarshalRequestAsJSON[armdatashare.ShareSubscriptionSynchronization](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		shareSubscriptionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("shareSubscriptionName")])
		if err != nil {
			return nil, err
		}
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		orderbyUnescaped, err := url.QueryUnescape(qp.Get("$orderby"))
		if err != nil {
			return nil, err
		}
		orderbyParam := getOptional(orderbyUnescaped)
		var options *armdatashare.ShareSubscriptionsClientListSynchronizationDetailsOptions
		if skipTokenParam != nil || filterParam != nil || orderbyParam != nil {
			options = &armdatashare.ShareSubscriptionsClientListSynchronizationDetailsOptions{
				SkipToken: skipTokenParam,
				Filter:    filterParam,
				Orderby:   orderbyParam,
			}
		}
		resp := s.srv.NewListSynchronizationDetailsPager(resourceGroupNameParam, accountNameParam, shareSubscriptionNameParam, body, options)
		newListSynchronizationDetailsPager = &resp
		s.newListSynchronizationDetailsPager.add(req, newListSynchronizationDetailsPager)
		server.PagerResponderInjectNextLinks(newListSynchronizationDetailsPager, req, func(page *armdatashare.ShareSubscriptionsClientListSynchronizationDetailsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListSynchronizationDetailsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListSynchronizationDetailsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListSynchronizationDetailsPager) {
		s.newListSynchronizationDetailsPager.remove(req)
	}
	return resp, nil
}

func (s *ShareSubscriptionsServerTransport) dispatchNewListSynchronizationsPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListSynchronizationsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListSynchronizationsPager not implemented")}
	}
	newListSynchronizationsPager := s.newListSynchronizationsPager.get(req)
	if newListSynchronizationsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataShare/accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/shareSubscriptions/(?P<shareSubscriptionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listSynchronizations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		shareSubscriptionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("shareSubscriptionName")])
		if err != nil {
			return nil, err
		}
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		orderbyUnescaped, err := url.QueryUnescape(qp.Get("$orderby"))
		if err != nil {
			return nil, err
		}
		orderbyParam := getOptional(orderbyUnescaped)
		var options *armdatashare.ShareSubscriptionsClientListSynchronizationsOptions
		if skipTokenParam != nil || filterParam != nil || orderbyParam != nil {
			options = &armdatashare.ShareSubscriptionsClientListSynchronizationsOptions{
				SkipToken: skipTokenParam,
				Filter:    filterParam,
				Orderby:   orderbyParam,
			}
		}
		resp := s.srv.NewListSynchronizationsPager(resourceGroupNameParam, accountNameParam, shareSubscriptionNameParam, options)
		newListSynchronizationsPager = &resp
		s.newListSynchronizationsPager.add(req, newListSynchronizationsPager)
		server.PagerResponderInjectNextLinks(newListSynchronizationsPager, req, func(page *armdatashare.ShareSubscriptionsClientListSynchronizationsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListSynchronizationsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListSynchronizationsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListSynchronizationsPager) {
		s.newListSynchronizationsPager.remove(req)
	}
	return resp, nil
}

func (s *ShareSubscriptionsServerTransport) dispatchBeginSynchronize(req *http.Request) (*http.Response, error) {
	if s.srv.BeginSynchronize == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginSynchronize not implemented")}
	}
	beginSynchronize := s.beginSynchronize.get(req)
	if beginSynchronize == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataShare/accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/shareSubscriptions/(?P<shareSubscriptionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/synchronize`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdatashare.Synchronize](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		shareSubscriptionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("shareSubscriptionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginSynchronize(req.Context(), resourceGroupNameParam, accountNameParam, shareSubscriptionNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginSynchronize = &respr
		s.beginSynchronize.add(req, beginSynchronize)
	}

	resp, err := server.PollerResponderNext(beginSynchronize, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginSynchronize.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginSynchronize) {
		s.beginSynchronize.remove(req)
	}

	return resp, nil
}