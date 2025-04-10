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
	"net/http"
	"net/url"
	"regexp"

	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// DaprComponentResiliencyPoliciesServer is a fake server for instances of the armappcontainers.DaprComponentResiliencyPoliciesClient type.
type DaprComponentResiliencyPoliciesServer struct {
	// CreateOrUpdate is the fake for method DaprComponentResiliencyPoliciesClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, environmentName string, componentName string, name string, daprComponentResiliencyPolicyEnvelope armappcontainers.DaprComponentResiliencyPolicy, options *armappcontainers.DaprComponentResiliencyPoliciesClientCreateOrUpdateOptions) (resp azfake.Responder[armappcontainers.DaprComponentResiliencyPoliciesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method DaprComponentResiliencyPoliciesClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, environmentName string, componentName string, name string, options *armappcontainers.DaprComponentResiliencyPoliciesClientDeleteOptions) (resp azfake.Responder[armappcontainers.DaprComponentResiliencyPoliciesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method DaprComponentResiliencyPoliciesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, environmentName string, componentName string, name string, options *armappcontainers.DaprComponentResiliencyPoliciesClientGetOptions) (resp azfake.Responder[armappcontainers.DaprComponentResiliencyPoliciesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method DaprComponentResiliencyPoliciesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, environmentName string, componentName string, options *armappcontainers.DaprComponentResiliencyPoliciesClientListOptions) (resp azfake.PagerResponder[armappcontainers.DaprComponentResiliencyPoliciesClientListResponse])
}

// NewDaprComponentResiliencyPoliciesServerTransport creates a new instance of DaprComponentResiliencyPoliciesServerTransport with the provided implementation.
// The returned DaprComponentResiliencyPoliciesServerTransport instance is connected to an instance of armappcontainers.DaprComponentResiliencyPoliciesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDaprComponentResiliencyPoliciesServerTransport(srv *DaprComponentResiliencyPoliciesServer) *DaprComponentResiliencyPoliciesServerTransport {
	return &DaprComponentResiliencyPoliciesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armappcontainers.DaprComponentResiliencyPoliciesClientListResponse]](),
	}
}

// DaprComponentResiliencyPoliciesServerTransport connects instances of armappcontainers.DaprComponentResiliencyPoliciesClient to instances of DaprComponentResiliencyPoliciesServer.
// Don't use this type directly, use NewDaprComponentResiliencyPoliciesServerTransport instead.
type DaprComponentResiliencyPoliciesServerTransport struct {
	srv          *DaprComponentResiliencyPoliciesServer
	newListPager *tracker[azfake.PagerResponder[armappcontainers.DaprComponentResiliencyPoliciesClientListResponse]]
}

// Do implements the policy.Transporter interface for DaprComponentResiliencyPoliciesServerTransport.
func (d *DaprComponentResiliencyPoliciesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DaprComponentResiliencyPoliciesClient.CreateOrUpdate":
		resp, err = d.dispatchCreateOrUpdate(req)
	case "DaprComponentResiliencyPoliciesClient.Delete":
		resp, err = d.dispatchDelete(req)
	case "DaprComponentResiliencyPoliciesClient.Get":
		resp, err = d.dispatchGet(req)
	case "DaprComponentResiliencyPoliciesClient.NewListPager":
		resp, err = d.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DaprComponentResiliencyPoliciesServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/managedEnvironments/(?P<environmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/daprComponents/(?P<componentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resiliencyPolicies/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armappcontainers.DaprComponentResiliencyPolicy](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	environmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("environmentName")])
	if err != nil {
		return nil, err
	}
	componentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("componentName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, environmentNameParam, componentNameParam, nameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DaprComponentResiliencyPolicy, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DaprComponentResiliencyPoliciesServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if d.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/managedEnvironments/(?P<environmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/daprComponents/(?P<componentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resiliencyPolicies/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	environmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("environmentName")])
	if err != nil {
		return nil, err
	}
	componentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("componentName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Delete(req.Context(), resourceGroupNameParam, environmentNameParam, componentNameParam, nameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DaprComponentResiliencyPoliciesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/managedEnvironments/(?P<environmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/daprComponents/(?P<componentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resiliencyPolicies/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	environmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("environmentName")])
	if err != nil {
		return nil, err
	}
	componentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("componentName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Get(req.Context(), resourceGroupNameParam, environmentNameParam, componentNameParam, nameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DaprComponentResiliencyPolicy, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DaprComponentResiliencyPoliciesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := d.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/managedEnvironments/(?P<environmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/daprComponents/(?P<componentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resiliencyPolicies`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		environmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("environmentName")])
		if err != nil {
			return nil, err
		}
		componentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("componentName")])
		if err != nil {
			return nil, err
		}
		resp := d.srv.NewListPager(resourceGroupNameParam, environmentNameParam, componentNameParam, nil)
		newListPager = &resp
		d.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armappcontainers.DaprComponentResiliencyPoliciesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		d.newListPager.remove(req)
	}
	return resp, nil
}
