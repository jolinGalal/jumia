package factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestEthiopiaState ...
func TestEthiopiaState(t *testing.T) {
	var phone = "(251) 411168450"
	State := NewEthiopia(phone).State()
	assert.Equal(t, State, true)

}

// TestEthiopiaStateFail ...
func TestEthiopiaStateFail(t *testing.T) {
	var phone = "(237) 6971515"
	State := NewEthiopia(phone).State()
	assert.Equal(t, State, false)

}

// TestEthiopiaStateEmty ...
func TestEthiopiaStateEmty(t *testing.T) {
	var phone = ""
	State := NewEthiopia(phone).State()
	assert.Equal(t, State, false)

}
