package country

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMozambiqueState ...
func TestMozambiqueState(t *testing.T) {
	var phone = "(258) 847651504"
	State := newMozambique(phone).State()
	assert.Equal(t, State, true)

}

// TestMozambiqueStateFail ...
func TestMozambiqueStateFail(t *testing.T) {
	var phone = "(237) 6971515"
	State := newMozambique(phone).State()
	assert.Equal(t, State, false)

}

// TestMozambiqueStateEmty ...
func TestMozambiqueStateEmty(t *testing.T) {
	var phone = ""
	State := newMozambique(phone).State()
	assert.Equal(t, State, false)

}
