package main

import (
	"fmt"
	"log"
)

type RestaurantA struct {
	OrderSpciyPizza
}

type RestaurantB struct {
	OrderVegetarianPizza
}

func (a *RestaurantA) CreateAPizza() PizzaFactory {
	return a.OrderSpciyPizza
}

func (b *RestaurantB) CreateAPizza() PizzaFactory {
	return b.OrderVegetarianPizza
}

// Restaurant為PizzaFactory的抽象工廠
type Restaurant interface {
	CreateAPizza() PizzaFactory
}

// 定義pizza可以操作的基本種類
type Pizza interface {
	AddSpicy()
	RemoveMeat()
	Result() string
}

type PizzaFactory interface {
	CreatePizza() Pizza
}

type PizzaBase struct {
	spicy      bool
	vegetarian bool
}

func (p *PizzaBase) AddSpicy() {
	p.spicy = true
}

func (p *PizzaBase) RemoveMeat() {
	p.vegetarian = true
}

// 定義好pizza的種類
type SpicyPizza struct {
	*PizzaBase
}

func (o *SpicyPizza) Result() string {
	o.AddSpicy()
	log.Println("the pizza is add spicy")
	return fmt.Sprintf("spciy: %v; vegetarian: %v\n", o.spicy, o.vegetarian)
}

type VegetarianPizza struct {
	*PizzaBase
}

func (o *VegetarianPizza) Result() string {
	o.RemoveMeat()
	log.Println("the pizza is not included meat")
	return fmt.Sprintf("spciy: %v; vegetarian: %v\n", o.spicy, o.vegetarian)
}

// 定義pizza的生成方法
type OrderSpciyPizza struct{}

func (OrderSpciyPizza) CreatePizza() Pizza {
	return &SpicyPizza{
		PizzaBase: &PizzaBase{},
	}
}

type OrderVegetarianPizza struct{}

func (OrderVegetarianPizza) CreatePizza() Pizza {
	return &VegetarianPizza{
		PizzaBase: &PizzaBase{},
	}
}

func main() {
	// var pizza PizzaFactory
	// pizza = OrderSpciyPizza{}
	// log.Println(MakePizza(pizza))

	// pizza = OrderVegetarianPizza{}
	// log.Println(MakePizza(pizza))
	var restaurant Restaurant
	restaurant = &RestaurantA{}
	log.Println(MakePizza(restaurant.CreateAPizza()))

	restaurant = &RestaurantB{}
	log.Println(MakePizza(restaurant.CreateAPizza()))

}

func MakePizza(pizza PizzaFactory) string {
	return pizza.CreatePizza().Result()
}
