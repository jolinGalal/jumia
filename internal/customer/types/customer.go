package types

import "github.com/jolinGalal/jumia/internal/models/customer"

var CustomerModel = customer.New()

// CustomerSort ...
var CustomerSort = struct {
	CustomerID    string
	CustomerName  string
	CustomerPhone string
}{
	CustomerID:    "CustomerID",
	CustomerName:  "CustomerName",
	CustomerPhone: "CustomerPhone",
}

// CustomerSortFields map sorting field from the api to the database
var CustomerSortFields = map[string]string{
	CustomerSort.CustomerID:    CustomerModel.CustomerID(),
	CustomerSort.CustomerName:  CustomerModel.CustomerName(),
	CustomerSort.CustomerPhone: CustomerModel.CustomerPhone(),
}

// SortDirection ...
var SortDirection = struct {
	Asc  string
	Desc string
}{
	Asc:  "asc",
	Desc: "desc",
}
