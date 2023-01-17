package main

import "fmt"

// FlyBehavior 飞行行为的接口
type FlyBehavior interface {
	Fly()
}

// FlyWithWings 用翅膀飞行, 实现FlyBehavior接口
type FlyWithWings struct{}

func (f *FlyWithWings) Fly() {
	fmt.Println("I'm flying!")
}

// FlyNoWay 不能飞, 实现FlyBehavior接口
type FlyNoWay struct{}

func (f *FlyNoWay) Fly() {
	fmt.Println("I can't fly")
}

// FlyRocketPowered 用火箭飞行, 实现FlyBehavior接口
type FlyRocketPowered struct{}

func (f *FlyRocketPowered) Fly() {
	fmt.Println("I'm flying with a rocket!")
}
