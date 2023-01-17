package main

import "fmt"

// Duck 鸭子的超类
type Duck struct {
	flyBehavior   FlyBehavior   // 飞行行为
	quackBehavior QuackBehavior // 呱呱叫行为
}

// SetFlyBehavior 修改飞行行为
func (d *Duck) SetFlyBehavior(fb FlyBehavior) {
	d.flyBehavior = fb
}

// SetQuackBehavior 修改呱呱叫行为
func (d *Duck) SetQuackBehavior(qb QuackBehavior) {
	d.quackBehavior = qb
}

func (d *Duck) PerformFly() {
	d.flyBehavior.Fly() // 委托给行为类
}

func (d *Duck) PerformQuack() {
	d.quackBehavior.Quack() // 委托给行为类
}

func (d *Duck) Display() {
	fmt.Println("I'm a duck")
}

func (d *Duck) Swim() {
	fmt.Println("All ducks float, even decoys!")
}
