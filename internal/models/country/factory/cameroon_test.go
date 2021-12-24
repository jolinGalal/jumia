package factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCameroonState ...
func TestCameroonState(t *testing.T) {
	var phone = "(237) 678000952"
	State := NewCameroon(phone).State()
	assert.Equal(t, State, true)

}

// TestCameroonStateFail ...
func TestCameroonStateFail(t *testing.T) {
	var phone = "(237) 6971515"
	State := NewCameroon(phone).State()
	assert.Equal(t, State, false)

}

// TestCameroonStateEmty ...
func TestCameroonStateEmty(t *testing.T) {
	var phone = ""
	State := NewCameroon(phone).State()
	assert.Equal(t, State, false)

}
