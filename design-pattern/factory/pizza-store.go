package main

// ******* 实现工厂方法模式的几个关键点 *******
// 1. 定义一个工厂方法抽象接口，比如下文的 PizzaFactory。
// 2. 工厂方法中，返回需要创建的对象/接口，比如 Pizza。其中，工厂方法通常命名为 Create。
// 3. 按照具体需要，定义工厂方法抽象接口的具体实现对象，比如 NYPizzaStore 和 ChicagoPizzaStore。
// 4. 客户端使用时，依赖工厂方法抽象接口。
// 5. 在客户端初始化阶段，完成具体工厂对象的依赖注入。

const (
	PIZZA_CHEESE    = "cheese"
	PIZZA_PEPPERONI = "pepperoni"
)

// PizzaFactory 工厂抽象接口
type PizzaFactory interface {
	CreatePizza(name string) Pizza // 创建工厂对象的方法
}

// ***** 实现具体的工厂 *******

// NYPizzaStore 纽约披萨店
type NYPizzaStore struct{}

func (n *NYPizzaStore) CreatePizza(name string) Pizza {
	ingredientFactory := &NYPizzaIngredientFactory{} // 抽象工厂接口具体实现
	var pizza Pizza
	switch name {
	case PIZZA_CHEESE:
		pizza = NewCheesePizza(ingredientFactory)
		pizza.SetName("New York Cheese Pizza")
	case PIZZA_PEPPERONI:
		pizza = NewPepperoniPizza(ingredientFactory)
		pizza.SetName("New York Pepperoni Pizza")
	}
	return pizza
}

// ChicagoPizzaStore 芝加哥披萨店
type ChicagoPizzaStore struct{}

func (c *ChicagoPizzaStore) CreatePizza(name string) Pizza {
	ingredientFactory := &ChicagoPizzaIngredientFactory{} // 抽象工厂接口具体实现
	var pizza Pizza
	switch name {
	case PIZZA_CHEESE:
		pizza = NewCheesePizza(ingredientFactory)
		pizza.SetName("Chicago Cheese Pizza")
	case PIZZA_PEPPERONI:
		pizza = NewPepperoniPizza(ingredientFactory)
		pizza.SetName("Chicago Pepperoni Pizza")
	}
	return pizza
}

// ******* 客户端使用工厂方法 *******

// PizzaStore 披萨店
type PizzaStore struct {
	factory PizzaFactory // 客户端依赖Factory抽象接口
}

func NewPizzaStore(factory PizzaFactory) *PizzaStore {
	return &PizzaStore{factory: factory}
}

func (p *PizzaStore) OrderPizza(name string) Pizza {
	pizza := p.factory.CreatePizza(name)
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()

	return pizza
}
