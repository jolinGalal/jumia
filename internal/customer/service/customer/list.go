package customer

import (
	"context"
	"fmt"

	customers "github.com/jolinGalal/jumia/internal/customer/gen/customers"
	"github.com/jolinGalal/jumia/internal/customer/types"
	"github.com/jolinGalal/jumia/internal/models/country"
	"github.com/jolinGalal/jumia/pkg/utils"
)

// list customers
func (s *customerssrvc) List(ctx context.Context, p *customers.ListPayload) (res *customers.ListResult, err error) {
	res = &customers.ListResult{}
	sortKey, _ := types.CustomerSortFields[p.SortKey]

	var query = s.db.Table(types.CustomerModel.GetTableName()).
		Where(s.countryCondition(p)).
		Order(fmt.Sprintf("%s %s", sortKey, p.SortDirection))

	var totalCount int64

	if p.State == types.State.All {
		err = utils.Paging(query, &p.PageNumber, &p.PageSize).
			Find(&res.Customers).
			Error
		if err != nil {
			return nil, err
		}

		err = query.Count(&totalCount).Error
		if err != nil {
			return nil, err
		}
	} else {
		err = query.Find(&res.Customers).Error
		if err != nil {
			return nil, err
		}

		res.Customers = s.removeCustomer(p, res.Customers)
		totalCount = int64(len(res.Customers))
		low, high := utils.Intervals(&p.PageNumber, &p.PageSize, len(res.Customers))
		if low != -1 {
			res.Customers = res.Customers[low:high]
		} else {
			res.Customers = res.Customers[len(res.Customers):len(res.Customers)]
		}

	}

	s.mapCustomerData(&res.Customers, p)
	paging := utils.GetPagination(&p.PageNumber, &p.PageSize, &totalCount)
	res.Pagination = &customers.Pagination{
		CurrentPage: paging.GetCurrentPage(),
		PageSize:    paging.GetPageSize(),
		TotalPages:  paging.GetTotalPages(),
		TotalCount:  paging.GetCount(),
	}
	return res, nil
}

// countryCondition ...
func (s *customerssrvc) countryCondition(p *customers.ListPayload) string {
	if p.Country == "all" {
		return ""
	}
	var code = types.CountryModel.CountryCodeByName(p.Country)
	if code == "" {
		return ""
	}
	return fmt.Sprintf("trim(%s) like '%s%s'", types.CustomerModel.CustomerPhone(), code, "%")

}

// Paging ...
func (s *customerssrvc) mapCustomerData(customers *[]*customers.ListCustomerResp, p *customers.ListPayload) {
	for index, value := range *customers {
		countryObj := country.New(&value.Phone)
		(*customers)[index].Country = countryObj.Name()
		(*customers)[index].State = countryObj.State()
		(*customers)[index].CountryCode = countryObj.Code()
	}

}

// removeCustomer ...
func (s *customerssrvc) removeCustomer(p *customers.ListPayload, res []*customers.ListCustomerResp) []*customers.ListCustomerResp {
	for index := 0; index < len(res); index++ {
		countryObj := country.New(&res[index].Phone)
		if countryObj.State() != p.State {
			res = append(res[:index], res[index+1:]...)
			index--
		}
	}
	return res
}
