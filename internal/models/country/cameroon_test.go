package country

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCameroonState ...
func TestCameroonState(t *testing.T) {
	var phone = "(237) 697151594"
	State := newCameroon(phone).State()
	assert.Equal(t, State, true)

}

// TestCameroonStateFail ...
func TestCameroonStateFail(t *testing.T) {
	var phone = "(237) 6971515"
	State := newCameroon(phone).State()
	assert.Equal(t, State, false)

}

// TestCameroonStateEmty ...
func TestCameroonStateEmty(t *testing.T) {
	var phone = ""
	State := newCameroon(phone).State()
	assert.Equal(t, State, false)

}
