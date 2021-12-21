package models

import (
	"fmt"
	"sync"

	"github.com/jolinGalal/jumia/internal/models/country"
)

// Customer ...
type Customer struct {
	ID      int
	Name    string
	Phone   string
	country country.CountryI
}

// CustomerI  customer interface
type CustomerI interface {
	GetTableName() string
	CustomerID() string
	CustomerName() string
	CustomerPhone() string
	GetCountry() string
	GetState() bool
}

// CustomerLock ...
var CustomerLock = &sync.Mutex{}

// customerInstance ...
var customerInstance *Customer

// NewCustomer ...
func NewCustomer() CustomerI {
	if customerInstance == nil {
		CustomerLock.Lock()
		defer CustomerLock.Unlock()
		if customerInstance == nil {
			customerInstance = &Customer{}
		}
	}

	return customerInstance
}

// GetTableName ...
func (c *Customer) GetTableName() string {
	return "customer"
}

// CustomerID ...
func (c *Customer) CustomerID() string {
	return fmt.Sprintf("%s.id", c.GetTableName())
}

// CustomerName ...
func (c *Customer) CustomerName() string {
	return fmt.Sprintf("%s.name", c.GetTableName())
}

// CustomerPhone ...
func (c *Customer) CustomerPhone() string {
	return fmt.Sprintf("%s.phone", c.GetTableName())
}

// Country ...
func (c *Customer) GetCountry() string {
	if c.country == nil {
		c.country = country.New(&c.Phone)
	}
	if c.country == nil {
		return "invalid code"
	}
	return c.country.Name()

}

// SetState ...
func (c *Customer) GetState() bool {
	if c.country == nil {
		c.country = country.New(&c.Phone)
	}
	if c.country == nil {
		return false
	}
	return c.country.State()

}
