package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPizzaFactory(t *testing.T) {
	var restaurant Restaurant
	restaurant = &RestaurantA{}
	assert.Equal(t, "spciy: true; vegetarian: false\n", restaurant.CreateAPizza().CreatePizza().Result())

	restaurant = &RestaurantB{}
	assert.Equal(t, "spciy: false; vegetarian: true\n", restaurant.CreateAPizza().CreatePizza().Result())
}
