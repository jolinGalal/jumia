package country

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMoroccoState ...
func TestMoroccoState(t *testing.T) {
	var phone = "(212) 698054317"
	State := newMorocco(phone).State()
	assert.Equal(t, State, true)

}

// TestMoroccoStateFail ...
func TestMoroccoStateFail(t *testing.T) {
	var phone = "(237) 6971515"
	State := newMorocco(phone).State()
	assert.Equal(t, State, false)

}

// TestMoroccoStateEmty ...
func TestMoroccoStateEmty(t *testing.T) {
	var phone = ""
	State := newMorocco(phone).State()
	assert.Equal(t, State, false)

}
