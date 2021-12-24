package customer

import "fmt"

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
