package main

import "fmt"

func main() {
	var pizzaStore *PizzaStore = NewPizzaStore(&NYPizzaStore{})
	pizzaStore.OrderPizza(PIZZA_CHEESE)
	pizzaStore.OrderPizza(PIZZA_PEPPERONI)

	fmt.Println()

	pizzaStore = NewPizzaStore(&ChicagoPizzaStore{})
	pizzaStore.OrderPizza(PIZZA_PEPPERONI)
	pizzaStore.OrderPizza(PIZZA_CHEESE)
}
