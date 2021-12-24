package factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMoroccoState ...
func TestMoroccoState(t *testing.T) {
	var phone = "(212) 698054317"
	State := NewMorocco(phone).State()
	assert.Equal(t, State, true)

}

// TestMoroccoStateFail ...
func TestMoroccoStateFail(t *testing.T) {
	var phone = "(237) 6971515"
	State := NewMorocco(phone).State()
	assert.Equal(t, State, false)

}

// TestMoroccoStateEmty ...
func TestMoroccoStateEmty(t *testing.T) {
	var phone = ""
	State := NewMorocco(phone).State()
	assert.Equal(t, State, false)

}
