package main

import "fmt"

// ******* 适配器模式 *******
// 将一个类的接口，转换成客户期望的另一个接口。

// 示例：用火鸡对象来冒充鸭子对象（将火鸡对象转换成鸭子对象）。

type Duck interface {
	Quack()
	Fly()
}

type Turkey interface {
	Gobble()
	Fly()
}

type WildTurkey struct{}

func (w *WildTurkey) Gobble() {
	fmt.Println("Gobble gobble")
}

func (w *WildTurkey) Fly() {
	fmt.Println("I'm flying a short distance")
}

// TurkeyAdapter 火鸡适配器，需要实现想转换成的类型的接口，鸭子接口。
type TurkeyAdapter struct {
	turkey Turkey // 适配对象
}

func NewTurkeyAdapter(turkey Turkey) *TurkeyAdapter {
	return &TurkeyAdapter{turkey: turkey}
}

func (t *TurkeyAdapter) Quack() {
	t.turkey.Gobble() // 委托给turkey
}

func (t *TurkeyAdapter) Fly() {
	for i := 0; i < 5; i++ {
		t.turkey.Fly() // 委托给turkey
	}
}
