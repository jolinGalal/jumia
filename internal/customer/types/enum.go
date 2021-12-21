package types

import "github.com/jolinGalal/jumia/internal/models"

var customer = models.NewCustomer()

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
	CustomerSort.CustomerID:    customer.CustomerID(),
	CustomerSort.CustomerName:  customer.CustomerName(),
	CustomerSort.CustomerPhone: customer.CustomerPhone(),
}

// SortDirection ...
var SortDirection = struct {
	Asc  string
	Desc string
}{
	Asc:  "asc",
	Desc: "desc",
}
