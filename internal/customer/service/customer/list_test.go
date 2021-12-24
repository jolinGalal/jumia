package customer

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	customers "github.com/jolinGalal/jumia/internal/customer/gen/customers"
	"github.com/stretchr/testify/assert"
)

var (
	customerCount = `SELECT count(*) FROM "customer"`
	customerQuery = `SELECT * FROM "customer" ORDER BY customer.id asc LIMIT 20 OFFSET 0`
)

//TestList ...
func TestList(t *testing.T) {
	newCustomerService()
	p := &customers.ListPayload{
		Country:       "all",
		State:         "all",
		SortKey:       "CustomerID",
		SortDirection: "asc",
		PageNumber:    1,
		PageSize:      20,
	}

	rows := sqlmock.NewRows([]string{"ID", "Name", "Phone"}).
		AddRow(31, "EMILE CHRISTIAN KOUKOU DIKANDA HONORE ", "(237) 697151594")

	obj.Mock.ExpectQuery(regexp.QuoteMeta(customerQuery)).WillReturnRows(rows)

	rows2 := sqlmock.NewRows([]string{"count"}).AddRow(1)
	obj.Mock.ExpectQuery(regexp.QuoteMeta(customerCount)).WillReturnRows(rows2)

	customers, err := obj.SVC.List(context.Background(), p)

	assert.NotNil(t, customers)
	assert.NoError(t, err)
}

//TestEventFail ...
func TestEventFail(t *testing.T) {
	newCustomerService()
	p := &customers.ListPayload{
		Country:       "all",
		State:         "all",
		SortKey:       "CustomerID",
		SortDirection: "asc",
		PageNumber:    1,
		PageSize:      20,
	}
	obj.Mock.ExpectQuery(regexp.QuoteMeta(customerQuery)).WillReturnError(assert.AnError)

	customers, err := obj.SVC.List(context.Background(), p)

	assert.Nil(t, customers)
	assert.Error(t, err)
}
