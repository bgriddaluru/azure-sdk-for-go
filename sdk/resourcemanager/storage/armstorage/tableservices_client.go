// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage

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

// TableServicesClient contains the methods for the TableServices group.
// Don't use this type directly, use NewTableServicesClient() instead.
type TableServicesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTableServicesClient creates a new instance of TableServicesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTableServicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TableServicesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TableServicesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// GetServiceProperties - Gets the properties of a storage account’s Table service, including properties for Storage Analytics
// and CORS (Cross-Origin Resource Sharing) rules.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - accountName - The name of the storage account within the specified resource group. Storage account names must be between
//     3 and 24 characters in length and use numbers and lower-case letters only.
//   - options - TableServicesClientGetServicePropertiesOptions contains the optional parameters for the TableServicesClient.GetServiceProperties
//     method.
func (client *TableServicesClient) GetServiceProperties(ctx context.Context, resourceGroupName string, accountName string, options *TableServicesClientGetServicePropertiesOptions) (TableServicesClientGetServicePropertiesResponse, error) {
	var err error
	const operationName = "TableServicesClient.GetServiceProperties"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getServicePropertiesCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return TableServicesClientGetServicePropertiesResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TableServicesClientGetServicePropertiesResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TableServicesClientGetServicePropertiesResponse{}, err
	}
	resp, err := client.getServicePropertiesHandleResponse(httpResp)
	return resp, err
}

// getServicePropertiesCreateRequest creates the GetServiceProperties request.
func (client *TableServicesClient) getServicePropertiesCreateRequest(ctx context.Context, resourceGroupName string, accountName string, _ *TableServicesClientGetServicePropertiesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/tableServices/{tableServiceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{tableServiceName}", url.PathEscape("default"))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getServicePropertiesHandleResponse handles the GetServiceProperties response.
func (client *TableServicesClient) getServicePropertiesHandleResponse(resp *http.Response) (TableServicesClientGetServicePropertiesResponse, error) {
	result := TableServicesClientGetServicePropertiesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TableServiceProperties); err != nil {
		return TableServicesClientGetServicePropertiesResponse{}, err
	}
	return result, nil
}

// List - List all table services for the storage account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - accountName - The name of the storage account within the specified resource group. Storage account names must be between
//     3 and 24 characters in length and use numbers and lower-case letters only.
//   - options - TableServicesClientListOptions contains the optional parameters for the TableServicesClient.List method.
func (client *TableServicesClient) List(ctx context.Context, resourceGroupName string, accountName string, options *TableServicesClientListOptions) (TableServicesClientListResponse, error) {
	var err error
	const operationName = "TableServicesClient.List"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return TableServicesClientListResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TableServicesClientListResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TableServicesClientListResponse{}, err
	}
	resp, err := client.listHandleResponse(httpResp)
	return resp, err
}

// listCreateRequest creates the List request.
func (client *TableServicesClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, _ *TableServicesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/tableServices"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TableServicesClient) listHandleResponse(resp *http.Response) (TableServicesClientListResponse, error) {
	result := TableServicesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListTableServices); err != nil {
		return TableServicesClientListResponse{}, err
	}
	return result, nil
}

// SetServiceProperties - Sets the properties of a storage account’s Table service, including properties for Storage Analytics
// and CORS (Cross-Origin Resource Sharing) rules.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - accountName - The name of the storage account within the specified resource group. Storage account names must be between
//     3 and 24 characters in length and use numbers and lower-case letters only.
//   - parameters - The properties of a storage account’s Table service, only properties for Storage Analytics and CORS (Cross-Origin
//     Resource Sharing) rules can be specified.
//   - options - TableServicesClientSetServicePropertiesOptions contains the optional parameters for the TableServicesClient.SetServiceProperties
//     method.
func (client *TableServicesClient) SetServiceProperties(ctx context.Context, resourceGroupName string, accountName string, parameters TableServiceProperties, options *TableServicesClientSetServicePropertiesOptions) (TableServicesClientSetServicePropertiesResponse, error) {
	var err error
	const operationName = "TableServicesClient.SetServiceProperties"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.setServicePropertiesCreateRequest(ctx, resourceGroupName, accountName, parameters, options)
	if err != nil {
		return TableServicesClientSetServicePropertiesResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TableServicesClientSetServicePropertiesResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TableServicesClientSetServicePropertiesResponse{}, err
	}
	resp, err := client.setServicePropertiesHandleResponse(httpResp)
	return resp, err
}

// setServicePropertiesCreateRequest creates the SetServiceProperties request.
func (client *TableServicesClient) setServicePropertiesCreateRequest(ctx context.Context, resourceGroupName string, accountName string, parameters TableServiceProperties, _ *TableServicesClientSetServicePropertiesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/tableServices/{tableServiceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{tableServiceName}", url.PathEscape("default"))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// setServicePropertiesHandleResponse handles the SetServiceProperties response.
func (client *TableServicesClient) setServicePropertiesHandleResponse(resp *http.Response) (TableServicesClientSetServicePropertiesResponse, error) {
	result := TableServicesClientSetServicePropertiesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TableServiceProperties); err != nil {
		return TableServicesClientSetServicePropertiesResponse{}, err
	}
	return result, nil
}
