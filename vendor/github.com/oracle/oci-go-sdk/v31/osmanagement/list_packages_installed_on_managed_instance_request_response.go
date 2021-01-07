// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/v31/common"
	"net/http"
)

// ListPackagesInstalledOnManagedInstanceRequest wrapper for the ListPackagesInstalledOnManagedInstance operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListPackagesInstalledOnManagedInstance.go.html to see an example of how to use ListPackagesInstalledOnManagedInstanceRequest.
type ListPackagesInstalledOnManagedInstanceRequest struct {

	// OCID for the managed instance
	ManagedInstanceId *string `mandatory:"true" contributesTo:"path" name:"managedInstanceId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The ID of the compartment in which to list resources. This parameter is optional and in some cases may have no effect.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListPackagesInstalledOnManagedInstanceSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListPackagesInstalledOnManagedInstanceSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPackagesInstalledOnManagedInstanceRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPackagesInstalledOnManagedInstanceRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPackagesInstalledOnManagedInstanceRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListPackagesInstalledOnManagedInstanceResponse wrapper for the ListPackagesInstalledOnManagedInstance operation
type ListPackagesInstalledOnManagedInstanceResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []InstalledPackageSummary instances
	Items []InstalledPackageSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of `InstalledPackageSummary`s. If this
	// header appears in the response, then this is a partial list of
	// `InstalledPackageSummary`s for the managed instance. Include
	// this value as the `page` parameter in a subsequent GET request
	// to get the next batch of managed instances.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPackagesInstalledOnManagedInstanceResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPackagesInstalledOnManagedInstanceResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPackagesInstalledOnManagedInstanceSortOrderEnum Enum with underlying type: string
type ListPackagesInstalledOnManagedInstanceSortOrderEnum string

// Set of constants representing the allowable values for ListPackagesInstalledOnManagedInstanceSortOrderEnum
const (
	ListPackagesInstalledOnManagedInstanceSortOrderAsc  ListPackagesInstalledOnManagedInstanceSortOrderEnum = "ASC"
	ListPackagesInstalledOnManagedInstanceSortOrderDesc ListPackagesInstalledOnManagedInstanceSortOrderEnum = "DESC"
)

var mappingListPackagesInstalledOnManagedInstanceSortOrder = map[string]ListPackagesInstalledOnManagedInstanceSortOrderEnum{
	"ASC":  ListPackagesInstalledOnManagedInstanceSortOrderAsc,
	"DESC": ListPackagesInstalledOnManagedInstanceSortOrderDesc,
}

// GetListPackagesInstalledOnManagedInstanceSortOrderEnumValues Enumerates the set of values for ListPackagesInstalledOnManagedInstanceSortOrderEnum
func GetListPackagesInstalledOnManagedInstanceSortOrderEnumValues() []ListPackagesInstalledOnManagedInstanceSortOrderEnum {
	values := make([]ListPackagesInstalledOnManagedInstanceSortOrderEnum, 0)
	for _, v := range mappingListPackagesInstalledOnManagedInstanceSortOrder {
		values = append(values, v)
	}
	return values
}

// ListPackagesInstalledOnManagedInstanceSortByEnum Enum with underlying type: string
type ListPackagesInstalledOnManagedInstanceSortByEnum string

// Set of constants representing the allowable values for ListPackagesInstalledOnManagedInstanceSortByEnum
const (
	ListPackagesInstalledOnManagedInstanceSortByTimecreated ListPackagesInstalledOnManagedInstanceSortByEnum = "TIMECREATED"
	ListPackagesInstalledOnManagedInstanceSortByDisplayname ListPackagesInstalledOnManagedInstanceSortByEnum = "DISPLAYNAME"
)

var mappingListPackagesInstalledOnManagedInstanceSortBy = map[string]ListPackagesInstalledOnManagedInstanceSortByEnum{
	"TIMECREATED": ListPackagesInstalledOnManagedInstanceSortByTimecreated,
	"DISPLAYNAME": ListPackagesInstalledOnManagedInstanceSortByDisplayname,
}

// GetListPackagesInstalledOnManagedInstanceSortByEnumValues Enumerates the set of values for ListPackagesInstalledOnManagedInstanceSortByEnum
func GetListPackagesInstalledOnManagedInstanceSortByEnumValues() []ListPackagesInstalledOnManagedInstanceSortByEnum {
	values := make([]ListPackagesInstalledOnManagedInstanceSortByEnum, 0)
	for _, v := range mappingListPackagesInstalledOnManagedInstanceSortBy {
		values = append(values, v)
	}
	return values
}
