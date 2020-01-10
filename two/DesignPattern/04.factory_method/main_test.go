package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPizzaFactory(t *testing.T) {
	var pizza PizzaFactory
	pizza = &OrderSpciyPizza{}
	assert.Equal(t, "spciy: true; vegetarian: false\n", pizza.CreatePizza().Result())

	pizza = &OrderVegetarianPizza{}
	assert.Equal(t, "spciy: false; vegetarian: true\n", pizza.CreatePizza().Result())
}
