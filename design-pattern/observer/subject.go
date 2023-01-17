package main

import "fmt"

// Subject 主题接口(发布者)
type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObservers()
}

// Subject 主题实现
type ConcreteSubject struct {
	Observers []Observer // 观察者(订阅者)列表
	changed   bool       // 状态数据是否改变
}

// RegisterObserver 注册观察者
func (c *ConcreteSubject) RegisterObserver(o Observer) {
	c.Observers = append(c.Observers, o)
}

// RemoveObserver 观察者取消注册
func (c *ConcreteSubject) RemoveObserver(o Observer) {
	for i, ob := range c.Observers {
		if ob == o {
			c.Observers = append(c.Observers[:i], c.Observers[i+1:]...)
			fmt.Printf("Removed a observer: %v\n", o)
		}
	}
}

func (c *ConcreteSubject) NotifyObservers() {

}

func (c *ConcreteSubject) SetChanged() {
	c.changed = true
}

func (c *ConcreteSubject) ClearChanged() {
	c.changed = false
}

func (c *ConcreteSubject) HasChanged() bool {
	return c.changed
}
