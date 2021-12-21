package country

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestUgandaState ...
func TestUgandaState(t *testing.T) {
	var phone = "(256) 775069443"
	State := newUganda(phone).State()
	assert.Equal(t, State, true)

}

// TestUgandaStateFail ...
func TestUgandaStateFail(t *testing.T) {
	var phone = "(256) 3142345678"
	State := newUganda(phone).State()
	assert.Equal(t, State, false)

}

// TestUgandaStateEmty ...
func TestUgandaStateEmty(t *testing.T) {
	var phone = ""
	State := newUganda(phone).State()
	assert.Equal(t, State, false)

}
