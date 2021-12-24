package country

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestCustomer ...
func TestCustomer(t *testing.T) {
	var phone = "(212) 698054317"
	var c = New(&phone)
	var country = c.Name()
	assert.Equal(t, country, "Morocco")
	var state = c.State()
	assert.Equal(t, state, "Valid")
	var code = c.Code()
	assert.Equal(t, code, "(212)")
}

//TestCustomerFail ...
func TestCustomerFail(t *testing.T) {
	var phone = "(212) 69805434317"
	var c = New(&phone)
	var country = c.Name()
	assert.Equal(t, country, "Morocco")
	var state = c.State()
	assert.Equal(t, state, "Not Valid")
	var code = c.Code()
	assert.Equal(t, code, "(212)")
}

// TestCustomerEmpty ...
func TestCustomerEmpty(t *testing.T) {
	var phone = ""
	var c = New(&phone)
	assert.NotNil(t, c)
}
