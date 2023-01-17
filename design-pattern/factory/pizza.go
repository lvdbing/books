package main

import "fmt"

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
	GetName() string
	SetName(name string)
}

type ConcretePizza struct {
	Name      string
	Cheese    Cheese
	Pepperoni Pepperoni
	Veggies   []Veggies
}

func (c *ConcretePizza) Prepare() {
	fmt.Println("Preparing " + c.Name)
	fmt.Println("Tossing dough...")
	fmt.Println("Adding sauce...")
	fmt.Println("Adding toppings:")
	for _, v := range c.Veggies {
		fmt.Printf("\t%v\n", v)
	}
}

func (c *ConcretePizza) Bake() {
	fmt.Println("Bake for 25 minutes at 350")
}

func (c *ConcretePizza) Cut() {
	fmt.Println("Cutting the pizza into diagonal slices")
}

func (c *ConcretePizza) Box() {
	fmt.Println("Place pizza in official PizzaStore box")
}

func (c *ConcretePizza) GetName() string {
	return c.Name
}

func (c *ConcretePizza) SetName(name string) {
	c.Name = name
}

type CheesePizza struct {
	ConcretePizza
	ingredientFactory PizzaIngredientFactory // 抽象工厂接口
}

func NewCheesePizza(factory PizzaIngredientFactory) *CheesePizza {
	c := &CheesePizza{}
	c.ingredientFactory = factory
	return c
}

func (c *CheesePizza) Prepare() {
	fmt.Println("Preparing " + c.Name)

	c.Cheese = c.ingredientFactory.CreateCheese()
	c.Veggies = c.ingredientFactory.CreateVeggies()

	fmt.Printf("Adding %v...\n", c.Cheese)
	fmt.Println("Adding veggies:")
	for _, v := range c.Veggies {
		fmt.Printf("\t%v\n", v)
	}
}

type PepperoniPizza struct {
	ConcretePizza
	ingredientFactory PizzaIngredientFactory // 抽象工厂接口
}

func NewPepperoniPizza(factory PizzaIngredientFactory) *PepperoniPizza {
	p := &PepperoniPizza{}
	p.ingredientFactory = factory
	return p
}

func (p *PepperoniPizza) Prepare() {
	fmt.Println("Preparing " + p.Name)

	p.Pepperoni = p.ingredientFactory.CreatePepperoni()
	p.Veggies = p.ingredientFactory.CreateVeggies()

	fmt.Printf("Adding %v...\n", p.Pepperoni)
	fmt.Println("Adding veggies:")
	for _, v := range p.Veggies {
		fmt.Printf("\t%v\n", v)
	}
}
