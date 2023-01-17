package main

import "fmt"

// ******* 模板方法模式 *******
// 定义：在一个方法中定义一个算法的骨架，而将一些步骤延迟到子类中。模板方法使得子类可以在不改变算法结构的情况下，重新定义算法中的某些步骤。

// Beverage 定义算法骨架的接口
type Beverage interface {
	PrepareRecipe()
	Brew()
	AddCondiments()
}

// ******* 抽象类 *******
// CaffeineBeverage 咖啡因饮料
type CaffeineBeverage struct {
	beverage Beverage
}

// ******* 具体实现方法 *******

// PrepareRecipe 制作饮料方法
func (c *CaffeineBeverage) PrepareRecipe() {
	c.BoilWater()
	c.Brew()
	c.PourInCup()
	c.AddCondiments()

}

// BoilWater 冲水
func (c *CaffeineBeverage) BoilWater() {
	fmt.Println("Boiling water")
}

// PourInCup 装杯
func (c *CaffeineBeverage) PourInCup() {
	fmt.Println("Pouring into cup")
}

// ******* 抽象方法 *******
// 抽象方法由具体的子类实现，使用时需要调用子类具体实现

// Brew 冲泡方法
func (c *CaffeineBeverage) Brew() {
	c.beverage.Brew()
}

// AddCondiments 加调料
func (c *CaffeineBeverage) AddCondiments() {
	c.beverage.AddCondiments()
}

type Tea struct {
	CaffeineBeverage
}

func (t *Tea) Brew() {
	fmt.Println("Steeping the tea")
}

func (t *Tea) AddCondiments() {
	fmt.Println("Adding Lemon")
}

type Coffee struct {
	CaffeineBeverage
}

func (c *Coffee) Brew() {
	fmt.Println("Dripping Coffee through filter")
}

func (c *Coffee) AddCondiments() {
	fmt.Println("Adding Sugar and Milk")
}
