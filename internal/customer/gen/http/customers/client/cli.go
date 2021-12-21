// Code generated by goa v3.5.3, DO NOT EDIT.
//
// customers HTTP client CLI support package
//
// Command:
// $ goa gen github.com/jolinGalal/jumia/internal/customer/design

package client

import (
	"fmt"
	"strconv"

	customers "github.com/jolinGalal/jumia/internal/customer/gen/customers"
	goa "goa.design/goa/v3/pkg"
)

// BuildListPayload builds the payload for the customers list endpoint from CLI
// flags.
func BuildListPayload(customersListCountry string, customersListState string, customersListSortDirection string, customersListSortKey string, customersListPageNumber string, customersListPageSize string) (*customers.ListPayload, error) {
	var err error
	var country string
	{
		if customersListCountry != "" {
			country = customersListCountry
			if !(country == "all" || country == "cameroon" || country == "ethiopia" || country == "morocco" || country == "mozambique" || country == "uganda") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("country", country, []interface{}{"all", "cameroon", "ethiopia", "morocco", "mozambique", "uganda"}))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var state string
	{
		if customersListState != "" {
			state = customersListState
			if !(state == "all" || state == "Valid" || state == "NotValid") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("state", state, []interface{}{"all", "Valid", "NotValid"}))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var sortDirection string
	{
		if customersListSortDirection != "" {
			sortDirection = customersListSortDirection
			if !(sortDirection == "asc" || sortDirection == "desc") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("sortDirection", sortDirection, []interface{}{"asc", "desc"}))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var sortKey string
	{
		if customersListSortKey != "" {
			sortKey = customersListSortKey
			if !(sortKey == "CustomerID" || sortKey == "CustomerName" || sortKey == "CustomerPhone") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("sortKey", sortKey, []interface{}{"CustomerID", "CustomerName", "CustomerPhone"}))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var pageNumber int
	{
		if customersListPageNumber != "" {
			var v int64
			v, err = strconv.ParseInt(customersListPageNumber, 10, 64)
			pageNumber = int(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for pageNumber, must be INT")
			}
		}
	}
	var pageSize int
	{
		if customersListPageSize != "" {
			var v int64
			v, err = strconv.ParseInt(customersListPageSize, 10, 64)
			pageSize = int(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for pageSize, must be INT")
			}
		}
	}
	v := &customers.ListPayload{}
	v.Country = country
	v.State = state
	v.SortDirection = sortDirection
	v.SortKey = sortKey
	v.PageNumber = pageNumber
	v.PageSize = pageSize

	return v, nil
}
