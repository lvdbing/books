package main

import "fmt"

// ******* 各种鸭子的实现 *******

// MallardDuck 绿头鸭
type MallardDuck struct {
	Duck
}

func NewMallardDuck() *MallardDuck {
	d := &MallardDuck{}
	d.flyBehavior = &FlyWithWings{}
	d.quackBehavior = &Quack{}
	return d
}

func (m *MallardDuck) Display() {
	fmt.Println("I'm a real Mallard duck")
}

// ModelDuck 模型鸭
type ModelDuck struct {
	Duck
}

func NewModelDuck() *ModelDuck {
	d := &ModelDuck{}
	d.flyBehavior = &FlyNoWay{}
	d.quackBehavior = &Quack{}
	return d
}

func (m *ModelDuck) Display() {
	fmt.Println("I'm a model duck")
}

// RedheadDuck 红头鸭
type RedheadDuck struct {
	Duck
}

func (d *RedheadDuck) New() *RedheadDuck {
	d.flyBehavior = &FlyWithWings{}
	d.quackBehavior = &Quack{}
	return d
}

// RubberDuck 橡皮鸭
type RubberDuck struct {
	Duck
}

func (d RubberDuck) New() *RubberDuck {
	nd := &RubberDuck{}
	nd.flyBehavior = &FlyNoWay{}
	nd.quackBehavior = &MuteQuack{}
	return nd
}

func (d RubberDuck) Display() {
	fmt.Println("I'm a rubber duck")
}

// DecoyDuck 诱饵鸭
type DecoyDuck struct {
	Duck
}

func (d *DecoyDuck) New() *DecoyDuck {
	d.flyBehavior = &FlyNoWay{}
	d.quackBehavior = &Squeak{}
	return d
}

func (d *DecoyDuck) Display() {
	fmt.Println("I'm a decoy duck")
}
