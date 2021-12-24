package design

import (
	"github.com/jolinGalal/jumia/internal/customer/types"
	"goa.design/goa/v3/dsl"
)

var _ = dsl.Service("customers", func() {
	dsl.Description("The customer service exposes endpoints.")

	dsl.HTTP(func() {
		dsl.Path("/")
	})

	dsl.Method("list", func() {
		dsl.Description("list customers")

		// The list customers endpoint is secured via jwt auth

		dsl.Payload(func() {
			dsl.Field(1, "country", dsl.String, func() {
				dsl.Description("country")
				dsl.Enum("all", types.Countrylst[0], types.Countrylst[1], types.Countrylst[2],
					types.Countrylst[3], types.Countrylst[4])
				dsl.Default("all")
			})
			dsl.Field(2, "state", dsl.String, func() {
				dsl.Description("state")
				dsl.Enum(types.State.All, types.State.Valid, types.State.NotValid)
				dsl.Default(types.State.All)

			})
			dsl.Field(3, "sort_key", dsl.String, func() {
				dsl.Description("")
				dsl.Enum(types.CustomerSort.CustomerID, types.CustomerSort.CustomerName,
					types.CustomerSort.CustomerPhone)
				dsl.Default(types.CustomerSort.CustomerID)
			})
			dsl.Field(4, "sort_direction", dsl.String, func() {
				dsl.Enum(types.SortDirection.Asc, types.SortDirection.Desc)
				dsl.Default(types.SortDirection.Desc)
			})
			dsl.Field(5, "page_number", dsl.Int, func() {
				dsl.Default(1)
				dsl.Minimum(1)
			})
			dsl.Field(6, "page_size", dsl.Int, func() {
				dsl.Default(20)
				dsl.Minimum(1)
			})
		})

		dsl.Result(func() {
			dsl.Field(1, "customers", dsl.ArrayOf(ListCustomerResp), "events")
			dsl.Field(2, "pagination", Pagination, "pagination")
			dsl.Required("customers", "pagination")
		})

		dsl.HTTP(func() {
			dsl.GET("")
			dsl.Params(func() {
				dsl.Param("country")
				dsl.Param("state")
				dsl.Param("sort_direction")
				dsl.Param("sort_key")
				dsl.Param("page_number")
				dsl.Param("page_size")

			})
			dsl.Response(dsl.StatusOK)
		})

	})

})

// ListCustomerResp ...
var ListCustomerResp = dsl.Type("ListCustomerResp", func() {
	dsl.Field(1, "ID", dsl.Int, "customer ID")
	dsl.Field(2, "Name", dsl.String, "customer name")
	dsl.Field(3, "Phone", dsl.String, "customer phone")
	dsl.Field(4, "Country", dsl.String, "customer country")
	dsl.Field(5, "State", dsl.String, "phone state")
	dsl.Field(6, "CountryCode", dsl.String, "country code")

	dsl.Required("ID", "Name", "Phone", "Country", "State", "CountryCode")

})

// Pagination is
var Pagination = dsl.ResultType("Pagination", func() {
	dsl.Field(1, "current_page", dsl.Int, func() {
		dsl.Description("The current page")
		dsl.Example(1)
	})
	dsl.Field(1, "page_size", dsl.Int, func() {
		dsl.Description("Max number of records per page ")
		dsl.Example(3)
	})
	dsl.Field(1, "total_pages", dsl.Int, func() {
		dsl.Description("Total pages")
		dsl.Example(11)
	})
	dsl.Field(1, "total_count", dsl.Int64, func() {
		dsl.Description("Total records count")
		dsl.Example(33)
	})
})
