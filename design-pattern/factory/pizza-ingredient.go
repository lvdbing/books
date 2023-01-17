package main

// ******* 实现抽象工厂模式的几个关键点 *******

// 1. 定义抽象工厂接口，里面包含创建各个产品的工厂方法定义。
// 2. 定义抽象工厂接口的实现类。
// 3. 在客户端程序中依赖抽象工厂接口，通过接口来完成产品的创建。
// 4. 在客户端程序初始化时，将抽象工厂接口的具体实现依赖注入进去。

// PizzaIngredientFactory 创建披萨原料的工厂接口
type PizzaIngredientFactory interface {
	CreateDough() Dough
	CreateSauce() Sauce
	CreateCheese() Cheese
	CreateVeggies() []Veggies
	CreatePepperoni() Pepperoni
	CreateClam() Clam
}

type NYPizzaIngredientFactory struct{}

func (f *NYPizzaIngredientFactory) CreateDough() Dough {
	return "New York Dough"
}

func (f *NYPizzaIngredientFactory) CreateSauce() Sauce {
	return "New York Sauce"
}

func (f *NYPizzaIngredientFactory) CreateCheese() Cheese {
	return "New York Cheese"
}

func (f *NYPizzaIngredientFactory) CreateVeggies() []Veggies {

	return []Veggies{"New York Garlic", "New York Onion", "New York Mushroom", "New York Red Pepper"}
}

func (f *NYPizzaIngredientFactory) CreatePepperoni() Pepperoni {
	return "New York Pepperoni"
}

func (f *NYPizzaIngredientFactory) CreateClam() Clam {
	return "New York Clam"
}

type ChicagoPizzaIngredientFactory struct{}

func (f *ChicagoPizzaIngredientFactory) CreateDough() Dough {
	return "Chicago Dough"
}

func (f *ChicagoPizzaIngredientFactory) CreateSauce() Sauce {
	return "Chicago Sauce"
}

func (f *ChicagoPizzaIngredientFactory) CreateCheese() Cheese {
	return "Chicago Cheese"
}

func (f *ChicagoPizzaIngredientFactory) CreateVeggies() []Veggies {

	return []Veggies{"Chicago Garlic", "Chicago Onion", "Chicago Mushroom", "Chicago Red Pepper"}
}

func (f *ChicagoPizzaIngredientFactory) CreatePepperoni() Pepperoni {
	return "Chicago Pepperoni"
}

func (f *ChicagoPizzaIngredientFactory) CreateClam() Clam {
	return "Chicago Clam"
}

type Dough interface{}

type Sauce interface{}

type Cheese interface{}

type Veggies interface{}

type Pepperoni interface{}

type Clam interface{}
