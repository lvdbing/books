package main

type Beverage interface {
	GetDescription() string
	Cost() int
}

// Espresso 浓缩咖啡
type Espresso struct{}

func (e *Espresso) GetDescription() string {
	return "Espresso Coffee"
}

func (e *Espresso) Cost() int {
	return 199
}

// HouseBlend 综合咖啡
type HouseBlend struct{}

func (h *HouseBlend) GetDescription() string {
	return "House Blend Coffee"
}

func (h *HouseBlend) Cost() int {
	return 89
}

// DarkRoast 深焙咖啡
type DarkRoast struct{}

func (d *DarkRoast) GetDescription() string {
	return "Dark Roast Coffee"
}

func (d *DarkRoast) Cost() int {
	return 99
}

// Decaf 低咖啡因咖啡
type Decaf struct{}

func (d *Decaf) GetDescription() string {
	return "Decaf Coffee"
}

func (d *Decaf) Cost() int {
	return 105
}
