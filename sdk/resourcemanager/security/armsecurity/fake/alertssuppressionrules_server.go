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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"net/http"
	"net/url"
	"regexp"
)

// AlertsSuppressionRulesServer is a fake server for instances of the armsecurity.AlertsSuppressionRulesClient type.
type AlertsSuppressionRulesServer struct {
	// Delete is the fake for method AlertsSuppressionRulesClient.Delete
	// HTTP status codes to indicate success: http.StatusNoContent
	Delete func(ctx context.Context, alertsSuppressionRuleName string, options *armsecurity.AlertsSuppressionRulesClientDeleteOptions) (resp azfake.Responder[armsecurity.AlertsSuppressionRulesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method AlertsSuppressionRulesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, alertsSuppressionRuleName string, options *armsecurity.AlertsSuppressionRulesClientGetOptions) (resp azfake.Responder[armsecurity.AlertsSuppressionRulesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method AlertsSuppressionRulesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armsecurity.AlertsSuppressionRulesClientListOptions) (resp azfake.PagerResponder[armsecurity.AlertsSuppressionRulesClientListResponse])

	// Update is the fake for method AlertsSuppressionRulesClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, alertsSuppressionRuleName string, alertsSuppressionRule armsecurity.AlertsSuppressionRule, options *armsecurity.AlertsSuppressionRulesClientUpdateOptions) (resp azfake.Responder[armsecurity.AlertsSuppressionRulesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewAlertsSuppressionRulesServerTransport creates a new instance of AlertsSuppressionRulesServerTransport with the provided implementation.
// The returned AlertsSuppressionRulesServerTransport instance is connected to an instance of armsecurity.AlertsSuppressionRulesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAlertsSuppressionRulesServerTransport(srv *AlertsSuppressionRulesServer) *AlertsSuppressionRulesServerTransport {
	return &AlertsSuppressionRulesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armsecurity.AlertsSuppressionRulesClientListResponse]](),
	}
}

// AlertsSuppressionRulesServerTransport connects instances of armsecurity.AlertsSuppressionRulesClient to instances of AlertsSuppressionRulesServer.
// Don't use this type directly, use NewAlertsSuppressionRulesServerTransport instead.
type AlertsSuppressionRulesServerTransport struct {
	srv          *AlertsSuppressionRulesServer
	newListPager *tracker[azfake.PagerResponder[armsecurity.AlertsSuppressionRulesClientListResponse]]
}

// Do implements the policy.Transporter interface for AlertsSuppressionRulesServerTransport.
func (a *AlertsSuppressionRulesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AlertsSuppressionRulesClient.Delete":
		resp, err = a.dispatchDelete(req)
	case "AlertsSuppressionRulesClient.Get":
		resp, err = a.dispatchGet(req)
	case "AlertsSuppressionRulesClient.NewListPager":
		resp, err = a.dispatchNewListPager(req)
	case "AlertsSuppressionRulesClient.Update":
		resp, err = a.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AlertsSuppressionRulesServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if a.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Security/alertsSuppressionRules/(?P<alertsSuppressionRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	alertsSuppressionRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("alertsSuppressionRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Delete(req.Context(), alertsSuppressionRuleNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AlertsSuppressionRulesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Security/alertsSuppressionRules/(?P<alertsSuppressionRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	alertsSuppressionRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("alertsSuppressionRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), alertsSuppressionRuleNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AlertsSuppressionRule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AlertsSuppressionRulesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := a.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Security/alertsSuppressionRules`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		alertTypeUnescaped, err := url.QueryUnescape(qp.Get("AlertType"))
		if err != nil {
			return nil, err
		}
		alertTypeParam := getOptional(alertTypeUnescaped)
		var options *armsecurity.AlertsSuppressionRulesClientListOptions
		if alertTypeParam != nil {
			options = &armsecurity.AlertsSuppressionRulesClientListOptions{
				AlertType: alertTypeParam,
			}
		}
		resp := a.srv.NewListPager(options)
		newListPager = &resp
		a.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armsecurity.AlertsSuppressionRulesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		a.newListPager.remove(req)
	}
	return resp, nil
}

func (a *AlertsSuppressionRulesServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Security/alertsSuppressionRules/(?P<alertsSuppressionRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armsecurity.AlertsSuppressionRule](req)
	if err != nil {
		return nil, err
	}
	alertsSuppressionRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("alertsSuppressionRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Update(req.Context(), alertsSuppressionRuleNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AlertsSuppressionRule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}