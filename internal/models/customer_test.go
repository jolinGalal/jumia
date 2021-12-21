package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestCustomer ...
func TestCustomer(t *testing.T) {
	var c = Customer{ID: 1, Name: "hello", Phone: "(212) 698054317"}
	var country = c.GetCountry()
	assert.Equal(t, country, "Morocco")
	var state = c.GetState()
	assert.Equal(t, state, true)
}

//TestCustomerFail ...
func TestCustomerFail(t *testing.T) {
	var c = Customer{ID: 1, Name: "hello", Phone: "(212) 69805434317"}
	var country = c.GetCountry()
	assert.Equal(t, country, "Morocco")
	var state = c.GetState()
	assert.Equal(t, state, false)
}

// TestCustomerEmpty ...
func TestCustomerEmpty(t *testing.T) {
	var c = Customer{ID: 1, Name: "hello"}
	var country = c.GetCountry()
	assert.Equal(t, country, "invalid code")
	var state = c.GetState()
	assert.Equal(t, state, false)
}
