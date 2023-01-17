package main

import "fmt"

// QuackBehavior 呱呱叫行为的接口
type QuackBehavior interface {
	Quack()
}

// Quack 呱呱叫, 实现QuackBehavior接口
type Quack struct{}

func (q *Quack) Quack() {
	fmt.Println("Quack")
}

// MuteQuack 不会叫, 实现QuackBehavior接口
type MuteQuack struct{}

func (mq *MuteQuack) Quack() {
	fmt.Println("<< Silence >>")
}

// Squeak 吱吱叫, 实现QuackBehavior接口
type Squeak struct{}

func (s *Squeak) Quack() {
	fmt.Println("Squeak")
}
