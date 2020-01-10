package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFacade(t *testing.T) {
	f := NewModuleInterface()
	assert.Equal(t, "TestATestB", f.Test())
}
