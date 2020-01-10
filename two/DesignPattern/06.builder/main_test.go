package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuilder(t *testing.T) {
	builder1 := &Builder1{}
	director1 := NewDirector(builder1)
	director1.Construct()

	assert.Equal(t, "123", builder1.GetResult())

	builder2 := &Builder2{}
	director2 := NewDirector(builder2)
	director2.Construct()

	assert.Equal(t, 6, builder2.GetResult())
}
