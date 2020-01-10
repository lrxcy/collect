package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleton(t *testing.T) {
	ins1 := GetInstance()
	ins2 := GetInstance()
	assert.Equal(t, ins1, ins2)
}
