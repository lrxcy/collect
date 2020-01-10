package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPowerAdaptee(t *testing.T) {
	f := NewPlug()
	ff := NewPowerAdaptee(f)
	assert.Equal(t, "device is charging", ff.Charge())
}
