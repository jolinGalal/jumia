// Code generated by goa v3.5.3, DO NOT EDIT.
//
// customers HTTP server types
//
// Command:
// $ goa gen github.com/jolinGalal/jumia/internal/customer/design

package server

import (
	customers "github.com/jolinGalal/jumia/internal/customer/gen/customers"
)

// ListResponseBody is the type of the "customers" service "list" endpoint HTTP
// response body.
type ListResponseBody struct {
	// events
	Customers []*ListCustomerRespResponseBody `form:"customers" json:"customers" xml:"customers"`
	// pagination
	Pagination *PaginationResponseBody `form:"pagination" json:"pagination" xml:"pagination"`
}

// ListCustomerRespResponseBody is used to define fields on response body types.
type ListCustomerRespResponseBody struct {
	// customer ID
	ID int `form:"ID" json:"ID" xml:"ID"`
	// customer name
	Name int `form:"Name" json:"Name" xml:"Name"`
	// customer phone
	Phone int `form:"Phone" json:"Phone" xml:"Phone"`
	// customer country
	Country int `form:"Country" json:"Country" xml:"Country"`
	// phone state
	State int `form:"State" json:"State" xml:"State"`
	// country code
	CountryCode int `form:"CountryCode" json:"CountryCode" xml:"CountryCode"`
}

// PaginationResponseBody is used to define fields on response body types.
type PaginationResponseBody struct {
	// The current page
	CurrentPage *int `form:"current_page,omitempty" json:"current_page,omitempty" xml:"current_page,omitempty"`
	// Max number of records per page
	PageSize *int `form:"page_size,omitempty" json:"page_size,omitempty" xml:"page_size,omitempty"`
	// Total pages
	TotalPages *int `form:"total_pages,omitempty" json:"total_pages,omitempty" xml:"total_pages,omitempty"`
	// Total records count
	TotalCount *int64 `form:"total_count,omitempty" json:"total_count,omitempty" xml:"total_count,omitempty"`
}

// NewListResponseBody builds the HTTP response body from the result of the
// "list" endpoint of the "customers" service.
func NewListResponseBody(res *customers.ListResult) *ListResponseBody {
	body := &ListResponseBody{}
	if res.Customers != nil {
		body.Customers = make([]*ListCustomerRespResponseBody, len(res.Customers))
		for i, val := range res.Customers {
			body.Customers[i] = marshalCustomersListCustomerRespToListCustomerRespResponseBody(val)
		}
	}
	if res.Pagination != nil {
		body.Pagination = marshalCustomersPaginationToPaginationResponseBody(res.Pagination)
	}
	return body
}

// NewListPayload builds a customers service list endpoint payload.
func NewListPayload(country string, state string, sortDirection string, sortKey string, pageNumber int, pageSize int) *customers.ListPayload {
	v := &customers.ListPayload{}
	v.Country = country
	v.State = state
	v.SortDirection = sortDirection
	v.SortKey = sortKey
	v.PageNumber = pageNumber
	v.PageSize = pageSize

	return v
}
