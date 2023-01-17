package main

import "fmt"

type person interface {
	SayHello(name string)
	Run() string
}

type hero struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Speed int    `json:"speed"`
	pri   string
}

func (h *hero) SayHello(name string) {
	fmt.Printf("Hello %s, I am %s\n", name, h.Name)
}

func (h *hero) Run() string {
	fmt.Printf("I am running at speed %d\n", h.Speed)
	return "I am running!"
}
