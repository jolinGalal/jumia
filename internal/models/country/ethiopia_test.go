package country

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestEthiopiaState ...
func TestEthiopiaState(t *testing.T) {
	var phone = "(251) 911168450"
	State := newEthiopia(phone).State()
	assert.Equal(t, State, true)

}

// TestEthiopiaStateFail ...
func TestEthiopiaStateFail(t *testing.T) {
	var phone = "(237) 6971515"
	State := newEthiopia(phone).State()
	assert.Equal(t, State, false)

}

// TestEthiopiaStateEmty ...
func TestEthiopiaStateEmty(t *testing.T) {
	var phone = ""
	State := newEthiopia(phone).State()
	assert.Equal(t, State, false)

}
