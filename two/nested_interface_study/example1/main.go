package main

import (
	"fmt"
)

type InterfaceCheckout interface {
	GetID() int
	GetItems() []InterfaceCartItem
}

type InterfaceCartItem interface {
	GetProduct() string
	GetID() int
}

type fakeCheckout struct {
	InterfaceCheckout
}

func (fakeCheckout) GetItems() []InterfaceCartItem {
	return []InterfaceCartItem{fakeItem{}, anotherFakeItem{}}
}

type fakeItem struct {
	InterfaceCartItem
}

func (fakeItem) GetProduct() string {
	return "This is the end"
}

type anotherFakeItem struct {
	InterfaceCartItem
}

func (anotherFakeItem) GetProduct() string {
	return "This is another end"
}

func getRates(checkoutI InterfaceCheckout) {
	for _, item := range checkoutI.GetItems() {
		fmt.Printf("%v\n", item.GetProduct())
	}
}

func main() {
	fc := fakeCheckout{}
	getRates(fc)
}
